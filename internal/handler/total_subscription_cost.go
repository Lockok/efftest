package handler

import (
	"log/slog"
	"net/http"

	"github.com/Lockok/efftest/internal/dto"
	"github.com/google/uuid"
)

// totalCost godoc
// @Summary Calculate total subscription cost
// @Description Calculates the total cost of subscriptions for the selected period with optional filters by user ID and subscription title.
// @Tags subscriptions
// @Produce json
// @Param start query string true "Period start in YYYY-MM format" example(2026-06)
// @Param end query string true "Period end in YYYY-MM format" example(2026-07)
// @Param user_id query string false "User UUID"
// @Param title query string false "Subscription title"
// @Success 200 {object} TotalCostResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/total [get]
func (h *SubscriptionHandler) totalCost(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	req := dto.TotalCostRequest{
		PeriodStart: query.Get("start"),
		PeriodEnd:   query.Get("end"),
		Title:       query.Get("title"),
	}

	if userIDValue := query.Get("user_id"); userIDValue != "" {
		userID, err := uuid.Parse(userIDValue)
		if err != nil {
			slog.Warn("invalid user_id in total cost request", "user_id", userIDValue, "error", err)
			writeError(w, http.StatusBadRequest, "invalid user_id")
			return
		}
		req.UserID = &userID
	}

	total, err := h.service.TotalCost(r.Context(), req)
	if err != nil {
		h.writeServiceError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, TotalCostResponse{TotalCost: total})
}
