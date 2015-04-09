package logging

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type payload struct {
	Ts      time.Time `json:"ts"`
	Level   string    `json:"level"`
	Service string    `json:"service"`
	Logger  string    `json:"logger"`
	Message string    `json:"message"`
}

// JSONLogger defines a named logger that represents a service and outputs
// JSON messages to a writer
type JSONLogger struct {
	service string
	name    string
	writer  io.Writer
}

// Log emits a message a given level. It uses fmt.Sprintf with format and args
// to build the message
func (l *JSONLogger) Log(level string, format string, args ...interface{}) {
	payload := payload{
		time.Now(),
		level,
		l.service,
		l.name,
		fmt.Sprintf(format, args...),
	}
	enc, _ := json.Marshal(payload)
	fmt.Fprintf(l.writer, "%s\n", enc)
}

// Critical logs a message with "critical" level
func (l *JSONLogger) Critical(format string, args ...interface{}) {
	l.Log("critical", format, args...)
}

// Error logs a message with "error" level
func (l *JSONLogger) Error(format string, args ...interface{}) {
	l.Log("error", format, args...)
}

// Warning logs a message with "warning" level
func (l *JSONLogger) Warning(format string, args ...interface{}) {
	l.Log("warning", format, args...)
}

// Info logs a message with "info" level
func (l *JSONLogger) Info(format string, args ...interface{}) {
	l.Log("info", format, args...)
}

// Debug logs a message with "debug" level
func (l *JSONLogger) Debug(format string, args ...interface{}) {
	l.Log("debug", format, args...)
}

// Err logs an error object with "error" level
func (l *JSONLogger) Err(err error) {
	l.Error("%v", err)
}

// NewJSONLogger constructs a json logger which outputs to stdout
func NewJSONLogger(service string, name string) *JSONLogger {
	return NewJSONLoggerWithWriter(service, name, os.Stdout)
}

// NewJSONLoggerWithWriter constructs a json logger with a defined output stream
// to write log entries to
func NewJSONLoggerWithWriter(service string, name string, out io.Writer) *JSONLogger {
	return &JSONLogger{service, name, out}
}
