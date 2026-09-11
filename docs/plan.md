# AnagataSentinel - Development Plan

> Roadmap peningkatan dari enterprise template menjadi production-ready application

---

## Phase 1: Core Infrastructure (Prioritas Tinggi)

### 1.1 Config System
- [x] Buat `internal/config/config.go`
- [x] Support YAML/TOML config file
- [x] Environment variable override
- [ ] Hot-reload saat config berubah
- [x] Default config generation saat first run
- [x] Validasi config dengan struct tags

### 1.2 Structured Logging
- [x] Ganti `log.Println` → `log/slog` (stdlib) atau `zerolog`
- [x] Output ke console + file (`logs/app.log`)
- [x] Log rotation (size-based atau time-based)
- [x] Log level configurable (debug/info/warn/error)
- [x] Structured fields (request_id, user_id, dll)

### 1.3 Error Handling UI
- [x] Tampilkan startup error via Walk dialog (Windows)
- [x] Error boundary di frontend (Svelte)
- [ ] Graceful degradation jika stage gagal
- [x] Retry mechanism untuk transient errors

---

## Phase 2: Developer Experience (Prioritas Sedang)

### 2.1 CI/CD Pipeline
- [ ] `.github/workflows/ci.yml` — lint + build + test
- [ ] `.github/workflows/release.yml` — build + upload assets
- [ ] Dependabot untuk dependency updates
- [ ] Conventional commits enforcement

### 2.2 Code Quality
- [x] `.golangci.yml` — Go linter config
- [x] `frontend/eslint.config.js` — ESLint config
- [x] `frontend/.prettierrc` — Prettier config
- [x] Pre-commit hooks (husky + lint-staged)

### 2.3 Build Versioning
- [x] Inject version via `ldflags` saat build
- [x] Tampilkan versi di About dialog
- [x] Auto-increment patch version (via git tags)
- [x] Git tag based release (`make version`)

### 2.4 Database Migrations
- [x] Integrasi `golang-migrate/migrate`
- [x] Migrations folder: `internal/database/migrations/`
- [x] Versioned migrations (000001_init.up.sql)
- [x] Auto-migrate saat startup

---

## Phase 3: User Experience (Prioritas Rendah)

### 3.1 Dark/Light Mode
- [x] Toggle di header/sidebar
- [x] Persist preference ke localStorage
- [x] System preference detection
- [x] Smooth transition animation (Tailwind `transition-colors`)

### 3.2 Internationalization (i18n)
- [x] Integrasi `svelte-i18n`
- [x] English + Indonesian
- [x] Locale file: `frontend/src/locales/`
- [x] Date/time/number formatting

### 3.3 Testing
- [x] Go unit tests (`*_test.go`) — 21 tests across all packages
- [x] Go integration tests
- [x] Frontend unit tests (Vitest)
- [x] Frontend component tests (Testing Library)
- [ ] E2E tests (Playwright)

### 3.4 Auto-Updater
- [x] Check GitHub releases periodically (1 hour interval)
- [ ] Download & prompt update
- [ ] Silent update option
- [ ] Update changelog display

---

## Phase 4: Features (Future)

### 4.1 Dashboard
- [x] Real-time monitoring widgets (StatusCard component)
- [x] Chart/graph (Canvas-based TrafficChart, no external dependencies)
- [x] Alert notification system (AlertList component)
- [x] Status overview cards
- [x] Sidebar navigation with tab switching

### 4.2 Plugin System
- [ ] Plugin interface definition
- [ ] Dynamic plugin loading
- [ ] Plugin manifest validation
- [ ] Plugin marketplace (optional)

### 4.3 Multi-Window
- [ ] Window management system
- [ ] Detached windows support
- [ ] Window state persistence

---

## Tech Stack Decisions

| Component | Choice | Reason |
|---|---|---|
| Config | `log/slog` + YAML | Stdlib, zero dependency |
| Logging | `log/slog` | Stdlib, structured, fast |
| Migrations | `golang-migrate` | Industry standard |
| CI/CD | GitHub Actions | Free for public repos |
| Linting | golangci-lint + ESLint | Industry standard |
| Testing | Go testing + Vitest | Simple, fast |
| i18n | svelte-i18n | Best for Svelte |

---

## Progress Log

| Date | Milestone | Status |
|---|---|---|
| 2026-09-10 | Enterprise template | ✅ Done |
| 2026-09-11 | Phase 1: Config + Logging + Error UI | ✅ Done |
| 2026-09-11 | Phase 2: Code Quality + Versioning + Migrations | ✅ Done |
| 2026-09-11 | Phase 3: Dark Mode + i18n + Testing + Updater | ✅ Done |
| - | Phase 4: Features | 🔲 Pending |
