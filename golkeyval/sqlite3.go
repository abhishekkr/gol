package golkeyval

import (
	"database/sql"
	"fmt"
	"os"

	golerror "github.com/abhishekkr/gol/golerror"
	_ "github.com/mattn/go-sqlite3"
)

/*
Sqlite3DB struct with sqlite3 required details.
*/
type Sqlite3DB struct {
	DBPath    string
	GolDB     *sql.DB
	TableName string
}

/*
init registers sqlite3 to DBEngines.
*/
func init() {
	RegisterDBEngine("sqlite3", new(Sqlite3DB))
}

/*
Configure populates SQLite required configs.
*/
func (sqlite3 *Sqlite3DB) Configure(cfg map[string]string) {
	sqlite3.DBPath = cfg["DBPath"]
	sqlite3.TableName = cfg["TableName"]
}

/*
CreateDB creates a sqlite3 db at provided DBPath.
It will create Table with the name provided during Configure(),
so multiple key-val can be kept isolated at same db-path.
Create fixed fields FieldKey and FieldValue for key-name and key-value in it.
*/
func (sqlite3 *Sqlite3DB) CreateDB() {
	var errDB error
	sqlite3.GolDB, errDB = sql.Open("sqlite3", sqlite3.DBPath)
	if errDB != nil {
		errMsg := fmt.Sprintf("DB %s Creation failed. %q", sqlite3.DBPath, errDB)
		golerror.Boohoo(errMsg, true)
	}
	sqlStmt := fmt.Sprintf("create table IF NOT EXISTS %s (FieldKey text not null primary key, FieldValue text); delete from %s;", sqlite3.TableName, sqlite3.TableName)
	_, errDB = sqlite3.GolDB.Exec(sqlStmt)
	if errDB != nil {
		golerror.Boohoo(errDB.Error(), true)
		return
	}
}

/*
CloseDB closes a db given handle.
*/
func (sqlite3 *Sqlite3DB) CloseDB() {
	sqlite3.GolDB.Close()
}

/*
CloseAndDeleteDB closes and deletes a db given handle and DBPath.
Useful in use and throw implementations. And also tests.
*/
func (sqlite3 *Sqlite3DB) CloseAndDeleteDB() {
	sqlite3.CloseDB()
	if os.RemoveAll(sqlite3.DBPath) != nil {
		panic("Fail: Temporary DB files are still present at: " + sqlite3.DBPath)
	}
}

/*
PushKeyVal creates a key-val in provided DB handle.
*/
func (sqlite3 *Sqlite3DB) PushKeyVal(key string, val string) bool {
	sqlCmd := fmt.Sprintf("insert into %s(FieldKey, FieldValue) values('%s', '%s')", sqlite3.TableName, key, val)
	_, errDB := sqlite3.GolDB.Exec(sqlCmd)
	if errDB != nil {
		golerror.Boohoo("Key "+key+" insertion failed. It's value was "+val, false)
		return false
	}
	return true
}

/*
GetVal fetches value of key from provided db handle.
*/
func (sqlite3 *Sqlite3DB) GetVal(key string) string {
	var val string
	sqlCmd := fmt.Sprintf("select FieldValue from %s where FieldKey = ?", sqlite3.TableName)
	rows, errDB := sqlite3.GolDB.Query(sqlCmd, key)
	if errDB != nil {
		golerror.Boohoo("Statement failed for Key "+key+" query.", false)
		return ""
	}
	for rows.Next() {
		if val != "" {
			golerror.Boohoo("Multiple Key "+key+" found. Should be Primary.", false)
			return val
		}
		if errDB = rows.Scan(&val); errDB != nil {
			golerror.Boohoo("Key "+key+" query failed.", false)
			return ""
		}
	}
	if errDB = rows.Err(); errDB != nil {
		golerror.Boohoo("Key "+key+" query failed, rows have error.", false)
	}
	rows.Close()
	return string(val)
}

/*
DelKey deletes a key from provided DB handle.
*/
func (sqlite3 *Sqlite3DB) DelKey(key string) bool {
	sqlCmd := fmt.Sprintf("delete from %s where FieldKey = '%s'", sqlite3.TableName, key)
	_, errDB := sqlite3.GolDB.Exec(sqlCmd)
	if errDB != nil {
		golerror.Boohoo("Key "+key+" query failed.", false)
		return false
	}
	return true
}
