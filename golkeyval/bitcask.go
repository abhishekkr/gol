package golkeyval

import (
	"fmt"
	"os"

	bitcask "git.mills.io/prologic/bitcask"

	golerror "github.com/abhishekkr/gol/golerror"
)

/*
Bitcask struct with required bitcask details.
*/
type Bitcask struct {
	DBPath string
	GolDB  *bitcask.Bitcask
}

/*
init registers bitcask (GO implemented Bitcask) to DBEngines.
*/
func init() {
	RegisterDBEngine("bitcask", new(Bitcask))
}

/*
Configure populates Bitcask required configs.
*/
func (b *Bitcask) Configure(cfg map[string]string) {
	b.DBPath = cfg["DBPath"]
}

/*
CreateDB creates a bitcask directory at provided DBPath.
*/
func (b *Bitcask) CreateDB() {
	var dbErr error

	b.GolDB, dbErr = bitcask.Open(b.DBPath)
	if dbErr != nil {
		errMsg := fmt.Sprintf("Bitcask '%s' Creation failed. %q", b.DBPath, dbErr)
		golerror.Boohoo(errMsg, true)
	}
}

/*
CloseDB syncs and closes a db given handle.
*/
func (b *Bitcask) CloseDB() {
	if err := b.GolDB.Sync(); err != nil {
		golerror.Boohoo("Failed closing DB after SYNC gives error.", false)
	}
	b.GolDB.Close()
}

/*
CloseAndDeleteDB closes and deletes a db given handle and DBPath.
Useful in use and throw implementations. And also tests.
*/
func (b *Bitcask) CloseAndDeleteDB() {
	b.GolDB.Close()
	if os.RemoveAll(b.DBPath) != nil {
		panic("Fail: Temporary DB files are still present at: " + b.DBPath)
	}
}

/*
PushKeyVal pushes key-val in provided DB handle.
*/
func (b *Bitcask) PushKeyVal(key string, val string) bool {
	if err := b.GolDB.Put([]byte(key), []byte(val)); err != nil {
		golerror.Boohoo("Key "+key+" insertion failed. It's value was "+val, false)
		return false
	}
	return true
}

/*
GetVal return string-ified value of key from provided db handle.
*/
func (b *Bitcask) GetVal(key string) string {
	data, err := b.GolDB.Get([]byte(key))
	if err != nil {
		golerror.Boohoo("Key "+key+" query failed.", false)
		return ""
	}
	return string(data)
}

/*
DelKey deletes key from provided DB handle.
*/
func (b *Bitcask) DelKey(key string) bool {
	if err := b.GolDB.Delete([]byte(key)); err != nil {
		golerror.Boohoo("Key "+key+" query failed.", false)
		return false
	}
	return true
}
