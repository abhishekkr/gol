package golkeyvalTSDS

import (
	"time"

	golhashmap "github.com/abhishekkr/gol/golhashmap"
	golkeyval "github.com/abhishekkr/gol/golkeyval"
	golkeyvalNS "github.com/abhishekkr/gol/golkeyvalNS"
)

/*
DBEngine interface enables adapter feature for a key-val oriented DB.
*/
type TSDSDBEngine interface {
	Configure(namespaceType golkeyvalNS.NSDBEngine)
	PushTSDS(key string, val string, tyme time.Time) bool
	PushNowTSDS(key string, val string) bool
	ReadTSDS(key string) golhashmap.HashMap
	DeleteTSDS(key string) bool

	// golkeyval proxy func
	PushNS(key string, val string) bool
	ReadNSRecursive(key string) golhashmap.HashMap
	DeleteNSRecursive(key string) bool

	PushKeyVal(key string, val string) bool
	GetVal(key string) string
	DelKey(key string) bool
}

/*
DBEngines acts as map for all available db-engines.
*/
var TSDSDBEngines = make(map[string]TSDSDBEngine)

/*
RegisterDBEngine gets used by adapters to register themselves.
*/
func RegisterTSDSDBEngine(name string, nsDbEngine TSDSDBEngine) {
	TSDSDBEngines[name] = nsDbEngine
}

/*
GetDBEngine gets used by client to fetch a required db-engine.
*/
func GetTSDSDBEngine(name string) TSDSDBEngine {
	return TSDSDBEngines[name]
}

/*
GetDBEngine gets used by client to fetch a 'namespace' db-engine.
parameter: Config with DBEngine, NSEngine, DBPath
*/
func GetNamespaceEngine(config map[string]string) TSDSDBEngine {
	db := golkeyval.GetDBEngine(config["DBEngine"])
	db.Configure(config)
	db.CreateDB()

	ns := golkeyvalNS.GetNSDBEngine(config["NSEngine"])
	ns.Configure(db)

	tsds := GetTSDSDBEngine("namespace")
	tsds.Configure(ns)
	return tsds
}
