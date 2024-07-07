package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Create Non-Formatted Logs
func (log *Logger) Debug(value ...interface{}) {
	log.debug.Println(value...)
}
func (log *Logger) Info(value ...interface{}) {
	log.info.Println(value...)
}
func (log *Logger) Warning(value ...interface{}) {
	log.warning.Println(value...)
}
func (log *Logger) Error(value ...interface{}) {
	log.err.Println(value...)
}

// Create Non-Format Enabled Logs
func (log *Logger) DebugF(format string, value ...interface{}) {
	log.debug.Printf(format, value...)
}
func (log *Logger) InfoF(format string, value ...interface{}) {
	log.info.Printf(format, value...)
}
func (log *Logger) WarningF(format string, value ...interface{}) {
	log.warning.Printf(format, value...)
}
func (log *Logger) ErrorF(format string, value ...interface{}) {
	log.err.Printf(format, value...)
}
