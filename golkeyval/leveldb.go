package golkeyval

import (
	"fmt"
	"os"

	levigo "github.com/jmhodges/levigo"

	golerror "github.com/abhishekkr/gol/golerror"
)

/*
LevelDB struct with required leveldb details.
*/
type LevelDB struct {
	DBPath string
	GolDB  *levigo.DB
}

/*
init registers sqlite3 to DBEngines.
*/
func init() {
	RegisterDBEngine("leveldb", new(LevelDB))
}

/*
Configure populates LevelDB required configs.
*/
func (levelDB *LevelDB) Configure(cfg map[string]string) {
	levelDB.DBPath = cfg["DBPath"]
}

/*
CreateDB creates a leveldb db at provided DBPath.
*/
func (levelDB *LevelDB) CreateDB() {
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
func (levelDB *LevelDB) CloseDB() {
	levelDB.GolDB.Close()
}

/*
CloseAndDeleteDB closes and deletes a db given handle and DBPath.
Useful in use and throw implementations. And also tests.
*/
func (levelDB *LevelDB) CloseAndDeleteDB() {
	levelDB.CloseDB()
	if os.RemoveAll(levelDB.DBPath) != nil {
		panic("Fail: Temporary DB files are still present at: " + levelDB.DBPath)
	}
}

/*
PushKeyVal pushes key-val in provided DB handle.
*/
func (levelDB *LevelDB) PushKeyVal(key string, val string) bool {
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
func (levelDB *LevelDB) GetVal(key string) string {
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
func (levelDB *LevelDB) DelKey(key string) bool {
	writer := levigo.NewWriteOptions()
	defer writer.Close()

	err := levelDB.GolDB.Delete(writer, []byte(key))
	if err != nil {
		golerror.Boohoo("Key "+key+" query failed.", false)
		return false
	}
	return true
}
