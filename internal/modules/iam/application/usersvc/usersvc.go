package usersvc

import (
	"context"

	"github.com/RealBirdMan91/blog/internal/modules/iam/application/ports"
	user "github.com/RealBirdMan91/blog/internal/modules/iam/domain"
)

type Service struct {
	repo   user.Repository
	hasher ports.Hasher
}

func NewService(repo user.Repository, hasher ports.Hasher) *Service {
	return &Service{repo: repo, hasher: hasher}
}

func (s *Service) Register(ctx context.Context, emailRaw, passwordPlain, avatarRaw string) (*user.User, error) {
	hash, err := s.hasher.Hash(passwordPlain)
	if err != nil {
		return nil, err
	}

	u, err := user.NewUser(emailRaw, hash, avatarRaw)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}
