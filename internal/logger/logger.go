package logger

import (
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"AnagataSentinel/internal/config"
)

var log *slog.Logger
var closer io.Closer

// Setup initializes the global logger based on config.
func Setup(cfg *config.LoggingConfig) {
	// Close previous logger if any
	if closer != nil {
		closer.Close()
		closer = nil
	}

	level := parseLevel(cfg.Level)

	opts := &slog.HandlerOptions{
		Level: level,
	}

	var writers []io.Writer
	writers = append(writers, os.Stdout)

	if cfg.File != "" {
		if err := os.MkdirAll(filepath.Dir(cfg.File), 0755); err == nil {
			f, err := os.OpenFile(cfg.File,
				os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err == nil {
				writers = append(writers, f)
				closer = f
			}
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
	if closer != nil {
		closer.Close()
		closer = nil
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
