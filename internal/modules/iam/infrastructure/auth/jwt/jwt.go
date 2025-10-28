package jwt

import (
	"errors"
	"time"

	"github.com/RealBirdMan91/blog/internal/modules/iam/application/ports"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

type HS256Signer struct {
	secret []byte
	ttl    time.Duration
}

func NewHS256(secret []byte, ttl time.Duration) *HS256Signer {
	return &HS256Signer{secret: secret, ttl: ttl}
}

func (s *HS256Signer) Sign(userID, email string) (string, error) {
	now := time.Now().UTC()
	claims := jwtv4.MapClaims{
		"sub":   userID,
		"email": email,
		"iat":   now.Unix(),
		"exp":   now.Add(s.ttl).Unix(),
	}
	tok := jwtv4.NewWithClaims(jwtv4.SigningMethodHS256, claims)
	return tok.SignedString(s.secret)
}

var ErrInvalidToken = errors.New("invalid token")

func (s *HS256Signer) Verify(token string) (*ports.TokenClaims, error) {
	parsed, err := jwtv4.Parse(token, func(t *jwtv4.Token) (any, error) {
		if _, ok := t.Method.(*jwtv4.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return s.secret, nil
	})
	if err != nil || !parsed.Valid {
		return nil, ErrInvalidToken
	}
	mc, ok := parsed.Claims.(jwtv4.MapClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	sub, _ := mc["sub"].(string)
	email, _ := mc["email"].(string)
	expF, _ := mc["exp"].(float64)
	if sub == "" || expF == 0 {
		return nil, ErrInvalidToken
	}

	return &ports.TokenClaims{
		UserID:  sub,
		Email:   email,
		ExpUnix: int64(expF),
	}, nil
}

// Compile-time checks:
var _ ports.TokenSigner = (*HS256Signer)(nil)
var _ ports.TokenVerifier = (*HS256Signer)(nil)
