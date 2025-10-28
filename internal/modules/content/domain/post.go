package content

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	id        uuid.UUID
	body      Body
	authorID  uuid.UUID
	likes     Likes
	views     Views
	createdAt time.Time
	updatedAt time.Time
}

func NewPost(body string, authorID uuid.UUID) (*Post, error) {
	bo, err := NewBody(body)
	if err != nil {
		return nil, err
	}
	li, err := NewLikes(0)
	if err != nil {
		return nil, err
	}
	vi, err := NewViews(0)
	if err != nil {
		return nil, err
	}

	return &Post{
		id:   uuid.New(),
		body: bo,
		authorID:  authorID,
		likes:     li,
		views:     vi,
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
	}, nil
}

func (p *Post) ID() uuid.UUID        { return p.id }
func (p *Post) AuthorID() uuid.UUID  { return p.authorID }
func (p *Post) Body() Body           { return p.body }
func (p *Post) Likes() Likes         { return p.likes }
func (p *Post) Views() Views         { return p.views }
func (p *Post) CreatedAt() time.Time { return p.createdAt }
func (p *Post) UpdatedAt() time.Time { return p.updatedAt }
