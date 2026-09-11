//go:build windows

package splash

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
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	var window *walk.MainWindow
	var status *walk.Label

	err := (MainWindow{
		AssignTo: &window,

		Title: "Anagata Sentinel",

		Size: Size{
			Width:  360,
			Height: 180,
		},

		MinSize: Size{
			Width:  360,
			Height: 180,
		},

		MaxSize: Size{
			Width:  360,
			Height: 180,
		},

		Background: SolidColorBrush{Color: walk.RGB(32, 32, 32)},

		Visible: false,

		Layout: VBox{
			Margins: Margins{Left: 0, Top: 0, Right: 0, Bottom: 0},
			Spacing: 0,
		},

		Children: []Widget{
			VSpacer{Size: 28},

			// --- Icon ---
			Label{
				Text:      "\uE737",
				Alignment: AlignHCenterVCenter,
				Font: Font{
					Family:    "Segoe MDL2 Assets",
					PointSize: 28,
				},
				TextColor: walk.RGB(255, 107, 0),
			},

			VSpacer{Size: 6},

			// --- Title ---
			Label{
				Text:      "ANAGATA SENTINEL",
				Alignment: AlignHCenterVCenter,
				Font: Font{
					Family:    "Segoe UI Variable",
					PointSize: 14,
					Bold:      true,
				},
				TextColor: walk.RGB(255, 255, 255),
			},

			VSpacer{Size: 4},

			// --- Subtitle ---
			Label{
				Text:      "Security Monitoring System",
				Alignment: AlignHCenterVCenter,
				Font: Font{
					Family:    "Segoe UI Variable",
					PointSize: 8,
				},
				TextColor: walk.RGB(120, 120, 120),
			},

			VSpacer{Size: 16},

			// --- Status ---
			Label{
				AssignTo: &status,
				Text:     "Initializing...",
				Alignment: AlignHCenterVCenter,
				Font: Font{
					Family:    "Segoe UI Variable",
					PointSize: 8,
				},
				TextColor: walk.RGB(180, 180, 180),
			},

			VSpacer{Size: 16},
		},
	}).Create()

	if err != nil {
		s.ready <- err
		return
	}

	s.window = window
	s.status = status

	removeWindowFrame(window)
	centerWindow(window)

	window.Show()

	s.ready <- nil

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
			window.SetVisible(false)
		})
	})
}

func removeWindowFrame(window *walk.MainWindow) {
	hwnd := window.Handle()

	style := uint32(win.GetWindowLong(hwnd, win.GWL_STYLE))
	style &^= win.WS_CAPTION
	style &^= win.WS_THICKFRAME
	style &^= win.WS_MINIMIZEBOX
	style &^= win.WS_MAXIMIZEBOX
	style &^= win.WS_SYSMENU
	style |= win.WS_POPUP
	win.SetWindowLong(hwnd, win.GWL_STYLE, int32(style))

	exStyle := uint32(win.GetWindowLong(hwnd, win.GWL_EXSTYLE))
	exStyle |= win.WS_EX_TOOLWINDOW
	exStyle |= win.WS_EX_TOPMOST
	win.SetWindowLong(hwnd, win.GWL_EXSTYLE, int32(exStyle))

	win.SetWindowPos(hwnd, win.HWND_TOPMOST,
		0, 0, 0, 0,
		win.SWP_NOMOVE|win.SWP_NOSIZE|win.SWP_FRAMECHANGED|win.SWP_NOACTIVATE,
	)
}

func centerWindow(window *walk.MainWindow) {
	width := 360
	height := 180

	screenWidth := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	screenHeight := int(win.GetSystemMetrics(win.SM_CYSCREEN))

	x := (screenWidth - width) / 2
	y := (screenHeight - height) / 2

	_ = window.SetBoundsPixels(walk.Rectangle{
		X: x, Y: y, Width: width, Height: height,
	})
}
