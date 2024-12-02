package main

import (
	"io"
	"log"
	"os"
	"time"
)

var (
	WarningLog *log.Logger
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	DebugLog   *log.Logger
)

func init() {
	currentDate := time.Now().Format("2006-01-02")
	logFileName := currentDate + ".log"

	logFile, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	multiwriter := io.MultiWriter(os.Stdout, logFile)

	InfoLog = log.New(multiwriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(multiwriter, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(multiwriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLog = log.New(multiwriter, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	// InfoLog.SetOutput(logFile)
	// WarningLog.SetOutput(logFile)
	// ErrorLog.SetOutput(logFile)
}

func main() {
	InfoLog.Println("Starting the application...")
	WarningLog.Println("WARNING!!!")
	ErrorLog.Println("Some error has occurred...")
	DebugLog.Println("Debug...")
}
