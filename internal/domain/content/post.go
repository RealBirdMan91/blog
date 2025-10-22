package content

import (
	"errors"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

var (
	ErrEmptyBody   = errors.New("post body is empty")
	ErrAuthorIDNil = errors.New("author id is nil")
	ErrBodyTooLong = errors.New("post body too long")
)

const maxBodyRunes = 20_000

type Post struct {
	id        uuid.UUID
	body      string
	authorID  uuid.UUID
	likes     int
	views     int
	createdAt time.Time
	updatedAt time.Time
}

func NewPost(body string, authorID uuid.UUID) (*Post, error) {
	if strings.TrimSpace(body) == "" { // nur hier trimmen
		return nil, ErrEmptyBody
	}
	if utf8.RuneCountInString(body) > maxBodyRunes {
		return nil, ErrBodyTooLong
	}
	if authorID == uuid.Nil {
		return nil, ErrAuthorIDNil
	}
	return &Post{
		id:        uuid.New(),
		body:      body,
		authorID:  authorID,
		likes:     0,
		views:     0,
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
	}, nil
}
