package repositories

import (
	"fmt"
	"path/filepath"

	utils "song-library/internal/utils"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (r *Repository) RunMigrations(cfg *utils.Config) error {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	migrationsPath := "file://" + filepath.ToSlash(filepath.Join(cfg.Database.MigrationsPath, "migrations"))

	r.log.Debug("Migrations path:", migrationsPath)

	m, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		return err
	}

	version, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return err
	}

	r.log.Debug("Current migration version:", version)

	if dirty {
		r.log.Debug("Dirty migration detected, starting rollback...")
		if err := m.Force(int(version) - 1); err != nil {
			return err
		}
	}

	if err := m.Steps(1); err != nil {
		if err == migrate.ErrNoChange {
			r.log.Debug("No migrations to apply")
			return nil
		}
		r.log.Debug("Rolling back last migration...")
		if rollbackErr := m.Down(); rollbackErr != nil {
			return fmt.Errorf("error while rolling back last migration: %w (initial error: %v)", rollbackErr, err)
		}
		return err
	}

	return nil
}
