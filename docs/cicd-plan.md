# CI/CD Pipeline — Implementation Plan

**Date:** 2026-09-12
**Feature:** Jenkins CI/CD Pipeline

---

## Overview

Membuat Jenkins pipeline untuk AnagataSentinel:
1. **CI Pipeline** — Lint, test, build setiap commit
2. **Windows Agent** — Karena Go dependencies hanya support Windows

---

## Kesiapan Saat Ini

| Kategori | Status | Keterangan |
|----------|--------|------------|
| Makefile | ✅ READY | `make build`, `make test`, `make lint`, `make check` |
| Go tests | ✅ READY | 12 test files across all packages |
| Frontend tests | ✅ READY | 102 tests dengan Vitest |
| Pre-commit hooks | ✅ READY | Husky + lint-staged |
| Linter config | ✅ READY | `.golangci.yml` dengan 20 linters |
| Type checking | ✅ READY | `svelte-check` |
| Version injection | ✅ READY | ldflags (version, commit, time) |
| **CI pipeline** | ✅ DONE | `Jenkinsfile` created |

---

## File yang Dibuat

| # | File | Fungsi |
|---|------|--------|
| 1 | `Jenkinsfile` | Jenkins pipeline definition |

---

## Pipeline Structure

```
Jenkins Pipeline
        ↓
┌───────────────────────────────────────────┐
│  Stage 1: Checkout                        │
│  ├── Clone repository                     │
│  └── Set environment variables            │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Stage 2: Setup                           │
│  ├── Verify Go, Node, pnpm               │
│  ├── Install Wails CLI                    │
│  └── pnpm install --frozen-lockfile       │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Stage 3: Lint (parallel)                 │
│  ├── Go: golangci-lint                    │
│  ├── Frontend: ESLint                     │
│  └── Frontend: Prettier check             │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Stage 4: Test (parallel)                 │
│  ├── Go: go test ./internal/...           │
│  ├── Frontend: Vitest                     │
│  └── Frontend: svelte-check               │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Stage 5: Build                           │
│  └── wails build + version injection      │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Stage 6: Package                         │
│  └── Archive AnagataSentinel.exe          │
└───────────────────────────────────────────┘
```

---

## Jenkins Setup Requirements

### Prerequisites

| Tool | Version | Purpose |
|------|---------|---------|
| **Jenkins** | Latest LTS | CI server |
| **Go** | 1.25.0+ | Build backend |
| **Node.js** | 20 LTS | Build frontend |
| **pnpm** | 9.x | Package manager frontend |
| **Wails CLI** | v2.15.0 | Build desktop app |
| **Git** | Latest | Source control |
| **golangci-lint** | Latest | Go linter |

### Jenkins Plugins

| Plugin | Purpose |
|--------|---------|
| **Git** | Clone repository |
| **Pipeline** | Jenkinsfile support |
| **Pipeline: Stage View** | Visualisasi stage |
| **NodeJS** | Node.js management |

### Jenkins Tools Configuration

**Jenkins → Manage Jenkins → Tools**

| Tool | Configuration |
|------|---------------|
| **Go** | Install Go 1.25.0, set GOROOT |
| **NodeJS** | Install Node 20, set npm/pnpm |
| **Git** | Default git path |

---

## Pipeline Flow Diagram

```
Developer push ke main
        ↓
Jenkins trigger pipeline
        ↓
┌───────────────────────────────────────────┐
│  Stage: Checkout                          │
│  ├── Clone repo                           │
│  └── Set version from git tags            │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Stage: Setup                             │
│  ├── Verify Go, Node, pnpm               │
│  ├── Install Wails CLI                    │
│  └── pnpm install --frozen-lockfile       │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Stage: Lint (parallel)                   │
│  ├── golangci-lint (20 linters)           │
│  ├── ESLint (frontend)                    │
│  └── Prettier check                       │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Stage: Test (parallel)                   │
│  ├── go test ./internal/...               │
│  ├── Vitest (102 tests)                   │
│  └── svelte-check (type safety)           │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Stage: Build                             │
│  └── wails build + version injection      │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Stage: Package                           │
│  └── Archive AnagataSentinel.exe          │
└───────────────────────────────────────────┘
        ↓
    Build artifact tersimpan di Jenkins
```

---

## Checklist Implementasi

| # | Task | Status |
|---|------|--------|
| 1 | Buat `Jenkinsfile` | ✅ Done |
| 2 | Install Jenkins di Windows server | ⬜ Pending |
| 3 | Install plugins (Git, Pipeline, NodeJS) | ⬜ Pending |
| 4 | Configure Go, Node, pnpm tools | ⬜ Pending |
| 5 | Setup Windows agent | ⬜ Pending |
| 6 | Test pipeline | ⬜ Pending |

---

## Notes

- **Windows Agent Wajib** — Karena Go dependencies (`lxn/walk`, `lxn/win`) hanya support Windows
- **WebView2 Runtime** — Pastikan terinstall di agent (biasanya sudah ada di Windows 10/11)
- **Parallel Stages** — Lint dan Test jalan parallel untuk speed
- **Version Injection** — Menggunakan git tags via `git describe --tags`
