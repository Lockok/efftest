package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Lockok/efftest/internal/config"
	"github.com/Lockok/efftest/internal/handler"
	"github.com/Lockok/efftest/internal/repository/postgres"
	"github.com/Lockok/efftest/internal/service"
	"github.com/Lockok/efftest/internal/storage"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	cfg, err := config.Load()
	if err != nil {
		logger.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	logger.Info("connecting to postgres", "host", cfg.DB.Host, "port", cfg.DB.Port, "database", cfg.DB.Name)

	pool, err := storage.NewPostgres(cfg.DB)
	if err != nil {
		logger.Error("failed to connect to postgres", "error", err)
		os.Exit(1)
	}

	defer pool.Close()
	logger.Info("postgres connection established", "host", cfg.DB.Host, "port", cfg.DB.Port, "database", cfg.DB.Name)

	repo := postgres.NewSubscriptionRepository(pool)
	subscriptionService := service.NewSubscriptionService(repo)
	subscriptionHandler := handler.NewSubscriptionHandler(subscriptionService)

	addr := ":" + cfg.HTTP.Port
	logger.Info("http server listening", "addr", addr)
	if err := http.ListenAndServe(addr, subscriptionHandler.Routes()); err != nil {
		logger.Error("http server stopped", "error", err)
		os.Exit(1)
	}
}
