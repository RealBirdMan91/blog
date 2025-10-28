package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/RealBirdMan91/blog/internal/application"
	"github.com/RealBirdMan91/blog/internal/graph"
	"github.com/RealBirdMan91/blog/internal/interfaces/httpauth"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app, err := application.NewApplication(application.Config{
		JWTSecret: "dev-insecure-secret",
		JWTTTL:    24 * time.Hour,
	})
	if err != nil {
		panic(err)
	}
	defer app.Close()

	//http und graph server defer?
	srv := graph.NewGraphQLServer(graph.Deps{
		UserService: app.Users(),
		AuthService: app.Auth(),
		PostService: app.Post(),
	})

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", httpauth.Middleware(app.Verifier(), true)(srv))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux, IdleTimeout: time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
