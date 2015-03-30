package logging

import (
	"encoding/json"
	"testing"
)

type writeTarget struct {
	Timestamp string `json:"ts"`
	Service   string `json:"service"`
	Logger    string `json:"logger"`
	Level     string `json:"level"`
	Message   string `json:"message"`
}

func (w *writeTarget) Write(p []byte) (int, error) {
	err := json.Unmarshal(p, w)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

func TestJSONLogger(t *testing.T) {
	target := new(writeTarget)
	service := "myapp"
	name := "testlogger"
	message := "hello world"
	logger := NewJSONLoggerWithWriter(service, name, target)
	logger.Log("info", "hello %s", "world")
	if target.Service != service {
		t.Errorf("Expected service: %s. Got: %s.", service, target.Service)
	}
	if target.Logger != name {
		t.Errorf("Expected logger: %s. Got: %s.", name, target.Logger)
	}
	if target.Message != message {
		t.Errorf("Expected message: %s. Got: %s.", message, target.Message)
	}
}

func TestJSONLogCritical(t *testing.T) {
	target := new(writeTarget)
	logger := NewJSONLoggerWithWriter("myapp", "default", target)
	logger.Critical("test")
	if target.Level != "critical" {
		t.Errorf("Expected level: %s. Got: %s.", "critical", target.Level)
	}
}

func TestJSONLogError(t *testing.T) {
	target := new(writeTarget)
	logger := NewJSONLoggerWithWriter("myapp", "default", target)
	logger.Error("test")
	if target.Level != "error" {
		t.Errorf("Expected level: %s. Got: %s.", "error", target.Level)
	}
}

func TestJSONLogWarning(t *testing.T) {
	target := new(writeTarget)
	logger := NewJSONLoggerWithWriter("myapp", "default", target)
	logger.Warning("test")
	if target.Level != "warning" {
		t.Errorf("Expected level: %s. Got: %s.", "warning", target.Level)
	}
}

func TestJSONLogInfo(t *testing.T) {
	target := new(writeTarget)
	logger := NewJSONLoggerWithWriter("myapp", "default", target)
	logger.Info("test")
	if target.Level != "info" {
		t.Errorf("Expected level: %s. Got: %s.", "info", target.Level)
	}
}

func TestJSONLogDebug(t *testing.T) {
	target := new(writeTarget)
	logger := NewJSONLoggerWithWriter("myapp", "default", target)
	logger.Debug("test")
	if target.Level != "debug" {
		t.Errorf("Expected level: %s. Got: %s.", "debug", target.Level)
	}
}
