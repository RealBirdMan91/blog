package postgres

import (
	"context"
	"database/sql"

	content "github.com/RealBirdMan91/blog/internal/modules/content/domain"
)

type PostgresPostRepo struct{ db *sql.DB }

func NewPostgresPostRepo(db *sql.DB) *PostgresPostRepo { return &PostgresPostRepo{db: db} }

var _ content.Repository = (*PostgresPostRepo)(nil)

func (r *PostgresPostRepo) Create(ctx context.Context, p *content.Post) error {
	const query = `
		INSERT INTO posts (id, author_id, body, likes, views, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.ExecContext(ctx, query,
		p.ID(),
		p.AuthorID(),
		p.Body().String(),
		p.Likes().Int(),
		p.Views().Int(),
		p.CreatedAt(),
		p.UpdatedAt(),
	)
	return err
}
