package lambdautils

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// EnvLogLevel is the environment variable that contains the minimum log level.
const EnvLogLevel = "LOG_LEVEL"

// Log* constants contain logging levels.
const (
	LogNone = iota
	LogError
	LogNotice
	LogInfo
	LogDebug
)

// LeveledLogger is a simple leveled logger that writes logs to STDOUT.
type LeveledLogger struct {
	ErrorLogger  *log.Logger
	NoticeLogger *log.Logger
	InfoLogger   *log.Logger
	DebugLogger  *log.Logger
}

// NewLogger returns a LeveledLogger that writs logs to os.Stdout or
// ioutil.Discard depending on the passed minimum log level.
func NewLogger(level int) *LeveledLogger {

	w := make([]io.Writer, 5)
	for i := range w {
		if i <= level {
			w[i] = os.Stdout
		} else {
			w[i] = ioutil.Discard
		}
	}

	// Return a Logger
	return &LeveledLogger{
		ErrorLogger:  log.New(w[LogError], "ERROR\t", 0),
		NoticeLogger: log.New(w[LogNotice], "NOTICE\t", 0),
		InfoLogger:   log.New(w[LogInfo], "INFO\t", 0),
		DebugLogger:  log.New(w[LogDebug], "DEBUG\t", 0),
	}
}

// Error writes an error level log.
func (l LeveledLogger) Error(format string, v ...interface{}) {
	l.ErrorLogger.Printf(format, v...)
}

// Notice writes an notice level log.
func (l LeveledLogger) Notice(format string, v ...interface{}) {
	l.NoticeLogger.Printf(format, v...)
}

// Info writes an info level log
func (l LeveledLogger) Info(format string, v ...interface{}) {
	l.InfoLogger.Printf(format, v...)
}

// Debug writes a debug level log.
func (l LeveledLogger) Debug(format string, v ...interface{}) {
	l.DebugLogger.Printf(format, v...)
}

// Panic writes an error level log and panics.
func (l *LeveledLogger) Panic(format string, v ...interface{}) {
	l.ErrorLogger.Panicf(format, v...)
}

// LogLevel reads the log level from the EnvLogLevel environment variable and
// defaults to the info level if it is not set.
func LogLevel() (i int) {
	var err error
	if s := os.Getenv(EnvLogLevel); s == "" {
		i = LogInfo
	} else {
		if i, err = strconv.Atoi(s); err != nil {
			panic(fmt.Errorf("invalid value for environemnt variable %q: %v", EnvLogLevel, s))
		}
	}
	return
}
