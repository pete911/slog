package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

type Server struct {
	logger     *slog.Logger
	middleware LogMiddleware
	port       int
}

func NewServer(logger *slog.Logger, port int) (Server, error) {
	if port < 0 || port > 65353 {
		return Server{}, fmt.Errorf("invalid port number %d, port has to be in range 0 - 65353", port)
	}
	return Server{
		logger:     logger.With("component", "server"),
		middleware: NewLogMiddleware(logger),
		port:       port,
	}, nil
}

func (s Server) Handler(w http.ResponseWriter, r *http.Request) {
	s.logger.InfoContext(r.Context(), "received request")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	if _, err := w.Write([]byte("OK")); err != nil {
		s.logger.ErrorContext(r.Context(), fmt.Sprintf("write response: %v", err))
	}
	s.logger.InfoContext(r.Context(), "returned response")
}

func (s Server) Run() error {
	mux := http.NewServeMux()
	mux.Handle("/", s.middleware.Handler(http.HandlerFunc(s.Handler)))

	s.logger.Info(fmt.Sprintf("starting server on port %d", s.port))
	address := fmt.Sprintf(":%d", s.port)
	if err := http.ListenAndServe(address, mux); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	s.logger.Info("server closed")
	return nil
}
