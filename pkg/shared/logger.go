// Author: Mitch Allen
// File: logger.go

package shared

import (
	"fmt"
	"time"
)

// Logger provides simple logging functionality shared across packages
type Logger struct {
	prefix string
}

// NewLogger creates a new logger with the given prefix
func NewLogger(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

// Info logs an informational message
func (l *Logger) Info(message string) {
	l.log("INFO", message)
}

// Error logs an error message
func (l *Logger) Error(message string) {
	l.log("ERROR", message)
}

// log formats and prints a log message
func (l *Logger) log(level, message string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s] [%s] [%s] %s\n", timestamp, level, l.prefix, message)
}
