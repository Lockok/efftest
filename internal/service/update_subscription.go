package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/Lockok/efftest/internal/domain"
	"github.com/Lockok/efftest/internal/dto"
	"github.com/google/uuid"
)

func (s *subscriptionService) Update(ctx context.Context, id int64, req dto.UpdateSubscriptionRequest) (*domain.Subscription, error) {
	if id <= 0 {
		return nil, ErrInvalidSubscriptionID
	}

	sub, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get subscription before update: %w", err)
	}

	if req.Title != nil {
		title := strings.TrimSpace(*req.Title)
		if title == "" {
			return nil, ErrInvalidSubscriptionTitle
		}
		sub.Title = title
	}

	if req.Price != nil {
		if *req.Price <= 0 {
			return nil, ErrInvalidSubscriptionPrice
		}
		sub.Price = *req.Price
	}

	if req.UserID != nil {
		if *req.UserID == uuid.Nil {
			return nil, ErrInvalidSubscriptionUser
		}
		sub.UserID = *req.UserID
	}

	if req.StartDate != nil {
		startDate, err := parseSubscriptionDate(*req.StartDate)
		if err != nil {
			return nil, err
		}
		sub.StartDate = startDate
	}

	if req.EndDate != nil {
		endDateValue := strings.TrimSpace(*req.EndDate)
		if endDateValue == "" {
			sub.EndDate = nil
		} else {
			endDate, err := parseSubscriptionDate(endDateValue)
			if err != nil {
				return nil, err
			}
			sub.EndDate = &endDate
		}
	}

	if sub.EndDate != nil && sub.EndDate.Before(sub.StartDate) {
		return nil, ErrInvalidSubscriptionRange
	}

	if err := s.repo.Update(ctx, sub); err != nil {
		return nil, fmt.Errorf("update subscription: %w", err)
	}

	return sub, nil
}
