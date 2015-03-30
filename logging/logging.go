package logging

// Logger interface wraps a single Log method that takes a level, a format
// string and arguments to fill it
type Logger interface {
	Log(level string, format string, args ...interface{})
}
