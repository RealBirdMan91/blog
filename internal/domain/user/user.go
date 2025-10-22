package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id                   uuid.UUID
	email                Email
	password             PasswordHash
	avatar               AvatarURL
	verified             bool
	createdAt, updatedAt time.Time
}

func (u *User) ID() uuid.UUID          { return u.id }
func (u *User) Email() Email           { return u.email }
func (u *User) Password() PasswordHash { return u.password }
func (u *User) Avatar() AvatarURL      { return u.avatar }
func (u *User) Verified() bool         { return u.verified }
func (u *User) CreatedAt() time.Time   { return u.createdAt }
func (u *User) UpdatedAt() time.Time   { return u.updatedAt }
