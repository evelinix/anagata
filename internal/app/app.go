package app

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"AnagataSentinel/internal/config"
	"AnagataSentinel/internal/database"
	apperrors "AnagataSentinel/internal/errors"
	"AnagataSentinel/internal/logger"
	"AnagataSentinel/internal/splash"
	"AnagataSentinel/internal/version"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx    context.Context
	config *config.Config

	splash *splash.NativeSplash

	startupError  *apperrors.AppError
	backendReady chan struct{}
	domReady     chan struct{}
	initOnce     sync.Once
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
