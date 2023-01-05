package logging

import (
	"Vitae/config"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	DEBUG   = "DEBUG"
	INFO    = "INFO"
	WARNING = "WARNING"
	ERROR   = "ERROR"
	TRACE   = "TRACE"
	FATAL   = "FATAL"
	PANIC   = "PANIC"
)

// Default for development mode
var DEFAULT string
var logger zerolog.Logger

func init() {
	mode := os.Getenv(config.EnvGinMode)
	if mode != "release" {
		DEFAULT = DEBUG
		logger = getConsoleLogger()
	}
	if mode == "release" {
		DEFAULT = INFO
		logger = getServiceLogger()
	}
}

func getConsoleLogger() zerolog.Logger {
	writer := zerolog.ConsoleWriter{
		Out: os.Stderr,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
		FormatCaller: func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("- %s", i)
		},
	}
	return zerolog.New(writer).Level(getLogLevel()).With().Timestamp().Caller().Logger()
}

func getServiceLogger() zerolog.Logger {
	// UNIMPLEMENTED
	return log.Logger
}

func getLogLevel() zerolog.Level {
	level := os.Getenv(config.EnvLogLevel)
	if len(level) == 0 {
		level = DEFAULT
	}
	switch level {
	case DEBUG:
		return zerolog.DebugLevel
	case INFO:
		return zerolog.InfoLevel
	case WARNING:
		return zerolog.WarnLevel
	case ERROR:
		return zerolog.ErrorLevel
	case TRACE:
		return zerolog.TraceLevel
	case FATAL:
		return zerolog.FatalLevel
	case PANIC:
		return zerolog.PanicLevel
	default:
		return zerolog.NoLevel
	}
}

func withOption(event *zerolog.Event) *zerolog.Event {
	return event.Timestamp()
}

func logWithMessageAndFields(level *zerolog.Event, message string, fields ...map[string]interface{}) {
	if level == nil {
		return
	}
	if fields != nil {
		withOption(level).Fields(fields).Msg(message)
	} else {
		withOption(level).Msg(message)
	}
}

// Debug prints DEBUG level logs with timestamp
// The debug message is a string
// Can add additional fields as a map of string
func Debug(message string, additional ...map[string]interface{}) {
	logWithMessageAndFields(logger.Debug(), message, additional...)
}

// Info prints INFO level logs with timestamp
// The message is a string
// Can add additional fields as a map of string
func Info(message string, additional ...map[string]interface{}) {
	logWithMessageAndFields(logger.Info(), message, additional...)
}

// Warning prints WARN level logs with timestamp
// The message is a string
// Can add additional fields as a map of string
func Warning(message string, additional ...map[string]interface{}) {
	logWithMessageAndFields(logger.Warn(), message, additional...)
}

// Error prints ERROR level logs with timestamp
// The message is a string
// Can add additional fields as a map of string
func Error(message string, additional ...map[string]interface{}) {
	logWithMessageAndFields(logger.Error(), message, additional...)
}

// Trace prints TRACE level logs with timestamp
// The message is a string
// Can add additional fields as a map of string
func Trace(message string, additional ...map[string]interface{}) {
	logWithMessageAndFields(logger.Trace(), message, additional...)
}

// Fatal prints FATAL level logs with timestamp
// The message is a string
// Can add additional fields as a map of string
func Fatal(message string, additional ...map[string]interface{}) {
	logWithMessageAndFields(logger.Fatal(), message, additional...)
}

// Panic prints PANIC level logs with timestamp
// The message is a string
// Can add additional fields as a map of string
func Panic(message string, additional ...map[string]interface{}) {
	logWithMessageAndFields(logger.Panic(), message, additional...)
}
