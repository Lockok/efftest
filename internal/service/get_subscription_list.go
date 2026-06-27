package service

import (
	"context"
	"fmt"

	"github.com/Lockok/efftest/internal/domain"
	"github.com/google/uuid"
)

func (s *subscriptionService) ListByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Subscription, error) {
	if userID == uuid.Nil {
		return nil, ErrInvalidSubscriptionUser
	}

	subscriptions, err := s.repo.ListByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("list subscriptions: %w", err)
	}

	return subscriptions, nil
}
