package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Lockok/efftest/internal/config"
	"github.com/Lockok/efftest/internal/handler"
	"github.com/Lockok/efftest/internal/middleware"
	"github.com/Lockok/efftest/internal/repository/postgres"
	"github.com/Lockok/efftest/internal/server"
	"github.com/Lockok/efftest/internal/service"
	"github.com/Lockok/efftest/internal/storage"
)

// @title Subscriptions API
// @version 1.0
// @description REST API FOR MANAGING USER ONLINE SUBSCRIPTIONS AND CALCULATING SUBSCRIPTION COSTS.
// @host localhost:8080
// @BasePath /
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

	logger.Info("postgres connection established", "host", cfg.DB.Host, "port", cfg.DB.Port, "database", cfg.DB.Name)

	mux := http.NewServeMux()

	repo := postgres.NewSubscriptionRepository(pool)

	healthService := service.NewHealthService(pool)
	subscriptionService := service.NewSubscriptionService(repo)

	healthHandler := handler.NewHealthHandler(healthService)
	subscriptionHandler := handler.NewSubscriptionHandler(subscriptionService)

	healthHandler.Routes(mux)
	subscriptionHandler.Routes(mux)

	handler := middleware.RequestID()(
		middleware.Logging(logger)(
			middleware.Recovery(logger)(mux),
		),
	)

	addr := ":" + cfg.HTTP.Port

	srv := server.NewServer(addr, handler, logger)

	go func() {
		if err := srv.Run(); err != nil {
			logger.Error("failed to start HTTP server", "error:", err)
			os.Exit(1)
		}
	}()

	logger.Info("waiting for shutdown signal")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	logger.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("failed to shutdown HTTP server", "error:", err)
	}

	logger.Info("closing postgres connection")
	pool.Close()
}
