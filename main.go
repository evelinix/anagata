package main

import (
	"embed"
	"log/slog"

	"AnagataSentinel/internal/app"
	"AnagataSentinel/internal/config"
	"AnagataSentinel/internal/logger"
	"AnagataSentinel/internal/reporter"
	"AnagataSentinel/internal/version"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		panic(err)
	}

	logger.Setup(&cfg.Logging)
	slog.Info("application starting",
		"app", cfg.App.Name,
		"version", version.Version,
		"commit", version.GitCommit[:min(8, len(version.GitCommit))],
	)

	var rep *reporter.Reporter
	if cfg.Report.Enabled {
		rep = reporter.New(cfg.Report.DSN)
	} else {
		rep = reporter.New("")
	}

	a := app.NewApp()

	defer func() {
		if rep != nil {
			rep.Flush()
		}
	}()

	err = wails.Run(&options.App{
		Title:  cfg.App.Name,
		Width:  cfg.Window.Width,
		Height: cfg.Window.Height,

		MinWidth:  cfg.Window.MinWidth,
		MinHeight: cfg.Window.MinHeight,

		DisableResize: false,

		StartHidden: true,

		BackgroundColour: &options.RGBA{
			R: 248,
			G: 250,
			B: 252,
			A: 255,
		},

		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		OnStartup:  a.Startup,
		OnDomReady: a.DOMReady,
		OnShutdown: a.Shutdown,

		Bind: []interface{}{
			a,
		},
	})

	if err != nil {
		slog.Error("failed to run application", "error", err)
		panic(err)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
