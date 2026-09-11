package splash

import (
	"os"
	"testing"
)

func TestNewNativeSplash(t *testing.T) {
	s := NewNativeSplash()
	if s == nil {
		t.Fatal("NewNativeSplash returned nil")
	}
	if s.platformSplash == nil {
		t.Fatal("platformSplash not initialized")
	}
}

func TestSplashSetStatusBeforeStart(t *testing.T) {
	s := NewNativeSplash()
	s.SetStatus("Should not panic")
}

func TestSplashCloseBeforeStart(t *testing.T) {
	s := NewNativeSplash()
	s.Close()
}

func requireDesktop(t *testing.T) {
	t.Helper()
	if os.Getenv("WALK_GUI_TEST") == "" {
		t.Skip("set WALK_GUI_TEST=1 to run GUI tests (requires interactive desktop)")
	}
}

func TestSplashStart(t *testing.T) {
	requireDesktop(t)
	s := NewNativeSplash()
	err := s.Start()
	if err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	s.Close()
}

func TestSplashSetStatus(t *testing.T) {
	requireDesktop(t)
	s := NewNativeSplash()
	if err := s.Start(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	s.SetStatus("Testing...")
	s.SetStatus("")
	s.SetStatus("Loading configuration...")
	s.Close()
}

func TestSplashCloseIdempotent(t *testing.T) {
	requireDesktop(t)
	s := NewNativeSplash()
	if err := s.Start(); err != nil {
		t.Fatalf("Start failed: %v", err)
	}
	s.Close()
	s.Close()
	s.Close()
}
