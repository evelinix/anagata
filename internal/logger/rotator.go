package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Rotator is a writer that rotates log files based on size.
// Active log: {base}.log
// Rotated logs: {base_dir}/{year}/{month}_{date}.log
type Rotator struct {
	mu       sync.Mutex
	file     *os.File
	basePath string
	maxSize  int64
	written  int64
}

// NewRotator creates a new Rotator.
// basePath is the active log file path (e.g. logs/app.log).
// maxSizeMB is the maximum file size in MB before rotation.
func NewRotator(basePath string, maxSizeMB int) (*Rotator, error) {
	if maxSizeMB <= 0 {
		maxSizeMB = 100
	}

	if err := os.MkdirAll(filepath.Dir(basePath), 0755); err != nil {
		return nil, fmt.Errorf("create log dir: %w", err)
	}

	f, err := os.OpenFile(basePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("open log file: %w", err)
	}

	info, _ := f.Stat()
	var written int64
	if info != nil {
		written = info.Size()
	}

	return &Rotator{
		file:     f,
		basePath: basePath,
		maxSize:  int64(maxSizeMB) * 1024 * 1024,
		written:  written,
	}, nil
}

func (r *Rotator) Write(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.written+int64(len(p)) > r.maxSize {
		if err := r.rotate(); err != nil {
			return 0, err
		}
	}

	n, err = r.file.Write(p)
	r.written += int64(n)
	return
}

func (r *Rotator) rotate() error {
	now := time.Now()

	// Close current file
	r.file.Close()

	// Build rotated path: logs/2026/09_11.log
	dir := filepath.Join(filepath.Dir(r.basePath), fmt.Sprintf("%d", now.Year()))
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create rotation dir: %w", err)
	}

	rotated := filepath.Join(dir, fmt.Sprintf("%02d_%02d.log", now.Month(), now.Day()))

	// If rotated file already exists, append a counter
	if _, err := os.Stat(rotated); err == nil {
		for i := 1; i < 100; i++ {
			candidate := filepath.Join(dir, fmt.Sprintf("%02d_%02d_%d.log", now.Month(), now.Day(), i))
			if _, err := os.Stat(candidate); os.IsNotExist(err) {
				rotated = candidate
				break
			}
		}
	}

	// Rename current file to rotated path
	if err := os.Rename(r.basePath, rotated); err != nil {
		return fmt.Errorf("rotate log: %w", err)
	}

	// Open new active log file
	f, err := os.OpenFile(r.basePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("reopen log file: %w", err)
	}

	r.file = f
	r.written = 0
	return nil
}

func (r *Rotator) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.file != nil {
		return r.file.Close()
	}
	return nil
}

// Compile-time check that Rotator implements io.WriteCloser.
var _ io.WriteCloser = (*Rotator)(nil)
