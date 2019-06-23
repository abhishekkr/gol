package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/abhishekkr/gol/gollog"

	gollog_example "github.com/abhishekkr/gol/_tests_behavioral_/gollog/helpers"
)

func main() {
	logfile, err := ioutil.TempFile("", "*-delete-me.log")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(logfile.Name())

	gollog.Debugf("gollog file would be %s", logfile.Name())
	gollog.LogOnce(logfile.Name(), "first message")

	logr := gollog.OpenLogFile(logfile.Name())
	defer gollog.CloseLogFile(logr)
	gollog.LogIt(logr, "sample logging")

	gollog_example.PassAndLog(gollog.LogIt, logr)
	gollog.Debug("finito gollog test")
	gollog.Infof("golog allows %s, %s, %s", "logrus", "gin-specific-logrus", "file-based-logging")
}
