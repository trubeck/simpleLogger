package simpleLogger

import (
	"fmt"
	"log"
	"os"
)

type simpleLogger struct {
	debug bool
}

// Create: returns a new simpleLogger and sets debug flag accordingly
func Create(debug bool) *simpleLogger {
	return &simpleLogger{debug: debug}
}

func (l *simpleLogger) Trace(v ...interface{}) {
	if l.debug {
		trace := log.New(os.Stdout, "\x1B[36mTRACE: \x1B[0m", log.Ldate|log.Ltime|log.Lshortfile)
		trace.Output(2, fmt.Sprintln(v...))
	}
}

func (l *simpleLogger) Info(v ...interface{}) {
	info := log.New(os.Stdout, "\x1B[32mINFO: \x1B[0m", log.Ldate|log.Ltime|log.Lshortfile)
	info.Output(2, fmt.Sprintln(v...))
}

func (l *simpleLogger) Warning(v ...interface{}) {
	warning := log.New(os.Stdout, "\x1B[35mWARN: \x1B[0m", log.Ldate|log.Ltime|log.Lshortfile)
	warning.Output(2, fmt.Sprintln(v...))
}

func (l *simpleLogger) Error(v ...interface{}) {
	_error := log.New(os.Stderr, "\x1B[31mERROR: \x1B[0m", log.Ldate|log.Ltime|log.Lshortfile)
	_error.Output(2, fmt.Sprintln(v...))
}
