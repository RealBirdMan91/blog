package authsvc

import (
	"context"
	"errors"

	"github.com/RealBirdMan91/blog/internal/application/ports"
	"github.com/RealBirdMan91/blog/internal/domain/user"
)

type Service struct {
	repo   user.Repository
	hasher ports.Hasher
	signer ports.TokenSigner
}

func New(repo user.Repository, hasher ports.Hasher, signer ports.TokenSigner) *Service {
	return &Service{repo: repo, hasher: hasher, signer: signer}
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func (s *Service) Login(ctx context.Context, emailRaw, password string) (string, error) {
	em, err := user.NewEmail(emailRaw)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	u, err := s.repo.ByEmail(ctx, em)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if !s.hasher.Verify(u.Password().Hash(), password) {
		return "", ErrInvalidCredentials
	}
	return s.signer.Sign(u.ID().String(), u.Email().String())
}
