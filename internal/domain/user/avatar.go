package user

import (
	"errors"
	"net/url"
	"strings"
)

var ErrInvalidAvatar = errors.New("invalid avatar")

// AvatarURL ist optional: leer = unset; falls gesetzt: http/https + Host.
type AvatarURL struct{ s string }

func NewAvatarURL(raw string) (AvatarURL, error) {
	s := strings.TrimSpace(raw)
	if s == "" {
		return AvatarURL{}, nil // optional
	}
	u, err := url.Parse(s)
	if err != nil || u.Host == "" {
		return AvatarURL{}, ErrInvalidAvatar
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return AvatarURL{}, ErrInvalidAvatar
	}
	return AvatarURL{s: u.String()}, nil
}

func (a AvatarURL) String() string { return a.s }
func (a AvatarURL) IsZero() bool   { return a.s == "" }
