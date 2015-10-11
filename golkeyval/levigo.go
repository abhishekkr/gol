package golkeyval

import (
	"fmt"
	"os"

	levigo "github.com/jmhodges/levigo"

	golerror "github.com/abhishekkr/gol/golerror"
)

/*
Levigo struct with required leveldb details.
*/
type Levigo struct {
	DBPath string
	GolDB  *levigo.DB
}

/*
init registers levigo (C API-d Levigo lib) to DBEngines.
*/
func init() {
	RegisterDBEngine("levigo", new(Levigo))
}

/*
Configure populates Levigo required configs.
*/
func (levelDB *Levigo) Configure(cfg map[string]string) {
	levelDB.DBPath = cfg["DBPath"]
}

/*
CreateDB creates a leveldb db at provided DBPath.
*/
func (levelDB *Levigo) CreateDB() {
	var errDB error
	opts := levigo.NewOptions()
	opts.SetCache(levigo.NewLRUCache(1 << 10))
	opts.SetCreateIfMissing(true)
	levelDB.GolDB, errDB = levigo.Open(levelDB.DBPath, opts)
	if errDB != nil {
		errMsg := fmt.Sprintf("DB %s Creation failed. %q", levelDB.DBPath, errDB)
		golerror.Boohoo(errMsg, true)
	}
}

/*
CloseDB closes a db given handle.
*/
func (levelDB *Levigo) CloseDB() {
	levelDB.GolDB.Close()
}

/*
CloseAndDeleteDB closes and deletes a db given handle and DBPath.
Useful in use and throw implementations. And also tests.
*/
func (levelDB *Levigo) CloseAndDeleteDB() {
	levelDB.CloseDB()
	if os.RemoveAll(levelDB.DBPath) != nil {
		panic("Fail: Temporary DB files are still present at: " + levelDB.DBPath)
	}
}

/*
PushKeyVal pushes key-val in provided DB handle.
*/
func (levelDB *Levigo) PushKeyVal(key string, val string) bool {
	writer := levigo.NewWriteOptions()
	defer writer.Close()

	keyname := []byte(key)
	value := []byte(val)
	err := levelDB.GolDB.Put(writer, keyname, value)
	if err != nil {
		golerror.Boohoo("Key "+key+" insertion failed. It's value was "+val, false)
		return false
	}
	return true
}

/*
GetVal return string-ified value of key from provided db handle.
*/
func (levelDB *Levigo) GetVal(key string) string {
	reader := levigo.NewReadOptions()
	defer reader.Close()

	data, err := levelDB.GolDB.Get(reader, []byte(key))
	if err != nil {
		golerror.Boohoo("Key "+key+" query failed.", false)
		return ""
	}
	return string(data)
}

/*
DelKey deletes key from provided DB handle.
*/
func (levelDB *Levigo) DelKey(key string) bool {
	writer := levigo.NewWriteOptions()
	defer writer.Close()

	err := levelDB.GolDB.Delete(writer, []byte(key))
	if err != nil {
		golerror.Boohoo("Key "+key+" query failed.", false)
		return false
	}
	return true
}
