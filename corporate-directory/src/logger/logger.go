package logger

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	DEFAULT_LOG_LEVEL = logrus.DebugLevel
	FUNC_DEPTH        = 10
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
		if pc, file, line, ok := runtime.Caller(FUNC_DEPTH); ok {
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

func Print(format string, data ...interface{}) {
	log.Printf(format, data...)
}

func Trace(format string, data ...interface{}) {
	log.Tracef(format, data...)
}

func Debug(format string, data ...interface{}) {
	log.Debugf(format, data...)
}

func Info(format string, data ...interface{}) {
	log.Infof(format, data...)
}

func Warn(format string, data ...interface{}) {
	log.Warnf(format, data...)
}

func Error(format string, data ...interface{}) {
	log.Errorf(format, data...)
}

func Fatal(format string, data ...interface{}) {
	log.Fatalf(format, data...)
}

func Panic(format string, data ...interface{}) {
	log.Panicf(format, data...)
}
