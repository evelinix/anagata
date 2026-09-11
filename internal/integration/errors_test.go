package integration

import (
	"errors"
	"fmt"
	"testing"

	apperrors "AnagataSentinel/internal/errors"
)

func TestErrorWrapPreservesStage(t *testing.T) {
	original := fmt.Errorf("connection refused")

	wrapped := apperrors.Wrap("boot database", original)

	if wrapped.Stage != "boot database" {
		t.Errorf("expected stage 'boot database', got '%s'", wrapped.Stage)
	}
	if wrapped.Level != apperrors.LevelError {
		t.Errorf("expected level ERROR, got %s", wrapped.Level)
	}
	if !wrapped.Retriable {
		t.Error("expected retriable to be true")
	}
}

func TestErrorWrapPreservesCause(t *testing.T) {
	original := fmt.Errorf("disk full")

	wrapped := apperrors.Wrap("run migrations", original)

	if !errors.Is(wrapped, original) {
		t.Error("wrapped error should be unwrappable to original")
	}
	if wrapped.Unwrap() != original {
		t.Error("Unwrap() should return original error")
	}
}

func TestErrorWrapMessageFormat(t *testing.T) {
	original := fmt.Errorf("timeout")

	wrapped := apperrors.Wrap("open database", original)

	expected := "Failed to open database: timeout"
	if wrapped.Error() != expected {
		t.Errorf("expected error message '%s', got '%s'", expected, wrapped.Error())
	}
}

func TestAppErrorChaining(t *testing.T) {
	type stage struct {
		name string
		err  error
	}

	stages := []stage{
		{"load config", fmt.Errorf("yaml parse error")},
		{"setup logger", fmt.Errorf("permission denied")},
		{"init database", fmt.Errorf("encryption failed")},
	}

	var lastErr *apperrors.AppError
	for _, s := range stages {
		lastErr = apperrors.Wrap(s.name, s.err)
	}

	// Verify the last error has the correct stage
	if lastErr.Stage != "init database" {
		t.Errorf("expected stage 'init database', got '%s'", lastErr.Stage)
	}

	// Verify we can unwrap to the original error
	unwrapped := errors.Unwrap(lastErr)
	if unwrapped == nil {
		t.Error("Unwrap() should not return nil")
	}
	if unwrapped.Error() != "encryption failed" {
		t.Errorf("expected unwrapped error 'encryption failed', got '%s'", unwrapped.Error())
	}
}

func TestAppErrorNilCause(t *testing.T) {
	wrapped := apperrors.Wrap("some stage", nil)

	// When Err is nil, Error() returns just the message
	if wrapped.Stage != "some stage" {
		t.Errorf("expected stage 'some stage', got '%s'", wrapped.Stage)
	}
}

func TestAppErrorLevelPropagation(t *testing.T) {
	tests := []struct {
		name  string
		level apperrors.ErrorLevel
		want  string
	}{
		{"info", apperrors.LevelInfo, "INFO"},
		{"warning", apperrors.LevelWarning, "WARNING"},
		{"error", apperrors.LevelError, "ERROR"},
		{"fatal", apperrors.LevelFatal, "FATAL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := apperrors.New(tt.level, "test", "message", nil, false)
			if e.Level != tt.level {
				t.Errorf("expected level %v, got %v", tt.level, e.Level)
			}
			if e.Level.String() != tt.want {
				t.Errorf("expected String() '%s', got '%s'", tt.want, e.Level.String())
			}
		})
	}
}
