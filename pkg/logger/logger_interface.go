package logger

// LogLevel represents the logging level.
type LogLevel string

// Available log levels.
const (
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
	FatalLevel LogLevel = "fatal"
)

// Logger defines the interface for logging operations.
type Logger interface {
	Debug(message interface{}, args ...interface{})
	Info(message interface{}, args ...interface{})
	Warn(message interface{}, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

// New creates a new Logger instance.
func New(level LogLevel) Logger {
	return newZerologLogger(level)
}
