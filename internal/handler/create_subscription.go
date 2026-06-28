package handler

import (
	"net/http"

	"github.com/Lockok/efftest/internal/dto"
)

// Create godoc
// @Summary Create subscription
// @Description Creates a new user subscription.
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param request body dto.CreateSubscriptionRequest true "Subscription data"
// @Success 201 {object} SubscriptionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions [post]
func (h *SubscriptionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateSubscriptionRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	sub, err := h.service.Create(r.Context(), req)
	if err != nil {
		h.writeServiceError(w, r, err)
		return
	}

	writeJSON(w, http.StatusCreated, subscriptionToResponse(sub))
}
