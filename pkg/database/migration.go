package database

import (
	"errors"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations menjalankan database migrations
func RunMigrations(databaseURL string) error {
	// Retry logic karena DB mungkin belum siap saat container app start
	var m *migrate.Migrate
	var err error

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		m, err = migrate.New(
			"file://migrations",
			databaseURL,
		)
		if err == nil {
			break
		}
		log.Printf("Gagal connect ke migration (percobaan %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("Database migrations berhasil dijalankan!")
	return nil
}
