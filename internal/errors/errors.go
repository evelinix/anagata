package errors

import "fmt"

type ErrorLevel int

const (
	LevelInfo ErrorLevel = iota
	LevelWarning
	LevelError
	LevelFatal
)

func (l ErrorLevel) String() string {
	switch l {
	case LevelInfo:
		return "INFO"
	case LevelWarning:
		return "WARNING"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

type AppError struct {
	Level     ErrorLevel
	Stage     string
	Message   string
	Err       error
	Retriable bool
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func New(level ErrorLevel, stage, message string, err error, retriable bool) *AppError {
	return &AppError{
		Level:     level,
		Stage:     stage,
		Message:   message,
		Err:       err,
		Retriable: retriable,
	}
}

func Wrap(stage string, err error) *AppError {
	return &AppError{
		Level:     LevelError,
		Stage:     stage,
		Message:   fmt.Sprintf("Failed to %s", stage),
		Err:       err,
		Retriable: true,
	}
}
