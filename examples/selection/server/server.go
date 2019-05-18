package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/samlitowitz/graphqlc-gen-gqlgen/examples/selection"
)

func main() {
	http.Handle("/", handler.Playground("Selection Demo", "/query"))
	http.Handle("/query", handler.GraphQL(selection.NewExecutableSchema(selection.Config{Resolvers: &selection.Resolver{}})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
