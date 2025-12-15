package xctx

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const (
	requestIDContextKey contextKey = "request_id"
)

func WithRequestID(ctx context.Context, requestID uuid.UUID) context.Context {
	return context.WithValue(ctx, requestIDContextKey, requestID)
}

func GetRequestID(ctx context.Context) uuid.UUID {
	return ctx.Value(requestIDContextKey).(uuid.UUID)
}

func TryGetRequestID(ctx context.Context) (uuid.UUID, bool) {
	requestID, ok := ctx.Value(requestIDContextKey).(uuid.UUID)
	return requestID, ok
}
