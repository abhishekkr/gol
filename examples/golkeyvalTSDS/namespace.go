package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	golkeyval "../../golkeyval"
	golkeyvalNS "../../golkeyvalNS"
	golkeyvalTSDS "../../golkeyvalTSDS"
	golassert "github.com/abhishekkr/gol/golassert"
	golfilesystem "github.com/abhishekkr/gol/golfilesystem"
	golhashmap "github.com/abhishekkr/gol/golhashmap"
)

var (
	dbpath = "/tmp/delete-this-levigoTSDS"
	db     golkeyval.DBEngine
	ns     golkeyvalNS.NSDBEngine
	tsds   golkeyvalTSDS.TSDSDBEngine
)

func setUpDB() golkeyvalTSDS.TSDSDBEngine {
	var cfg = make(map[string]string)
	cfg["DBPath"] = dbpath
	db = golkeyval.GetDBEngine("leveldb")
	db.Configure(cfg)
	db.CreateDB()

	ns = golkeyvalNS.GetNSDBEngine("delimited")
	ns.Configure(db)

	tsds = golkeyvalTSDS.GetTSDSDBEngine("namespace")
	tsds.Configure(ns)
	return tsds
}

func closeAndDeleteDB() {
	db.CloseAndDeleteDB()
}

func setupTestData() {
	ns.PushNS("upstate:2014:January:2:12:1:20", "down")
	ns.PushNS("2014:January:2:12:1:20:upstate", "down")

	ns.PushNS("upstate:2014:January:2:12:11:20", "up")
	ns.PushNS("2014:January:2:12:11:20:upstate", "up")
}

func testTimeKeyPart() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)

	expectedVal := "2014:January:2:12:10:1"
	resultVal := new(golkeyvalTSDS.Namespace).TimeKeyPart(anytime)
	golassert.AssertEqual(expectedVal, resultVal)
}

func testKeyNameSpaceWithTime() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)

	expectedVal := "upstate:2014:January:2:12:10:1"
	resultVal := new(golkeyvalTSDS.Namespace).KeyNameSpaceWithTime("upstate", anytime)
	golassert.AssertEqual(expectedVal, resultVal)
}

func testTimeNameSpaceWithKey() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)

	expectedVal := "2014:January:2:12:10:1:upstate"
	resultVal := new(golkeyvalTSDS.Namespace).TimeNameSpaceWithKey("upstate", anytime)
	golassert.AssertEqual(expectedVal, resultVal)
}

func testKeyAndTimeBothNameSpace() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)

	expectedVal1, expectedVal2 := "upstate:2014:January:2:12:10:1", "2014:January:2:12:10:1:upstate"
	resultVal1, resultVal2 := new(golkeyvalTSDS.Namespace).KeyAndTimeBothNameSpace("upstate", anytime)
	golassert.AssertEqual(expectedVal1, resultVal1)
	golassert.AssertEqual(expectedVal2, resultVal2)
}

func testReadTSDS() {
	setUpDB()
	setupTestData()

	expectedValArr := []string{"upstate:2014:January:2:12:1:20,down",
		"upstate:2014:January:2:12:11:20,up"}
	resultVal := golhashmap.HashMapToCSV(ns.ReadNSRecursive("upstate:2014:January"))
	resultValArr := strings.Split(resultVal, "\n")
	golassert.AssertEqualStringArray(expectedValArr, resultValArr)

	expectedVal := "upstate:2014:January:2:12:11:20,up"
	resultVal = golhashmap.HashMapToCSV(ns.ReadNSRecursive("upstate:2014:January:2:12:11"))
	golassert.AssertEqual(expectedVal, resultVal)

	expectedVal = ""
	resultVal = golhashmap.HashMapToCSV(ns.ReadNSRecursive("upstate:2014:February:2:12:11"))
	golassert.AssertEqual(expectedVal, resultVal)

	closeAndDeleteDB()
}

func testPushTSDS() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)
	setUpDB()
	setupTestData()

	if !tsds.PushTSDS("upstate", "up", anytime) {
		panic("PushTSDS creation failed for upstate:2014:January:2:12:10:1")
	}

	expectedVal := "upstate:2014:January:2:12:10:1,up"
	resultVal := golhashmap.HashMapToCSV(ns.ReadNSRecursive("upstate:2014:January:2:12:10:1"))
	golassert.AssertEqual(expectedVal, resultVal)

	closeAndDeleteDB()
}

/*
func testPushTSDSBaseKey() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)
	setUpDB()
	setupTestData()

	if !tsds.PushTSDS_BaseKey("upstate", "up", anytime) {
		panic("PushTSDS_BaseKey creation failed for upstate:2014:January:2:12:10:1")
	}

	expectedVal := "upstate:2014:January:2:12:10:1,up"
	resultVal := golhashmap.HashMapToCSV(ns.ReadNSRecursive("upstate:2014:January:2:12:10:1"))
	golassert.AssertEqual(expectedVal, resultVal)

	closeAndDeleteDB()
}

func testPushTSDSBaseTime() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)
	setUpDB()
	setupTestData()

	if !tsds.PushTSDS_BaseTime("upstate", "up", anytime) {
		panic("PushTSDS_BaseTime creation failed for upstate:2014:January:2:12:10:1")
	}
	expectedVal := "2014:January:2:12:10:1:upstate,up"
	resultVal := golhashmap.HashMapToCSV(ns.ReadNSRecursive("2014:January:2:12:10:1:upstate"))
	golassert.AssertEqual(expectedVal, resultVal)

	closeAndDeleteDB()
}

func testPushTSDSBaseBoth() {
	anytime := time.Date(2014, 1, 2, 12, 10, 1, 0, time.UTC)
	setUpDB()
	setupTestData()

	if !tsds.PushTSDS_BaseBoth("upstate", "up", anytime) {
		panic("PushTSDS_BaseBoth creation failed for upstate:2014:January:2:12:10:1")
	}

	expectedVal := "2014:January:2:12:10:1:upstate,up"
	resultVal := golhashmap.HashMapToCSV(ns.ReadNSRecursive("2014:January:2:12:10:1:upstate"))
	golassert.AssertEqual(expectedVal, resultVal)

	expectedVal = "upstate:2014:January:2:12:10:1,up"
	resultVal = golhashmap.HashMapToCSV(ns.ReadNSRecursive("upstate:2014:January:2:12:10:1"))
	golassert.AssertEqual(expectedVal, resultVal)

	closeAndDeleteDB()
}
*/
func testPushNowTSDS() {
	setUpDB()
	setupTestData()

	if !tsds.PushNowTSDS("testPushNowTSDS", "up") {
		panic("PushNowTSDS creation failed.")
	}

	resultVal := golhashmap.HashMapToCSV(ns.ReadNSRecursive("testPushNowTSDS"))
	if len(resultVal) == 1 {
		panic("Fail: Wrong count of Key creation")
	}

	closeAndDeleteDB()
}

/*
func testPushNowTSDSBaseKey() {
	setUpDB()
	setupTestData()

	if !tsds.PushNowTSDS_BaseKey("PushNowTSDS_BaseKey", "up") {
		panic("PushNowTSDS_BaseKey creation failed.")
	}

	resultVal := golhashmap.HashMapToCSV(ns.ReadNSRecursive("PushNowTSDS_BaseKey"))
	if len(resultVal) == 1 {
		panic("Fail: Wrong count of Key creation")
	}

	closeAndDeleteDB()
}

func testPushNowTSDSBaseTime() {
	setUpDB()
	setupTestData()

	if !tsds.PushNowTSDS_BaseTime("PushNowTSDS_BaseTime", "up") {
		panic("PushNowTSDS_BaseTime creation failed.")
	}

	resultVal := golhashmap.HashMapToCSV(ns.ReadNSRecursive("PushNowTSDS_BaseTime"))
	if len(resultVal) == 1 {
		panic("Fail: Wrong count of Key creation")
	}

	closeAndDeleteDB()
}

func testPushNowTSDSBaseBoth() {
	setUpDB()
	setupTestData()

	if !tsds.PushNowTSDS_BaseBoth("PushNowTSDS_BaseBoth", "up") {
		panic("PushNowTSDS_BaseBoth creation failed.")
	}

	resultVal := golhashmap.HashMapToCSV(ns.ReadNSRecursive("PushNowTSDS_BaseBoth"))
	if len(resultVal) == 1 {
		panic("Fail: Wrong count of Key creation")
	}

	closeAndDeleteDB()
}
*/
func testDeleteTSDS() {
	setUpDB()
	setupTestData()

	if !tsds.DeleteTSDS("upstate:2014:January:2:12") {
		panic("Fail: Deletion of upstate:2014:January:2:12 failed")
	}
	expectedVal := ""
	resultVal := golhashmap.HashMapToCSV(ns.ReadNSRecursive("upstate:2014:January:2:12"))
	golassert.AssertEqual(expectedVal, resultVal)

	closeAndDeleteDB()
}

func main() {
	fmt.Println("Your DB is referenced at", dbpath)
	if golfilesystem.PathExists(dbpath) {
		if os.RemoveAll(dbpath) != nil {
			panic("Fail: Temporary DB files are still present at: " + dbpath)
		}
	}
	runtime.GOMAXPROCS(runtime.NumCPU())

	testTimeKeyPart()
	testKeyNameSpaceWithTime()
	testTimeNameSpaceWithKey()
	testKeyAndTimeBothNameSpace()

	testReadTSDS()
	testPushTSDS()
	/*
		testPushTSDSBaseKey()
		testPushTSDSBaseTime()
		testPushTSDSBaseBoth()
	*/
	testPushNowTSDS()
	/*
		testPushNowTSDSBaseKey()
		testPushNowTSDSBaseTime()
		testPushNowTSDSBaseBoth()
	*/
	testDeleteTSDS()
	fmt.Println("nothing to panic...")
}
