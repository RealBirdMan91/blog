package resolvers

import (
	"github.com/RealBirdMan91/blog/internal/application/services/authsvc"
	"github.com/RealBirdMan91/blog/internal/application/services/postsvc"
	"github.com/RealBirdMan91/blog/internal/application/services/usersvc"
)

//go:generate go run github.com/99designs/gqlgen generate
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService *usersvc.Service
	AuthService *authsvc.Service
	PostService *postsvc.Service
}
