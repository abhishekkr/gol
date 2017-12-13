package goldb

var (
	MigrationSQLDefault = map[string]string{
		"init": "CREATE TABLE IF NOT EXISTS migrations (uuid STRING PRIMARY KEY) ;",
		"up":   "INSERT INTO migrations (uuid) VALUES ('%s') ;",
		"down": "DELETE FROM migrations WHERE uuid = '%s' ;",
	}

	MigrationSQL = map[string]map[string]string{
		"postgresql": MigrationSQLDefault,
		"postgres":   MigrationSQLDefault,
	}
)
