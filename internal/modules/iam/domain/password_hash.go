package user

import (
	"errors"
)

var ErrInvalidPasswordHash = errors.New("invalid password hash")

type PasswordHash struct {
	v string
}

func NewPasswordHash(hash string) (PasswordHash, error) {

	if len(hash) < 20 {
		return PasswordHash{}, ErrInvalidPasswordHash
	}
	return PasswordHash{v: hash}, nil
}

// Absichtlich kein Stringer, um Logging zu vermeiden.
func (p PasswordHash) Hash() string { return p.v }
func (p PasswordHash) IsZero() bool { return p.v == "" }
