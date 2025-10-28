package user

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, u *User) error
	ByEmail(ctx context.Context, email Email) (*User, error)
}
