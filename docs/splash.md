Bisa. Dan kalau requirement-nya **splash benar-benar native, bukan `<Splash />` di frontend**, arsitekturnya sebaiknya seperti ini:

```text
┌──────────────────────────────────────────────┐
│                 OS starts EXE                │
└──────────────────────┬───────────────────────┘
                       │
                       ▼
              Native Splash Window
                 ┌─────────────┐
                 │   ANAGATA   │
                 │             │
                 │ Initializing│
                 │    ████     │
                 └─────────────┘
                       │
                       │ Go startup
                       ▼
              Backend initialization
                       │
          ┌────────────┼─────────────┐
          ▼            ▼             ▼
       Config       Database       Services
          │            │             │
          └────────────┼─────────────┘
                       ▼
                Wails WebView loads
                       │
                       ▼
                 OnDomReady()
                       │
             ┌─────────┴─────────┐
             │                   │
       Close native splash   Show main window
             │                   │
             └─────────┬─────────┘
                       ▼
                  Application
```

Wails v2 memang menyediakan `StartHidden`, `OnStartup`, dan `OnDomReady`. `OnStartup` dipanggil setelah frontend dibuat tetapi **sebelum `index.html` dimuat**, sedangkan `OnDomReady` dipanggil setelah frontend selesai dimuat. ([Wails][1])

Untuk splash yang **tidak bergantung pada frontend**, kita buat satu window native Windows menggunakan Go/Win32 melalui [`github.com/lxn/walk`](https://github.com/lxn/walk). Walk adalah toolkit GUI Windows native untuk Go, jadi splash ini bukan WebView. ([GitHub][2])

> Saya buat versi **Windows-native terlebih dahulu**, karena ini yang paling clean untuk Wails Windows. Struktur ini tetap memungkinkan kita menambahkan GTK-native splash untuk Fedora/Linux kemudian.

---

# 1. Struktur project

Misalnya project Wails kamu:

```text
my-app/
│
├── main.go
├── app.go
│
├── splash.go
├── splash_windows.go
├── splash_other.go
│
├── frontend/
│   ├── src/
│   ├── index.html
│   └── ...
│
├── build/
│   ├── appicon.png
│   └── windows/
│
├── go.mod
├── go.sum
└── wails.json
```

Kita pisahkan `splash_windows.go` dengan build tag `_windows.go`.

---

# 2. Install native Windows GUI dependency

Dari root project:

```powershell
go get github.com/lxn/walk
```

Walk memang merupakan Windows GUI toolkit dan menggunakan Win32 API di bawahnya. ([GitHub][2])

---

# 3. `main.go`

Ini konfigurasi Wails utamanya.

```go
package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	app := NewApp()

	err := wails.Run(&options.App{
		Title: "Anagata Application",

		// Main window
		Width: 1280,
		Height: 800,

		MinWidth: 1024,
		MinHeight: 640,

		DisableResize: false,

		// IMPORTANT:
		// Main Wails window tidak langsung ditampilkan.
		StartHidden: true,

		// Main window background.
		BackgroundColour: &options.RGBA{
			R: 248,
			G: 250,
			B: 252,
			A: 255,
		},

		// Frontend assets
		AssetServer: &assetserver.Options{
			Assets: assets,
		},

		// Lifecycle
		OnStartup:  app.Startup,
		OnDomReady: app.DOMReady,
		OnShutdown: app.Shutdown,

		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
```

Yang paling penting di sini:

```go
StartHidden: true,
```

Jadi window Wails utama tidak muncul ketika proses baru dimulai. Wails memang menyediakan opsi `StartHidden` untuk menahan window sampai `WindowShow()` dipanggil. ([Wails][1])

---

# 4. `app.go`

Sekarang lifecycle aplikasinya.

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context

	splash *NativeSplash

	startupError error
}

func NewApp() *App {
	return &App{}
}

// Startup:
//
// Dipanggil Wails sebelum index.html frontend dimuat.
//
// Di sinilah native splash dibuat dan backend initialization
// dijalankan.
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	log.Println("[BOOT] Application startup")

	// --------------------------------------------------
	// 1. Create native splash
	// --------------------------------------------------

	a.splash = NewNativeSplash()

	if err := a.splash.Start(); err != nil {
		log.Printf("[BOOT] Splash error: %v", err)
	}

	// --------------------------------------------------
	// 2. Backend initialization
	// --------------------------------------------------

	a.initializeBackend()
}

// initializeBackend menjalankan seluruh initialization
// yang harus selesai sebelum aplikasi digunakan.
func (a *App) initializeBackend() {
	log.Println("[BOOT] Initializing backend")

	// -----------------------------------------------
	// Config
	// -----------------------------------------------

	a.splashStatus("Loading configuration...")
	time.Sleep(300 * time.Millisecond)

	// -----------------------------------------------
	// Database
	// -----------------------------------------------

	a.splashStatus("Initializing database...")
	time.Sleep(500 * time.Millisecond)

	// -----------------------------------------------
	// Security
	// -----------------------------------------------

	a.splashStatus("Initializing security...")
	time.Sleep(400 * time.Millisecond)

	// -----------------------------------------------
	// Services
	// -----------------------------------------------

	a.splashStatus("Starting services...")
	time.Sleep(500 * time.Millisecond)

	// -----------------------------------------------
	// Finalization
	// -----------------------------------------------

	a.splashStatus("Preparing application...")
	time.Sleep(300 * time.Millisecond)

	log.Println("[BOOT] Backend initialization complete")
}

// Helper untuk update text pada native splash.
func (a *App) splashStatus(message string) {
	if a.splash == nil {
		return
	}

	a.splash.SetStatus(message)
}

// DOM Ready:
//
// Dipanggil setelah frontend Wails selesai load.
func (a *App) DOMReady(ctx context.Context) {
	log.Println("[BOOT] DOM ready")

	// Pastikan frontend sudah benar-benar siap
	// sebelum menampilkan main window.

	runtime.WindowShow(ctx)

	// Close native splash.
	if a.splash != nil {
		a.splash.Close()
		a.splash = nil
	}

	log.Println("[BOOT] Application ready")
}

// Shutdown
func (a *App) Shutdown(ctx context.Context) {
	log.Println("[BOOT] Application shutdown")

	if a.splash != nil {
		a.splash.Close()
		a.splash = nil
	}
}
```

---

# 5. `splash.go`

Ini abstraction supaya `app.go` tidak perlu tahu apakah splash menggunakan Windows, Linux, atau macOS.

```go
package main

type NativeSplash struct {
	platformSplash platformSplash
}

type platformSplash interface {
	Start() error
	SetStatus(message string)
	Close()
}

func NewNativeSplash() *NativeSplash {
	return &NativeSplash{
		platformSplash: newPlatformSplash(),
	}
}

func (s *NativeSplash) Start() error {
	return s.platformSplash.Start()
}

func (s *NativeSplash) SetStatus(message string) {
	s.platformSplash.SetStatus(message)
}

func (s *NativeSplash) Close() {
	s.platformSplash.Close()
}
```

---

# 6. `splash_windows.go`

Nah, ini bagian native-nya.

```go
//go:build windows

package main

import (
	"runtime"
	"sync"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

type windowsSplash struct {
	window *walk.MainWindow
	status *walk.Label

	ready chan error
	once  sync.Once
}

func newPlatformSplash() platformSplash {
	return &windowsSplash{
		ready: make(chan error, 1),
	}
}

func (s *windowsSplash) Start() error {
	go s.run()

	return <-s.ready
}

func (s *windowsSplash) run() {
	// Native Windows controls/window harus dibuat pada
	// OS thread yang sesuai.
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	var window *walk.MainWindow
	var status *walk.Label

	background, err := walk.NewSolidColorBrush(
		walk.RGB(15, 23, 42),
	)
	if err != nil {
		s.ready <- err
		return
	}

	defer background.Dispose()

	// --------------------------------------------------
	// Create native splash
	// --------------------------------------------------

	err = (MainWindow{
		AssignTo: &window,

		Title: "Anagata",

		Size: walk.Size{
			Width: 520,
			Height: 300,
		},

		MinSize: walk.Size{
			Width: 520,
			Height: 300,
		},

		MaxSize: walk.Size{
			Width: 520,
			Height: 300,
		},

		Background: background,

		Visible: false,

		Layout: VBox{
			Margins: Margins{
				HNear: 40,
				VNear: 30,
				HFar: 40,
				VFar: 30,
			},

			Spacing: 8,
		},

		Children: []Widget{
			VSpacer{},

			Label{
				Text: "ANAGATA",

				Alignment: AlignHCenterVCenter,

				Font: Font{
					Family:    "Segoe UI",
					PointSize: 26,
					Bold:      true,
				},

				TextColor: walk.RGB(
					255,
					255,
					255,
				),
			},

			Label{
				Text: "Application",

				Alignment: AlignHCenterVCenter,

				Font: Font{
					Family:    "Segoe UI",
					PointSize: 10,
				},

				TextColor: walk.RGB(
					148,
					163,
					184,
				),
			},

			VSpacer{
				Size: 12,
			},

			Label{
				AssignTo: &status,

				Text: "Starting...",

				Alignment: AlignHCenterVCenter,

				Font: Font{
					Family:    "Segoe UI",
					PointSize: 9,
				},

				TextColor: walk.RGB(
					203,
					213,
					225,
				),
			},

			ProgressBar{
				MarqueeMode: true,

				MinSize: walk.Size{
					Width: 300,
					Height: 4,
				},
			},

			VSpacer{},
		},
	}).Create()

	if err != nil {
		s.ready <- err
		return
	}

	// Simpan reference supaya goroutine lain bisa
	// meng-update status.
	s.window = window
	s.status = status

	// --------------------------------------------------
	// Remove standard Windows frame
	// --------------------------------------------------

	removeWindowFrame(window)

	// --------------------------------------------------
	// Center splash
	// --------------------------------------------------

	centerWindow(window)

	// --------------------------------------------------
	// Show splash
	// --------------------------------------------------

	window.Show()

	// Signal bahwa splash sudah benar-benar dibuat.
	s.ready <- nil

	// --------------------------------------------------
	// Native message loop
	// --------------------------------------------------

	window.Run()
}

func (s *windowsSplash) SetStatus(message string) {
	window := s.window
	status := s.status

	if window == nil || status == nil {
		return
	}

	window.Synchronize(func() {
		if status != nil {
			status.SetText(message)
		}
	})
}

func (s *windowsSplash) Close() {
	s.once.Do(func() {
		window := s.window

		if window == nil {
			return
		}

		window.Synchronize(func() {
			_ = window.Close()
		})
	})
}

// ------------------------------------------------------
// Remove title bar / border
// ------------------------------------------------------

func removeWindowFrame(window *walk.MainWindow) {
	hwnd := window.Handle()

	style := uint32(
		win.GetWindowLong(
			hwnd,
			win.GWL_STYLE,
		),
	)

	style &^= win.WS_CAPTION
	style &^= win.WS_THICKFRAME
	style &^= win.WS_MINIMIZEBOX
	style &^= win.WS_MAXIMIZEBOX
	style &^= win.WS_SYSMENU

	style |= win.WS_POPUP

	win.SetWindowLong(
		hwnd,
		win.GWL_STYLE,
		int32(style),
	)

	exStyle := uint32(
		win.GetWindowLong(
			hwnd,
			win.GWL_EXSTYLE,
		),
	)

	exStyle |= win.WS_EX_TOOLWINDOW
	exStyle |= win.WS_EX_TOPMOST

	win.SetWindowLong(
		hwnd,
		win.GWL_EXSTYLE,
		int32(exStyle),
	)

	win.SetWindowPos(
		hwnd,
		win.HWND_TOPMOST,
		0,
		0,
		0,
		0,
		win.SWP_NOMOVE|
			win.SWP_NOSIZE|
			win.SWP_FRAMECHANGED|
			win.SWP_NOACTIVATE,
	)
}

// ------------------------------------------------------
// Center splash on primary monitor
// ------------------------------------------------------

func centerWindow(window *walk.MainWindow) {
	width := 520
	height := 300

	screenWidth := int(
		win.GetSystemMetrics(
			win.SM_CXSCREEN,
		),
	)

	screenHeight := int(
		win.GetSystemMetrics(
			win.SM_CYSCREEN,
		),
	)

	x := (screenWidth - width) / 2
	y := (screenHeight - height) / 2

	_ = window.SetBoundsPixels(
		walk.Rectangle{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		},
	)
}
```

### Kenapa ada `runtime.LockOSThread()`?

Karena window native Windows dibuat menggunakan Win32 dan Walk sendiri juga mengharuskan operasi window initialization dilakukan pada OS thread yang sesuai. Source Walk secara eksplisit menggunakan `runtime.LockOSThread()` saat inisialisasi window. ([GitHub][3])

Dan karena kita menjalankan message loop sendiri:

```go
window.Run()
```

splash tetap responsif selama proses startup berlangsung.

---

# 7. `splash_other.go`

Supaya project tetap bisa di-compile pada Linux/macOS:

```go
//go:build !windows

package main

type otherPlatformSplash struct{}

func newPlatformSplash() platformSplash {
	return &otherPlatformSplash{}
}

func (s *otherPlatformSplash) Start() error {
	return nil
}

func (s *otherPlatformSplash) SetStatus(message string) {
	// No-op
}

func (s *otherPlatformSplash) Close() {
	// No-op
}
```

Jadi:

```text
Windows
    ↓
Native Win32 Splash

Linux
    ↓
No-op sementara

macOS
    ↓
No-op sementara
```

Ini sengaja saya buat begitu karena **native splash lintas platform bukan satu API Wails v2**. Wails v2 sendiri menggunakan WebView sebagai frontend window, sementara native window API runtime-nya ditujukan untuk mengontrol window utama. ([Wails][4])

Untuk Fedora/Linux, nanti `splash_linux.go` bisa kita implementasikan menggunakan GTK native. Wails Linux sendiri memang sudah bergantung pada GTK/WebKitGTK. ([Wails][5])

---

# 8. Hasil lifecycle

Sekarang startup menjadi:

```text
main()
 │
 ├── NewApp()
 │
 └── wails.Run()
       │
       ▼
   Wails creates
   main window
       │
       │ StartHidden=true
       ▼
   OnStartup()
       │
       ├── Create Native Splash
       │       │
       │       └── ANAGATA
       │
       ├── Loading configuration
       │
       ├── Initializing database
       │
       ├── Initializing security
       │
       ├── Starting services
       │
       └── Preparing application
       │
       ▼
   Startup returns
       │
       ▼
   Wails loads index.html
       │
       ▼
   Frontend initialized
       │
       ▼
   OnDomReady()
       │
       ├── runtime.WindowShow()
       │
       └── NativeSplash.Close()
       │
       ▼
   MAIN APPLICATION
```

Jadi **frontend sama sekali tidak perlu tahu tentang splash**.

Tidak ada:

```ts
if (loading) {
    return <Splash />
}
```

Tidak ada:

```html
<div id="splash">
```

Tidak ada:

```css
#splash {
    ...
}
```

Splash hidup di sisi **Go/native Windows**.

---

# 9. Jangan pakai `time.Sleep()` untuk production

Bagian ini:

```go
time.Sleep(300 * time.Millisecond)
```

di atas cuma simulasi supaya kelihatan alurnya.

Production sebaiknya:

```go
func (a *App) initializeBackend() {
	a.splashStatus("Loading configuration...")

	if err := a.loadConfig(); err != nil {
		a.startupError = err
		return
	}

	a.splashStatus("Initializing database...")

	if err := a.initDatabase(); err != nil {
		a.startupError = err
		return
	}

	a.splashStatus("Initializing security...")

	if err := a.initSecurity(); err != nil {
		a.startupError = err
		return
	}

	a.splashStatus("Starting services...")

	if err := a.initServices(); err != nil {
		a.startupError = err
		return
	}

	a.splashStatus("Ready")
}
```

Ini jauh lebih cocok untuk aplikasi serius.

---

# 10. Lebih bagus lagi: gunakan boot stages

Untuk aplikasi enterprise, saya justru akan membuat startup seperti:

```go
type BootStage string

const (
	BootConfig     BootStage = "config"
	BootEnvironment BootStage = "environment"
	BootSecurity   BootStage = "security"
	BootDatabase   BootStage = "database"
	BootServices   BootStage = "services"
	BootFrontend   BootStage = "frontend"
	BootReady      BootStage = "ready"
)
```

Kemudian:

```go
func (a *App) initializeBackend() error {
	if err := a.bootConfig(); err != nil {
		return err
	}

	if err := a.bootEnvironment(); err != nil {
		return err
	}

	if err := a.bootSecurity(); err != nil {
		return err
	}

	if err := a.bootDatabase(); err != nil {
		return err
	}

	if err := a.bootServices(); err != nil {
		return err
	}

	return nil
}
```

Splash tinggal menampilkan:

```text
ANAGATA

Initializing...

✓ Configuration
✓ Environment
✓ Security
● Database
○ Services

████████████████░░░░░░
```

Tetapi semua itu tetap **native**.

---

# 11. Satu perbaikan penting untuk `DOMReady`

Saya sarankan jangan langsung:

```go
runtime.WindowShow(ctx)
a.splash.Close()
```

kalau kita ingin transisi yang lebih halus.

Buat:

```go
func (a *App) DOMReady(ctx context.Context) {
	runtime.WindowShow(ctx)

	if a.splash != nil {
		a.splash.Close()
		a.splash = nil
	}
}
```

Ini sudah aman karena `OnDomReady` adalah lifecycle point ketika frontend selesai dimuat, dan dokumentasi Wails memang menyebut runtime calls lebih tepat dilakukan di `OnDomReady` daripada `OnStartup`. ([Wails][6])

---

# 12. Kalau ingin splash hanya 1 window, ada opsi lain

Ada satu alternatif yang **lebih sederhana**, tetapi bukan native splash terpisah:

```go
StartHidden: false
```

kemudian:

```go
OnStartup
    ↓
window background
    ↓
OnDomReady
    ↓
frontend
```

Namun itu tetap bergantung pada **WebView/Wails window**.

Untuk requirement kamu:

> **Splash → Main Window tanpa frontend menampilkan splash**

maka desain yang saya pilih adalah:

```text
             WAILS PROCESS
                   │
       ┌───────────┴───────────┐
       │                       │
       ▼                       ▼
Native Splash             Wails Main Window
  Win32                       WebView2
       │                       │
       │                       │ hidden
       │                       │
       └──────────┬────────────┘
                  │
             Boot complete
                  │
                  ▼
           DOMReady callback
                  │
        ┌─────────┴─────────┐
        ▼                   ▼
   Destroy Splash       Show Wails
                            │
                            ▼
                       Application
```

**Ini jauh lebih proper daripada membuat splash di React/Svelte/Vue**, terutama kalau tujuanmu adalah menghilangkan kesan *blank screen* saat startup.

Wails memang mendukung `StartHidden`, `OnStartup`, dan `OnDomReady` untuk pola lifecycle seperti ini. ([Wails][1])

### Catatan penting

Kode di atas membuat **Windows-native splash** dengan Walk. Kalau target aplikasi kamu adalah **Windows + Fedora/Linux**, saya sarankan langkah berikutnya bukan memakai fallback `no-op`, tetapi membuat:

```text
splash_windows.go   → Win32 / Walk
splash_linux.go     → GTK3
splash_darwin.go    → Cocoa
splash.go           → common interface
```

Sehingga hasil akhirnya benar-benar:

```text
Windows  → native Win32 splash
Linux    → native GTK splash
macOS    → native Cocoa splash
```

tanpa **satu baris pun frontend** untuk mekanisme splash.

[1]: https://v2.wails.io/docs/reference/options/?utm_source=chatgpt.com "Options | Wails"
[2]: https://github.com/lxn/walk?utm_source=chatgpt.com "GitHub - lxn/walk: A Windows GUI toolkit for the Go Programming Language · GitHub"
[3]: https://github.com/lxn/walk/blob/master/window.go "walk/window.go at master · lxn/walk · GitHub"
[4]: https://v2.wails.io/docs/reference/runtime/window/?utm_source=chatgpt.com "Window | Wails"
[5]: https://v2.wails.io/docs/guides/linux-distro-support/?utm_source=chatgpt.com "Linux Distro Support | Wails"
[6]: https://v2.wails.io/docs/reference/runtime/intro/?utm_source=chatgpt.com "Introduction | Wails"
