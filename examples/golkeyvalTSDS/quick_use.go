package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/abhishekkr/gol/golassert"
	"github.com/abhishekkr/gol/golkeyval"
	"github.com/abhishekkr/gol/golkeyvalTSDS"
)

var (
	dbpath = flag.String("db", "/tmp/LevigoTSDS01", "the path to your db")
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Your DB is referenced at", *dbpath)

	var cfg = make(map[string]string)
	cfg["DBPath"] = *dbpath
	db := golkeyval.GetDBEngine("leveldb")
	db.Configure(cfg)
	db.CreateDB()

	base_key := "127.0.0.1:ping"
	base_state := "up"

	time_now := time.Now()
	golassert.AssertEqual(true, golkeyvalTSDS.PushTSDS(base_key, base_state, time_now, db))

	tsds_read01 := golkeyvalTSDS.ReadTSDS(base_key, db)
	tsds_key01 := golkeyvalTSDS.KeyNameSpaceWithTime(base_key, time_now)
	golassert.AssertEqual(base_state, tsds_read01[tsds_key01])

	golassert.AssertEqual(true, golkeyvalTSDS.DeleteTSDS(tsds_key01, db))

	tsds_read02 := golkeyvalTSDS.ReadTSDS(base_key, db)
	golassert.AssertEqual("", tsds_read02[tsds_key01])

	golassert.AssertEqual(true, golkeyvalTSDS.DeleteTSDS(tsds_key01, db))

	golassert.AssertEqual(true, golkeyvalTSDS.PushNowTSDS(base_key, base_state, db))

	tsds_read03 := golkeyvalTSDS.ReadTSDS(base_key, db)
	if len(tsds_read03) != 1 {
		panic("Push of base_key should have just created on time entry at this point.")
	}
	golassert.AssertEqualStringArray([]string{base_state}, tsds_read03.Values())

	golassert.AssertEqual(true, golkeyvalTSDS.DeleteTSDS(tsds_key01, db))

	fmt.Println("No need to Panic with Gol KeyVal TSDS feature.")
	os.Remove(*dbpath)
}
