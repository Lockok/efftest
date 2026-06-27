package handler

import (
	"log/slog"
	"net/http"

	"github.com/Lockok/efftest/internal/dto"
	"github.com/google/uuid"
)

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

	writeJSON(w, http.StatusOK, totalCostResponse{TotalCost: total})
}