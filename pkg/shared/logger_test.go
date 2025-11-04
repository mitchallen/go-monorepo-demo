// Author: Mitch Allen
// File: logger_test.go

package shared

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger("test")
	if logger == nil {
		t.Error("NewLogger() returned nil")
	}
	if logger.prefix != "test" {
		t.Errorf("Expected prefix 'test', got '%s'", logger.prefix)
	}
}

func TestLoggerInfo(t *testing.T) {
	logger := NewLogger("test")
	// This test just ensures Info() doesn't panic
	logger.Info("test message")
}

func TestLoggerError(t *testing.T) {
	logger := NewLogger("test")
	// This test just ensures Error() doesn't panic
	logger.Error("test error")
}
