package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/Lockok/efftest/internal/dto"
	"github.com/Lockok/efftest/internal/repository"
)

func (s *subscriptionService) TotalCost(ctx context.Context, req dto.TotalCostRequest) (int, error) {
	periodStart, err := parseSubscriptionDate(req.PeriodStart)
	if err != nil {
		return 0, err
	}

	periodEnd, err := parseSubscriptionDate(req.PeriodEnd)
	if err != nil {
		return 0, err
	}

	if periodEnd.Before(periodStart) {
		return 0, ErrInvalidSubscriptionRange
	}

	filter := repository.SubscriptionCostFilter{
		PeriodStart: periodStart,
		PeriodEnd:   periodEnd,
		UserID:      req.UserID,
		Title:       strings.TrimSpace(req.Title),
	}

	total, err := s.repo.TotalCost(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("calculate subscriptions total cost: %w", err)
	}

	return total, nil
}
