package storage

import (
	"context"


	"github.com/Lockok/efftest/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(cfg config.DBConfig) (*pgxpool.Pool, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.ConnectionString())
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}