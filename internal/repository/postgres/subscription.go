package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type subscriptionRepository struct {
	pool *pgxpool.Pool
}

func NewSubscriptionRepository(pool *pgxpool.Pool) *subscriptionRepository {
	return &subscriptionRepository{
		pool: pool,
	}
}
