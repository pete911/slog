package main

import (
	"context"
	"log/slog"
	"os"
)

type ContextHandler struct {
	*slog.JSONHandler
}

func NewContextHandler(level slog.Level) *ContextHandler {
	opts := &slog.HandlerOptions{Level: level}
	return &ContextHandler{
		slog.NewJSONHandler(os.Stderr, opts),
	}
}

func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ContextHandler{JSONHandler: h.JSONHandler.WithAttrs(attrs).(*slog.JSONHandler)}
}

func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return &ContextHandler{JSONHandler: h.JSONHandler.WithGroup(name).(*slog.JSONHandler)}
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	requestGroup := slog.Group("request",
		"method", getRequestMethod(ctx),
		"path", getRequestPath(ctx),
		"id", GetRequestId(ctx))
	cr := r.Clone()
	cr.AddAttrs(requestGroup)
	return h.JSONHandler.Handle(ctx, cr)
}
