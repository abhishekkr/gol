package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	golhashmap "github.com/abhishekkr/gol/golhashmap"
	golkeyval "github.com/abhishekkr/gol/golkeyval"
	golkeyvalNS "github.com/abhishekkr/gol/golkeyvalNS"
)

var (
	dbpath = flag.String("db", "/tmp/LevigoNS00", "the path to your db")
)

func everySecondOfHour(hour int, check string, db golkeyval.DBEngine) {
	for sec := 0; sec < 3600; sec++ {
		nukey := fmt.Sprintf("127.0.0.1:%s:2013:10:26:%d:%d",
			check, hour, sec)
		if sec%500 != 0 {
			continue
		}
		val := "up"
		if sec%1000 == 0 {
			val = "down"
		}
		golkeyvalNS.PushNS(nukey, val, db)
	}
	fmt.Printf("Hour %d done. Enter 'yes' anytime to end Push action.\n", hour)
}

func witeMap(db golkeyval.DBEngine) {
	for hour := 0; hour < 24; hour++ {
		go everySecondOfHour(hour, "status", db)
	}
	for {
		var quit string
		fmt.Scanf("%s", &quit)
		if quit == "y" || quit == "yes" {
			return
		}
		time.Sleep(10 * time.Second)
	}
}

func readMap(key string, db golkeyval.DBEngine) {
	var hmap map[string]string
	hmap = make(map[string]string)
	hmap = golkeyvalNS.ReadNS(key, db)
	fmt.Println("Total Child Keys found:", len(hmap))
	for k, v := range hmap {
		fmt.Printf("%s => %s\n", k, v)
	}
}

func printMapRecursive(m golhashmap.HashMap) {
	for k, v := range m {
		fmt.Println("val for key:", k, v)
	}
}

func main() {
	flag.Parse()
	startTime := time.Now()
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Your DB is referenced at", *dbpath)
	createStartTime := time.Now()

	var cfg = make(map[string]string)
	cfg["DBPath"] = *dbpath
	db := golkeyval.GetDBEngine("leveldb")
	db.Configure(cfg)
	db.CreateDB()

	witeMap(db)
	fmt.Println("Writing is over.")
	readMap("127.0.0.1:status:2013:10:26:12", db)
	result := golkeyvalNS.ReadNSRecursive("127.0.0.1:status", db)
	readStartTime := time.Now()
	printMapRecursive(result)
	readMap("127.0.0.1:status:2013:10:26", db)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("\n\nStatistics:\n\tStarted at: %q\n", startTime)
	fmt.Printf("\tCreating DB: %q\n", createStartTime)
	fmt.Printf("\tReading DB: %q\n\tRead For an Hour: %q\n", readStartTime, time.Now())
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println(len(result))
	os.Remove(*dbpath)
}
