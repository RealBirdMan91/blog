package graph

import appuser "github.com/RealBirdMan91/blog/internal/application/services/user"

//go:generate go run github.com/99designs/gqlgen generate
// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService *appuser.Service
}
