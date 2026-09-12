package reporter

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime/debug"
	"sync"
	"time"
)

type Level int

const (
	LevelError Level = iota
	LevelWarning
	LevelInfo
)

type Report struct {
	Level     Level
	Message   string
	Err       error
	Extra     map[string]interface{}
	Timestamp time.Time
}

type Reporter struct {
	dsn        string
	filePath   string
	mu         sync.Mutex
	buffer     []*Report
	flushTimer *time.Timer
}

func New(dsn string) *Reporter {
	r := &Reporter{
		dsn:      dsn,
		filePath: getReportPath(),
	}

	r.setupPanicHandler()

	if dsn != "" {
		slog.Info("error reporter initialized", "provider", "sentry")
	} else {
		slog.Info("error reporter initialized", "provider", "file", "path", r.filePath)
	}

	return r
}

func (r *Reporter) ReportErrorf(format string, args ...interface{}) {
	r.report(LevelError, fmt.Sprintf(format, args...), nil)
}

func (r *Reporter) ReportError(err error, message string, extra map[string]interface{}) {
	r.report(LevelError, message, err, extra)
}

func (r *Reporter) ReportWarning(message string, extra map[string]interface{}) {
	r.report(LevelWarning, message, nil, extra)
}

func (r *Reporter) ReportInfo(message string, extra map[string]interface{}) {
	r.report(LevelInfo, message, nil, extra)
}

func (r *Reporter) report(level Level, message string, err error, extra ...map[string]interface{}) {
	var extraMap map[string]interface{}
	if len(extra) > 0 {
		extraMap = extra[0]
	}

	report := &Report{
		Level:     level,
		Message:   message,
		Err:       err,
		Extra:     extraMap,
		Timestamp: time.Now(),
	}

	r.mu.Lock()
	r.buffer = append(r.buffer, report)
	r.mu.Unlock()

	slog.Error("error reported",
		"level", level.String(),
		"message", message,
		"error", err,
	)

	if r.dsn != "" {
		r.sendToSentry(report)
	} else {
		r.writeToFile(report)
	}
}

func (r *Reporter) setupPanicHandler() {
	go func() {
		defer func() {
			if rec := recover(); rec != nil {
				stack := debug.Stack()
				slog.Error("panic recovered",
					"error", rec,
					"stack", string(stack),
				)

				r.report(LevelError, fmt.Sprintf("panic: %v", rec), nil,
					map[string]interface{}{
						"stack": string(stack),
					},
				)
			}
		}()
	}()
}

func (r *Reporter) sendToSentry(report *Report) {
	slog.Debug("would send to sentry", "message", report.Message)
}

func (r *Reporter) writeToFile(report *Report) {
	r.mu.Lock()
	defer r.mu.Unlock()

	f, err := os.OpenFile(r.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		slog.Error("failed to open report file", "error", err)
		return
	}
	defer f.Close()

	entry := fmt.Sprintf("[%s] [%s] %s",
		report.Timestamp.Format(time.RFC3339),
		report.Level.String(),
		report.Message,
	)

	if report.Err != nil {
		entry += fmt.Sprintf(" error=%s", report.Err.Error())
	}

	entry += "\n"

	if _, err := f.WriteString(entry); err != nil {
		slog.Error("failed to write report", "error", err)
	}
}

func (r *Reporter) Flush() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.buffer) == 0 {
		return
	}

	slog.Info("flushing error reports", "count", len(r.buffer))
	r.buffer = nil
}

func (l Level) String() string {
	switch l {
	case LevelError:
		return "ERROR"
	case LevelWarning:
		return "WARNING"
	case LevelInfo:
		return "INFO"
	default:
		return "UNKNOWN"
	}
}

func getReportPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}

	dir := filepath.Join(home, ".anagata", "reports")
	os.MkdirAll(dir, 0755)

	return filepath.Join(dir, fmt.Sprintf("reports-%s.log", time.Now().Format("2006-01-02")))
}
