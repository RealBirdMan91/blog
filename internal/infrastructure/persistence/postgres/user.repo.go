package postgres

import (
	"context"
	"database/sql"

	duser "github.com/RealBirdMan91/blog/internal/domain/user"
)

type PostgresUsersRepo struct{ db *sql.DB }

func NewPostgresUsersRepo(db *sql.DB) *PostgresUsersRepo { return &PostgresUsersRepo{db: db} }

func (r *PostgresUsersRepo) Create(ctx context.Context, u *duser.User) error {
	const query = `
		INSERT INTO users (id, email, password_hash, avatar_url, verified, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	var avatar sql.NullString
	if !u.Avatar().IsZero() {
		avatar = sql.NullString{String: u.Avatar().String(), Valid: true}
	}

	_, err := r.db.ExecContext(ctx, query,
		u.ID(),
		u.Email().String(),
		u.Password().Hash(),
		avatar,
		u.Verified(),
		u.CreatedAt(),
		u.UpdatedAt(),
	)
	return err
}
