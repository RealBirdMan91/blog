package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID
	email     Email
	password  PasswordHash
	avatar    AvatarURL
	verified  bool
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(emailRaw, passwordHash, avatarRaw string) (*User, error) {
	em, err := NewEmail(emailRaw)
	if err != nil {
		return nil, err
	}
	pw, err := NewPasswordHash(passwordHash)
	if err != nil {
		return nil, err
	}

	av, err := NewAvatarURL(avatarRaw)
	if err != nil {
		return nil, err
	}

	return &User{
		id:        uuid.New(),
		email:     em,
		password:  pw,
		avatar:    av,
		verified:  false,
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
	}, nil
}

func (u *User) ID() uuid.UUID          { return u.id }
func (u *User) Email() Email           { return u.email }
func (u *User) Password() PasswordHash { return u.password }
func (u *User) Avatar() AvatarURL      { return u.avatar }
func (u *User) Verified() bool         { return u.verified }
func (u *User) CreatedAt() time.Time   { return u.createdAt }
func (u *User) UpdatedAt() time.Time   { return u.updatedAt }
