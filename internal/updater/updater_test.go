package updater

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestNewChecker(t *testing.T) {
	c := NewChecker("user", "repo")

	if c.repoOwner != "user" {
		t.Errorf("expected repoOwner 'user', got '%s'", c.repoOwner)
	}

	if c.repoName != "repo" {
		t.Errorf("expected repoName 'repo', got '%s'", c.repoName)
	}

	if c.client == nil {
		t.Error("http client not initialized")
	}
}

func TestFindAssetForPlatform(t *testing.T) {
	c := NewChecker("user", "repo")

	release := &Release{
		TagName: "v1.0.0",
		Assets: []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
			Size               int64  `json:"size"`
		}{
			{Name: "AnagataSentinel-windows-amd64.exe", BrowserDownloadURL: "https://example.com/win.exe", Size: 1000},
			{Name: "AnagataSentinel-linux-amd64", BrowserDownloadURL: "https://example.com/linux", Size: 2000},
			{Name: "AnagataSentinel-darwin-arm64", BrowserDownloadURL: "https://example.com/mac", Size: 3000},
		},
	}

	url, name, size := c.findAssetForPlatform(release)
	if url == "" {
		t.Error("expected to find an asset for current platform")
	}
	if name == "" {
		t.Error("expected asset name to be non-empty")
	}
	if size == 0 {
		t.Error("expected asset size to be non-zero")
	}
}

func TestFindAssetForPlatformNotFound(t *testing.T) {
	c := NewChecker("user", "repo")

	release := &Release{
		TagName: "v1.0.0",
		Assets: []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
			Size               int64  `json:"size"`
		}{
			{Name: "other-binary.zip", BrowserDownloadURL: "https://example.com/other.zip", Size: 500},
		},
	}

	url, _, _ := c.findAssetForPlatform(release)
	if url != "" {
		t.Error("expected empty URL for non-matching platform")
	}
}

func TestDownloadProgress(t *testing.T) {
	p := DownloadProgress{
		BytesDownloaded: 500,
		TotalBytes:      1000,
		Percent:         50,
		Status:          "downloading",
	}

	if p.Percent != 50 {
		t.Errorf("expected percent 50, got %d", p.Percent)
	}
	if p.Status != "downloading" {
		t.Errorf("expected status 'downloading', got '%s'", p.Status)
	}
}

func TestApplyUpdate_MissingBinary(t *testing.T) {
	c := NewChecker("user", "repo")

	err := c.ApplyUpdate("/nonexistent/path/binary")
	if err == nil {
		t.Error("expected error for missing binary, got nil")
	}
}

func TestGetDownloadedFilePath_NotFound(t *testing.T) {
	c := NewChecker("user", "repo")

	release := &Release{
		TagName: "v1.0.0",
		Assets: []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
			Size               int64  `json:"size"`
		}{
			{Name: "AnagataSentinel-windows-amd64.exe", BrowserDownloadURL: "https://example.com/win.exe", Size: 1000},
		},
	}

	_, found := c.GetDownloadedFilePath(release)
	if found {
		t.Error("expected not found for non-existent download")
	}
}

func TestGetDownloadedFilePath_Found(t *testing.T) {
	c := NewChecker("user", "repo")

	tmpDir := filepath.Join(os.TempDir(), "AnagataSentinel-updates")
	os.MkdirAll(tmpDir, 0o755)
	defer os.RemoveAll(tmpDir)

	var fileName string
	if runtime.GOOS == "windows" {
		fileName = "AnagataSentinel-windows-amd64.exe"
	} else {
		fileName = "AnagataSentinel-linux-amd64"
	}

	fakeFile := filepath.Join(tmpDir, fileName)
	os.WriteFile(fakeFile, []byte("fake"), 0o755)
	defer os.Remove(fakeFile)

	release := &Release{
		TagName: "v1.0.0",
		Assets: []struct {
			Name               string `json:"name"`
			BrowserDownloadURL string `json:"browser_download_url"`
			Size               int64  `json:"size"`
		}{
			{Name: fileName, BrowserDownloadURL: "https://example.com/" + fileName, Size: 1000},
		},
	}

	path, found := c.GetDownloadedFilePath(release)
	if !found {
		t.Error("expected to find downloaded file")
	}
	if path != fakeFile {
		t.Errorf("expected path %s, got %s", fakeFile, path)
	}
}
