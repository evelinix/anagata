package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"AnagataSentinel/internal/config"

	sqlite "gosqlite.org"
	crypto "gosqlite.org/vfs/crypto"
)

var db *sql.DB

func Init(cfg *config.Config) error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("get executable path: %w", err)
	}

	dataDir := filepath.Join(filepath.Dir(exePath), cfg.Database.Path)

	if err := os.MkdirAll(filepath.Dir(dataDir), 0755); err != nil {
		return fmt.Errorf("create data directory: %w", err)
	}

	salt := []byte("anagata-sentinel-v1-salt")
	key, err := crypto.DeriveKey([]byte(cfg.Database.Password), salt, crypto.Adiantum)
	if err != nil {
		return fmt.Errorf("derive encryption key: %w", err)
	}

	sqldb, err := crypto.Open(sqlite.Config{
		Path: dataDir,
		Pragmas: sqlite.Pragmas{
			JournalMode: sqlite.JournalWAL,
			BusyTimeout: 5 * time.Second,
			ForeignKeys: true,
		},
		MaxOpenConns: 1,
	}, crypto.Options{
		Key: key,
	})
	if err != nil {
		return fmt.Errorf("open encrypted database: %w", err)
	}

	db = sqldb.DB

	slog.Info("database opened", "path", dataDir)

	if err := runMigrations(); err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	return nil
}

func runMigrations() error {
	// Create migrations tracking table
	createTable := `
	CREATE TABLE IF NOT EXISTS schema_migrations (
		version INTEGER PRIMARY KEY,
		applied_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(createTable); err != nil {
		return fmt.Errorf("create schema_migrations table: %w", err)
	}

	// Get applied migrations
	applied, err := getAppliedMigrations()
	if err != nil {
		return fmt.Errorf("get applied migrations: %w", err)
	}

	// Find migration files
	migrations, err := filepath.Glob("internal/database/migrations/*.up.sql")
	if err != nil {
		// Try embedded path
		migrations, err = findMigrations()
		if err != nil {
			slog.Warn("no migration files found, skipping migrations")
			return nil
		}
	}

	sort.Strings(migrations)

	for _, path := range migrations {
		version := extractVersion(path)
		if version == 0 {
			continue
		}

		if applied[version] {
			continue
		}

		slog.Info("applying migration", "version", version)

		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read migration %d: %w", version, err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			return fmt.Errorf("execute migration %d: %w", version, err)
		}

		// Record migration
		if _, err := db.Exec("INSERT INTO schema_migrations (version) VALUES (?)", version); err != nil {
			return fmt.Errorf("record migration %d: %w", version, err)
		}
	}

	slog.Info("migrations complete")
	return nil
}

func getAppliedMigrations() (map[int]bool, error) {
	applied := make(map[int]bool)

	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var version int
		if err := rows.Scan(&version); err != nil {
			return nil, err
		}
		applied[version] = true
	}

	return applied, rows.Err()
}

func extractVersion(path string) int {
	base := filepath.Base(path)
	parts := strings.SplitN(base, "_", 2)
	if len(parts) == 0 {
		return 0
	}

	var version int
	fmt.Sscanf(parts[0], "%d", &version)
	return version
}

func findMigrations() ([]string, error) {
	// Try common locations
	locations := []string{
		"internal/database/migrations",
		filepath.Join(filepath.Dir(os.Args[0]), "migrations"),
	}

	for _, loc := range locations {
		migrations, err := filepath.Glob(filepath.Join(loc, "*.up.sql"))
		if err == nil && len(migrations) > 0 {
			return migrations, nil
		}
	}

	return nil, fmt.Errorf("migration files not found")
}

func Close() {
	if db != nil {
		db.Close()
		db = nil
	}
}
