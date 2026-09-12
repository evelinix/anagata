package updater

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"AnagataSentinel/internal/version"
)

type Release struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Body    string `json:"body"`
	HTMLURL string `json:"html_url"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
		Size               int64  `json:"size"`
	} `json:"assets"`
}

type DownloadProgress struct {
	BytesDownloaded int64  `json:"bytes_downloaded"`
	TotalBytes      int64  `json:"total_bytes"`
	Percent         int    `json:"percent"`
	Status          string `json:"status"`
	FilePath        string `json:"file_path,omitempty"`
	Error           string `json:"error,omitempty"`
}

type Checker struct {
	repoOwner string
	repoName  string
	client    *http.Client
}

func NewChecker(repoOwner, repoName string) *Checker {
	return &Checker{
		repoOwner: repoOwner,
		repoName:  repoName,
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Checker) CheckForUpdate() (*Release, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", c.repoOwner, c.repoName)

	resp, err := c.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("check github releases: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github api returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	var release Release
	if err := json.Unmarshal(body, &release); err != nil {
		return nil, fmt.Errorf("parse release: %w", err)
	}

	return &release, nil
}

func (c *Checker) IsUpdateAvailable() (bool, *Release, error) {
	release, err := c.CheckForUpdate()
	if err != nil {
		return false, nil, err
	}

	if release.TagName != version.Version {
		slog.Info("update available",
			"current", version.Version,
			"latest", release.TagName,
		)
		return true, release, nil
	}

	return false, nil, nil
}

func (c *Checker) StartPeriodicCheck(interval time.Duration, onUpdate func(*Release)) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			available, release, err := c.IsUpdateAvailable()
			if err != nil {
				slog.Warn("update check failed", "error", err)
				continue
			}

			if available && onUpdate != nil {
				onUpdate(release)
			}
		}
	}()
}

func (c *Checker) findAssetForPlatform(release *Release) (string, string, int64) {
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	for _, asset := range release.Assets {
		name := strings.ToLower(asset.Name)

		osMatch := false
		archMatch := false

		switch goos {
		case "windows":
			osMatch = strings.Contains(name, "windows") || strings.Contains(name, ".exe")
		case "linux":
			osMatch = strings.Contains(name, "linux")
		case "darwin":
			osMatch = strings.Contains(name, "darwin") || strings.Contains(name, "macos")
		}

		switch goarch {
		case "amd64":
			archMatch = strings.Contains(name, "amd64") || strings.Contains(name, "x86_64")
		case "arm64":
			archMatch = strings.Contains(name, "arm64") || strings.Contains(name, "aarch64")
		}

		if osMatch && archMatch {
			return asset.BrowserDownloadURL, asset.Name, asset.Size
		}
	}

	return "", "", 0
}

func (c *Checker) DownloadUpdate(release *Release, destDir string, onProgress func(DownloadProgress)) (*string, error) {
	downloadURL, fileName, totalSize := c.findAssetForPlatform(release)
	if downloadURL == "" {
		return nil, fmt.Errorf("no compatible asset found for %s/%s", runtime.GOOS, runtime.GOARCH)
	}

	if err := os.MkdirAll(destDir, 0o755); err != nil {
		return nil, fmt.Errorf("create download directory: %w", err)
	}

	destPath := filepath.Join(destDir, fileName)

	onProgress(DownloadProgress{
		Status:     "downloading",
		TotalBytes: totalSize,
	})

	resp, err := c.client.Get(downloadURL)
	if err != nil {
		return nil, fmt.Errorf("download update: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download returned status %d", resp.StatusCode)
	}

	out, err := os.Create(destPath)
	if err != nil {
		return nil, fmt.Errorf("create file: %w", err)
	}
	defer out.Close()

	var downloaded int64
	buf := make([]byte, 32*1024)
	for {
		n, readErr := resp.Body.Read(buf)
		if n > 0 {
			if _, writeErr := out.Write(buf[:n]); writeErr != nil {
				return nil, fmt.Errorf("write file: %w", writeErr)
			}
			downloaded += int64(n)

			percent := 0
			if totalSize > 0 {
				percent = int((downloaded * 100) / totalSize)
			}

			onProgress(DownloadProgress{
				BytesDownloaded: downloaded,
				TotalBytes:      totalSize,
				Percent:         percent,
				Status:          "downloading",
			})
		}
		if readErr != nil {
			if readErr == io.EOF {
				break
			}
			return nil, fmt.Errorf("read response: %w", readErr)
		}
	}

	onProgress(DownloadProgress{
		BytesDownloaded: downloaded,
		TotalBytes:      totalSize,
		Percent:         100,
		Status:          "completed",
		FilePath:        destPath,
	})

	slog.Info("update downloaded", "path", destPath, "size", downloaded)
	return &destPath, nil
}

func (c *Checker) ApplyUpdate(newBinaryPath string) error {
	currentExec, err := os.Executable()
	if err != nil {
		return fmt.Errorf("get current executable: %w", err)
	}
	currentExec, err = filepath.EvalSymlinks(currentExec)
	if err != nil {
		return fmt.Errorf("resolve current executable: %w", err)
	}

	if _, err := os.Stat(newBinaryPath); err != nil {
		return fmt.Errorf("new binary not found: %w", err)
	}

	if runtime.GOOS == "windows" {
		return c.applyUpdateWindows(currentExec, newBinaryPath)
	}
	return c.applyUpdateUnix(currentExec, newBinaryPath)
}

func (c *Checker) applyUpdateWindows(currentExec, newBinaryPath string) error {
	scriptPath := filepath.Join(filepath.Dir(newBinaryPath), "update.bat")

	script := fmt.Sprintf(`@echo off
timeout /t 2 /nobreak >nul 2>&1
copy /Y "%s" "%s" >nul 2>&1
if %%errorlevel%% equ 0 (
    del /f /q "%s" >nul 2>&1
    start "" "%s"
)
del /f /q "%s" >nul 2>&1
`, newBinaryPath, currentExec, newBinaryPath, currentExec, scriptPath)

	if err := os.WriteFile(scriptPath, []byte(script), 0o755); err != nil {
		return fmt.Errorf("create update script: %w", err)
	}

	slog.Info("applying update", "from", newBinaryPath, "to", currentExec)

	cmd := exec.Command("cmd", "/C", scriptPath)
	cmd.SysProcAttr = nil
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start update script: %w", err)
	}

	return nil
}

func (c *Checker) applyUpdateUnix(currentExec, newBinaryPath string) error {
	scriptPath := filepath.Join(filepath.Dir(newBinaryPath), "update.sh")

	script := fmt.Sprintf(`#!/bin/sh
sleep 2
cp -f "%s" "%s"
chmod +x "%s"
"%s" &
rm -f "%s"
rm -f "%s"
`, newBinaryPath, currentExec, currentExec, currentExec, scriptPath, newBinaryPath)

	if err := os.WriteFile(scriptPath, []byte(script), 0o755); err != nil {
		return fmt.Errorf("create update script: %w", err)
	}

	slog.Info("applying update", "from", newBinaryPath, "to", currentExec)

	cmd := exec.Command("/bin/sh", scriptPath)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start update script: %w", err)
	}

	return nil
}

func (c *Checker) GetDownloadedFilePath(release *Release) (string, bool) {
	downloadURL, fileName, _ := c.findAssetForPlatform(release)
	if downloadURL == "" {
		return "", false
	}

	candidate := filepath.Join(os.TempDir(), "AnagataSentinel-updates", fileName)
	if _, err := os.Stat(candidate); err == nil {
		return candidate, true
	}
	return "", false
}
