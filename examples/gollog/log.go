package main

import (
	"flag"

	"github.com/abhishekkr/gol/gollog"

	gollog_example "./helpers"
)

var (
	logfile = flag.String("log-file", "gol.log", "to dump run-logs to")
)

func main() {
	flag.Parse()
	gollog.LogOnce(*logfile, "first message")

	logr := gollog.OpenLogFile(*logfile)
	defer gollog.CloseLogFile(logr)
	gollog.LogIt(logr, "sample logging")

	gollog_example.PassAndLog(gollog.LogIt, logr)
}
