package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type HS256Signer struct {
	secret []byte
	ttl    time.Duration
}

func NewHS256(secret []byte, ttl time.Duration) *HS256Signer {
	return &HS256Signer{secret: secret, ttl: ttl}
}

func (s *HS256Signer) Sign(userID string, email string) (string, error) {
	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"iat":   now.Unix(),
		"exp":   now.Add(s.ttl).Unix(),
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tok.SignedString(s.secret)
}
