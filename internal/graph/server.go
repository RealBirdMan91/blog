package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/RealBirdMan91/blog/internal/graph/resolvers"
	"github.com/RealBirdMan91/blog/internal/modules/content/application/postsvc"
	"github.com/RealBirdMan91/blog/internal/modules/iam/application/authsvc"
	"github.com/RealBirdMan91/blog/internal/modules/iam/application/usersvc"
	"github.com/vektah/gqlparser/v2/ast"
)

type Deps struct {
	UserService *usersvc.Service
	AuthService *authsvc.Service
	PostService *postsvc.Service
}

func NewGraphQLServer(d Deps) *handler.Server {
	srv := handler.New(resolvers.NewExecutableSchema(resolvers.Config{
		Resolvers: &resolvers.Resolver{UserService: d.UserService, AuthService: d.AuthService, PostService: d.PostService},
	}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	return srv
}
