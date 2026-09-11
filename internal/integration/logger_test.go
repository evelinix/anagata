package integration

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"AnagataSentinel/internal/config"
	"AnagataSentinel/internal/logger"
)

func TestConfigToLoggerSetup(t *testing.T) {
	// Create a temp directory for the log file
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "test.log")

	// Create a LoggingConfig
	cfg := &config.LoggingConfig{
		Level:      "debug",
		File:       logFile,
		MaxSizeMB:  10,
		MaxBackups: 5,
		Compress:   false,
	}

	// Setup logger with the config
	logger.Setup(cfg)

	// Verify logger is configured
	l := logger.Get()
	if l == nil {
		t.Fatal("logger.Get() returned nil after Setup")
	}

	// Write a test log message
	l.Info("integration test message", "key", "value")

	// Verify the log file was created and contains the message
	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, "integration test message") {
		t.Errorf("log file should contain test message, got: %s", content)
	}
	if !strings.Contains(content, "key=value") {
		t.Errorf("log file should contain structured fields, got: %s", content)
	}

	// Cleanup
	logger.Close()
}

func TestLoggerDebugUsesTextHandler(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "debug.log")

	cfg := &config.LoggingConfig{
		Level: "debug",
		File:  logFile,
	}

	logger.Setup(cfg)
	l := logger.Get()
	l.Debug("debug message")

	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	// Debug level should use TextHandler (key=value format)
	content := string(data)
	if !strings.Contains(content, "debug message") {
		t.Errorf("log file should contain debug message, got: %s", content)
	}

	logger.Close()
}

func TestLoggerInfoUsesJSONHandler(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "info.log")

	cfg := &config.LoggingConfig{
		Level: "info",
		File:  logFile,
	}

	logger.Setup(cfg)
	l := logger.Get()
	l.Info("info message")

	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	// Info level should use JSONHandler (JSON format)
	content := string(data)
	if !strings.Contains(content, "info message") {
		t.Errorf("log file should contain info message, got: %s", content)
	}
	// JSON format should have quotes around keys
	if !strings.Contains(content, `"msg"`) {
		t.Errorf("JSON handler should have msg key, got: %s", content)
	}

	logger.Close()
}

func TestLoggerGlobalSlogPropagation(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "global.log")

	cfg := &config.LoggingConfig{
		Level: "info",
		File:  logFile,
	}

	// Setup logger (this sets slog.SetDefault)
	logger.Setup(cfg)

	// Use slog.Info directly (simulating how database/updater use slog)
	slog.Info("global slog message", "component", "database")

	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}

	content := string(data)
	if !strings.Contains(content, "global slog message") {
		t.Errorf("log file should contain global message, got: %s", content)
	}
	// JSON format uses "component":"database" not component=database
	if !strings.Contains(content, `"component":"database"`) {
		t.Errorf("log file should contain component field, got: %s", content)
	}

	logger.Close()
}
