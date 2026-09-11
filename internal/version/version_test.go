package version

import "testing"

func TestInfo(t *testing.T) {
	Version = "1.2.3"
	GitCommit = "abc123456789"
	BuildTime = "2026-09-11"
	GoVersion = "go1.25.0"

	info := Info()

	if info == "" {
		t.Error("Info() returned empty string")
	}

	t.Logf("Version info: %s", info)
}

func TestMin(t *testing.T) {
	if min(1, 2) != 1 {
		t.Error("min(1, 2) should be 1")
	}
	if min(5, 3) != 3 {
		t.Error("min(5, 3) should be 3")
	}
}
