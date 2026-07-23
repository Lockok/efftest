package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func RequestID() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.New().String()

			w.Header().Set("X-Request-ID", requestID)

			ctx := r.Context()
			ctx = context.WithValue(ctx, requestIDKey, requestID)

			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
