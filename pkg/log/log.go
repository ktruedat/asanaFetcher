package log

// Logger represents the logger interface, used for logging messages in the application.
type Logger interface {
	Info(msg string, args ...any)
	Error(msg string, err error, args ...any)
	Warning(msg string, args ...any)
	Debug(msg string, args ...any)
}
