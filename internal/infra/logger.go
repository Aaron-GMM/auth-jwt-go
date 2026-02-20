package infra

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p+": ", log.Ldate|log.Ltime|log.Lshortfile)
	return &Logger{
		debug:   log.New(os.Stdout, p+">[DEBUG]: ", logger.Flags()),
		info:    log.New(os.Stdout, p+">[INFO]: ", logger.Flags()),
		warning: log.New(os.Stdout, p+">[WARNING]: ", logger.Flags()),
		error:   log.New(os.Stderr, p+">[ERROR]: ", logger.Flags()),
		writer:  writer,
	}
}

// create no formatted logs
func (L *Logger) Debug(v ...interface{}) {
	L.debug.Println(v)
}
func (L *Logger) Info(v ...interface{}) {
	L.debug.Println(v)
}
func (L *Logger) Warning(v ...interface{}) {
	L.warning.Println(v)
}
func (L *Logger) Error(v ...interface{}) {
	L.error.Println(v)
}

// create formatted logs
func (L *Logger) Debugf(format string, v ...interface{}) {
	L.debug.Printf(format, v)
}
func (L *Logger) InforF(format string, v ...interface{}) {
	L.info.Printf(format, v)
}
func (L *Logger) WarningF(format string, v ...interface{}) {
	L.warning.Printf(format, v)
}
func (L *Logger) ErrorF(format string, v ...interface{}) {
	L.error.Printf(format, v)
}
