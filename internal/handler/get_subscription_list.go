package handler

import (
	"log/slog"
	"net/http"

	"github.com/google/uuid"
)

// ListByUserID godoc
// @Summary List subscriptions by user
// @Description Returns subscriptions filtered by user ID.
// @Tags subscriptions
// @Produce json
// @Param user_id query string true "User UUID"
// @Success 200 {array} SubscriptionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions [get]
func (h *SubscriptionHandler) ListByUserID(w http.ResponseWriter, r *http.Request) {
	userIDValue := r.URL.Query().Get("user_id")
	userID, err := uuid.Parse(userIDValue)
	if err != nil {
		slog.Warn("invalid user_id in subscription list request", "user_id", userIDValue, "error", err)
		writeError(w, http.StatusBadRequest, "invalid user_id")
		return
	}

	subscriptions, err := h.service.ListByUserID(r.Context(), userID)
	if err != nil {
		h.writeServiceError(w, r, err)
		return
	}

	response := make([]SubscriptionResponse, 0, len(subscriptions))
	for i := range subscriptions {
		response = append(response, subscriptionToResponse(&subscriptions[i]))
	}

	writeJSON(w, http.StatusOK, response)
}
