package service

import (
	"context"
	"fmt"

	"github.com/Lockok/efftest/internal/domain"
	"github.com/Lockok/efftest/internal/dto"
)

func (s *subscriptionService) Create(ctx context.Context, req dto.CreateSubscriptionRequest) (*domain.Subscription, error) {
	sub, err := subscriptionFromCreateRequest(req)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, sub); err != nil {
		return nil, fmt.Errorf("create subscription: %w", err)
	}

	return sub, nil
}
