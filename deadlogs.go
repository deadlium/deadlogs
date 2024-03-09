package deadlogs

import (
	"github.com/fatih/color"
	"log"
	"os"
	"runtime"
)

// Enhanced logging levels
const (
	InfoLevel    = "INFO:"
	SuccessLevel = "SUCCESS:"
	ErrorLevel   = "ERROR:"
	WarnLevel    = "WARN:"
	DebugLevel   = "DEBUG:"
)

func Info(message string) {
	log.Printf("%s %s\n", color.BlueString(InfoLevel), message)
}

func Success(message string) {
	log.Printf("%s %s\n", color.GreenString(SuccessLevel), message)
}

func Error(message string) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s %s \n(file: \x1b]8;;file://%s\x1b\\%s:%d\x1b]8;;\x1b\\, line: %d)\n", color.RedString(ErrorLevel), message, file, file, line, line)
	os.Exit(1)
}

func Warn(message string) {
	log.Printf("%s %s\n", color.YellowString(WarnLevel), message)
}

func Debug(message string) {
	log.Printf("%s %s\n", color.WhiteString(DebugLevel), message)
}
