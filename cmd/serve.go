/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ashkanamani/chatbot/internal/repository"
	"github.com/ashkanamani/chatbot/internal/repository/postgres"
	"github.com/ashkanamani/chatbot/internal/repository/redis"
	"github.com/ashkanamani/chatbot/internal/service"
	"github.com/ashkanamani/chatbot/internal/telegram"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve command starts the application",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {

	// setup repositories
	redisClient, err := redis.NewRedisClient(os.Getenv("REDIS_URL"))
	if err != nil {
		slog.Error("could not connect to redis", "error", err.Error())
		os.Exit(1)
	}
	slog.Info("connected to redis successfully.")
	accountRedisRepository := repository.NewAccountRedisRepository(redisClient)

	postgresClient, err := postgres.NewPostgresConnection(os.Getenv("POSTGRES_ADDRESS"))
	if err != nil {
		slog.Error("could not connect to postgres", "error", err.Error())
		os.Exit(1)
	}
	slog.Info("connected to postgres successfully.")
	accountPostgresRepository := repository.NewAccountPostgresRepository(postgresClient, "accounts")

	// setup services
	accountService := service.NewAccountService(accountRedisRepository, accountPostgresRepository)

	// setup app
	app := service.NewApp(accountService)

	// setup telegram
	tg, err := telegram.NewTelegram(app, os.Getenv("TELEGRAM_API_TOKEN"))
	if err != nil {
		slog.Error("could not setup telegram", "error", err.Error())
		os.Exit(1)
	}
	tg.Start()

}
func init() {
	rootCmd.AddCommand(serveCmd)
}
