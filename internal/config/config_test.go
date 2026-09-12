package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := defaultConfig()

	if cfg.App.Name != "AnagataSentinel" {
		t.Errorf("expected name 'AnagataSentinel', got '%s'", cfg.App.Name)
	}
	if cfg.Window.Width != 1280 {
		t.Errorf("expected width 1280, got %d", cfg.Window.Width)
	}
	if cfg.Database.Password != "0123456789" {
		t.Errorf("expected password '0123456789', got '%s'", cfg.Database.Password)
	}
	if cfg.Logging.Level != "info" {
		t.Errorf("expected log level 'info', got '%s'", cfg.Logging.Level)
	}
	if !cfg.Splash.Enabled {
		t.Error("expected splash enabled")
	}
}

func TestLoadDefaults(t *testing.T) {
	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.App.Name != "AnagataSentinel" {
		t.Errorf("expected name 'AnagataSentinel', got '%s'", cfg.App.Name)
	}
}

func TestEnvOverrides(t *testing.T) {
	os.Setenv("ANAGATA_DEBUG", "true")
	os.Setenv("ANAGATA_DB_PASSWORD", "secret123")
	os.Setenv("ANAGATA_LOG_LEVEL", "debug")
	defer os.Unsetenv("ANAGATA_DEBUG")
	defer os.Unsetenv("ANAGATA_DB_PASSWORD")
	defer os.Unsetenv("ANAGATA_LOG_LEVEL")

	cfg := defaultConfig()
	applyEnvOverrides(cfg)

	if !cfg.App.Debug {
		t.Error("expected debug to be true")
	}
	if cfg.Database.Password != "secret123" {
		t.Errorf("expected password 'secret123', got '%s'", cfg.Database.Password)
	}
	if cfg.Logging.Level != "debug" {
		t.Errorf("expected log level 'debug', got '%s'", cfg.Logging.Level)
	}
}

func TestConfigDir(t *testing.T) {
	dir, err := Dir()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	exe, err := os.Executable()
	if err != nil {
		t.Skip("cannot get executable path")
	}

	if portable {
		expected := filepath.Dir(exe)
		if dir != expected {
			t.Errorf("expected dir '%s', got '%s'", expected, dir)
		}
	} else {
		home, _ := os.UserHomeDir()
		expected := filepath.Join(home, "AppData", "Roaming", "AnagataSentinel")
		if dir != expected {
			t.Errorf("expected dir '%s', got '%s'", expected, dir)
		}
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Error("config directory was not created")
	}
}
