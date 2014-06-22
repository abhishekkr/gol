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

// LogIt logs any string passed to it, into log file-handle
func LogIt(log *os.File, msg string) {
	n, err := io.WriteString(log, msg)
	if err != nil {
		fmt.Println(n, err)
	}
}
