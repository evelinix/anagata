package logger

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRotator_RotateOnSize(t *testing.T) {
	dir := t.TempDir()
	activeLog := filepath.Join(dir, "app.log")

	r, err := NewRotator(activeLog, 0) // 0 = 100MB default
	if err != nil {
		t.Fatalf("NewRotator: %v", err)
	}
	defer r.Close()

	// Write data smaller than max size
	data := []byte("test log line\n")
	if _, err := r.Write(data); err != nil {
		t.Fatalf("Write: %v", err)
	}

	// Verify active log exists
	if _, err := os.Stat(activeLog); os.IsNotExist(err) {
		t.Fatal("active log should exist")
	}

	// Force rotation by setting max size to 1 byte
	r.maxSize = 1
	if _, err := r.Write(data); err != nil {
		t.Fatalf("Write (trigger rotation): %v", err)
	}

	// Verify rotated file exists in year/month_date.log format
	yearDir := filepath.Join(dir, "2026")
	if _, err := os.Stat(yearDir); os.IsNotExist(err) {
		t.Fatalf("year directory should exist at %s", yearDir)
	}

	entries, _ := os.ReadDir(yearDir)
	found := false
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".log") {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("rotated log file should exist in year directory")
	}

	// Verify active log exists with only the new data (13 bytes)
	info, _ := os.Stat(activeLog)
	if info == nil {
		t.Fatal("active log should exist after rotation")
	}
	if info.Size() > int64(len(data)) {
		t.Fatalf("active log should only contain new data, got %d bytes", info.Size())
	}
}

func TestRotator_CounterOnExistingRotated(t *testing.T) {
	dir := t.TempDir()
	activeLog := filepath.Join(dir, "app.log")

	r, err := NewRotator(activeLog, 0)
	if err != nil {
		t.Fatalf("NewRotator: %v", err)
	}
	defer r.Close()

	data := []byte("line\n")

	// First rotation
	r.maxSize = 1
	r.Write(data)

	// Second rotation (same day should use counter)
	r.written = r.maxSize + 1
	r.Write(data)

	// Should have app.log and at least 2 rotated files
	yearDir := filepath.Join(dir, "2026")
	entries, _ := os.ReadDir(yearDir)
	count := 0
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".log") {
			count++
		}
	}
	if count < 2 {
		t.Fatalf("expected at least 2 rotated files, got %d", count)
	}
}

func TestRotator_Close(t *testing.T) {
	dir := t.TempDir()
	activeLog := filepath.Join(dir, "app.log")

	r, err := NewRotator(activeLog, 100)
	if err != nil {
		t.Fatalf("NewRotator: %v", err)
	}

	if err := r.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}

	// Double close should not panic
	r.Close()
}

func TestRotator_InitialSize(t *testing.T) {
	dir := t.TempDir()
	activeLog := filepath.Join(dir, "app.log")

	// Pre-populate with some data
	os.WriteFile(activeLog, []byte("existing data\n"), 0644)

	r, err := NewRotator(activeLog, 100)
	if err != nil {
		t.Fatalf("NewRotator: %v", err)
	}
	defer r.Close()

	if r.written == 0 {
		t.Fatal("Rotator should detect existing file size")
	}
}

func TestRotator_WriteCloserInterface(t *testing.T) {
	dir := t.TempDir()
	r, err := NewRotator(filepath.Join(dir, "app.log"), 100)
	if err != nil {
		t.Fatalf("NewRotator: %v", err)
	}
	defer r.Close()

	// Verify it satisfies io.WriteCloser
	var _ interface {
		Write(p []byte) (n int, err error)
		Close() error
	} = r
}
