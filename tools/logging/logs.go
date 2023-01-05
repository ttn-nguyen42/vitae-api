package logging

import (
	"Vitae/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

const (
	DEBUG   string = "DEBUG"
	INFO           = "INFO"
	WARNING        = "WARNING"
	ERROR          = "ERROR"
	TRACE          = "TRACE"
	FATAL          = "FATAL"
	PANIC          = "PANIC"
)

var DEFAULT = "DEBUG"

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	mode := os.Getenv(config.EnvGinMode)
	if mode == "release" {
		DEFAULT = INFO
	} else {
		DEFAULT = DEBUG
	}
	zerolog.SetGlobalLevel(getLogLevel())
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
	logWithMessageAndFields(log.Debug(), message, additional...)
}

// Info prints INFO level logs with timestamp
// The message is a string
// Can add additional fields as a map of string
func Info(message string, additional ...map[string]interface{}) {
	logWithMessageAndFields(log.Info(), message, additional...)
}

// Warning prints WARN level logs with timestamp
// The message is a string
// Can add additional fields as a map of string
func Warning(message string, additional ...map[string]interface{}) {
	logWithMessageAndFields(log.Warn(), message, additional...)
}

// Error prints ERROR level logs with timestamp
// The message is a string
// Can add additional fields as a map of string
func Error(message string, additional ...map[string]interface{}) {
	logWithMessageAndFields(log.Error(), message, additional...)
}

// Trace prints TRACE level logs with timestamp
// The message is a string
// Can add additional fields as a map of string
func Trace(message string, additional ...map[string]interface{}) {
	logWithMessageAndFields(log.Trace(), message, additional...)
}

// Fatal prints FATAL level logs with timestamp
// The message is a string
// Can add additional fields as a map of string
func Fatal(message string, additional ...map[string]interface{}) {
	logWithMessageAndFields(log.Fatal(), message, additional...)
}

// Panic prints PANIC level logs with timestamp
// The message is a string
// Can add additional fields as a map of string
func Panic(message string, additional ...map[string]interface{}) {
	logWithMessageAndFields(log.Panic(), message, additional...)
}
