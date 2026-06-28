package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Lockok/efftest/internal/domain"
	"github.com/Lockok/efftest/internal/dto"
	"github.com/Lockok/efftest/internal/repository"
	"github.com/google/uuid"
)

var (
	ErrInvalidSubscriptionID    = errors.New("invalid subscription id")
	ErrInvalidSubscriptionTitle = errors.New("invalid subscription title")
	ErrInvalidSubscriptionPrice = errors.New("invalid subscription price")
	ErrInvalidSubscriptionUser  = errors.New("invalid subscription user")
	ErrInvalidSubscriptionDate  = errors.New("invalid subscription start date")
	ErrInvalidSubscriptionRange = errors.New("invalid subscription date range")
)

type subscriptionService struct {
	repo repository.SubscriptionRepository
}

type SubscriptionService interface {
	Create(ctx context.Context, req dto.CreateSubscriptionRequest) (*domain.Subscription, error)
	GetByID(ctx context.Context, id int64) (*domain.Subscription, error)
	ListByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Subscription, error)
	Update(ctx context.Context, id int64, req dto.UpdateSubscriptionRequest) (*domain.Subscription, error)
	Delete(ctx context.Context, id int64) error
	TotalCost(ctx context.Context, req dto.TotalCostRequest) (int, error)
}

func NewSubscriptionService(repo repository.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{
		repo: repo,
	}
}

func subscriptionFromCreateRequest(req dto.CreateSubscriptionRequest) (*domain.Subscription, error) {
	return buildSubscription(0, req.Title, req.Price, req.UserID, req.StartDate, req.EndDate)
}

func buildSubscription(id int64, title string, price int, userID uuid.UUID, startDateValue string, endDateValue string) (*domain.Subscription, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, ErrInvalidSubscriptionTitle
	}

	if price <= 0 {
		return nil, ErrInvalidSubscriptionPrice
	}

	if userID == uuid.Nil {
		return nil, ErrInvalidSubscriptionUser
	}

	startDate, err := parseSubscriptionDate(startDateValue)
	if err != nil {
		return nil, err
	}

	var endDate *time.Time
	if strings.TrimSpace(endDateValue) != "" {
		parsedEndDate, err := parseSubscriptionDate(endDateValue)
		if err != nil {
			return nil, err
		}
		if parsedEndDate.Before(startDate) {
			return nil, ErrInvalidSubscriptionRange
		}
		endDate = &parsedEndDate
	}

	return domain.NewSubscription(id, title, price, userID, startDate, endDate), nil
}

func parseSubscriptionDate(value string) (time.Time, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Time{}, ErrInvalidSubscriptionDate
	}

	layouts := []string{
		"2006-01",
		"01-2006",
		time.DateOnly,
	}

	var parseErr error
	for _, layout := range layouts {
		parsed, err := time.Parse(layout, value)
		if err == nil {
			return parsed, nil
		}
		parseErr = err
	}

	return time.Time{}, fmt.Errorf("%w: %v", ErrInvalidSubscriptionDate, parseErr)
}
