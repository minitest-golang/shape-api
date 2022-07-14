package utils

import (
	"fmt"
	"log"
	"os"
)

var (
	debugLog   *log.Logger
	infoLog    *log.Logger
	warningLog *log.Logger
	errorLog   *log.Logger
	fatalLog   *log.Logger
)

func LoggerInit() (f *os.File) {
	fmt.Printf("AppMode: %s\n", AppMode)
	if AppMode == "debug" {
		return nil
	}
	logFile := os.Getenv("LOG_FILE")
	if logFile == "" {
		return
	}
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to init logger: %s", err.Error())
		return nil
	}

	debugLog = log.New(f, "[DEBUG] ", log.Ldate|log.Ltime)
	infoLog = log.New(f, "[INFO] ", log.Ldate|log.Ltime)
	warningLog = log.New(f, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog = log.New(f, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	fatalLog = log.New(f, "[FATAL] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog.Println("Started logger.")
	return
}

func DebugLog(format string, v ...interface{}) {
	if AppMode != "debug" {
		return
	}
	if debugLog != nil {
		debugLog.Printf(format, v...)
	} else {
		log.Printf(format, v...)
	}
}

func InfoLog(format string, v ...interface{}) {
	if infoLog != nil {
		infoLog.Printf(format, v...)
	} else {
		log.Printf(format, v...)
	}
}
func WarningLog(format string, v ...interface{}) {
	if warningLog != nil {
		warningLog.Printf(format, v...)
	} else {
		log.Printf(format, v...)
	}
}
func ErrorLog(format string, v ...interface{}) {
	if errorLog != nil {
		errorLog.Printf(format, v...)
	} else {
		log.Printf(format, v...)
	}
}
func FatalLog(format string, v ...interface{}) {
	if fatalLog != nil {
		fatalLog.Printf(format, v...)
		os.Exit(1)
	} else {
		log.Fatalf(format, v...)
	}
}
