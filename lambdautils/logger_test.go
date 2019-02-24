package lambdautils_test

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/cpliakas/aws-sam-golang-example/lambdautils"
)

func TestLogger(t *testing.T) {
	logger := lambdautils.NewLogger(lambdautils.LogDebug)

	tests := []struct {
		log  *log.Logger
		fn   func(string, ...interface{})
		want string
	}{
		{logger.ErrorLogger, logger.Error, "ERROR\tfoo bar\n"},
		{logger.NoticeLogger, logger.Notice, "NOTICE\tfoo bar\n"},
		{logger.InfoLogger, logger.Info, "INFO\tfoo bar\n"},
		{logger.DebugLogger, logger.Debug, "DEBUG\tfoo bar\n"},
	}

	for _, tt := range tests {
		var buf bytes.Buffer
		tt.log.SetOutput(&buf)
		tt.fn("foo %s", "bar")
		have := buf.String()
		if have != tt.want {
			t.Errorf("have %q, want %q", have, tt.want)
		}
	}
}

func TestLoggerPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic")
		}
	}()

	logger := lambdautils.NewLogger(lambdautils.LogNone)
	logger.Panic("Ruh-roh!")
}

func TestLogLevel(t *testing.T) {
	defer os.Setenv(lambdautils.EnvLogLevel, os.Getenv(lambdautils.EnvLogLevel))

	tests := []struct {
		want int
	}{
		{lambdautils.LogNone},
		{lambdautils.LogError},
		{lambdautils.LogNotice},
		{lambdautils.LogInfo},
		{lambdautils.LogDebug},
	}

	for _, tt := range tests {
		os.Setenv(lambdautils.EnvLogLevel, strconv.Itoa(tt.want))
		have := lambdautils.LogLevel()
		if have != tt.want {
			t.Errorf("have %v, want %v", have, tt.want)
		}
	}
}

func TestLogLevelEmptyVar(t *testing.T) {
	defer os.Setenv(lambdautils.EnvLogLevel, os.Getenv(lambdautils.EnvLogLevel))

	os.Setenv(lambdautils.EnvLogLevel, "")
	want := lambdautils.LogInfo

	have := lambdautils.LogLevel()
	if have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestLogLevelPanic(t *testing.T) {
	defer os.Setenv(lambdautils.EnvLogLevel, os.Getenv(lambdautils.EnvLogLevel))

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic")
		}
	}()

	os.Setenv(lambdautils.EnvLogLevel, "not an integer")
	lambdautils.LogLevel()
}
