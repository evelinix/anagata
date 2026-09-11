# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Dark/Light mode with system preference detection and localStorage persistence
- Internationalization (i18n) with svelte-i18n (EN + ID locales)
- Language toggle button in UI
- Auto-updater with GitHub release check (periodic, 1 hour interval)
- CheckForUpdate() method exposed to frontend
- internal/updater package with GitHub API integration
- Vitest frontend testing with @testing-library/svelte
- Go unit tests for all internal packages (21 tests total)
- ESLint flat config with Svelte 5 + TypeScript
- Prettier configuration for consistent formatting
- Pre-commit hooks with husky + lint-staged

### Changed
- Config system with YAML support and env variable override
- Structured logging with slog (JSON file + text console output)
- Error handling UI with Walk dialog (Retry/Close)
- Build versioning with ldflags (version, commit, build time, go version)
- Database migration system with versioned SQL files
- Config file location changed to executable directory (portable deployment)

### Removed
- Removed `%APPDATA%` dependency for config location
- Removed root-level logo files (cleaned up)

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
