package content

import "context"

type Repository interface {
	Create(ctx context.Context, p *Post) error
}
