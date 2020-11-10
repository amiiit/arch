package main

import (
	sql "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gitlab.com/amiiit/arco/auth"
	"gitlab.com/amiiit/arco/user"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"gitlab.com/amiiit/arco/graph"
	"gitlab.com/amiiit/arco/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	conn, err := sql.Connect("postgres", "user=test dbname=arco_test sslmode=disable")
	if err != nil {
		panic(err)
	}

	userRepo := user.UserRepository{
		DB: conn,
	}
	authRepo := auth.Repository{DB: conn}
	userService := user.UserService{
		Repo: userRepo,
	}
	graphResolver := graph.Resolver{
		UserService:    userService,
		UserRepository: userRepo,
	}

	router := chi.NewRouter()
	router.Use(auth.Middleware(authRepo, userRepo))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphResolver}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
