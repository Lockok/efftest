package domain

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          int64
	Title 		string
	Price       int
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     *time.Time
}

func NewSubscription(id int64, title string, price int, userID uuid.UUID, startDate time.Time, endDate *time.Time) *Subscription {
	return &Subscription{
		ID:          id,
		Title:       title,
		Price:       price,
		UserID:      userID,
		StartDate:   startDate,
		EndDate:     endDate,
	}
}
