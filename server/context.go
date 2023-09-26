package main

import (
	"context"
	"crypto/rand"
	"fmt"
)

const (
	requestMethodContextKey contextKey = "method"
	requestPathContextKey   contextKey = "path"
	requestIDContextKey     contextKey = "id"
)

type contextKey string

// GetContext returns context with 'method', 'path' and 'request_id' set. This context should be set in the middleware
// and can be retrieved by logger when the context is passed in.
func GetContext(ctx context.Context, method, path string) context.Context {
	ctx = context.WithValue(ctx, requestMethodContextKey, method)
	ctx = context.WithValue(ctx, requestPathContextKey, path)
	return context.WithValue(ctx, requestIDContextKey, generateRequestId())
}

func getRequestMethod(ctx context.Context) string {
	return getContextValue(ctx, requestMethodContextKey)
}

func getRequestPath(ctx context.Context) string {
	return getContextValue(ctx, requestPathContextKey)
}

func GetRequestId(ctx context.Context) string {
	return getContextValue(ctx, requestIDContextKey)
}

func getContextValue(ctx context.Context, key contextKey) string {
	if v := ctx.Value(key); v != nil {
		if stringVal, ok := v.(string); ok {
			return stringVal
		}
	}
	return ""
}

func generateRequestId() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "00000000-0000-0000-0000-000000000000"
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
