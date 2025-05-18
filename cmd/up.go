package cmd

import (
	"fmt"
	"log"

	"github.com/chiragsoni81245/net-sentinel/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3" // Required for Sqlite3
	_ "github.com/golang-migrate/migrate/v4/source/file"       // Required for file-based migrations

	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply database migrations",
	Run: func(cmd *cobra.Command, args []string) {
        // Generate application configuration
        config, err := config.GetConfig(configPath)
        if err != nil {
            log.Fatal(err)
        }
        
        migrationsPath := fmt.Sprintf("file://%s", config.Server.MigrationsPath)
		m, err := migrate.New(migrationsPath, config.Database.URI)
		if err != nil {
			log.Fatalf("Failed to initialize migration: %v", err)
		}

		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Migration failed: %v", err)
		}

		fmt.Println("Migrations applied successfully!")
	},
}

func init() {
	migrateCmd.AddCommand(upCmd)

	upCmd.Flags().StringVar(&configPath, "config", "", "Path to the config file (required)")
	upCmd.MarkFlagRequired("config")
}

