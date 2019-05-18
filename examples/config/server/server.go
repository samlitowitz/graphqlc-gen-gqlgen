package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	todo "github.com/samlitowitz/graphqlc-gen-gqlgen/examples/config"
)

func main() {
	http.Handle("/", handler.Playground("Todo", "/query"))
	http.Handle("/query", handler.GraphQL(
		todo.NewExecutableSchema(todo.New()),
	))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
