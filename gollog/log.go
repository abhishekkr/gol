package gollog

import (
	"fmt"
	"io"
	"os"
)

// Logfile sends back handle of opened logfile, remember to defer F.Close() at usage.
func OpenLogFile(logFile string) *os.File {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
	}
	// at callee: defer f.Close()
	return f
}

// close LogFile
func CloseLogFile(log *os.File) {
	log.Close()
}

// just write to given file handle
func LogIt(fyl *os.File, lyn string) {
	lyn = fmt.Sprintf("%s\n", lyn)
	n, err := io.WriteString(fyl, lyn)
	if err != nil {
		fmt.Println(n, err)
	}
}

// Open, Log, Close
func LogOnce(logfile string, msg string) {
	logr := OpenLogFile(logfile)
	defer logr.Close()
	LogIt(logr, msg)
}
