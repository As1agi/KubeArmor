package log

import (
	"os"
	"testing"

	"go.uber.org/zap"
)

// test if Logger switches to Debug Mode when the environment varible is set to true
func TestLoggerDebugMode(t *testing.T) {
	os.Setenv("DEBUG", "true")

	initLogger()

	if !zapLogger.Desugar().Core().Enabled(zap.DebugLevel) {
		t.Errorf("Expected Logger to be in debug mode but it wasn't")
	}
	os.Unsetenv("DEBUG")
}

// test if logger is in INFO mode by default
func TestDefaultLevel(t *testing.T) {
	//
	os.Unsetenv("DEBUG")

	initLogger()

	if !zapLogger.Desugar().Core().Enabled(zap.InfoLevel) {
		t.Error("Default Logger should  be in INFO mode")
	}

	if zapLogger.Desugar().Core().Enabled(zap.DebugLevel) {
		t.Error("Default Logger should not be in DEBUG mode")
	}
}

// tests if that DEBUG = True and Debug = TRUE works (case insensitivity for the variable value)
func TestCaseInsensitity(t *testing.T) {
	inputs := []string{"true", "True", "tRue", "trUE"}

	for _, input := range inputs {
		os.Setenv("DEBUG", input)
		defer os.Unsetenv("DEBUG")

		initLogger()

		if !zapLogger.Desugar().Core().Enabled(zap.DebugLevel) {
			t.Errorf("expected DEBUG mode to be enable with input : %s", input)
		}

	}
}
