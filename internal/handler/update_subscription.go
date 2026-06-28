package handler

import (
	"net/http"

	"github.com/Lockok/efftest/internal/dto"
)

// Update godoc
// @Summary Partially update subscription
// @Description Updates only fields provided in the request body and returns the full updated subscription.
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "Subscription ID"
// @Param request body dto.UpdateSubscriptionRequest true "Fields to update"
// @Success 200 {object} SubscriptionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [patch]
func (h *SubscriptionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	var req dto.UpdateSubscriptionRequest
	if !decodeJSON(w, r, &req) {
		return
	}

	sub, err := h.service.Update(r.Context(), id, req)
	if err != nil {
		h.writeServiceError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, subscriptionToResponse(sub))
}
