# Report: Auto-Updater (Download & Prompt)

**Date:** 2026-09-12
**Feature:** Auto-Updater - Download & Prompt with Auto-Restart

---

## Summary

Implementasi lengkap auto-updater untuk AnagataSentinel dengan fitur:
- Periodic check GitHub releases (1 hour interval)
- Modal dialog UI untuk notifikasi update
- Progress bar saat download
- Auto-restart dengan binary baru setelah download selesai

---

## Changes Made

### Go Backend

#### `internal/updater/updater.go`
- **Added `ApplyUpdate(filePath string) error`** — Membuat batch script wrapper yang:
  1. Menunggu 2 detik (app process exit)
  2. Copy binary baru → replace binary yang sedang jalan
  3. Jalankan app lagi
  4. Execute batch script → return ke Go
- **Added `applyUpdateWindows(currentExec, newBinaryPath string) error`** — Windows-specific implementation
- **Added `applyUpdateUnix(currentExec, newBinaryPath string) error`** — Unix-specific implementation
- **Added `GetDownloadedFilePath(release *Release) (string, bool)`** — Check apakah file update sudah di-download

#### `internal/updater/updater_test.go`
- **Added `TestApplyUpdate_MissingBinary`** — Test error handling untuk missing binary
- **Added `TestGetDownloadedFilePath_NotFound`** — Test case file tidak ditemukan
- **Added `TestGetDownloadedFilePath_Found`** — Test case file ditemukan

#### `internal/app/app.go`
- **Added `ApplyUpdate() (map[string]interface{}, error)`** — Exported method untuk frontend:
  - Check pending release & downloaded file exists
  - Call `updater.ApplyUpdate`
  - Emit event `update-applied` ke frontend

### Frontend Bindings

#### `frontend/wailsjs/go/app/App.js`
- Added `ApplyUpdate()` function
- Added `DownloadUpdate()` function
- Added `GetPendingRelease()` function

#### `frontend/wailsjs/go/app/App.d.ts`
- Added type declarations for all new functions

### Frontend UI

#### `frontend/src/App.svelte`
- **Replaced banner with modal dialog** — Popup center-screen dengan:
  - Header: "Update Available" + version info
  - Body: Release notes / changelog
  - Progress bar saat downloading
  - Tombol "Download" → "Restart Now" → "Later"
  - Tombol dismiss
- **Fixed Wails event listeners** — Register `runtime.EventsOn()` untuk:
  - `update-available` — Trigger modal muncul
  - `update-progress` — Update progress bar
  - `update-applied` — Handle restart state
- **Added state management:**
  - `showUpdateModal` — Control modal visibility
  - `updateStep` — Track download state: idle → downloading → downloaded → applying
  - `updateError` — Display error messages
- **Added "Check for Update" button** — Klik version number di footer untuk manual check

#### `frontend/src/lib/components/icons/IconDownload.svelte` (NEW)
- SVG icon untuk download button

#### `frontend/src/lib/components/icons/IconCheck.svelte` (NEW)
- SVG icon untuk restart/apply button

#### `frontend/src/lib/components/icons/index.ts`
- Added exports for `IconDownload` and `IconCheck`

### i18n

#### `frontend/src/locales/en.json`
```json
"update": {
  "title": "Update Available",
  "version": "Version {{version}} is available",
  "description": "A new version of AnagataSentinel is ready to install.",
  "checkForUpdate": "Check for Update",
  "download": "Download Update",
  "downloading": "Downloading...",
  "completed": "Download complete",
  "restartNow": "Restart Now",
  "later": "Later",
  "dismiss": "Dismiss",
  "error": "Update failed",
  "checking": "Checking for updates...",
  "upToDate": "You're up to date"
}
```

#### `frontend/src/locales/id.json`
```json
"update": {
  "title": "Update Tersedia",
  "version": "Versi {{version}} tersedia",
  "description": "Versi baru AnagataSentinel siap diinstal.",
  "checkForUpdate": "Periksa Update",
  "download": "Unduh Update",
  "downloading": "Mengunduh...",
  "completed": "Unduhan selesai",
  "restartNow": "Mulai Ulang Sekarang",
  "later": "Nanti",
  "dismiss": "Tutup",
  "error": "Update gagal",
  "checking": "Memeriksa update...",
  "upToDate": "Versi terbaru"
}
```

---

## Update Flow

```
┌─────────────────────────────────────────────────────────┐
│  1. CHECK FOR UPDATE                                    │
│     - Periodic check setiap 1 jam                       │
│     - Manual check: klik version di footer              │
│     - Emit event "update-available" ke frontend         │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  2. SHOW MODAL                                          │
│     - Tampilkan modal dialog dengan info version        │
│     - Release notes / changelog                         │
│     - Tombol: Download / Later / Dismiss                │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  3. DOWNLOAD                                            │
│     - User klik "Download Update"                       │
│     - Progress bar tampil (0% → 100%)                   │
│     - Emit event "update-progress" setiap chunk         │
└─────────────────────────────────────────────────────────┘
                          ↓
┌─────────────────────────────────────────────────────────┐
│  4. RESTART                                             │
│     - Download selesai → tombol "Restart Now" muncul    │
│     - User klik "Restart Now"                           │
│     - Go jalankan batch script wrapper                  │
│     - App restart dengan binary baru                    │
└─────────────────────────────────────────────────────────┘
```

---

## Testing

### Go Tests
```bash
$ go test ./internal/...
ok   AnagataSentinel/internal/app      0.039s
ok   AnagataSentinel/internal/config   (cached)
ok   AnagataSentinel/internal/database (cached)
ok   AnagataSentinel/internal/errors   (cached)
ok   AnagataSentinel/internal/integration (cached)
ok   AnagataSentinel/internal/logger   (cached)
ok   AnagataSentinel/internal/splash   (cached)
ok   AnagataSentinel/internal/updater  0.039s
ok   AnagataSentinel/internal/version  (cached)
```

### Frontend Tests
```bash
$ pnpm test --run
 Test Files  11 passed (11)
      Tests  102 passed (102)
```

### Svelte Check
```bash
$ pnpm svelte-check
svelte-check found 0 errors and 0 warnings
```

---

## Files Changed

| File | Aksi | Lines Changed |
|------|------|---------------|
| `internal/updater/updater.go` | Modified | +120 lines |
| `internal/updater/updater_test.go` | Modified | +60 lines |
| `internal/app/app.go` | Modified | +45 lines |
| `frontend/wailsjs/go/app/App.js` | Modified | +15 lines |
| `frontend/wailsjs/go/app/App.d.ts` | Modified | +6 lines |
| `frontend/src/App.svelte` | Modified | +200 lines (modal) |
| `frontend/src/locales/en.json` | Modified | +10 lines |
| `frontend/src/locales/id.json` | Modified | +10 lines |
| `frontend/src/lib/components/icons/IconDownload.svelte` | Created | +9 lines |
| `frontend/src/lib/components/icons/IconCheck.svelte` | Created | +7 lines |
| `frontend/src/lib/components/icons/index.ts` | Modified | +2 lines |
| `docs/plan.md` | Modified | Updated checklist |

---

## Known Limitations

1. **No semver comparison** — Uses raw string comparison (`release.TagName != version.Version`). Any tag difference triggers update notification.

2. **Unauthenticated GitHub API** — Subject to 60 requests/hour rate limit.

3. **No download cleanup** — Downloaded files accumulate in temp directory.

4. **Windows-specific** — Batch script approach works on Windows. Unix uses shell script.

---

## Next Steps (Future)

- [ ] Silent update option (auto-download + prompt restart later)
- [ ] Update changelog display (parse release notes)
- [ ] Semver comparison for proper version ordering
- [ ] Download cleanup on app startup
- [ ] Rate limiting / caching for GitHub API
