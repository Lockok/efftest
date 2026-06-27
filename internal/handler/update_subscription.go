package handler

import (
	"net/http"

	"github.com/Lockok/efftest/internal/dto"
)

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
