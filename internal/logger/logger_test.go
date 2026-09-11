package logger

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"AnagataSentinel/internal/config"
)

func TestSetup(t *testing.T) {
	dir := t.TempDir()

	cfg := &config.LoggingConfig{
		Level: "debug",
		File:  filepath.Join(dir, "test.log"),
	}

	Setup(cfg)

	l := Get()
	if l == nil {
		t.Fatal("logger is nil")
	}

	l.Info("test message", "key", "value")

	Close()

	data, err := os.ReadFile(cfg.File)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	if !bytes.Contains(data, []byte("test message")) {
		t.Errorf("log file does not contain expected message: %s", data)
	}
	if !bytes.Contains(data, []byte("key=value")) {
		t.Errorf("log file does not contain structured field: %s", data)
	}
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected slog.Level
	}{
		{"debug", slog.LevelDebug},
		{"info", slog.LevelInfo},
		{"warn", slog.LevelWarn},
		{"error", slog.LevelError},
		{"", slog.LevelInfo},
		{"invalid", slog.LevelInfo},
	}

	for _, tt := range tests {
		got := parseLevel(tt.input)
		if got != tt.expected {
			t.Errorf("parseLevel(%q) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func TestGetDefault(t *testing.T) {
	log = nil

	l := Get()
	if l == nil {
		t.Fatal("default logger is nil")
	}
}
