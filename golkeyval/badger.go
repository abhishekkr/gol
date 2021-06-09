package golkeyval

import (
	"fmt"
	"log"
	"os"
	"strconv"

	badger "github.com/dgraph-io/badger/v3"

	golerror "github.com/abhishekkr/gol/golerror"
)

/*
Badger struct with required badger details.
*/
type Badger struct {
	DBPath          string
	DetectConflicts bool
	NumGoroutines   int
	LogLevel        badger.Logger
	GolDB           *badger.DB
}

/*
init registers badger (GO implemented Badger) to DBEngines.
*/
func init() {
	RegisterDBEngine("badger", new(Badger))
}

/*
Configure populates Badger required configs.
*/
func (b *Badger) Configure(cfg map[string]string) {
	b.LogLevel = defaultLogger("INFO")
	for key, val := range cfg {
		switch key {
		case "DBPath":
			b.DBPath = val
		case "DetectConflicts":
			b.DetectConflicts, _ = strconv.ParseBool(val) // false by default
		case "NumGoroutines":
			numGoroutines, _ := strconv.ParseInt(val, 10, 64)
			b.NumGoroutines = int(numGoroutines)
		case "LogLevel":
			b.LogLevel = defaultLogger(val)
		}
	}
	if b.NumGoroutines < 1 {
		b.NumGoroutines = 8
	}
}

/*
CreateDB creates a Badger db at provided DBPath.
*/
func (b *Badger) CreateDB() {
	var dbErr error

	opts := badger.DefaultOptions(b.DBPath)
	opts.DetectConflicts = b.DetectConflicts
	opts.NumGoroutines = b.NumGoroutines
	opts.Logger = b.LogLevel

	b.GolDB, dbErr = badger.Open(opts)
	if dbErr != nil {
		errMsg := fmt.Sprintf("DB %s Creation failed. %q", b.DBPath, dbErr)
		golerror.Boohoo(errMsg, true)
	}
}

/*
CloseDB closes a db given handle.
*/
func (b *Badger) CloseDB() {
	b.GolDB.Close()
}

/*
CloseAndDeleteDB closes and deletes a db given handle and DBPath.
Useful in use and throw implementations. And also tests.
*/
func (b *Badger) CloseAndDeleteDB() {
	b.CloseDB()
	if os.RemoveAll(b.DBPath) != nil {
		panic("Fail: Temporary DB files are still present at: " + b.DBPath)
	}
}

/*
PushKeyVal pushes key-val in provided DB handle.
*/
func (b *Badger) PushKeyVal(key string, val string) bool {
	err := b.GolDB.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(val))
	})

	if err != nil {
		golerror.Boohoo("Key "+key+" insertion failed. It's value was "+val, false)
		return false
	}
	return true
}

/*
GetVal return string-ified value of key from provided db handle.
*/
func (b *Badger) GetVal(key string) string {
	var data []byte
	err := b.GolDB.View(func(txn *badger.Txn) error {
		var errGet error
		item, errGet := txn.Get([]byte(key))
		if errGet != nil {
			return errGet
		}

		data, errGet = item.ValueCopy(data)
		if errGet != nil {
			return errGet
		}
		return nil
	})

	if err != nil {
		golerror.Boohoo("Key "+key+" query failed.", false)
		return ""
	}
	return string(data)
}

/*
DelKey deletes key from provided DB handle.
*/
func (b *Badger) DelKey(key string) bool {
	err := b.GolDB.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(key))
	})

	if err != nil {
		golerror.Boohoo("Key "+key+" query failed.", false)
		return false
	}
	return true
}

/* Logger implementation configurable by log-level */
type loggingLevel int

const (
	DEBUG loggingLevel = iota
	INFO
	WARNING
	ERROR
)

type defaultLog struct {
	*log.Logger
	level loggingLevel
}

func defaultLogger(logLevel string) *defaultLog {
	level := map[string]loggingLevel{
		"DEBUG":   DEBUG,
		"INFO":    INFO,
		"WARNING": WARNING,
		"ERROR":   ERROR,
	}[logLevel]
	return &defaultLog{Logger: log.New(os.Stderr, "badger ", log.LstdFlags), level: level}
}

func (l *defaultLog) Errorf(f string, v ...interface{}) {
	if l.level <= ERROR {
		l.Printf("ERROR: "+f, v...)
	}
}

func (l *defaultLog) Warningf(f string, v ...interface{}) {
	if l.level <= WARNING {
		l.Printf("WARNING: "+f, v...)
	}
}

func (l *defaultLog) Infof(f string, v ...interface{}) {
	if l.level <= INFO {
		l.Printf("INFO: "+f, v...)
	}
}

func (l *defaultLog) Debugf(f string, v ...interface{}) {
	if l.level <= DEBUG {
		l.Printf("DEBUG: "+f, v...)
	}
}
