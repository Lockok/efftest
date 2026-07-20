package handler

import (
	"log/slog"
	"net/http"

	"github.com/Lockok/efftest/internal/service"
)

type HealthHandler struct {
	service service.HealthService
}

func NewHealthHandler(service service.HealthService) *HealthHandler {
	return &HealthHandler{
		service: service,
	}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (h *HealthHandler) Ready(w http.ResponseWriter, r *http.Request) {	
	ctx := r.Context()
	err := h.service.Ready(ctx)
	if err != nil {
		slog.Error("readliness check failed", "error", err)

		writeJSON(w, http.StatusServiceUnavailable, map[string]string{
			"status": "not ready",
		})
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})

}