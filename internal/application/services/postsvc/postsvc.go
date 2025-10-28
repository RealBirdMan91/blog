package postsvc

import (
	"context"
	"errors"

	"github.com/RealBirdMan91/blog/internal/domain/content"
	"github.com/RealBirdMan91/blog/internal/interfaces/authctx"
)

var ErrUnauthenticated = errors.New("unauthenticated")

type Service struct {
	repo content.Repository
}

func NewService(repo content.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreatePost(ctx context.Context, body string) (*content.Post, error) {
	userID, ok := authctx.UserIDFrom(ctx)
	if !ok {
		return nil, ErrUnauthenticated
	}

	p, err := content.NewPost(body, userID)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}
