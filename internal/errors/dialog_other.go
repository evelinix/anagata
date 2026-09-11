//go:build !windows

package errors

// ShowDialog displays an error dialog on non-Windows platforms.
// Returns true if user clicked Retry.
func ShowDialog(err *AppError) bool {
	return false
}
