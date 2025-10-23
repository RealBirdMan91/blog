package authsvc

import (
	"context"
	"errors"

	"github.com/RealBirdMan91/blog/internal/domain/user"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type Hasher interface {
	Verify(hash user.PasswordHash, plaintext string) bool
}
type TokenSigner interface {
	Sign(userID string, email string) (string, error)
}

type Service struct {
	users  user.Repository
	hasher Hasher
	signer TokenSigner
}

func New(users user.Repository, hasher Hasher, signer TokenSigner) *Service {
	return &Service{users: users, hasher: hasher, signer: signer}
}

func (s *Service) Login(ctx context.Context, emailRaw, password string) (string, error) {
	em, err := user.NewEmail(emailRaw)
	if err != nil {
		return "", ErrInvalidCredentials
	}
	u, err := s.users.ByEmail(ctx, em)
	if err != nil {
		return "", ErrInvalidCredentials
	}
	if !s.hasher.Verify(u.Password(), password) {
		return "", ErrInvalidCredentials
	}
	return s.signer.Sign(u.ID().String(), u.Email().String())
}
