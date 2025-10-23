package postgres

import (
	"context"
	"database/sql"

	"github.com/RealBirdMan91/blog/internal/domain/content"
)

type PostgresPostRepo struct{ db *sql.DB }

func NewPostgresPostRepo(db *sql.DB) *PostgresPostRepo { return &PostgresPostRepo{db: db} }

var _ content.Repository = (*PostgresPostRepo)(nil)

func (r *PostgresPostRepo) Create(ctx context.Context, p *content.Post) error {
	const query = `
		INSERT INTO posts (id, body, likes, views, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		p.ID(),
		p.Body().String(),
		p.Likes().Int(),
		p.Views().Int(),
		p.CreatedAt(),
		p.UpdatedAt(),
	)
	return err
}
