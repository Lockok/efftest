package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Lockok/efftest/internal/repository"
)

func (r *subscriptionRepository) TotalCost(ctx context.Context, filter repository.SubscriptionCostFilter) (int, error) {
	query := `
		SELECT COALESCE(SUM(price), 0)
		FROM subscriptions
		WHERE start_date <= $2
			AND (end_date IS NULL OR end_date >= $1)
			AND ($3::uuid IS NULL OR user_id = $3)
			AND ($4::text = '' OR title = $4);
	`

	var total int
	err := r.pool.QueryRow(ctx, query,
		filter.PeriodStart,
		filter.PeriodEnd,
		filter.UserID,
		filter.Title,
	).Scan(&total)
	if err != nil {
		slog.Error(
			"failed to calculate subscriptions total cost",
			"period_start", filter.PeriodStart,
			"period_end", filter.PeriodEnd,
			"user_id", filter.UserID,
			"title", filter.Title,
			"error", err,
		)
		return 0, fmt.Errorf("failed to calculate subscriptions total cost: %w", err)
	}

	return total, nil
}
