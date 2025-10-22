package appuser

import (
	"context"

	domain "github.com/RealBirdMan91/blog/internal/domain/user"
)

type Hasher interface {
	Hash(plaintext string) (domain.PasswordHash, error)
}

type Service struct {
	repo   domain.Repository
	hasher Hasher
}

func NewService(repo domain.Repository, hasher Hasher) *Service {
	return &Service{repo: repo, hasher: hasher}
}

func (s *Service) Register(
	ctx context.Context,
	emailRaw string,
	passwordPlain string,
	avatarRaw string,
) (*domain.User, error) {
	hash, err := s.hasher.Hash(passwordPlain) // infra dependency -> app service
	if err != nil {
		return nil, err
	}

	u, err := domain.NewUser(emailRaw, hash.Hash(), avatarRaw)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}
