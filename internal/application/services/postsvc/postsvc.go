package postsvc

import (
	"context"

	"github.com/RealBirdMan91/blog/internal/domain/content"
)

type Service struct {
	repo content.Repository
}

func NewService(repo content.Repository) *Service{
	return &Service{repo: repo}
}

func (s *Service) CreatePost(ctx context.Context, body string) (*content.Post, error) {
	p, err := content.NewPost(body)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}
