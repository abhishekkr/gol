package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	golassert "github.com/abhishekkr/gol/golassert"
	golfilesystem "github.com/abhishekkr/gol/golfilesystem"
	golhashmap "github.com/abhishekkr/gol/golhashmap"
	golkeyval "github.com/abhishekkr/gol/golkeyval"
	golkeyvalNS "github.com/abhishekkr/gol/golkeyvalNS"
	golkeyvalTSDS "github.com/abhishekkr/gol/golkeyvalTSDS"
)

var (
	dbpath = "/tmp/delete-this-levigoTSDS"
)

func setUpDB() golkeyval.DBEngine {
	var cfg = make(map[string]string)
	cfg["DBPath"] = dbpath
	db := golkeyval.GetDBEngine("leveldb")
	db.Configure(cfg)
	db.CreateDB()
	return db

}

func closeAndDeleteDB(db golkeyval.DBEngine) {
	db.CloseAndDeleteDB()
}

func setupTestData(db golkeyval.DBEngine) {
	golkeyvalNS.PushNS("upstate:2014:January:2:12:1:20", "down", db)
	golkeyvalNS.PushNS("2014:January:2:12:1:20:upstate", "down", db)

	golkeyvalNS.PushNS("upstate:2014:January:2:12:11:20", "up", db)
	golkeyvalNS.PushNS("2014:January:2:12:11:20:upstate", "up", db)
}

func TestTimeKeyPart() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)

	expected_val := "2014:January:2:12:10:1"
	result_val := golkeyvalTSDS.TimeKeyPart(anytime)
	golassert.AssertEqual(expected_val, result_val)
}

func TestKeyNameSpaceWithTime() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)

	expected_val := "upstate:2014:January:2:12:10:1"
	result_val := golkeyvalTSDS.KeyNameSpaceWithTime("upstate", anytime)
	golassert.AssertEqual(expected_val, result_val)
}

func TestTimeNameSpaceWithKey() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)

	expected_val := "2014:January:2:12:10:1:upstate"
	result_val := golkeyvalTSDS.TimeNameSpaceWithKey("upstate", anytime)
	golassert.AssertEqual(expected_val, result_val)
}

func TestKeyAndTimeBothNameSpace() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)

	expected_val1, expected_val2 := "upstate:2014:January:2:12:10:1", "2014:January:2:12:10:1:upstate"
	result_val1, result_val2 := golkeyvalTSDS.KeyAndTimeBothNameSpace("upstate", anytime)
	golassert.AssertEqual(expected_val1, result_val1)
	golassert.AssertEqual(expected_val2, result_val2)
}

func TestReadTSDS() {
	db := setUpDB()
	setupTestData(db)

	expected_val := "upstate:2014:January:2:12:1:20,down\nupstate:2014:January:2:12:11:20,up"
	result_val := golhashmap.HashMapToCSV(golkeyvalNS.ReadNSRecursive("upstate:2014:January", db))
	golassert.AssertEqual(expected_val, result_val)

	expected_val = "upstate:2014:January:2:12:11:20,up"
	result_val = golhashmap.HashMapToCSV(golkeyvalNS.ReadNSRecursive("upstate:2014:January:2:12:11", db))
	golassert.AssertEqual(expected_val, result_val)

	expected_val = ""
	result_val = golhashmap.HashMapToCSV(golkeyvalNS.ReadNSRecursive("upstate:2014:February:2:12:11", db))
	golassert.AssertEqual(expected_val, result_val)

	closeAndDeleteDB(db)
}

func TestPushTSDS() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)
	db := setUpDB()
	setupTestData(db)

	if !golkeyvalTSDS.PushTSDS("upstate", "up", anytime, db) {
		panic("PushTSDS creation failed for upstate:2014:January:2:12:10:1")
	}

	expected_val := "upstate:2014:January:2:12:10:1,up"
	result_val := golhashmap.HashMapToCSV(golkeyvalNS.ReadNSRecursive("upstate:2014:January:2:12:10:1", db))
	golassert.AssertEqual(expected_val, result_val)

	closeAndDeleteDB(db)
}

func TestPushTSDS_BaseKey() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)
	db := setUpDB()
	setupTestData(db)

	if !golkeyvalTSDS.PushTSDS_BaseKey("upstate", "up", anytime, db) {
		panic("PushTSDS_BaseKey creation failed for upstate:2014:January:2:12:10:1")
	}

	expected_val := "upstate:2014:January:2:12:10:1,up"
	result_val := golhashmap.HashMapToCSV(golkeyvalNS.ReadNSRecursive("upstate:2014:January:2:12:10:1", db))
	golassert.AssertEqual(expected_val, result_val)

	closeAndDeleteDB(db)
}

func TestPushTSDS_BaseTime() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)
	db := setUpDB()
	setupTestData(db)

	if !golkeyvalTSDS.PushTSDS_BaseTime("upstate", "up", anytime, db) {
		panic("PushTSDS_BaseTime creation failed for upstate:2014:January:2:12:10:1")
	}
	expected_val := "2014:January:2:12:10:1:upstate,up"
	result_val := golhashmap.HashMapToCSV(golkeyvalNS.ReadNSRecursive("2014:January:2:12:10:1:upstate", db))
	golassert.AssertEqual(expected_val, result_val)

	closeAndDeleteDB(db)
}

func TestPushTSDS_BaseBoth() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)
	db := setUpDB()
	setupTestData(db)

	if !golkeyvalTSDS.PushTSDS_BaseBoth("upstate", "up", anytime, db) {
		panic("PushTSDS_BaseBoth creation failed for upstate:2014:January:2:12:10:1")
	}

	expected_val := "2014:January:2:12:10:1:upstate,up"
	result_val := golhashmap.HashMapToCSV(golkeyvalNS.ReadNSRecursive("2014:January:2:12:10:1:upstate", db))
	golassert.AssertEqual(expected_val, result_val)

	expected_val = "upstate:2014:January:2:12:10:1,up"
	result_val = golhashmap.HashMapToCSV(golkeyvalNS.ReadNSRecursive("upstate:2014:January:2:12:10:1", db))
	golassert.AssertEqual(expected_val, result_val)

	closeAndDeleteDB(db)
}

func TestPushNowTSDS() {
	db := setUpDB()
	setupTestData(db)

	if !golkeyvalTSDS.PushNowTSDS("TestPushNowTSDS", "up", db) {
		panic("PushNowTSDS creation failed.")
	}

	result_val := golhashmap.HashMapToCSV(golkeyvalNS.ReadNSRecursive("TestPushNowTSDS", db))
	if len(result_val) == 1 {
		panic("Fail: Wrong count of Key creation")
	}

	closeAndDeleteDB(db)
}

func TestPushNowTSDS_BaseKey() {
	db := setUpDB()
	setupTestData(db)

	if !golkeyvalTSDS.PushNowTSDS_BaseKey("PushNowTSDS_BaseKey", "up", db) {
		panic("PushNowTSDS_BaseKey creation failed.")
	}

	result_val := golhashmap.HashMapToCSV(golkeyvalNS.ReadNSRecursive("PushNowTSDS_BaseKey", db))
	if len(result_val) == 1 {
		panic("Fail: Wrong count of Key creation")
	}

	closeAndDeleteDB(db)
}

func TestPushNowTSDS_BaseTime() {
	db := setUpDB()
	setupTestData(db)

	if !golkeyvalTSDS.PushNowTSDS_BaseTime("PushNowTSDS_BaseTime", "up", db) {
		panic("PushNowTSDS_BaseTime creation failed.")
	}

	result_val := golhashmap.HashMapToCSV(golkeyvalNS.ReadNSRecursive("PushNowTSDS_BaseTime", db))
	if len(result_val) == 1 {
		panic("Fail: Wrong count of Key creation")
	}

	closeAndDeleteDB(db)
}

func TestPushNowTSDS_BaseBoth() {
	db := setUpDB()
	setupTestData(db)

	if !golkeyvalTSDS.PushNowTSDS_BaseBoth("PushNowTSDS_BaseBoth", "up", db) {
		panic("PushNowTSDS_BaseBoth creation failed.")
	}

	result_val := golhashmap.HashMapToCSV(golkeyvalNS.ReadNSRecursive("PushNowTSDS_BaseBoth", db))
	if len(result_val) == 1 {
		panic("Fail: Wrong count of Key creation")
	}

	closeAndDeleteDB(db)
}

func TestDeleteTSDS() {
	db := setUpDB()
	setupTestData(db)

	if !golkeyvalTSDS.DeleteTSDS("upstate:2014:January:2:12", db) {
		panic("Fail: Deletion of upstate:2014:January:2:12 failed")
	}
	expected_val := ""
	result_val := golhashmap.HashMapToCSV(golkeyvalNS.ReadNSRecursive("upstate:2014:January:2:12", db))
	golassert.AssertEqual(expected_val, result_val)

	closeAndDeleteDB(db)
}

func main() {
	fmt.Println("Your DB is referenced at", dbpath)
	if golfilesystem.PathExists(dbpath) {
		if os.RemoveAll(dbpath) != nil {
			panic("Fail: Temporary DB files are still present at: " + dbpath)
		}
	}
	runtime.GOMAXPROCS(runtime.NumCPU())

	TestTimeKeyPart()
	TestKeyNameSpaceWithTime()
	TestTimeNameSpaceWithKey()
	TestKeyAndTimeBothNameSpace()
	TestReadTSDS()
	TestPushTSDS()
	TestPushTSDS_BaseKey()
	TestPushTSDS_BaseTime()
	TestPushTSDS_BaseBoth()
	TestPushNowTSDS()
	TestPushNowTSDS_BaseKey()
	TestPushNowTSDS_BaseTime()
	TestPushNowTSDS_BaseBoth()
	TestDeleteTSDS()
}
