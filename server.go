package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/benedictus-danielle/go-gql-template-project/graph"
	"github.com/benedictus-danielle/go-gql-template-project/graph/generated"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	godotenv.Load()
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("SECRET")), nil)

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedHeaders: []string{"X-PINGOTHER", "Accept", "Authorization", "Content-Type"},
		AllowedMethods: []string{"POST", "OPTIONS"},
		MaxAge:         86400,
	}).Handler)

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	db, _ := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	fmt.Printf("%+v\n", db)

	router.Use(jwtauth.Verifier(tokenAuth))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.MultipartForm{})
	srv.AddTransport(transport.POST{})

	if os.Getenv("APP_ENV") == "local" {
		router.Use(middleware.Logger)
		srv.Use(extension.Introspection{})
		router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	}
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
