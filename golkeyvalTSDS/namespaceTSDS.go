package golkeyvalTSDS

import (
	"fmt"
	"time"

	golhashmap "github.com/abhishekkr/gol/golhashmap"
	golkeyvalNS "github.com/abhishekkr/gol/golkeyvalNS"
)

/*
 */
type Namespace struct {
	ns golkeyvalNS.NSDBEngine
}

/*
 */
func (namespace *Namespace) Configure(namespaceType golkeyvalNS.NSDBEngine) {
	namespace.ns = namespaceType
}

/*
register namespaceTSDS engine
*/
func init() {
	RegisterTSDSDBEngine("namespace", new(Namespace))
}

/* Return string of Namespace-d Time-Value for NS-Key */
func (namespace *Namespace) TimeKeyPart(key_time time.Time) string {
	return fmt.Sprintf("%v:%v:%v:%v:%v:%v", key_time.Year(), key_time.Month(),
		key_time.Day(), key_time.Hour(),
		key_time.Minute(), key_time.Second())
}

/* Return string of Namespace-d Key with Time-Namespace under it */
func (namespace *Namespace) KeyNameSpaceWithTime(key string, key_time time.Time) string {
	return fmt.Sprintf("%s:%s", key, namespace.TimeKeyPart(key_time))
}

/* Return string of Namespace-d Time-Namespace with Key under it */
func (namespace *Namespace) TimeNameSpaceWithKey(key string, key_time time.Time) string {
	return fmt.Sprintf("%s:%s", namespace.TimeKeyPart(key_time), key)
}

/* Return KeyNameSpaceWithTime and TimeNameSpaceWithKey as multi-return set */
func (namespace *Namespace) KeyAndTimeBothNameSpace(key string, key_time time.Time) (string, string) {
	time_ns := namespace.TimeKeyPart(key_time)
	return fmt.Sprintf("%s:%s", key, time_ns), fmt.Sprintf("%s:%s", time_ns, key)
}

/*
Returns Recursive-Namespaced data under given key, Proxy TSDS
The desried Time-frame shall be as namespace-d key
*/
func (namespace *Namespace) ReadTSDS(key string) golhashmap.HashMap {
	return namespace.ns.ReadNSRecursive(key)
}

/*
Returns Push status for TimeSeries data-store for a key-val for given timestamp
Default is Key Namespace-d with Time-Namespace under it
*/
func (namespace *Namespace) PushTSDS(key string, val string, key_time time.Time) bool {
	return namespace.PushTSDS_BaseKey(key, val, key_time)
}

/* Returns Push status for Key Namespace-d with Time-Namespace under it for given Timestamp */
func (namespace *Namespace) PushTSDS_BaseKey(key string, val string, key_time time.Time) bool {
	keytsds := namespace.KeyNameSpaceWithTime(key, key_time)
	return namespace.ns.PushNS(keytsds, val)
}

/* Returns Push status for Time-Namespace with Key Namespace-d under it for given Timestamp */
func (namespace *Namespace) PushTSDS_BaseTime(key string, val string, key_time time.Time) bool {
	timetsds := namespace.TimeNameSpaceWithKey(key, key_time)
	return namespace.ns.PushNS(timetsds, val)
}

/*
Returns Push status for TimeSeries data-store for a key-val for given timestamp
Both Key and Timestamp base key-vals are creaed
*/
func (namespace *Namespace) PushTSDS_BaseBoth(key string, val string, key_time time.Time) (status bool) {
	keytsds := namespace.KeyNameSpaceWithTime(key, key_time)
	status = namespace.ns.PushNS(keytsds, val)

	timetsds := namespace.TimeNameSpaceWithKey(key, key_time)
	status = status && namespace.ns.PushNS(timetsds, val)

	return status
}

/*
Returns Push status for TimeSeries data-store for a key-val for time of key-creation
Default is Key Namespace-d with Time-Namespace under it
*/
func (namespace *Namespace) PushNowTSDS(key string, val string) bool {
	return namespace.PushTSDS_BaseKey(key, val, time.Now())
}

/* Returns Push status for Key Namespace-d with Time-Namespace under it for time of key-creation */
func (namespace *Namespace) PushNowTSDS_BaseKey(key string, val string) bool {
	return namespace.PushTSDS_BaseKey(key, val, time.Now())
}

/* Returns Push status for Time-Namespace with Key Namespace-d under it for time of key-creation */
func (namespace *Namespace) PushNowTSDS_BaseTime(key string, val string) bool {
	return namespace.PushTSDS_BaseTime(key, val, time.Now())
}

/*
Returns Push status for TimeSeries data-store for a key-val for time of key-creation
Both Key and Timestamp base key-vals are creaed
*/
func (namespace *Namespace) PushNowTSDS_BaseBoth(key string, val string) (status bool) {
	return namespace.PushTSDS_BaseBoth(key, val, time.Now())
}

/*
Returns status to delete Recursive-Namespaced data under given key, Proxy TSDS
The desried Time-frame shall be as namespace-d key
*/
func (namespace *Namespace) DeleteTSDS(key string) bool {
	return namespace.ns.DeleteNSRecursive(key)
}

/** golkeyvalNS PROXY FUNC **/

func (namespace *Namespace) PushNS(key string, val string) bool {
	return namespace.ns.PushNS(key, val)
}
func (namespace *Namespace) ReadNSRecursive(key string) golhashmap.HashMap {
	return namespace.ns.ReadNSRecursive(key)
}
func (namespace *Namespace) DeleteNSRecursive(key string) bool {
	return namespace.ns.DeleteNSRecursive(key)
}

func (namespace *Namespace) PushKeyVal(key string, val string) bool {
	return namespace.ns.PushKeyVal(key, val)
}
func (namespace *Namespace) GetVal(key string) string {
	return namespace.ns.GetVal(key)
}
func (namespace *Namespace) DelKey(key string) bool {
	return namespace.ns.DelKey(key)
}
