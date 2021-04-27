package database

import (
	"naivegateway/internal/config"
	"naivegateway/internal/logger"

	"github.com/kyokomi/emoji/v2"
	"github.com/spf13/cobra"
)

var log = logger.Log
var cfg = config.GetConfig()

// NewCommand creates a new database command for the cli
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "database",
		Short: "Database management",
	}
	cmd.AddCommand(newCreateCommand())
	cmd.AddCommand(newUpCommand())
	cmd.AddCommand(newDownCommand())
	cmd.AddCommand(newListCommand())
	return cmd
}

func newCreateCommand() *cobra.Command {
	migrationName := ""
	migrationDirectory := ""
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create a new migration in example with the provided name",
		Args:  cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			runCustomMigrations(migrationDirectory, "create", migrationName, -1)
		},
	}
	cmd.Flags().StringVarP(&migrationName, "name", "n", "", "Migration name (required)")
	cmd.Flags().StringVarP(&migrationDirectory, "path", "p", "migrations", "Migration directory")
	cmd.MarkFlagRequired("name")
	return cmd
}

func newUpCommand() *cobra.Command {
	migrationDirectory := ""
	var target int
	cmd := &cobra.Command{
		Use:   "migrate",
		Args:  cobra.NoArgs,
		Short: "run any migrations that haven't been run yet",
		Run: func(cmd *cobra.Command, args []string) {
			runCustomMigrations(migrationDirectory, "up", "", target)
		},
	}
	cmd.Flags().StringVarP(&migrationDirectory, "path", "p", "migrations", "Migration directory")
	cmd.Flags().IntVarP(&target, "target", "t", 9999, "Target migration number")
	return cmd
}

func newDownCommand() *cobra.Command {
	migrationDirectory := ""
	cmd := &cobra.Command{
		Use:   "rollback",
		Short: "roll back the previous run batch of migrations",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			runCustomMigrations(migrationDirectory, "down", "", -1)
		},
	}
	cmd.Flags().StringVarP(&migrationDirectory, "path", "p", "migrations", "Migration directory")
	return cmd
}

func newListCommand() *cobra.Command {
	migrationDirectory := ""
	cmd := &cobra.Command{
		Use:   "list",
		Short: "lists all available migrations",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			migrations, currentPos := listMigrations(migrationDirectory)
			for idx := range migrations {
				if idx > currentPos {
					emoji.Printf(":gear: %d - %s\n", idx, migrations[idx])
				} else if idx == currentPos {
					emoji.Printf(":check_mark_button: -> %d - %s\n", idx, migrations[idx])
				} else {
					emoji.Printf(":check_mark_button: %d - %s\n", idx, migrations[idx])
				}
			}
		},
	}
	cmd.Flags().StringVarP(&migrationDirectory, "path", "p", "migrations", "Migration directory")
	return cmd
}
