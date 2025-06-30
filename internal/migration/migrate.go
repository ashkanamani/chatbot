package migration

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log/slog"
)

func RunMigrations(migrationsPath, dbURL string) error {
	m, err := migrate.New(fmt.Sprintf("file://%s", migrationsPath), dbURL)
	if err != nil {
		slog.Error("error", "err", err)
		return err
	}
	defer func() {
		err1, err2 := m.Close()
		if err1 != nil || err2 != nil {
			slog.Error("error while closing migration.", "source_err", err1, "database_err", err2)

		}
	}()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	slog.Info("Migrations successfully migrated")
	return nil
}
