package log

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var infoLogger = log.New(os.Stdout, "INFO:  ", log.LUTC|log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
var errorLogger = log.New(os.Stderr, "ERROR: ", log.LUTC|log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
var isDebug = false
var silence = false

// RecordedError write error to logs and return error object
func RecordedError(errorMessage string, args ...interface{}) error {
	message := fmt.Sprintf(errorMessage, args...)
	writeLog(errorLogger, "ERROR", message)
	return errors.New(message)
}

// Error write error to log
func Error(message string, args ...interface{}) {
	writeLog(errorLogger, "ERROR", message, args...)
}

// Warn write warn to log
func Warn(message string, args ...interface{}) {
	writeLog(infoLogger, "WARN", message, args...)
}

// Info write info to log
func Info(message string, args ...interface{}) {
	writeLog(infoLogger, "INFO", message, args...)
}

// Debug write debig to log
func Debug(infoMessage string, args ...interface{}) {
	if isDebug {
		writeLog(infoLogger, "DEBUG", infoMessage, args...)
	}
}

// EnableDebugMode Enable debug mode
func EnableDebugMode() {
	isDebug = true
	Debug(CAIC001D)
}

// EnableSilence will only print errors
func EnableSilence() {
	silence = true
}

func writeLog(logger *log.Logger, logLevel string, message string, args ...interface{}) {
	// -7 format ensures logs alignment, by padding spaces to log level to ensure 7 characters length.
	// 5 for longest log level, 1 for ':', and a space separator.
	if silence && logLevel != "ERROR" {
		return
	}

	logger.SetPrefix(fmt.Sprintf("%-7s", logLevel+":"))
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}
	logger.Output(3, message)
}
