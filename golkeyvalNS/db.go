package golkeyvalNS

import (
	"github.com/abhishekkr/gol/golhashmap"
	"github.com/abhishekkr/gol/golkeyval"
)

/*
DBEngine interface enables adapter feature for a key-val oriented DB.
*/
type NSDBEngine interface {
	Configure(db golkeyval.DBEngine)
	PushNS(key string, val string) bool
	ReadNS(key string) golhashmap.HashMap
	ReadNSRecursive(key string) golhashmap.HashMap
	DeleteNSKey(key string) bool
	DeleteNS(key string) bool
	DeleteNSRecursive(key string) bool
	// golkeyval proxy func
	PushKeyVal(key string, val string) bool
	GetVal(key string) string
	DelKey(key string) bool
}

/*
DBEngines acts as map for all available db-engines.
*/
var NSDBEngines = make(map[string]NSDBEngine)

/*
RegisterDBEngine gets used by adapters to register themselves.
*/
func RegisterNSDBEngine(name string, nsDbEngine NSDBEngine) {
	NSDBEngines[name] = nsDbEngine
}

/*
GetDBEngine gets used by client to fetch a required db-engine.
*/
func GetNSDBEngine(name string) NSDBEngine {
	return NSDBEngines[name]
}
