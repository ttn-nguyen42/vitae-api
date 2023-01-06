package logging

import (
	"Vitae/config"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/rs/zerolog"
)

func TestInit_ShouldSetLevelToInfoInRelease(t *testing.T) {
	os.Clearenv()
	os.Setenv(config.EnvGinMode, "release")

	setup()

	if DEFAULT != INFO {
		t.Errorf("Expected DEFAULT to be %v, got %v", INFO, DEFAULT)
	}
	if logger.GetLevel() != zerolog.InfoLevel {
		t.Errorf("Expected logger level to be %v, got %v", zerolog.InfoLevel, logger.GetLevel())
	}
}

func TestInit_ShouldBeDefinedLogLevelWhenNotInRelease(t *testing.T) {
	os.Clearenv()
	os.Setenv(config.EnvGinMode, "not release")
	os.Setenv(config.EnvLogLevel, "TRACE")

	setup()

	if DEFAULT != DEBUG {
		t.Errorf("Expected DEFAULT to be %v, got %v", DEBUG, DEFAULT)
	}
	if logger.GetLevel() != zerolog.TraceLevel {
		t.Errorf("Expected logger level to be %v, got %v", zerolog.TraceLevel, logger.GetLevel())
	}
}

func TestInit_ShouldSetLogLevelToDefaultWhenNotInReleaseAndNotDefineLogLevel(t *testing.T) {
	os.Clearenv()
	os.Setenv(config.EnvGinMode, "not release")
	os.Setenv(config.EnvLogLevel, "")

	setup()

	if DEFAULT != DEBUG {
		t.Errorf("Expected DEFAULT to be %v, got %v", DEBUG, DEFAULT)
	}
	if logger.GetLevel() != zerolog.DebugLevel {
		t.Errorf("Expected logger level to be %v, got %v", zerolog.DebugLevel, logger.GetLevel())
	}
}

func TestGetLogLevel(t *testing.T) {
	cases := []struct {
		Input string
		Expect zerolog.Level
	} {
		{DEBUG, zerolog.DebugLevel},
		{INFO, zerolog.InfoLevel},
		{WARNING, zerolog.WarnLevel},
		{ERROR, zerolog.ErrorLevel},
		{TRACE, zerolog.TraceLevel},
		{FATAL, zerolog.FatalLevel},
		{PANIC, zerolog.PanicLevel},
		{"NOT A LEVEL", zerolog.NoLevel},
	}
	for _, c := range cases {
		got := getLogLevel(c.Input)
		if got != c.Expect {
			t.Errorf("Expected %v will give %v, got %v", c.Input, c.Expect, got)
		}
	}
}

func TestLogWithMessageAndFields_ShouldLogWithOnlyMessage(t *testing.T) {
	os.Clearenv()
	os.Setenv(config.EnvLogLevel, DEBUG)
	os.Setenv(config.EnvGinMode, "NOT PRODUCTION")

	defaultStderr := os.Stderr
	read, write, _ := os.Pipe()
	os.Stderr = write
	setup()

	testMessage := "TEST MESSAGE"
	logWithMessageAndFields(logger.Debug(), testMessage)

	write.Close()
	out, _ := ioutil.ReadAll(read)
	got := string(out)
	os.Stderr = defaultStderr
	
	if got == "" {
		t.Errorf("No log was printed")
	}
	if !strings.Contains(got, testMessage) {
		t.Errorf("Log did not contain test message, got %v", got)
	}
	if !strings.Contains(got, DEBUG) {
		t.Errorf("Log did not contain correct level, expected %v, got %v", DEBUG, got)
	}
}
