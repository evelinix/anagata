package integration

import (
	"os"
	"path/filepath"
	"testing"

	"AnagataSentinel/internal/config"
	"AnagataSentinel/internal/logger"
)

func TestConfigToLoggerToFile(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "app.log")

	cfg := &config.LoggingConfig{
		Level:      "info",
		File:       logFile,
		MaxSizeMB:  50,
		MaxBackups: 3,
		Compress:   true,
	}

	// Setup logger
	logger.Setup(cfg)

	// Verify config is propagated
	l := logger.Get()
	if l == nil {
		t.Fatal("logger should not be nil")
	}

	// Write logs
	l.Info("test 1")
	l.Info("test 2")

	// Verify file exists
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Fatal("log file should exist")
	}

	logger.Close()
}

func TestConfigToLoggerRotation(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "app.log")

	cfg := &config.LoggingConfig{
		Level:     "info",
		File:      logFile,
		MaxSizeMB: 0, // 0 = 100MB default
	}

	logger.Setup(cfg)
	l := logger.Get()

	// Write enough data to trigger rotation (simulated)
	for i := 0; i < 100; i++ {
		l.Info("test message", "iteration", i)
	}

	// Verify log file exists
	if _, err := os.Stat(logFile); os.IsNotExist(err) {
		t.Fatal("log file should exist")
	}

	logger.Close()
}

func TestLoggerWithConfigLevels(t *testing.T) {
	tests := []struct {
		name     string
		level    string
		logDebug bool
		logInfo  bool
		logWarn  bool
		logError bool
	}{
		{"debug captures all", "debug", true, true, true, true},
		{"info captures info+", "info", false, true, true, true},
		{"warn captures warn+", "warn", false, false, true, true},
		{"error captures error", "error", false, false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir := t.TempDir()
			logFile := filepath.Join(tmpDir, "test.log")

			cfg := &config.LoggingConfig{
				Level: tt.level,
				File:  logFile,
			}

			logger.Setup(cfg)
			l := logger.Get()

			l.Debug("debug msg")
			l.Info("info msg")
			l.Warn("warn msg")
			l.Error("error msg")

			data, _ := os.ReadFile(logFile)
			content := string(data)

			if tt.logDebug && !containsString(content, "debug msg") {
				t.Error("expected debug message in log")
			}
			if tt.logInfo && !containsString(content, "info msg") {
				t.Error("expected info message in log")
			}
			if tt.logWarn && !containsString(content, "warn msg") {
				t.Error("expected warn message in log")
			}
			if tt.logError && !containsString(content, "error msg") {
				t.Error("expected error message in log")
			}

			logger.Close()
		})
	}
}

func TestConfigDefaultValues(t *testing.T) {
	// Load config (uses defaults when no config file exists)
	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Verify default values are set
	if cfg.App.Name != "AnagataSentinel" {
		t.Errorf("expected app name 'AnagataSentinel', got '%s'", cfg.App.Name)
	}
	if cfg.Window.Width != 1280 {
		t.Errorf("expected window width 1280, got %d", cfg.Window.Width)
	}
	if cfg.Database.Password != "0123456789" {
		t.Errorf("expected default password '0123456789', got '%s'", cfg.Database.Password)
	}
	if cfg.Logging.Level != "info" {
		t.Errorf("expected log level 'info', got '%s'", cfg.Logging.Level)
	}
}

func containsString(s, substr string) bool {
	return len(s) >= len(substr) && searchString(s, substr)
}

func searchString(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
