package handler

import (
	"net/http"

	"github.com/Lockok/efftest/internal/dto"
)

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
