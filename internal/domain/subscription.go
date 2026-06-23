package domain

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID        int64
	Title     string
	Price     int
	UserID    uuid.UUID
	StartDate time.Time
}