package user

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

func NewUser(emailRaw, passwordHash, avatarRaw string) (*User, error) {
	email, password, avatar, err := buildVOsFromRaw(emailRaw, passwordHash, avatarRaw)
	if err != nil {
		return nil, err
	}
	//time based UUID?
	//generate in client check then if exists in repo?
	return &User{
		id:        uuid.New(),
		email:     email,
		password:  password,
		avatar:    avatar,
		verified:  false,
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
	}, nil
}

func ReconstituteRaw(
	id uuid.UUID,
	emailRaw, passwordHash, avatarRaw string,
	verified bool,
	createdAt, updatedAt time.Time,
) (*User, error) {
	email, password, avatar, err := buildVOsFromRaw(emailRaw, passwordHash, avatarRaw)
	if err != nil {
		return nil, err
	}

	return reconstitute(id, email, password, avatar, verified, createdAt, updatedAt), nil
}

func reconstitute(
	id uuid.UUID,
	email Email,
	password PasswordHash,
	avatar AvatarURL,
	verified bool,
	createdAt, updatedAt time.Time,
) *User {
	return &User{
		id:        id,
		email:     email,
		password:  password,
		avatar:    avatar,
		verified:  verified,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func buildVOsFromRaw(emailStr, passHashStr, avatarStr string) (Email, PasswordHash, AvatarURL, error) {
	em, err := NewEmail(emailStr)
	if err != nil {
		return Email{}, PasswordHash{}, AvatarURL{}, err
	}
	ph, err := NewPasswordHash(passHashStr)
	if err != nil {
		return Email{}, PasswordHash{}, AvatarURL{}, err
	}
	av := AvatarURL{} // Zero = kein Avatar
	if s := strings.TrimSpace(avatarStr); s != "" {
		v, err := NewAvatarURL(s)
		if err != nil {
			return Email{}, PasswordHash{}, AvatarURL{}, err
		}
		av = v
	}
	return em, ph, av, nil
}
