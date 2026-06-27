package dto

import "github.com/google/uuid"

type CreateSubscriptionRequest struct {
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	UserID    uuid.UUID `json:"user_id"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date,omitempty"`
}

type UpdateSubscriptionRequest struct {
	Title     *string    `json:"title,omitempty"`
	Price     *int       `json:"price,omitempty"`
	UserID    *uuid.UUID `json:"user_id,omitempty"`
	StartDate *string    `json:"start_date,omitempty"`
	EndDate   *string    `json:"end_date,omitempty"`
}

type TotalCostRequest struct {
	PeriodStart string     `json:"period_start"`
	PeriodEnd   string     `json:"period_end"`
	UserID      *uuid.UUID `json:"user_id,omitempty"`
	Title       string     `json:"title,omitempty"`
}
