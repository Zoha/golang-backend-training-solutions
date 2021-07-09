package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Logger struct {
	logDestination io.Writer
}

func (l Logger) Log(a ...interface{}) {

	// use l.logDestination or io.Discard by default for log destination
	logDestination := l.logDestination
	if logDestination == nil {
		logDestination = io.Discard
	}

	fmt.Fprintln(logDestination, a...)
}

func (l *Logger) Begin() {
	l.logDestination = os.Stdout
	l.Log("BEGIN")
}

func (l *Logger) End() {
	l.Log("END")
	l.logDestination = ioutil.Discard
}

func main() {
	log := Logger{}

	log.Begin()

	msg := "Hello World!"
	log.Log("I said:", msg)

	log.End()
}
