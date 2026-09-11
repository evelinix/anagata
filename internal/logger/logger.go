package logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"AnagataSentinel/internal/config"
)

var log *slog.Logger
var rotator *Rotator

// Setup initializes the global logger based on config.
func Setup(cfg *config.LoggingConfig) {
	// Close previous logger if any
	if rotator != nil {
		rotator.Close()
		rotator = nil
	}

	level := parseLevel(cfg.Level)

	opts := &slog.HandlerOptions{
		Level: level,
	}

	var writers []io.Writer
	writers = append(writers, os.Stdout)

	if cfg.File != "" {
		if r, err := NewRotator(cfg.File, cfg.MaxSizeMB); err == nil {
			writers = append(writers, r)
			rotator = r
		}
	}

	writer := io.MultiWriter(writers...)

	var handler slog.Handler
	if level == slog.LevelDebug {
		handler = slog.NewTextHandler(writer, opts)
	} else {
		handler = slog.NewJSONHandler(writer, opts)
	}

	log = slog.New(handler)
	slog.SetDefault(log)
}

// Get returns the configured logger.
func Get() *slog.Logger {
	if log == nil {
		return slog.Default()
	}
	return log
}

// Close closes the log file if open.
func Close() {
	if rotator != nil {
		rotator.Close()
		rotator = nil
	}
}

func parseLevel(s string) slog.Level {
	switch s {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// ArchiveActiveLog moves the current active log to its rotated path.
// Called on graceful shutdown to ensure logs are properly archived.
func ArchiveActiveLog() {
	if rotator != nil {
		rotator.mu.Lock()
		rotator.rotate()
		rotator.mu.Unlock()
	}
}

// initLogDir ensures the logs directory exists relative to the executable.
func initLogDir(path string) {
	dir := filepath.Dir(path)
	os.MkdirAll(dir, 0755)
}
