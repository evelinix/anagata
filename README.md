# AnagataSentinel

> Security monitoring desktop application — Go backend + Svelte 5 frontend

## Author

**Eve Lin** — [GitHub](https://github.com/evelinix)

## Prerequisites

- [Go 1.25+](https://go.dev/dl/)
- [Wails v2](https://wails.io/docs/gettingstarted/installation)
- [Node.js 20+](https://nodejs.org/)
- [pnpm](https://pnpm.io/)

## Quick Start

```bash
# Install dependencies
make install

# Development mode
make dev

# Production build
make build
```

## Project Structure

```
├── main.go                          # Entry point (//go:embed)
├── internal/
│   ├── app/                         # App logic & lifecycle
│   ├── config/                      # YAML config system
│   ├── database/                    # Encrypted SQLite + migrations
│   ├── errors/                      # Error handling + Walk dialog
│   ├── logger/                      # Structured logging + rotation
│   ├── splash/                      # Native splash screen
│   ├── updater/                     # Auto-updater (GitHub releases)
│   └── version/                     # Build version info
├── frontend/
│   ├── src/
│   │   ├── lib/components/          # ANTS design system components
│   │   ├── lib/i18n/                # Internationalization + formatting
│   │   ├── lib/stores/              # Theme store
│   │   └── locales/                 # EN + ID translations
│   └── wailsjs/                     # Auto-generated bindings
├── build/                           # Build assets (icons, manifest)
├── go.mod
├── wails.json
└── Makefile
```

## Tech Stack

| Layer | Technology |
|---|---|
| Backend | Go 1.25, Wails v2.15, Encrypted SQLite (Adiantum) |
| Frontend | Svelte 5 (runes), Tailwind CSS v4, Vite 7 |
| Design | ANTS Design System (based on WinUI 3 / Fluent) |
| Icons | Custom SVG icon components |
| Package | pnpm |
| Testing | Go testing, Vitest + Testing Library |
| Linting | golangci-lint, ESLint, Prettier |

## Commands

| Command | Description |
|---|---|
| `make help` | Show all available commands |
| `make dev` | Development mode with hot-reload |
| `make build` | Production build |
| `make test` | Run all tests (Go + frontend + svelte-check) |
| `make lint` | Format and vet Go code |
| `make check` | Lint + build + type-check |
| `make clean` | Remove build artifacts |

## License

[MIT](LICENSE) — Eve Lin
