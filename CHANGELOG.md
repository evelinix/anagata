# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Config system with YAML support and env variable override
- Structured logging with slog (JSON file + text console output)
- Error handling UI with Walk dialog (Retry/Close)
- AppError type with level, stage, message, retriable flag
- Build versioning with ldflags (version, commit, build time, go version)
- Version() and GetVersionInfo() methods for frontend access
- golangci-lint configuration
- Database migration system with versioned SQL files
- `make version` command to show build info
- Unit tests for config, logger, errors, and version (13 tests total)

### Changed
- Config file location changed to executable directory (portable deployment)
- Replaced `log.Println`/`log.Printf` with structured `slog.Info`/`slog.Error`
- Logger integrates with config system (level, file path from `config.yaml`)
- Boot stages now log with structured fields (stage name, error)
- Makefile `build` target now uses `wails build` with ldflags

### Removed
- Removed `time.Sleep` from boot stages (config, security, services, finalize)
- Removed `%APPDATA%` dependency for config location

---

## [0.1.0] - 2026-09-10

### Added
- Initial enterprise template structure
- Go backend with `internal/` package organization
- Svelte 5 frontend with Tailwind CSS v4
- Native Win32 splash screen (walk)
- Encrypted SQLite database (Adiantum cipher)
- pnpm + Vite 7 build pipeline
- Makefile with build, dev, clean, lint, check, rename targets
- Git template support (`git clone --reference`)
- Project rename command (`make rename NEW=Name`)

### Architecture
- `main.go` — Entry point with `//go:embed`
- `internal/app/` — App lifecycle and boot stages
- `internal/database/` — Encrypted SQLite initialization
- `internal/splash/` — Platform-agnostic native splash screen
- `frontend/` — Svelte 5 + Tailwind CSS v4 + Oxanium font

---

[Unreleased]: https://github.com/youruser/AnagataSentinel/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/youruser/AnagataSentinel/releases/tag/v0.1.0
