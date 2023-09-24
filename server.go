package main

import (
	"log"
	"net/http"
	"os"
	"review-pull-request-back-end/db"
	"review-pull-request-back-end/graph"
	"review-pull-request-back-end/graph/services"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// connect to db
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	database := db.Open()
	defer database.Close()
	db.Migrate(database)

	service := services.New(database)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Srv: service,
	}}))

	mux := http.NewServeMux()
	handler := cors.Default().Handler(mux)
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
