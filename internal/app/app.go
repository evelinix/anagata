package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"
	"time"

	"AnagataSentinel/internal/config"
	"AnagataSentinel/internal/database"
	apperrors "AnagataSentinel/internal/errors"
	"AnagataSentinel/internal/logger"
	"AnagataSentinel/internal/splash"
	"AnagataSentinel/internal/updater"
	"AnagataSentinel/internal/version"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx    context.Context
	config *config.Config

	splash  *splash.NativeSplash
	checker *updater.Checker

	startupError   *apperrors.AppError
	backendReady   chan struct{}
	domReady       chan struct{}
	initOnce       sync.Once
	pendingRelease *updater.Release
	downloadDir    string
}

func NewApp() *App {
	return &App{
		backendReady: make(chan struct{}),
		domReady:     make(chan struct{}),
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	slog.Info("application startup")

	a.splash = splash.NewNativeSplash()

	if err := a.splash.Start(); err != nil {
		slog.Error("splash error", "error", err)
	}

	go func() {
		a.initializeBackend()
		close(a.backendReady)
	}()

	go func() {
		<-a.backendReady
		<-a.domReady
		a.showMainApp()
	}()
}

func (a *App) initializeBackend() {
	slog.Info("initializing backend")

	stages := []struct {
		name    string
		handler func() error
	}{
		{"Loading configuration", a.bootConfig},
		{"Initializing database", a.bootDatabase},
		{"Initializing security", a.bootSecurity},
		{"Starting services", a.bootServices},
		{"Preparing application", a.bootFinalize},
	}

	for _, stage := range stages {
		a.splash.SetStatus(stage.name + "...")

		if err := stage.handler(); err != nil {
			slog.Error("boot stage failed", "stage", stage.name, "error", err)
			a.startupError = apperrors.Wrap(stage.name, err)
			return
		}

		slog.Info("boot stage complete", "stage", stage.name)
	}

	slog.Info("backend initialization complete")
}

func (a *App) bootConfig() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}
	a.config = cfg
	logger.Setup(&cfg.Logging)
	return nil
}

func (a *App) bootDatabase() error {
	return database.Init(a.config)
}

func (a *App) bootSecurity() error {
	return nil
}

func (a *App) bootServices() error {
	if a.config != nil {
		a.checker = updater.NewChecker("AnagataSentinel", "AnagataSentinel")
		a.downloadDir = filepath.Join(os.TempDir(), "AnagataSentinel-updates")
		a.checker.StartPeriodicCheck(1*time.Hour, func(release *updater.Release) {
			slog.Info("update available", "version", release.TagName)
			a.pendingRelease = release
			runtime.EventsEmit(a.ctx, "update-available", map[string]interface{}{
				"version":     release.TagName,
				"name":        release.Name,
				"url":         release.HTMLURL,
				"description": release.Body,
			})
		})
	}
	return nil
}

func (a *App) bootFinalize() error {
	return nil
}

func (a *App) showMainApp() {
	a.initOnce.Do(func() {
		if a.startupError != nil {
			slog.Error("startup error", "stage", a.startupError.Stage, "error", a.startupError)
			retry := apperrors.ShowDialog(a.startupError)
			if retry {
				slog.Info("retrying failed stage", "stage", a.startupError.Stage)
				a.startupError = nil
				a.initOnce = sync.Once{}
				a.initializeBackend()
				close(a.backendReady)
				return
			}
			runtime.Quit(a.ctx)
			return
		}

		slog.Info("showing main application")

		runtime.WindowShow(a.ctx)

		if a.splash != nil {
			a.splash.Close()
			a.splash = nil
		}

		slog.Info("application ready")
	})
}

func (a *App) DOMReady(ctx context.Context) {
	slog.Info("DOM ready")
	close(a.domReady)
}

func (a *App) Shutdown(ctx context.Context) {
	slog.Info("application shutdown")

	database.Close()

	if a.splash != nil {
		a.splash.Close()
		a.splash = nil
	}
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Version() string {
	return version.Info()
}

func (a *App) GetVersionInfo() map[string]string {
	return map[string]string{
		"version":   version.Version,
		"commit":    version.GitCommit,
		"buildTime": version.BuildTime,
		"goVersion": version.GoVersion,
	}
}

func (a *App) CheckForUpdate() (map[string]interface{}, error) {
	if a.checker == nil {
		return map[string]interface{}{
			"available": false,
			"message":   "updater not initialized",
		}, nil
	}

	available, release, err := a.checker.IsUpdateAvailable()
	if err != nil {
		return nil, err
	}

	if !available {
		return map[string]interface{}{
			"available": false,
			"message":   "no update available",
		}, nil
	}

	a.pendingRelease = release

	return map[string]interface{}{
		"available":   true,
		"version":     release.TagName,
		"name":        release.Name,
		"url":         release.HTMLURL,
		"description": release.Body,
	}, nil
}

func (a *App) DownloadUpdate() (map[string]interface{}, error) {
	if a.checker == nil {
		return map[string]interface{}{
			"success": false,
			"error":   "updater not initialized",
		}, nil
	}

	if a.pendingRelease == nil {
		available, release, err := a.checker.IsUpdateAvailable()
		if err != nil {
			return map[string]interface{}{
				"success": false,
				"error":   err.Error(),
			}, nil
		}
		if !available {
			return map[string]interface{}{
				"success": false,
				"error":   "no update available",
			}, nil
		}
		a.pendingRelease = release
	}

	filePath, err := a.checker.DownloadUpdate(a.pendingRelease, a.downloadDir, func(progress updater.DownloadProgress) {
		runtime.EventsEmit(a.ctx, "update-progress", progress)
	})
	if err != nil {
		slog.Error("download update failed", "error", err)
		return map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}, nil
	}

	return map[string]interface{}{
		"success":   true,
		"file_path": *filePath,
		"version":   a.pendingRelease.TagName,
	}, nil
}

func (a *App) GetPendingRelease() map[string]interface{} {
	if a.pendingRelease == nil {
		return map[string]interface{}{
			"available": false,
		}
	}
	return map[string]interface{}{
		"available":   true,
		"version":     a.pendingRelease.TagName,
		"name":        a.pendingRelease.Name,
		"url":         a.pendingRelease.HTMLURL,
		"description": a.pendingRelease.Body,
	}
}

func (a *App) ApplyUpdate() (map[string]interface{}, error) {
	if a.checker == nil {
		return map[string]interface{}{
			"success": false,
			"error":   "updater not initialized",
		}, nil
	}

	if a.pendingRelease == nil {
		return map[string]interface{}{
			"success": false,
			"error":   "no pending release",
		}, nil
	}

	filePath, found := a.checker.GetDownloadedFilePath(a.pendingRelease)
	if !found || filePath == "" {
		return map[string]interface{}{
			"success": false,
			"error":   "downloaded file not found, please download again",
		}, nil
	}

	slog.Info("applying update", "version", a.pendingRelease.TagName, "path", filePath)

	if err := a.checker.ApplyUpdate(filePath); err != nil {
		slog.Error("apply update failed", "error", err)
		return map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		}, nil
	}

	runtime.EventsEmit(a.ctx, "update-applied", map[string]interface{}{
		"version": a.pendingRelease.TagName,
	})

	return map[string]interface{}{
		"success": true,
		"version": a.pendingRelease.TagName,
	}, nil
}
