# AGENTS.md

## What This Is

Wails v2 desktop app — Go backend + Svelte 5 frontend embedded in a native WebView. The frontend is bundled by Vite and baked into the Go binary via `//go:embed`.

## Key Commands

| Command | What it does |
|---|---|
| `wails dev` | Full dev mode: Go backend + Vite frontend with hot-reload |
| `wails build` | Production build: compiles Go binary with embedded frontend into `build/bin/` |
| `make install` | Install frontend deps (pnpm) |
| `make build` | Production build via Makefile |
| `make dev` | Development mode |
| `make clean` | Remove build artifacts and node_modules |
| `make lint` | Format and vet Go code |
| `make check` | Lint + build + svelte-check |
| `make rename NEW=Name` | Rename project across all files |
| `make version` | Show current version info |
| `cd frontend && pnpm lint` | ESLint frontend |
| `cd frontend && pnpm format` | Prettier format frontend |

## Architecture

- **Go entry**: `main.go` → `wails.Run()` → creates window, binds `App` struct methods
- **Config**: `internal/config/config.go` → YAML config at `%APPDATA%/AnagataSentinel/config.yaml`
- **Logger**: `internal/logger/logger.go` → structured slog (JSON file + text console)
- **Go app logic**: `internal/app/app.go` → `App` struct with lifecycle hooks + exported methods (e.g. `Greet`)
- **Go database**: `internal/database/database.go` → encrypted SQLite via `gosqlite.org/vfs/crypto`
- **Error handling**: `internal/errors/` → AppError type + Walk dialog for boot failures
- **Splash screen**: `internal/splash/splash.go` (interface) + `splash_windows.go` (Win32 native) + `splash_other.go` (no-op)
- **Frontend entry**: `frontend/index.html` → `frontend/src/main.ts` → mounts `App.svelte`
- **Wails bindings**: `frontend/wailsjs/go/app/App.js` — **auto-generated**, never edit manually
- **Frontend assets**: embedded from `frontend/dist/` into Go binary at compile time
- **Window**: 1280×800, `StartHidden: true`, background `rgba(248,250,252,1)`

## Config

Config file: `config.yaml` (sama directory dengan executable)

| Env Variable | Config Path | Default |
|---|---|---|
| `ANAGATA_DEBUG` | `app.debug` | `false` |
| `ANAGATA_DB_PASSWORD` | `database.password` | `0123456789` |
| `ANAGATA_LOG_LEVEL` | `logging.level` | `info` |

## Conventions

- **Svelte 5 runes**: uses `$state()`, not the older `let` reactive syntax
- **Tailwind CSS v4**: configured via `@theme` in `frontend/src/style.css`, no `tailwind.config.js`. Plugin is `@tailwindcss/vite` (not PostCSS)
- **Custom font**: Oxanium loaded as local woff2 in `style.css`
- **All npm packages are devDependencies** — nothing runs at runtime; everything is bundled
- **Go method changes** require `wails dev` or `wails build` to regenerate JS bindings
- **App icons**: source at `build/appicon.png` and `build/windows/icon.ico`, frontend logo at `frontend/src/assets/images/logo-universal.png`
- **Logging**: use `slog.Info`/`slog.Error` with structured fields, never `log.Println`
- **Config**: always use `config.Get()` after initial load, never read config file directly

## Gotchas

- `frontend/dist/` is in `.gitignore` — it's generated, not committed
- `build/bin/AnagataSentinel.exe` also shouldn't be tracked
- `wails.json` `frontend:dev:serverUrl: "auto"` — Wails auto-detects the Vite dev server URL
- Window background in `main.go` and `style.css` should match to avoid flash-of-white on startup
- Go tests exist for config and logger: `go test ./internal/config/ ./internal/logger/`
