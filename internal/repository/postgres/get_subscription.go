package postgres

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Lockok/efftest/internal/domain"
	"github.com/Lockok/efftest/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *subscriptionRepository) GetByID(ctx context.Context, id int64) (*domain.Subscription, error) {

	query := `
		SELECT id, title, price, user_id, start_date, end_date FROM subscriptions WHERE id = $1;
	`

	sub := &domain.Subscription{}
	var endDate pgtype.Date

	row := r.pool.QueryRow(ctx, query, id)

	err := row.Scan(&sub.ID, &sub.Title, &sub.Price, &sub.UserID, &sub.StartDate, &endDate)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			slog.Warn("subscription not found", "id", id)
			return nil, repository.ErrSubscriptionNotFound
		}
		slog.Error("failed to get subscription by id", "id", id, "error", err)
		return nil, fmt.Errorf("failed to get subscription by ID: %w", err)
	}

	if endDate.Valid {
		sub.EndDate = &endDate.Time
	}

	return sub, nil
}
