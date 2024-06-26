package storage

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)

type Migrator struct {
	migrate *migrate.Migrate
}

func (m *Migrator) Up() error {
	err := m.migrate.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("migration() up: %v", err)
		}
	}
	return nil
}

func (m *Migrator) Down() error {
	err := m.migrate.Down()
	if err != nil {
		if err != migrate.ErrNoChange {
			return fmt.Errorf("migration() down: %v", err)
		}
	}
	return nil
}

func NewMigrator() (*Migrator, error) {
	database := NewDatabase()
	migrationPath := viper.GetString("MIGRATION_PATH")
	migrationDSN := database.DSNMigrate()
	migrator, err := migrate.New(
		migrationPath,
		migrationDSN,
	)

	if err != nil {
		return nil, fmt.Errorf("migration() creating instance: %v", err)
	}
	return &Migrator{migrator}, nil
}
