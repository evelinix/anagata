package config

import (
	"fmt"
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

var cfg *Config

// Load reads config from file and applies env overrides.
// If no config file exists, defaults are used.
func Load() (*Config, error) {
	cfg = defaultConfig()

	path, err := configPath()
	if err != nil {
		return cfg, nil
	}

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

// Save writes the current config to file.
func Save() error {
	path, err := configPath()
	if err != nil {
		return err
	}

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
	path, err := configPath()
	if err != nil {
		return "", err
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	return dir, nil
}

func configPath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Join(filepath.Dir(exe), "config.yaml"), nil
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
}
