package version

import "fmt"

// Set via ldflags at build time.
var (
	Version   = "0.1.0"
	GitCommit = "unknown"
	BuildTime = "unknown"
	GoVersion = "unknown"
)

func Info() string {
	return fmt.Sprintf("%s (commit: %s, built: %s, go: %s)",
		Version, GitCommit[:min(8, len(GitCommit))], BuildTime, GoVersion)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
