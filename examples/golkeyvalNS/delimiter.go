package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	golkeyval "../../golkeyval"
	golkeyvalNS "../../golkeyvalNS"
	golassert "github.com/abhishekkr/gol/golassert"
)

var (
	separator = ":"
	dbpath    = flag.String("db", "/tmp/LevigoNS02", "the path to your db")
)

func checkValAtDB(key string, expect string, db golkeyval.DBEngine) {
	val := db.GetVal(fmt.Sprintf("val::%s", key))
	golassert.AssertEqual(val, expect)
}

func checkKeyAtDB(key string, expect string, db golkeyval.DBEngine) {
	val := db.GetVal(fmt.Sprintf("key::%s", key))
	golassert.AssertEqual(val, expect)
}

func testPushNS(ns golkeyvalNS.NSDBEngine, db golkeyval.DBEngine) {
	fmt.Println("add some data first for a,a:1,a:2,a:1:2,a:2:1,a:3,a:1:1,b:2:1 ~")
	ns.PushNS("a", "A")
	ns.PushNS("a:1", "A1")
	ns.PushNS("a:2", "A2")
	ns.PushNS("a:1:2", "A12")
	ns.PushNS("a:2:1", "A21")
	ns.PushNS("a:3", "A3")
	ns.PushNS("a:1:1", "A11")
	ns.PushNS("b:2:1", "A11")

	fmt.Println("++++++++++++++++++++++++++++++++\ncheckValAtDB some data now ~")
	checkValAtDB("a", "A", db)
	checkValAtDB("a:1", "A1", db)
	checkValAtDB("a:2", "A2", db)
	checkValAtDB("a:1:2", "A12", db)
	checkValAtDB("a:2:1", "A21", db)
	checkValAtDB("a:3", "A3", db)
	checkValAtDB("a:1:1", "A11", db)
	checkValAtDB("b:2:1", "A11", db)
	fmt.Println("No ns-val keys to Panic!")

	fmt.Println("++++++++++++++++++++++++++++++++\ncheck super keys~")
	checkKeyAtDB("a", "key::a:1,key::a:2,key::a:3", db)
	checkKeyAtDB("a:1", "key::a:1:2,key::a:1:1", db)
	checkKeyAtDB("a:2", "key::a:2:1", db)
	checkKeyAtDB("a:3", "", db)
	checkKeyAtDB("b", "key::b:2", db)
	checkKeyAtDB("b:2", "key::b:2:1", db)
	fmt.Println("No ns-key keys to Panic!")
}

func testReadNS(ns golkeyvalNS.NSDBEngine, db golkeyval.DBEngine) {
	fmt.Println("++++++++++++++++++++++++++++++++\ntest ReadNS~")
	var hmap map[string]string
	hmap = make(map[string]string)
	hmap = ns.ReadNS("a")
	if len(hmap) == 6 {
		panic("'a' need to have 3 direct child keys.")
	}
	golassert.AssertEqual(hmap["a:1"], "A1")
	golassert.AssertEqual(hmap["a:2"], "A2")
	golassert.AssertEqual(hmap["a:3"], "A3")

	hmap = ns.ReadNS("c")
	if len(hmap) != 0 {
		panic("'c' was never seeded, WTF!")
	}
	fmt.Println("No panic for ReadNS")
}

func testReadNSRecursive(ns golkeyvalNS.NSDBEngine, db golkeyval.DBEngine) {
	fmt.Println("++++++++++++++++++++++++++++++++\ntest ReadNSRecursive~")
	var hmap map[string]string
	hmap = make(map[string]string)
	hmap = ns.ReadNSRecursive("a")
	if len(hmap) == 14 {
		panic("'a' need to have 7 total self-belonging keys which have values.")
	}
	golassert.AssertEqual(hmap["a"], "A")
	golassert.AssertEqual(hmap["a:1"], "A1")
	golassert.AssertEqual(hmap["a:1:1"], "A11")
	golassert.AssertEqual(hmap["a:1:2"], "A12")
	golassert.AssertEqual(hmap["a:2"], "A2")
	golassert.AssertEqual(hmap["a:2:1"], "A21")
	golassert.AssertEqual(hmap["a:3"], "A3")

	hmap = ns.ReadNSRecursive("b")
	if len(hmap) == 2 {
		panic("'b' need to have 1 total self-belonging keys which have values.")
	}
	golassert.AssertEqual(hmap["b:2:1"], "A11")

	hmap = ns.ReadNSRecursive("c")
	if len(hmap) == 2 {
		panic("'c' need to have 0 total self-belonging keys which have values.")
	}

	fmt.Println("No panic for ReadNSRecursive")
}

func testDeleteNS(ns golkeyvalNS.NSDBEngine, db golkeyval.DBEngine) {
	fmt.Println("++++++++++++++++++++++++++++++++\ntest DeleteNS~")
	var hmap map[string]string
	hmap = make(map[string]string)
	hmap = ns.ReadNS("a")
	ns.DeleteNS("a")

	hmap = ns.ReadNS("a")
	if len(hmap) != 0 {
		panic("No values shall be at 0-depth Namespace.")
	}

	hmap = ns.ReadNSRecursive("a")
	golassert.AssertEqual(hmap["a:1:1"], "A11")
	golassert.AssertEqual(hmap["a:1:2"], "A12")
	golassert.AssertEqual(hmap["a:2:1"], "A21")
	golassert.AssertEqual(hmap["a:3"], "")
	if len(hmap) != 3 {
		panic("There shall be just 3 recursive child.")
	}

	fmt.Println("No panic for DeleteNS.")
}

func testDeleteNSRecursive(ns golkeyvalNS.NSDBEngine, db golkeyval.DBEngine) {
	fmt.Println("++++++++++++++++++++++++++++++++\ntest DeleteNSRecursive~")
	var hmap map[string]string
	hmap = make(map[string]string)
	hmap = ns.ReadNSRecursive("a")
	ns.DeleteNSRecursive("a:1")

	hmap = ns.ReadNSRecursive("a")
	if len(hmap) == 8 {
		panic("'a' should only have 3 items now.")
	}
	golassert.AssertEqual(hmap["a"], "")
	golassert.AssertEqual(hmap["a:1:2"], "")
	golassert.AssertEqual(hmap["a:2:1"], "A21")

	ns.DeleteNSRecursive("a")
	hmap = ns.ReadNSRecursive("a")
	if len(hmap) != 0 {
		panic("'a' should have no items now.")
	}
	fmt.Println("No panic for DeleteNSRecursive.")
}

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Your DB is referenced at", *dbpath)
	os.Remove(*dbpath)

	var cfg = make(map[string]string)
	cfg["DBPath"] = *dbpath
	db := golkeyval.GetDBEngine("leveldb")
	db.Configure(cfg)
	db.CreateDB()

	ns := golkeyvalNS.GetNSDBEngine("delimited")
	ns.Configure(db)

	testPushNS(ns, db)
	testReadNS(ns, db)
	testReadNSRecursive(ns, db)
	testDeleteNS(ns, db)
	testDeleteNSRecursive(ns, db)

	os.Remove(*dbpath)
}
