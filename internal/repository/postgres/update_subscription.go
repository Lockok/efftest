package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Lockok/efftest/internal/domain"
	"github.com/Lockok/efftest/internal/repository"
)

func (r *subscriptionRepository) Update(ctx context.Context, sub *domain.Subscription) error {
	query := `
		UPDATE subscriptions
		SET title = $1, price = $2, user_id = $3, start_date = $4, end_date = $5
		WHERE id = $6;
	`
	cmdTag, err := r.pool.Exec(ctx, query, sub.Title, sub.Price, sub.UserID, sub.StartDate, sub.EndDate, sub.ID)
	if err != nil {
		slog.Error("failed to update subscription", "id", sub.ID, "user_id", sub.UserID, "error", err)
		return fmt.Errorf("failed to update subscription: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		slog.Warn("subscription not found for update", "id", sub.ID)
		return repository.ErrSubscriptionNotFound
	}

	return nil
}
