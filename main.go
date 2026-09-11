package main

import (
	"embed"
	"log"

	"AnagataSentinel/internal/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	a := app.NewApp()

	err := wails.Run(&options.App{
		Title:  "AnagataSentinel",
		Width:  1280,
		Height: 800,

		MinWidth:  1024,
		MinHeight: 640,

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
		log.Fatal(err)
	}
}
