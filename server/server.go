package server

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/daichi1002/go-graphql/adapters/graph/generated"
	"github.com/daichi1002/go-graphql/adapters/graph/resolvers"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func Serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}}))
	// CORS対応
	handler := cors.Default().Handler(srv)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
