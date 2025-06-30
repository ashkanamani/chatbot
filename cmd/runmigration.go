/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ashkanamani/chatbot/internal/migration"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

// runmigrationCmd represents the runmigration command
var runmigrationCmd = &cobra.Command{
	Use:   "runmigration",
	Short: "A brief description of your command",
	Run:   runMigrations,
}

func runMigrations(cmd *cobra.Command, args []string) {
	postgresAddress := os.Getenv("POSTGRES_ADDRESS")
	err := migration.RunMigrations("internal/migration/sql", postgresAddress)
	if err != nil {
		slog.Error("failed to run migrations", "err", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(runmigrationCmd)

}
