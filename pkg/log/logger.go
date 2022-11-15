package log

import (
	"fmt"
	"io"
	"log"
)

type Console interface {
	WithPrefix(prefix string) Console
	Debug(msg string, args ...interface{})
	Log(msg string, args ...interface{})
}

type consoleLogger struct {
	debug  bool
	output *log.Logger
}

func NewConsole(writer io.Writer, debug bool) Console {
	return consoleLogger{
		debug:  debug,
		output: log.New(writer, "", 0),
	}
}

func (l consoleLogger) print(msg string, args ...interface{}) {
	l.output.Printf(msg, args...)
}

func (l consoleLogger) Log(msg string, args ...interface{}) {
	l.print(msg, args...)
}

func (l consoleLogger) Debug(msg string, args ...interface{}) {
	if !l.debug {
		return
	}
	l.print(msg, args...)
}

func (l consoleLogger) WithPrefix(prefix string) Console {
	currPrefix := l.output.Prefix()
	newPrefix := fmt.Sprintf("%s: %s", prefix, currPrefix)
	output := log.New(l.output.Writer(), newPrefix, l.output.Flags())
	l.output = output
	return l
}
