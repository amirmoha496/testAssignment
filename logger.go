package main

import (
	"log"
	"os"
)

var logger *Logger

func getLogger() *Logger {
	var errorLog *log.Logger
	var warnLog *log.Logger
	var infoLog *log.Logger
	var debugLog *log.Logger
	if logger == nil {
		errorLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		warnLog = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
		infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		debugLog = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

		l := Logger{debugLog: debugLog, infoLog: infoLog, warnLog: warnLog, errorLog: errorLog}
		logger = &l
	}
	return logger
}

//Logger Given class encapsulates the Logging operations at different levels
type Logger struct {
	debugLog *log.Logger
	infoLog  *log.Logger
	warnLog  *log.Logger
	errorLog *log.Logger
}

//Debug It logs the given message at DEBUG level
func (l Logger) Debug(message string) {
	l.debugLog.Println(message)
}

//Info It logs the given message at INFO level
func (l Logger) Info(message string) {
	l.infoLog.Println(message)
}

//Warn It logs the given message at WARN level
func (l Logger) Warn(message string) {
	l.warnLog.Println(message)
}

//Error It logs the given message at ERROR level
func (l Logger) Error(message string) {
	l.errorLog.Println(message)
}
