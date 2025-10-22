package usersvc

import (
	"context"

	"github.com/RealBirdMan91/blog/internal/domain/user"
)

type Hasher interface {
	Hash(plaintext string) (user.PasswordHash, error)
}

type Service struct {
	repo   user.Repository
	hasher Hasher
}

func NewService(repo user.Repository, hasher Hasher) *Service {
	return &Service{repo: repo, hasher: hasher}
}

func (s *Service) Register(
	ctx context.Context,
	emailRaw string,
	passwordPlain string,
	avatarRaw string,
) (*user.User, error) {
	hash, err := s.hasher.Hash(passwordPlain) // infra dependency -> app service
	if err != nil {
		return nil, err
	}

	u, err := user.NewUser(emailRaw, hash.Hash(), avatarRaw)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}
