package commands

import (
	"fmt"
	"users/storage"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Relational database migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		migrator, err := storage.NewMigrator()
		if err != nil {
			return err
		}
		err = migrator.Up()
		if err != nil {
			return err
		}

		fmt.Println("Successfully migrated files")

		return nil
	},
}
