package golkeyvalTSDS

import (
	"fmt"
	"time"

	golhashmap "github.com/abhishekkr/gol/golhashmap"
	golkeyval "github.com/abhishekkr/gol/golkeyval"
	golkeyvalNS "github.com/abhishekkr/gol/golkeyvalNS"
)

/* Return string of Namespace-d Time-Value for NS-Key */
func TimeKeyPart(key_time time.Time) string {
	return fmt.Sprintf("%v:%v:%v:%v:%v:%v", key_time.Year(), key_time.Month(),
		key_time.Day(), key_time.Hour(),
		key_time.Minute(), key_time.Second())
}

/* Return string of Namespace-d Key with Time-Namespace under it */
func KeyNameSpaceWithTime(key string, key_time time.Time) string {
	return fmt.Sprintf("%s:%s", key, TimeKeyPart(key_time))
}

/* Return string of Namespace-d Time-Namespace with Key under it */
func TimeNameSpaceWithKey(key string, key_time time.Time) string {
	return fmt.Sprintf("%s:%s", TimeKeyPart(key_time), key)
}

/* Return KeyNameSpaceWithTime and TimeNameSpaceWithKey as multi-return set */
func KeyAndTimeBothNameSpace(key string, key_time time.Time) (string, string) {
	time_ns := TimeKeyPart(key_time)
	return fmt.Sprintf("%s:%s", key, time_ns), fmt.Sprintf("%s:%s", time_ns, key)
}

/*
Returns Recursive-Namespaced data under given key, Proxy TSDS
The desried Time-frame shall be as namespace-d key
*/
func ReadTSDS(key string, db golkeyval.DBEngine) golhashmap.HashMap {
	return golkeyvalNS.ReadNSRecursive(key, db)
}

/*
Returns Push status for TimeSeries data-store for a key-val for given timestamp
Default is Key Namespace-d with Time-Namespace under it
*/
func PushTSDS(key string, val string, key_time time.Time, db golkeyval.DBEngine) bool {
	keytsds := KeyNameSpaceWithTime(key, key_time)
	return golkeyvalNS.PushNS(keytsds, val, db)
}

/* Returns Push status for Key Namespace-d with Time-Namespace under it for given Timestamp */
func PushTSDS_BaseKey(key string, val string, key_time time.Time, db golkeyval.DBEngine) bool {
	return PushTSDS(key, val, key_time, db)
}

/* Returns Push status for Time-Namespace with Key Namespace-d under it for given Timestamp */
func PushTSDS_BaseTime(key string, val string, key_time time.Time, db golkeyval.DBEngine) bool {
	timetsds := TimeNameSpaceWithKey(key, key_time)
	return golkeyvalNS.PushNS(timetsds, val, db)
}

/*
Returns Push status for TimeSeries data-store for a key-val for given timestamp
Both Key and Timestamp base key-vals are creaed
*/
func PushTSDS_BaseBoth(key string, val string, key_time time.Time, db golkeyval.DBEngine) (status bool) {
	keytsds := KeyNameSpaceWithTime(key, key_time)
	status = golkeyvalNS.PushNS(keytsds, val, db)

	timetsds := TimeNameSpaceWithKey(key, key_time)
	status = status && golkeyvalNS.PushNS(timetsds, val, db)

	return status
}

/*
Returns Push status for TimeSeries data-store for a key-val for time of key-creation
Default is Key Namespace-d with Time-Namespace under it
*/
func PushNowTSDS(key string, val string, db golkeyval.DBEngine) bool {
	keytsds := KeyNameSpaceWithTime(key, time.Now())
	return golkeyvalNS.PushNS(keytsds, val, db)
}

/* Returns Push status for Key Namespace-d with Time-Namespace under it for time of key-creation */
func PushNowTSDS_BaseKey(key string, val string, db golkeyval.DBEngine) bool {
	return PushNowTSDS(key, val, db)
}

/* Returns Push status for Time-Namespace with Key Namespace-d under it for time of key-creation */
func PushNowTSDS_BaseTime(key string, val string, db golkeyval.DBEngine) bool {
	timetsds := TimeNameSpaceWithKey(key, time.Now())
	return golkeyvalNS.PushNS(timetsds, val, db)
}

/*
Returns Push status for TimeSeries data-store for a key-val for time of key-creation
Both Key and Timestamp base key-vals are creaed
*/
func PushNowTSDS_BaseBoth(key string, val string, db golkeyval.DBEngine) (status bool) {
	key_time := time.Now()

	keytsds := KeyNameSpaceWithTime(key, key_time)
	status = golkeyvalNS.PushNS(keytsds, val, db)

	timetsds := TimeNameSpaceWithKey(key, key_time)
	status = status && golkeyvalNS.PushNS(timetsds, val, db)

	return status
}

/*
Returns status to delete Recursive-Namespaced data under given key, Proxy TSDS
The desried Time-frame shall be as namespace-d key
*/
func DeleteTSDS(key string, db golkeyval.DBEngine) bool {
	return golkeyvalNS.DeleteNSRecursive(key, db)
}
