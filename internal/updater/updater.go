package updater

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
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
	} `json:"assets"`
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
