package service

import (
	"context"
	"fmt"

	"github.com/Lockok/efftest/internal/domain"
)

func (s *subscriptionService) GetByID(ctx context.Context, id int64) (*domain.Subscription, error) {
	if id <= 0 {
		return nil, ErrInvalidSubscriptionID
	}

	sub, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get subscription: %w", err)
	}

	return sub, nil
}
