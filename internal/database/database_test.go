package database

import (
	"testing"
)

func TestExtractVersion(t *testing.T) {
	tests := []struct {
		path     string
		expected int
	}{
		{"000001_init.up.sql", 1},
		{"000010_add_users.up.sql", 10},
		{"000099_migration.up.sql", 99},
		{"invalid.sql", 0},
		{"", 0},
	}

	for _, tt := range tests {
		got := extractVersion(tt.path)
		if got != tt.expected {
			t.Errorf("extractVersion(%q) = %d, want %d", tt.path, got, tt.expected)
		}
	}
}
