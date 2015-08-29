package gollog

/*
grow with reference from:

http://www.goinggo.net/2013/11/using-log-package-in-go.html

http://golang.org/pkg/log/
https://golang.org/pkg/log/syslog/
https://github.com/op/go-logging
https://github.com/Sirupsen/logrus

*/

import (
	"fmt"
	"io"
	"os"
)

/*
Log to declare different streams to manage log

logInfo := gollog.Log{Level: "info", Thread: make(chan string)}
go logInfo.LogIt()

logInfo.Thread <- fmt.Sprintf("Message Recieved: %s", string(msg))
*/
type Log struct {
	Level  string
	Thread chan string
}

// start Log Action
func (l Log) Start() {
	for {
		msg := <-(l.Thread)
		fmt.Printf("[%s] %s", l.Level, msg)
	}
}

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
