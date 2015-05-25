package golkeyval

/*
DBEngine interface enables adapter feature for a key-val oriented DB.
*/
type DBEngine interface {
	Configure(cfg map[string]string)
	CreateDB()
	CloseDB()
	CloseAndDeleteDB()
	PushKeyVal(key string, val string) bool
	GetVal(key string) string
	DelKey(key string) bool
}

/*
DBEngines acts as map for all available db-engines.
*/
var DBEngines = make(map[string]DBEngine)

/*
RegisterDBEngine gets used by adapters to register themselves.
*/
func RegisterDBEngine(name string, dbEngine DBEngine) {
	DBEngines[name] = dbEngine
}

/*
GetDBEngine gets used by client to fetch a required db-engine.
*/
func GetDBEngine(name string) DBEngine {
	return DBEngines[name]
}
