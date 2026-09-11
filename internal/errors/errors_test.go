package errors

import (
	"errors"
	"testing"
)

func TestAppError(t *testing.T) {
	err := New(LevelError, "database", "connection failed", nil, true)

	if err.Error() != "connection failed" {
		t.Errorf("expected 'connection failed', got '%s'", err.Error())
	}
	if err.Stage != "database" {
		t.Errorf("expected stage 'database', got '%s'", err.Stage)
	}
	if err.Level != LevelError {
		t.Errorf("expected level ERROR, got %s", err.Level)
	}
	if !err.Retriable {
		t.Error("expected retriable")
	}
}

func TestAppErrorWithCause(t *testing.T) {
	cause := errors.New("timeout")
	err := New(LevelError, "config", "load failed", cause, false)

	if err.Error() != "load failed: timeout" {
		t.Errorf("unexpected error string: %s", err.Error())
	}
	if err.Unwrap() != cause {
		t.Error("Unwrap does not return cause")
	}
	if !errors.Is(err, cause) {
		t.Error("errors.Is should find cause")
	}
}

func TestWrap(t *testing.T) {
	cause := errors.New("connection refused")
	err := Wrap("database", cause)

	if err.Level != LevelError {
		t.Errorf("expected level ERROR, got %s", err.Level)
	}
	if err.Stage != "database" {
		t.Errorf("expected stage 'database', got '%s'", err.Stage)
	}
	if !err.Retriable {
		t.Error("expected retriable")
	}
	if err.Error() != "Failed to database: connection refused" {
		t.Errorf("unexpected error string: %s", err.Error())
	}
}

func TestErrorLevelString(t *testing.T) {
	tests := []struct {
		level    ErrorLevel
		expected string
	}{
		{LevelInfo, "INFO"},
		{LevelWarning, "WARNING"},
		{LevelError, "ERROR"},
		{LevelFatal, "FATAL"},
		{ErrorLevel(99), "UNKNOWN"},
	}

	for _, tt := range tests {
		if got := tt.level.String(); got != tt.expected {
			t.Errorf("ErrorLevel(%d).String() = %s, want %s", tt.level, got, tt.expected)
		}
	}
}
