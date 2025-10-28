package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	user "github.com/RealBirdMan91/blog/internal/modules/iam/domain"
	"github.com/google/uuid"
)

type PostgresUsersRepo struct{ db *sql.DB }

func NewPostgresUsersRepo(db *sql.DB) *PostgresUsersRepo { return &PostgresUsersRepo{db: db} }

var _ user.Repository = (*PostgresUsersRepo)(nil)

func (r *PostgresUsersRepo) Create(ctx context.Context, u *user.User) error {
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

func (r *PostgresUsersRepo) ByEmail(ctx context.Context, em user.Email) (*user.User, error) {
	const q = `SELECT id,email,password_hash,avatar_url,verified,created_at,updated_at
               FROM users WHERE email=$1 LIMIT 1`

	var (
		id        uuid.UUID
		emailStr  string
		passHash  string
		avatar    sql.NullString
		verified  bool
		createdAt time.Time
		updatedAt time.Time
	)

	err := r.db.QueryRowContext(ctx, q, em.String()).Scan(
		&id, &emailStr, &passHash, &avatar, &verified, &createdAt, &updatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, user.ErrNotFound
		}
		return nil, fmt.Errorf("users.byEmail: %w", err)
	}

	avatarStr := ""
	if avatar.Valid {
		avatarStr = avatar.String
	}

	u, err := user.ReconstituteRaw(id, emailStr, passHash, avatarStr, verified, createdAt, updatedAt)
	if err != nil {
		return nil, err
	}

	return u, nil

}
