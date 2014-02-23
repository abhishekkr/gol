package main

import (
  "flag"
  "github.com/abhishekkr/gol/gollog"

  "./helpers"
)

var (
  logfile     = flag.String("log-file", "gol.log", "to dump run-logs to")
)


func main(){
  flag.Parse()
  gollog.LogFile = *logfile
  gollog.Log_it("sample logging")

  gollog_example.PassAndLog(gollog.Log_it)
}
