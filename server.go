package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/RealBirdMan91/blog/internal/application"
	"github.com/RealBirdMan91/blog/internal/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app, err := application.NewApplication()
	if err != nil {
		panic(err)
	}
	defer app.Close()

	srv := graph.NewGraphQLServer(graph.Deps{
		UserService: app.Users,
	})

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

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
