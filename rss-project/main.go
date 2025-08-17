package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Sudhanshu-Taneja/rsagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load(".env")

	postString := os.Getenv("PORT")
	if postString == "" {
		log.Fatal("PORT environment variable not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable not set")
	}
	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", readinessHandler)
	v1Router.Get("/err", errorHandler)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))

	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Router.Post("/feedFollows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollows))
	v1Router.Get("/feedFollows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + postString,
	}

	log.Printf("Server starting on port %s", postString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
