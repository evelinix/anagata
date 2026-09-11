# AnagataSentinel

> Wails v2 desktop app — Go backend + Svelte 5 frontend

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
│   ├── database/                    # Encrypted SQLite
│   └── splash/                      # Native splash screen
├── frontend/
│   ├── src/                         # Svelte 5 + Tailwind CSS v4
│   └── wailsjs/                     # Auto-generated bindings
├── build/                           # Build assets (icons, manifest)
├── go.mod
├── wails.json
└── Makefile
```

## Tech Stack

- **Backend:** Go 1.25, Wails v2.15, Encrypted SQLite (Adiantum)
- **Frontend:** Svelte 5 (runes), Tailwind CSS v4, Vite 7, Oxanium font
- **Package:** pnpm

## Commands

| Command | Description |
|---|---|
| `make help` | Show all available commands |
| `make dev` | Development mode with hot-reload |
| `make build` | Production build |
| `make clean` | Remove build artifacts |
| `make lint` | Format and vet Go code |
| `make check` | Lint + build + type-check |
