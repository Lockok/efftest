package dto

import "github.com/google/uuid"

type CreateSubscriptionRequest struct {
	Title     string    `json:"title" example:"Netflix"`
	Price     int       `json:"price" example:"320"`
	UserID    uuid.UUID `json:"user_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	StartDate string    `json:"start_date" swaggertype:"string" format:"string" example:"2026-06"`
	EndDate   string    `json:"end_date,omitempty" swaggertype:"string" format:"string" example:"2026-07"`
} // @name CreateSubscriptionRequest

type UpdateSubscriptionRequest struct {
	Title     *string    `json:"title,omitempty" example:"Netflix"`
	Price     *int       `json:"price,omitempty" example:"320"`
	UserID    *uuid.UUID `json:"user_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440000"`
	StartDate *string    `json:"start_date,omitempty" swaggertype:"string" format:"string" example:"2026-06"`
	EndDate   *string    `json:"end_date,omitempty" swaggertype:"string" format:"string" example:"2026-07"`
} // @name UpdateSubscriptionRequest

type TotalCostRequest struct {
	PeriodStart string     `json:"period_start"`
	PeriodEnd   string     `json:"period_end"`
	UserID      *uuid.UUID `json:"user_id,omitempty"`
	Title       string     `json:"title,omitempty"`
}
