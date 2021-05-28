package logger

import (
	log "github.com/sirupsen/logrus"
)

const (
	DEFAULT_LOG_LEVEL = log.InfoLevel
)

func (u UTCFormatter) Format(e *log.Entry) ([]byte, error) {
	e.Time = e.Time.UTC()
	return u.Formatter.Format(e)
}

func init() {
	//log.SetFormatter(UTCFormatter{&log.JSONFormatter{}})
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(DEFAULT_LOG_LEVEL)
}

func Trace(data string) {
	log.Trace(data)
}

func Debug(data string) {
	log.Debug(data)
}

func Info(data interface{}) {
	log.Info(data)
}

func Warn(data string) {
	log.Warn(data)
}

func Error(data string) {
	log.Error(data)
}

func Fatal(data string) {
	log.Fatal(data)
}

func Panic(data string) {
	log.Panic(data)
}
