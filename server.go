package main

import (
	"database/sql"
	"fmt"
	"github.com/suzushin54/study-graphql-with-go/graph/services"
	"github.com/suzushin54/study-graphql-with-go/middlewares/auth"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"
	"github.com/suzushin54/study-graphql-with-go/graph"
	"github.com/suzushin54/study-graphql-with-go/internal"
)

const (
	defaultPort = "8080"
	dbFile      = "mygraphql.db"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("%s?_foreign_keys=on", dbFile))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	service := services.NewServices(db)

	srv := handler.NewDefaultServer(
		internal.NewExecutableSchema(
			internal.Config{
				Resolvers: &graph.Resolver{
					Srv: service,
				},
				Directives: graph.Directive,
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", auth.Auth(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
