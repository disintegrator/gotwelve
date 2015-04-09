package logging

import "errors"

func ExampleNewJSONLogger() {
	logger := NewJSONLogger("myapp", "example")
	logger.Critical("Hello critical message")
	logger.Error("Hello error message")
	logger.Warning("Hello warning message")
	logger.Info("Hello info message")
	logger.Debug("Hello debug message")
	logger.Err(errors.New("oops"))
}
