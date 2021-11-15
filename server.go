package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ftcRibeiro/simple-graphql/graph"
	"github.com/ftcRibeiro/simple-graphql/graph/generated"
	"github.com/ftcRibeiro/simple-graphql/postgres"
	"github.com/go-pg/pg/v10"
)

const defaultPort = "8080"

func main() {

	DB := postgres.New(&pg.Options{
		User:     "felipe",
		Password: "postgres",
		Database: "simple-graphql",
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		UserRepo: postgres.UserRepo{DB: DB},
		TodoRepo: postgres.TodoRepo{DB: DB},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
