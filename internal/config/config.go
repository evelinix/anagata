package config

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App      AppConfig      `yaml:"app"`
	Window   WindowConfig   `yaml:"window"`
	Database DatabaseConfig `yaml:"database"`
	Logging  LoggingConfig  `yaml:"logging"`
	Splash   SplashConfig   `yaml:"splash"`
	Report   ReportConfig   `yaml:"report"`
}

type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Debug   bool   `yaml:"debug"`
}

type WindowConfig struct {
	Width     int `yaml:"width"`
	Height    int `yaml:"height"`
	MinWidth  int `yaml:"min_width"`
	MinHeight int `yaml:"min_height"`
}

type DatabaseConfig struct {
	Path     string `yaml:"path"`
	Password string `yaml:"password"`
}

type LoggingConfig struct {
	Level      string `yaml:"level"`
	File       string `yaml:"file"`
	MaxSizeMB  int    `yaml:"max_size_mb"`
	MaxBackups int    `yaml:"max_backups"`
	Compress   bool   `yaml:"compress"`
}

type SplashConfig struct {
	Enabled  bool `yaml:"enabled"`
	Duration int  `yaml:"duration_ms"`
}

type ReportConfig struct {
	Enabled bool   `yaml:"enabled"`
	DSN     string `yaml:"dsn"`
}

var (
	cfg      *Config
	portable bool
	dataDir  string
)

// Load reads config from file and applies env overrides.
// Portable mode: config.yaml next to exe
// Installed mode: %APPDATA%/AnagataSentinel/
func Load() (*Config, error) {
	cfg = defaultConfig()
	portable = detectPortable()

	if portable {
		slog.Info("running in portable mode")
	} else {
		slog.Info("running in installed mode")
	}

	if err := setupDataDir(); err != nil {
		return cfg, nil
	}

	path := configPath()

	if _, err := os.Stat(path); err == nil {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read config: %w", err)
		}
		if err := yaml.Unmarshal(data, cfg); err != nil {
			return nil, fmt.Errorf("parse config: %w", err)
		}
	}

	applyEnvOverrides(cfg)

	return cfg, nil
}

// Get returns the loaded config. Must call Load() first.
func Get() *Config {
	return cfg
}

// IsPortable returns true if running in portable mode.
func IsPortable() bool {
	return portable
}

// DataDir returns the data directory (portable: exe dir, installed: %APPDATA%).
func DataDir() string {
	return dataDir
}

// Save writes the current config to file.
func Save() error {
	path := configPath()

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// Dir returns the config directory, creating it if needed.
func Dir() (string, error) {
	path := configPath()
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return dir, nil
}

// detectPortable checks if config.yaml exists next to the executable.
func detectPortable() bool {
	exe, err := os.Executable()
	if err != nil {
		return false
	}

	configPath := filepath.Join(filepath.Dir(exe), "config.yaml")
	_, err = os.Stat(configPath)
	return err == nil
}

// setupDataDir sets the data directory based on mode.
func setupDataDir() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}

	if portable {
		dataDir = filepath.Dir(exe)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		dataDir = filepath.Join(home, "AppData", "Roaming", "AnagataSentinel")
	}

	return os.MkdirAll(dataDir, 0755)
}

func configPath() string {
	return filepath.Join(dataDir, "config.yaml")
}

func applyEnvOverrides(cfg *Config) {
	if v := os.Getenv("ANAGATA_DEBUG"); v != "" {
		cfg.App.Debug = v == "true"
	}
	if v := os.Getenv("ANAGATA_DB_PASSWORD"); v != "" {
		cfg.Database.Password = v
	}
	if v := os.Getenv("ANAGATA_LOG_LEVEL"); v != "" {
		cfg.Logging.Level = v
	}
	if v := os.Getenv("ANAGATA_REPORT_DSN"); v != "" {
		cfg.Report.DSN = v
		cfg.Report.Enabled = true
	}
}
