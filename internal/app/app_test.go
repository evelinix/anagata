package app

import (
	"testing"
)

func TestNewApp(t *testing.T) {
	a := NewApp()

	if a == nil {
		t.Fatal("NewApp returned nil")
	}

	if a.backendReady == nil {
		t.Error("backendReady channel not initialized")
	}

	if a.domReady == nil {
		t.Error("domReady channel not initialized")
	}
}

func TestGreet(t *testing.T) {
	a := NewApp()

	result := a.Greet("World")
	expected := "Hello World, It's show time!"

	if result != expected {
		t.Errorf("Greet() = %q, want %q", result, expected)
	}
}

func TestVersion(t *testing.T) {
	a := NewApp()

	v := a.Version()
	if v == "" {
		t.Error("Version() returned empty string")
	}

	t.Logf("Version: %s", v)
}

func TestGetVersionInfo(t *testing.T) {
	a := NewApp()

	info := a.GetVersionInfo()
	if info == nil {
		t.Fatal("GetVersionInfo returned nil")
	}

	required := []string{"version", "commit", "buildTime", "goVersion"}
	for _, key := range required {
		if _, ok := info[key]; !ok {
			t.Errorf("missing key %q in version info", key)
		}
	}
}
