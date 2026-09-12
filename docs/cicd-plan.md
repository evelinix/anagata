# CI/CD Pipeline — Implementation Plan

**Date:** 2026-09-12
**Feature:** GitHub Actions CI/CD Pipeline

---

## Overview

Membuat 2 GitHub Actions workflows:
1. **CI Pipeline** — Lint, test, build setiap push/PR
2. **Release Pipeline** — Build & upload binaries saat tag release

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
| **CI pipeline** | ❌ MISSING | Belum ada `.github/workflows/` |
| **Release automation** | ❌ MISSING | Belum ada release workflow |

---

## File yang Perlu Dibuat

| # | File | Fungsi |
|---|------|--------|
| 1 | `.github/workflows/ci.yml` | CI pipeline (lint + test + build) |
| 2 | `.github/workflows/release.yml` | Release pipeline (build + upload) |

---

## Flow Diagram

### CI Flow (Setiap Push/PR)

```
Push ke main/develop atau Buka PR
        ↓
┌───────────────────────────────────────────┐
│  Job 1: LINT (parallel)                  │
│  ├── golangci-lint (20 linters)          │
│  ├── ESLint (frontend)                   │
│  └── Prettier check                      │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Job 2: TEST (parallel)                  │
│  ├── go test ./internal/...              │
│  ├── Vitest (102 tests)                  │
│  └── svelte-check (type safety)          │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Job 3: BUILD (needs lint+test pass)     │
│  ├── wails build                         │
│  └── Upload artifact (7 days retention)  │
└───────────────────────────────────────────┘
        ↓
    All Pass? → Bisa merge PR
```

### Release Flow (Saat Tag `v*`)

```
git tag v0.2.0 && git push --tags
        ↓
┌───────────────────────────────────────────┐
│  Build Windows AMD64                      │
│  └── wails build -platform windows/amd64 │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Build Windows ARM64                      │
│  └── wails build -platform windows/arm64 │
└───────────────────────────────────────────┘
        ↓
┌───────────────────────────────────────────┐
│  Create GitHub Release                    │
│  ├── Upload: AnagataSentinel.exe          │
│  └── Auto-generate release notes          │
└───────────────────────────────────────────┘
        ↓
    User bisa download dari Releases page
```

---

## Workflow 1: CI Pipeline (`.github/workflows/ci.yml`)

### Trigger
- Push ke branch `main` atau `develop`
- Pull request ke branch `main`

### Permissions
- `contents: read`

### Jobs

#### Job 1: `lint`
- Runs on: `windows-latest`
- Steps:
  1. Checkout code
  2. Setup Go 1.25
  3. Run golangci-lint
  4. Setup pnpm 9
  5. Setup Node 20
  6. Install frontend dependencies
  7. Run ESLint
  8. Run Prettier check

#### Job 2: `test`
- Runs on: `windows-latest`
- Steps:
  1. Checkout code
  2. Setup Go 1.25
  3. Run Go tests (`go test ./internal/... -v`)
  4. Setup pnpm 9
  5. Setup Node 20
  6. Install frontend dependencies
  7. Run Vitest (`pnpm test`)
  8. Run svelte-check

#### Job 3: `build`
- Needs: `lint` + `test` (harus pass dulu)
- Runs on: `windows-latest`
- Steps:
  1. Checkout code (full history untuk `git describe`)
  2. Setup Go 1.25
  3. Setup pnpm 9
  4. Setup Node 20
  5. Install Wails CLI
  6. Install frontend dependencies
  7. Build (`wails build`)
  8. Upload artifact (7 hari retention)

---

## Workflow 2: Release Pipeline (`.github/workflows/release.yml`)

### Trigger
- Push tag `v*` (contoh: `v0.2.0`, `v1.0.0`)

### Permissions
- `contents: write`

### Jobs

#### Job 1: `release`
- Runs on: `windows-latest`
- Steps:
  1. Checkout code (full history)
  2. Setup Go 1.25
  3. Setup pnpm 9
  4. Setup Node 20
  5. Install Wails CLI
  6. Install frontend dependencies
  7. Build Windows AMD64
  8. Build Windows ARM64
  9. Get version dari tag
  10. Create GitHub Release
  11. Upload binaries sebagai release assets

---

## Setup Requirements

### 1. GitHub Repository Settings

- Push repo ke GitHub (jika belum)
- Settings → Actions → General → "Allow all actions and reusable workflows"
- Settings → Actions → General → "Allow GitHub Actions to create and approve pull requests"

### 2. Branch Protection (Optional tapi Recommended)

- Settings → Branches → Add rule
- Branch name pattern: `main`
- ☑ Require a pull request before merging
- ☑ Require status checks to pass before merging
- Required status checks: `Lint`, `Test`, `Build`
- ☑ Require branches to be up to date before merging

### 3. Secrets (Tidak Wajib)

Tidak perlu secrets untuk initial setup. GitHub Actions punya `GITHUB_TOKEN` bawaan untuk upload release.

---

## Checklist Implementasi

| # | Task | File | Status |
|---|------|------|--------|
| 1 | Buat `.github/workflows/` directory | — | ⬜ |
| 2 | Buat `ci.yml` — lint + test + build jobs | `.github/workflows/ci.yml` | ⬜ |
| 3 | Buat `release.yml` — build + upload | `.github/workflows/release.yml` | ⬜ |
| 4 | Test CI pipeline (push ke main) | GitHub | ⬜ |
| 5 | Test release pipeline (tag v0.2.0) | GitHub | ⬜ |
| 6 | Update `docs/plan.md` | `docs/plan.md` | ⬜ |
| 7 | Update `CHANGELOG.md` | `CHANGELOG.md` | ⬜ |

---

## Estimasi Waktu

| Task | Estimasi |
|------|----------|
| Buat CI workflow | 30 menit |
| Buat Release workflow | 20 menit |
| Test & debug | 30 menit |
| **Total** | **~1.5 jam** |

---

## Notes

- **Platform:** Windows only (sesuai current project scope). macOS/Linux bisa ditambah nanti.
- **Go 1.25:** Pastikan GitHub Actions support Go 1.25.
- **Wails CLI:** Perlu install `wails` CLI di CI runner.
- **pnpm caching:** Menggunakan `pnpm/action-setup` + `actions/setup-node` dengan cache.
- **Artifact retention:** Build artifact disimpan 7 hari (bisa diatur).
