package main

import (
	"fmt"

	golkeyval "github.com/abhishekkr/gol/golkeyval"
)

var (
	dataToUse    map[string]string
	dbPathPrefix = "/tmp/dbs/golkeyval_"
	tableName    = "alpha"
)

/*
init populates data to be used for db tasks
*/
func init() {
	dataToUse = make(map[string]string)
	dataToUse["Bob"] = "Alice"
	dataToUse["Eve"] = "Trudy"
}

/*
TestGolKeyVal checks for following actions via golkeyval
	* fetch db-engine of desired tyoe
	* populate all required db properties for required engine
	* create and open a new db (with other details if required)
	* pushing key-val to db
	* query of existing keys
	* closing an open db
	* opening an existing db
	* deletion of a key-val
	* query of missing keys
	* closing and deleting db
*/
func TestGolKeyVal(engineName string, cfg map[string]string) {
	fmt.Printf("running for engine: %s\n\n", engineName)
	db := golkeyval.GetDBEngine(engineName)
	db.Configure(cfg)

	db.CreateDB()
	for _key := range dataToUse {
		if !db.PushKeyVal(_key, dataToUse[_key]) {
			db.CloseAndDeleteDB()
			panic(fmt.Sprintf("ERROR: PushKeyVal failed for %s: %s", _key, dataToUse[_key]))
		}
	}
	for _key := range dataToUse {
		_val := db.GetVal(_key)
		if _val != dataToUse[_key] {
			db.CloseAndDeleteDB()
			panic(fmt.Sprintf("ERROR: GetVal failed for %s: %s; got %s", _key, dataToUse[_key], _val))
		}
	}
	db.CloseDB()
	db.CreateDB()
	for _key := range dataToUse {
		if !db.DelKey(_key) {
			db.CloseAndDeleteDB()
			panic(fmt.Sprintf("ERROR: DelKey failed for %s: %s", _key, dataToUse[_key]))
		}
	}
	for _key := range dataToUse {
		if db.GetVal(_key) != "" {
			db.CloseAndDeleteDB()
			panic(fmt.Sprintf("ERROR: DelKey failed for %s: %s", _key, dataToUse[_key]))
		}
	}
	for _key := range dataToUse {
		if !db.PushKeyVal(_key, dataToUse[_key]) {
			db.CloseAndDeleteDB()
			panic(fmt.Sprintf("ERROR: PushKeyVal failed for %s: %s", _key, dataToUse[_key]))
		}
	}
	db.CloseDB()
}

/*
TestBadger runs tests for badger engine.
*/
func TestBadger() {
	cfg := make(map[string]string)
	cfg["DBPath"] = fmt.Sprintf("%sbadger", dbPathPrefix)
	cfg["DetectConflicts"] = "false"
	cfg["NumGoroutines"] = "8"
	cfg["LogLevel"] = "WARNING"
	TestGolKeyVal("badger", cfg)
}

/*
TestBitcask runs tests for bitcask engine.
*/
func TestBitcask() {
	cfg := make(map[string]string)
	cfg["DBPath"] = fmt.Sprintf("%sbitcask", dbPathPrefix)
	TestGolKeyVal("bitcask", cfg)
}

/*
TestSqlite3 runs tests for sqlite3 engine.
*/
func TestSqlite3() {
	cfg := make(map[string]string)
	cfg["DBPath"] = fmt.Sprintf("%ssqlite3", dbPathPrefix)
	cfg["TableName"] = tableName
	TestGolKeyVal("sqlite3", cfg)
}

/*
TestLevelDB runs tests for leveldb engine.
*/
func TestLevelDB() {
	cfg := make(map[string]string)
	cfg["DBPath"] = fmt.Sprintf("%sleveldb", dbPathPrefix)
	TestGolKeyVal("leveldb", cfg)
}

/*
yeah main
*/
func main() {
	TestBadger()
	TestBitcask()
	TestSqlite3()
	TestLevelDB()
	fmt.Println("pass not panic")
}
