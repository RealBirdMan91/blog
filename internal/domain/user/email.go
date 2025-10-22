package user

import (
	"errors"
	"net/mail"
	"strings"
)

var ErrInvalidEmail = errors.New("invalid email")

// Email ist ein Value Object: immer getrimmt + lowercase + gültig.
type Email struct{ v string }

func NewEmail(raw string) (Email, error) {
	s := strings.ToLower(strings.TrimSpace(raw))
	if s == "" {
		return Email{}, ErrInvalidEmail
	}
	if _, err := mail.ParseAddress(s); err != nil {
		return Email{}, ErrInvalidEmail
	}
	return Email{v: s}, nil
}

func (e Email) String() string { return e.v }
func (e Email) IsZero() bool   { return e.v == "" }
