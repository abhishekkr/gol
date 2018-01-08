package gollog

import (
	"log"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/abhishekkr/gol/golenv"
)

type Logrus struct {
	LogLevel int
}

var (
	/*
		O:"panic"
		1:"fatal"
		2:"error"
		3:"warn"
		4:"info"
		5:"debug"
	*/
	LogLevel = golenv.OverrideIfEnv("GOLLOG_LOG_LEVEL", "5")
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(logrusLevel())
}

func logrusLevel() logrus.Level {
	loglevelIdx, err := strconv.Atoi(LogLevel)
	if err != nil {
		log.Fatalln("GOLLOG_LOG_LEVEL config seems to panic me!")
	}

	return logrus.AllLevels[loglevelIdx]
}

func Debug(msg string) {
	logrus.Debug(msg)
}

func Info(msg string) {
	logrus.Info(msg)
}

func Warn(msg string) {
	logrus.Warn(msg)
}

func Err(msg string) {
	logrus.Error(msg)
}

func Panic(msg string) {
	logrus.Panic(msg)
}
