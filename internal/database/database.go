package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	sqlite "gosqlite.org"
	crypto "gosqlite.org/vfs/crypto"
)

const dbPassword = "0123456789"

var db *sql.DB

func Init() error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("get executable path: %w", err)
	}

	dataDir := filepath.Join(filepath.Dir(exePath), "data")

	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return fmt.Errorf("create data directory: %w", err)
	}

	dbPath := filepath.Join(dataDir, "sentinel.db")

	salt := []byte("anagata-sentinel-v1-salt")
	key, err := crypto.DeriveKey([]byte(dbPassword), salt, crypto.Adiantum)
	if err != nil {
		return fmt.Errorf("derive encryption key: %w", err)
	}

	sqldb, err := crypto.Open(sqlite.Config{
		Path: dbPath,
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

	log.Printf("[BOOT] Database opened: %s", dbPath)

	if err := runMigrations(); err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	return nil
}

func runMigrations() error {
	query := `
	CREATE TABLE IF NOT EXISTS app_meta (
		key   TEXT PRIMARY KEY,
		value TEXT NOT NULL
	);`

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("create app_meta table: %w", err)
	}

	log.Println("[BOOT] Migrations complete")
	return nil
}

func Close() {
	if db != nil {
		db.Close()
		db = nil
	}
}
