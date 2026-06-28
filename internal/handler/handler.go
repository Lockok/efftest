package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	_ "github.com/Lockok/efftest/docs"
	"github.com/Lockok/efftest/internal/domain"
	"github.com/Lockok/efftest/internal/repository"
	"github.com/Lockok/efftest/internal/service"
	"github.com/google/uuid"
	httpSwagger "github.com/swaggo/http-swagger"
)

const maxRequestBodySize = 1 << 20

type SubscriptionHandler struct {
	service service.SubscriptionService
}

type ErrorResponse struct {
	Error string `json:"error"`
} // @name ErrorResponse

type SubscriptionResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	UserID    uuid.UUID `json:"user_id"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date,omitempty"`
} // @name SubscriptionResponse

type TotalCostResponse struct {
	TotalCost int `json:"total_cost"`
} // @name TotalCostResponse

func NewSubscriptionHandler(service service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		service: service,
	}
}

func (h *SubscriptionHandler) Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /subscriptions", h.Create)
	mux.HandleFunc("GET /subscriptions", h.ListByUserID)
	mux.HandleFunc("GET /subscriptions/total", h.totalCost)
	mux.HandleFunc("GET /subscriptions/{id}", h.GetByID)
	mux.HandleFunc("PATCH /subscriptions/{id}", h.Update)
	mux.HandleFunc("DELETE /subscriptions/{id}", h.Delete)
	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}

func parseID(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "invalid subscription id")
		return 0, false
	}

	return id, true
}

func decodeJSON(w http.ResponseWriter, r *http.Request, dst any) bool {
	r.Body = http.MaxBytesReader(w, r.Body, maxRequestBodySize)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return false
	}

	return true
}

func (h *SubscriptionHandler) writeServiceError(w http.ResponseWriter, r *http.Request, err error) {
	switch {
	case errors.Is(err, repository.ErrSubscriptionNotFound):
		slog.Warn("subscription not found", "method", r.Method, "path", r.URL.Path, "error", err)
		writeError(w, http.StatusNotFound, "subscription not found")
	case errors.Is(err, service.ErrInvalidSubscriptionID):
		slog.Warn("invalid subscription id", "method", r.Method, "path", r.URL.Path, "error", err)
		writeError(w, http.StatusBadRequest, "invalid subscription id")
	case errors.Is(err, service.ErrInvalidSubscriptionTitle):
		slog.Warn("invalid subscription title", "method", r.Method, "path", r.URL.Path, "error", err)
		writeError(w, http.StatusBadRequest, "invalid subscription title")
	case errors.Is(err, service.ErrInvalidSubscriptionPrice):
		slog.Warn("invalid subscription price", "method", r.Method, "path", r.URL.Path, "error", err)
		writeError(w, http.StatusBadRequest, "invalid subscription price")
	case errors.Is(err, service.ErrInvalidSubscriptionUser):
		slog.Warn("invalid subscription user", "method", r.Method, "path", r.URL.Path, "error", err)
		writeError(w, http.StatusBadRequest, "invalid subscription user")
	case errors.Is(err, service.ErrInvalidSubscriptionDate):
		slog.Warn("invalid subscription date", "method", r.Method, "path", r.URL.Path, "error", err)
		writeError(w, http.StatusBadRequest, "invalid subscription date")
	case errors.Is(err, service.ErrInvalidSubscriptionRange):
		slog.Warn("invalid subscription date range", "method", r.Method, "path", r.URL.Path, "error", err)
		writeError(w, http.StatusBadRequest, "invalid subscription date range")
	default:
		slog.Error("request failed", "method", r.Method, "path", r.URL.Path, "error", err)
		writeError(w, http.StatusInternalServerError, "internal server error")
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{Error: message})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func subscriptionToResponse(sub *domain.Subscription) SubscriptionResponse {
	response := SubscriptionResponse{
		ID:        sub.ID,
		Title:     sub.Title,
		Price:     sub.Price,
		UserID:    sub.UserID,
		StartDate: formatDate(sub.StartDate),
	}

	if sub.EndDate != nil {
		response.EndDate = formatDate(*sub.EndDate)
	}

	return response
}

func formatDate(value time.Time) string {
	return value.Format("2006-01")
}
