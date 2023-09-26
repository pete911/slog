package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type LogMiddleware struct {
	logger *slog.Logger
}

func NewLogMiddleware(logger *slog.Logger) LogMiddleware {
	return LogMiddleware{logger: logger.With("component", "middleware")}
}

func (m LogMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		r = r.WithContext(GetContext(r.Context(), r.Method, r.URL.Path))

		next.ServeHTTP(w, r)
		duration := time.Since(startTime)

		durationMs := float64(duration.Nanoseconds()/1000) / 1000
		m.logger.InfoContext(r.Context(), fmt.Sprintf("request took %.2f milliseconds", durationMs))
	})
}
