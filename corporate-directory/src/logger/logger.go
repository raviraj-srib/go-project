package logger

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	DEFAULT_LOG_LEVEL = logrus.DebugLevel
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.SetReportCaller(true)
	var textFormatter = new(logrus.TextFormatter)
	textFormatter.TimestampFormat = "2006-01-02 15:04:05"
	textFormatter.ForceColors = true
	textFormatter.FullTimestamp = true
	textFormatter.CallerPrettyfier = func(f *runtime.Frame) (string, string) {
		if pc, file, line, ok := runtime.Caller(12); ok {
			file = file[strings.LastIndex(file, "/")+1:]
			funcName := runtime.FuncForPC(pc).Name()
			funcName = funcName[strings.LastIndex(funcName, ".")+1:]
			return fmt.Sprintf("[%s:%d]", file, line), "[" + funcName + "()]"
		}
		return "", ""
	}

	log.SetFormatter(textFormatter)

	log.SetLevel(DEFAULT_LOG_LEVEL)
}

func Print(fromat string, data ...interface{}) {
	log.Printf(fromat, data)
}

func Trace(fromat string, data ...interface{}) {
	log.Tracef(fromat, data)
}

func Debug(fromat string, data ...interface{}) {
	log.Debugf(fromat, data)
}

func Info(fromat string, data ...interface{}) {
	log.Infof(fromat, data)
}

func Warn(fromat string, data ...interface{}) {
	log.Warnf(fromat, data)
}

func Error(fromat string, data ...interface{}) {
	log.Errorf(fromat, data)
}

func Fatal(fromat string, data ...interface{}) {
	log.Fatalf(fromat, data)
}

func Panic(fromat string, data ...interface{}) {
	log.Panicf(fromat, data)
}
