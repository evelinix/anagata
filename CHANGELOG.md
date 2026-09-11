# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- **Log rotation** (`internal/logger/rotator.go`)
  - Size-based rotation (configurable via `max_size_mb`)
  - Active log: `logs/app.log`
  - Rotated logs: `logs/{year}/{month}_{date}.log`
  - Counter suffix for multiple rotations same day
  - `ArchiveActiveLog()` for graceful shutdown archiving
  - 5 unit tests (rotation, counter, close, initial size, interface)

- **ANTS Design System** (`frontend/src/style.css`)
  - Based on Microsoft WinUI 3 / Fluent Design
  - CSS custom properties for light/dark themes
  - Typography scale (display → caption)
  - Border radius, spacing, elevation tokens
  - Animation duration/easing tokens

- **ANTS Components** (`frontend/src/lib/components/`)
  - `Button` — 6 variants (primary, secondary, danger, ghost, outline, subtle), 3 sizes
  - `Input` — Error state, 3 sizes, label/placeholder support
  - `Card` — 4 variants (default, filled, outlined, acrylic)
  - `IconButton` — 4 variants, 3 sizes, icon-only (aria-label only)
  - `Typography` — 12 variants, 9 colors, dynamic HTML element, auto-weight for strong variants
  - `icons/` — 11 SVG icon components (IconBase + individual icons)

- **Date/time/number formatting** (`frontend/src/lib/i18n/format.ts`)
  - `formatDate`, `formatDateShort`, `formatTime`, `formatDateTime`
  - `formatNumber`, `formatCurrency`, `formatPercent`
  - `formatRelative` (relative time: "2 hours ago")
  - Uses `Intl.DateTimeFormat` / `NumberFormat` (browser-native)
  - 15 tests

- **Frontend test infrastructure**
  - Vitest + jsdom + @testing-library/svelte + @testing-library/jest-dom
  - `vite.config.ts` with `resolve.conditions: ['browser']`
  - `src/vitest-setup.ts` with cleanup + jest-dom matchers
  - `src/vite-env.d.ts` with Svelte type references + asset declarations
  - 81 total tests across 7 test files

- **Custom i18n system** (`frontend/src/lib/i18n/`)
  - Custom implementation with `CustomEvent` dispatch (replaced svelte-i18n)
  - EN + ID locales in `src/locales/`
  - Reactive via `$state` + `localeVersion` counter
  - `tl()` wrapper function for reactivity

- **Splash tests** (`internal/splash/splash_test.go`)
  - 6 tests (3 unconditional, 3 require `WALK_GUI_TEST=1`)
  - Tests: NewNativeSplash, SetStatusBeforeStart, CloseBeforeStart, Start, SetStatus, CloseIdempotent

### Changed
- **CSS prefix renamed**: `winui-` → `ants-`
  - All component classes: `ants-btn`, `ants-card`, `ants-input`, `ants-icon-btn`
  - All variant/size modifiers updated

- **Icon system**: Replaced `@tabler/icons-svelte` (incompatible with Svelte 5 runes) with custom SVG icon components
  - `IconBase.svelte` — shared SVG wrapper with size/stroke/color/style props
  - 11 individual icon components using IconBase
  - All icons support `style` prop for inline color overrides

- **Typography component** used throughout `App.svelte`
  - Replaced all inline style text with `<Typography>` component
  - Added `info` color token
  - Added `class`/`style` props for external styling

- **IconButton**: `label` prop now only sets `aria-label` (accessibility), no visible text rendered

- **App.svelte**: All icons now use custom SVG components (IconWorld, IconSun, IconMoon, etc.)

- **Project metadata** updated:
  - `package.json`: Added author, description, license, repository, homepage, bugs, keywords
  - `LICENSE`: Updated copyright holder to Eve Lin
  - `README.md`: Added author info, tech stack table, full project structure

### Removed
- `@tabler/icons-svelte` package (uses legacy `$$props`, incompatible with Svelte 5 runes mode)

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

[Unreleased]: https://github.com/evelinix/AnagataSentinel/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/evelinix/AnagataSentinel/releases/tag/v0.1.0
