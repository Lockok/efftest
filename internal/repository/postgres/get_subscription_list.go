package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Lockok/efftest/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *subscriptionRepository) ListByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Subscription, error) {
	var subscriptions []domain.Subscription

	query := `
		SELECT id, title, price, user_id, start_date, end_date
		FROM subscriptions
		WHERE user_id = $1;
	`
	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		slog.Error("failed to list subscriptions by user id", "user_id", userID, "error", err)
		return nil, fmt.Errorf("failed to list subscriptions by user ID: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var sub domain.Subscription
		var endDate pgtype.Date

		err := rows.Scan(&sub.ID, &sub.Title, &sub.Price, &sub.UserID, &sub.StartDate, &endDate)
		if err != nil {
			slog.Error("failed to scan subscription", "user_id", userID, "error", err)
			return nil, fmt.Errorf("failed to scan subscription: %w", err)
		}
		if endDate.Valid {
			sub.EndDate = &endDate.Time
		}

		subscriptions = append(subscriptions, sub)
	}

	if err := rows.Err(); err != nil {
		slog.Error("failed to iterate subscriptions", "user_id", userID, "error", err)
		return nil, fmt.Errorf("failed to iterate subscriptions: %w", err)
	}

	return subscriptions, nil
}
