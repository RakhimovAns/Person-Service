package logging

import (
	"log"
	"os"
)

type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
}

type logger struct {
	debugLog *log.Logger
	infoLog  *log.Logger
	warnLog  *log.Logger
	errorLog *log.Logger
	fatalLog *log.Logger
}

func New(level string) Logger {
	return &logger{
		debugLog: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLog:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		fatalLog: log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *logger) Debug(format string, args ...interface{}) {
	l.debugLog.Printf(format, args...)
}

func (l *logger) Info(format string, args ...interface{}) {
	l.infoLog.Printf(format, args...)
}

func (l *logger) Warn(format string, args ...interface{}) {
	l.warnLog.Printf(format, args...)
}

func (l *logger) Error(format string, args ...interface{}) {
	l.errorLog.Printf(format, args...)
}

func (l *logger) Fatal(format string, args ...interface{}) {
	l.fatalLog.Fatalf(format, args...)
}
