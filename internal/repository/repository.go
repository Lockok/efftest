package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Lockok/efftest/internal/domain"
	"github.com/google/uuid"
)

var (
	ErrSubscriptionNotFound = errors.New("subscription not found")
)

type SubscriptionCostFilter struct {
	PeriodStart time.Time
	PeriodEnd   time.Time
	UserID      *uuid.UUID
	Title       string
}

type SubscriptionRepository interface {
	Create(ctx context.Context, subscription *domain.Subscription) error
	GetByID(ctx context.Context, id int64) (*domain.Subscription, error)
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Subscription, error)
	Update(ctx context.Context, subscription *domain.Subscription) error
	Delete(ctx context.Context, id int64) error
	TotalCost(ctx context.Context, filter SubscriptionCostFilter) (int, error)
}
