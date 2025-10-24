package authctx

import (
	"context"

	"github.com/google/uuid"
)

type ctxKey int

const (
	userIDKey ctxKey = iota
	emailKey
)

func WithUser(ctx context.Context, userID uuid.UUID, email string) context.Context {
	ctx = context.WithValue(ctx, userIDKey, userID)
	ctx = context.WithValue(ctx, emailKey, email)
	return ctx
}

func UserIDFrom(ctx context.Context) (uuid.UUID, bool) {
	v := ctx.Value(userIDKey)
	if v == nil {
		return uuid.UUID{}, false
	}
	id, ok := v.(uuid.UUID)
	return id, ok
}

func EmailFrom(ctx context.Context) (string, bool) {
	v := ctx.Value(emailKey)
	if v == nil {
		return "", false
	}
	s, ok := v.(string)
	return s, ok
}
