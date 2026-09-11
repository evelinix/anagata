package app

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"AnagataSentinel/internal/database"
	"AnagataSentinel/internal/splash"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context

	splash *splash.NativeSplash

	startupError  error
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

	log.Println("[BOOT] Application startup")

	a.splash = splash.NewNativeSplash()

	if err := a.splash.Start(); err != nil {
		log.Printf("[BOOT] Splash error: %v", err)
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
	log.Println("[BOOT] Initializing backend")

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
			log.Printf("[BOOT] Stage '%s' failed: %v", stage.name, err)
			a.startupError = err
			return
		}

		log.Printf("[BOOT] Stage '%s' complete", stage.name)
	}

	log.Println("[BOOT] Backend initialization complete")
}

func (a *App) bootConfig() error {
	time.Sleep(2 * time.Second)
	return nil
}

func (a *App) bootDatabase() error {
	return database.Init()
}

func (a *App) bootSecurity() error {
	time.Sleep(2 * time.Second)
	return nil
}

func (a *App) bootServices() error {
	time.Sleep(2 * time.Second)
	return nil
}

func (a *App) bootFinalize() error {
	time.Sleep(2 * time.Second)
	return nil
}

func (a *App) showMainApp() {
	a.initOnce.Do(func() {
		log.Println("[BOOT] Showing main application")

		runtime.WindowShow(a.ctx)

		if a.splash != nil {
			a.splash.Close()
			a.splash = nil
		}

		log.Println("[BOOT] Application ready")
	})
}

func (a *App) DOMReady(ctx context.Context) {
	log.Println("[BOOT] DOM ready")
	close(a.domReady)
}

func (a *App) Shutdown(ctx context.Context) {
	log.Println("[BOOT] Application shutdown")

	database.Close()

	if a.splash != nil {
		a.splash.Close()
		a.splash = nil
	}
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
