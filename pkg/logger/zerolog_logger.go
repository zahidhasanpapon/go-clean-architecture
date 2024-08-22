package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// zero logLogger represents the logger structure using zerolog.
type zerologLogger struct {
	zl zerolog.Logger
}

// Ensure zerologLogger implements the Logger interface.
var _ Logger = (*zerologLogger)(nil)

// newZerologLogger creates a new zerologLogger instance.
func newZerologLogger(level LogLevel) *zerologLogger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	l := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Caller().
		Logger()

	zLevel := getZerologLevel(level)
	l = l.Level(zLevel)

	return &zerologLogger{zl: l}
}

// getZerologLevel converts our LogLevel to zerolog.Level.
func getZerologLevel(level LogLevel) zerolog.Level {
	switch strings.ToLower(string(level)) {
	case string(DebugLevel):
		return zerolog.DebugLevel
	case string(InfoLevel):
		return zerolog.InfoLevel
	case string(WarnLevel):
		return zerolog.WarnLevel
	case string(ErrorLevel):
		return zerolog.ErrorLevel
	case string(FatalLevel):
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}

// Debug logs a debug message.
func (l *zerologLogger) Debug(message interface{}, args ...interface{}) {
	l.log(l.zl.Debug(), message, args...)
}

// Info logs an info message.
func (l *zerologLogger) Info(message interface{}, args ...interface{}) {
	l.log(l.zl.Info(), message, args...)
}

// Warn logs a warning message.
func (l *zerologLogger) Warn(message interface{}, args ...interface{}) {
	l.log(l.zl.Warn(), message, args...)
}

// Error logs an error message.
func (l *zerologLogger) Error(message interface{}, args ...interface{}) {
	l.log(l.zl.Error(), message, args...)
}

// Fatal logs a fatal message and exits the program.
func (l *zerologLogger) Fatal(message interface{}, args ...interface{}) {
	l.log(l.zl.Fatal(), message, args...)
}

// log handles the actual logging based on the message type.
func (l *zerologLogger) log(event *zerolog.Event, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		event.Err(msg).Msg("")
	case string:
		if len(args) > 0 {
			event.Msgf(msg, args...)
		} else {
			event.Msg(msg)
		}
	default:
		event.Interface("message", message).Msg("")
	}
}
