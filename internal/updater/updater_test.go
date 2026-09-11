package updater

import (
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
