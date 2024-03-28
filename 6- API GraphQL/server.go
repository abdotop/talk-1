package main

import (
	"graphql/graph"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/abdotop/octopus"
	"github.com/abdotop/octopus/middleware/adaptor"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	app := octopus.New()

	app.Any("/", adaptor.HTTPHandler(playground.Handler("GraphQL playground", "/query")))
	app.Any("/query", adaptor.HTTPHandler(srv))

	app.Run("11.11.90.165:" + port)
}
