// +build !windows

package simpleLogger

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func (l *SimpleLogger) FatalGracefully(v ...interface{}) {
	if l.logfile != "" {
		f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		fatalGrace := log.New(f, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		fatalGrace.Output(2, fmt.Sprintln(v...))
		fatalGrace.Output(2, "Shutdown gracefully")
	}
	fatalGrace := log.New(os.Stderr, "\x1B[31mERROR: \x1B[0m", log.Ldate|log.Ltime|log.Lshortfile)
	fatalGrace.Output(2, fmt.Sprintln(v...))
	fatalGrace.Output(2, "Shutdown gracefully")

	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
}
