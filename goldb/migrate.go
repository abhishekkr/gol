package goldb

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strings"

	_ "github.com/lib/pq"
)

type DataStore struct {
	ConnectionURL string
	DB            *sql.DB
	DBType        string
	MigrationsDir string
	MigrationType string
}

func reverseList(migrations []string) []string {
	var revMigrations = make([]string, len(migrations))

	migrationsLen := len(migrations)

	revIdx := 0
	for idx := range migrations {
		if migrations[migrationsLen-idx-1] == "" {
			continue
		}
		revMigrations[revIdx] = migrations[migrationsLen-idx-1]
		revIdx += 1
	}
	return revMigrations
}

func (datastore *DataStore) migrationUUID(sqlFile string) string {
	migrationTokens := strings.Split(sqlFile, ".")
	return migrationTokens[0]
}

func (datastore *DataStore) migrationFiles(migrationsDir string) []string {
	sqlFiles, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		log.Fatal(err)
	}

	var migrationFiles = make([]string, len(sqlFiles))
	migrationFilesIdx := 0
	for _, sqlFile := range sqlFiles {
		if sqlFile.IsDir() {
			continue
		}
		migrationFiles[migrationFilesIdx] = sqlFile.Name()
		migrationFilesIdx += 1
	}

	return migrationFiles
}

func (datastore *DataStore) checkMigrationInDB(file string) bool {
	queryStmt := fmt.Sprintf("SELECT uuid FROM migrations WHERE uuid = '%s' ;",
		datastore.migrationUUID(file),
	)

	var uuid string
	err := datastore.DB.QueryRow(queryStmt).Scan(&uuid)

	return err == nil
}

func (datastore *DataStore) isMigrationRequired(file string, migrationType string) bool {
	fileTokens := strings.Split(file, ".")
	if len(fileTokens) != 3 {
		return false
	}
	if fileTokens[1] != migrationType {
		return false
	}

	existsInTable := datastore.checkMigrationInDB(file)
	if migrationType == "up" && existsInTable {
		return false
	}
	if migrationType == "down" && !existsInTable {
		return false
	}
	return true
}

func (datastore *DataStore) migrationFilesForType(allMigrations []string, migrationType string) []string {
	var requiredMigrations = make([]string, len(allMigrations))
	requiredMigrationsIdx := 0
	for _, sqlFile := range allMigrations {
		if datastore.isMigrationRequired(sqlFile, migrationType) {
			requiredMigrations[requiredMigrationsIdx] = sqlFile
			requiredMigrationsIdx += 1
		}
	}
	return requiredMigrations
}

func (datastore *DataStore) Connect() (err error) {
	datastore.DB, err = sql.Open("postgres", datastore.ConnectionURL)
	if err != nil {
		log.Printf("[error] failure connecting to the database: ", err)
	}
	return err
}

func (datastore *DataStore) RunSql(sql string) (err error) {
	if _, err := datastore.DB.Exec(sql); err != nil {
		log.Fatal(err)
	}
	return
}

func (datastore *DataStore) MigrateByType(migrationsDir string, migrationFile string, migrationType string) {
	if migrationType != "up" && migrationType != "down" {
		log.Fatalln("unsupported migration type", migrationType)
	}

	uuid := datastore.migrationUUID(migrationFile)

	migrationPath := path.Join(migrationsDir, migrationFile)
	sqlBytes, err := ioutil.ReadFile(migrationPath)
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = datastore.RunSql(string(sqlBytes))
	if err != nil {
		log.Fatalln(err)
		return
	}

	sql_stmt := fmt.Sprintf(MigrationSQL[datastore.DBType][migrationType],
		uuid)

	if err := datastore.RunSql(sql_stmt); err != nil {
		log.Fatalln(err)
	}
}

func (datastore *DataStore) InitMigrationsTable() {
	err := datastore.RunSql(MigrationSQL[datastore.DBType]["init"])
	if err != nil {
		log.Fatalln(err)
	}
}

func (datastore *DataStore) Migrate() {
	if datastore.DBType == "" {
		datastore.DBType = strings.Split(datastore.ConnectionURL, ":")[0]
	}

	datastore.Connect()
	defer datastore.DB.Close()
	datastore.InitMigrationsTable()

	allMigrations := datastore.migrationFiles(datastore.MigrationsDir)
	requiredMigrations := datastore.migrationFilesForType(allMigrations, datastore.MigrationType)

	if datastore.MigrationType == "down" {
		requiredMigrations = reverseList(requiredMigrations)
	}

	for _, sqlFile := range requiredMigrations {
		if sqlFile == "" {
			continue
		}
		log.Printf("[+] %s\n", sqlFile)
		datastore.MigrateByType(datastore.MigrationsDir, sqlFile, datastore.MigrationType)
	}
}

func Migrate() {
	migrationsDir := flag.String("dir", "./migrations", "dir to pick migrations from")
	dbConnectionUrl := flag.String("db", "postgresql://gol@127.0.0.1:26257/gol?sslmode=disable", "db connection url")
	migrationType := flag.String("type", "up", "migrate up or down")
	flag.Parse()

	datastore := DataStore{
		ConnectionURL: *dbConnectionUrl,
		MigrationsDir: *migrationsDir,
		MigrationType: *migrationType,
	}

	datastore.Migrate()
}
