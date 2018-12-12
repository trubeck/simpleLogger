package simpleLogger

import (
	"fmt"
	"log"
	"os"
)

type SimpleLogger struct {
	debug   bool
	logfile string
}

// Create: returns a new SimpleLogger and sets debug flag accordingly
func Create(debug bool, logfile string) *SimpleLogger {
	return &SimpleLogger{debug: debug, logfile: logfile}
}

func (l *SimpleLogger) Initilize() {
	if l.logfile != "" {
		f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		fmt.Fprintln(f, "")
		fmt.Fprintln(f, "#########################################################")
		fmt.Fprintln(f, "")
		trace := log.New(f, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		trace.Output(2, "Start logging")
	}

	fmt.Println("")
	fmt.Println("#########################################################")
	fmt.Println("")
	info := log.New(os.Stdout, "\x1B[32mINFO: \x1B[0m", log.Ldate|log.Ltime|log.Lshortfile)
	info.Output(2, "Start logging")

}

func (l *SimpleLogger) Trace(v ...interface{}) {
	if l.debug {
		if l.logfile != "" {
			f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				log.Fatalf("error opening file: %v", err)
			}
			defer f.Close()
			trace := log.New(f, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
			trace.Output(2, fmt.Sprintln(v...))
		}
		trace := log.New(os.Stdout, "\x1B[36mTRACE: \x1B[0m", log.Ldate|log.Ltime|log.Lshortfile)
		trace.Output(2, fmt.Sprintln(v...))
	}
}

func (l *SimpleLogger) Info(v ...interface{}) {

	if l.logfile != "" {
		f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		info := log.New(f, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		info.Output(2, fmt.Sprintln(v...))
	}

	info := log.New(os.Stdout, "\x1B[32mINFO: \x1B[0m", log.Ldate|log.Ltime|log.Lshortfile)
	info.Output(2, fmt.Sprintln(v...))
}

func (l *SimpleLogger) Warning(v ...interface{}) {

	if l.logfile != "" {
		f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		warning := log.New(f, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
		warning.Output(2, fmt.Sprintln(v...))
	}

	warning := log.New(os.Stdout, "\x1B[35mWARN: \x1B[0m", log.Ldate|log.Ltime|log.Lshortfile)
	warning.Output(2, fmt.Sprintln(v...))
}

func (l *SimpleLogger) Error(v ...interface{}) {

	if l.logfile != "" {
		f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		_error := log.New(f, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		_error.Output(2, fmt.Sprintln(v...))
	}

	_error := log.New(os.Stderr, "\x1B[31mERROR: \x1B[0m", log.Ldate|log.Ltime|log.Lshortfile)
	_error.Output(2, fmt.Sprintln(v...))
}

func (l *SimpleLogger) Fatal(v ...interface{}) {
	if l.logfile != "" {
		f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		fatal := log.New(f, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		fatal.Output(2, fmt.Sprintln(v...))
		fatal.Output(2, "Force shutdown")
	}

	fatal := log.New(os.Stderr, "\x1B[31mERROR: \x1B[0m", log.Ldate|log.Ltime|log.Lshortfile)
	fatal.Output(2, fmt.Sprintln(v...))
	fatal.Output(2, "Force shutdown")

	os.Exit(1)
}

func (l *SimpleLogger) Panic(v ...interface{}) {
	if l.logfile != "" {
		f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		fatalPanic := log.New(f, "PANIC: ", log.Ldate|log.Ltime|log.Lshortfile)
		fatalPanic.Output(2, fmt.Sprintln(v...))
	}

	s := fmt.Sprint(v...)

	fatalPanic := log.New(os.Stderr, "PANIC: ", log.Ldate|log.Ltime|log.Lshortfile)
	fatalPanic.Output(2, s)
	panic(s)
}
