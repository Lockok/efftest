package handler

import "net/http"

func (h *SubscriptionHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	sub, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		h.writeServiceError(w, r, err)
		return
	}

	writeJSON(w, http.StatusOK, subscriptionToResponse(sub))
}
