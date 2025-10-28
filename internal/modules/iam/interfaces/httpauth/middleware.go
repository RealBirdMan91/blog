package httpauth

import (
	"net/http"
	"strings"

	"github.com/RealBirdMan91/blog/internal/modules/iam/application/ports"
	"github.com/RealBirdMan91/blog/internal/shared/authctx"
	"github.com/google/uuid"
)

func Middleware(po ports.TokenVerifier, optional bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if !strings.HasPrefix(strings.ToLower(auth), "bearer ") {
				if optional {
					next.ServeHTTP(w, r)
					return
				}
				http.Error(w, "missing bearer token", http.StatusUnauthorized)
				return
			}
			token := strings.TrimSpace(auth[len("Bearer "):])
			if token == "" {
				http.Error(w, "empty bearer token", http.StatusUnauthorized)
				return
			}
			claims, err := po.Verify(token)
			if err != nil {
				if optional {
					next.ServeHTTP(w, r)
					return
				}
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			uid, err := uuid.Parse(claims.UserID)
			if err != nil {
				http.Error(w, "invalid sub", http.StatusUnauthorized)
				return
			}
			ctx := authctx.WithUser(r.Context(), uid, claims.Email)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
