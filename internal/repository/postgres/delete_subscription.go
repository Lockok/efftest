package postgres

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Lockok/efftest/internal/repository"
)

func (r *subscriptionRepository) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM subscriptions
		WHERE id = $1;
	`
	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		slog.Error("failed to delete subscription", "id", id, "error", err)
		return fmt.Errorf("failed to delete subscription: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		slog.Warn("subscription not found for delete", "id", id)
		return repository.ErrSubscriptionNotFound
	}

	return nil
}
