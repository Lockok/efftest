package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Lockok/efftest/internal/domain"
)

func (r *subscriptionRepository) Create(ctx context.Context, subscription *domain.Subscription) error {

	query := `
		INSERT INTO subscriptions (title, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	row := r.pool.QueryRow(ctx, query,
		subscription.Title,
		subscription.Price,
		subscription.UserID,
		subscription.StartDate,
		subscription.EndDate,
	)

	err := row.Scan(&subscription.ID)
	if err != nil {
		slog.Error(
			"failed to create subscription",
			"error", err,
			"title", subscription.Title,
			"user_id", subscription.UserID,
		)
		return fmt.Errorf("failed to create subscription: %w", err)
	}

	return nil
}
