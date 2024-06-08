package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"github.com/bootdotdev/projects/posts/internal/database"

	_ "github.com/lib/pq"
)

// apiConfig struct holds the database queries
type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Load environment variables from .env file
	godotenv.Load(".env")

	// Get the PORT from environment variables or log fatal error if not set
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	// Get the DATABASE_URL from environment variables or log fatal error if not set
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)

	// Initialize apiConfig with the database queries
	apiCfg := apiConfig{
		DB: dbQueries,
	}

	// Initialize a new router
	router := chi.NewRouter()

	// Use CORS middleware to handle cross-origin requests
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Create a sub-router for version 1 of the API
	v1Router := chi.NewRouter()

	// Define API routes and their handlers
	v1Router.Post("/users", apiCfg.handlerUsersCreate)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerUsersGet))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerFeedCreate))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowsGet))
	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowCreate))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerFeedFollowDelete))

	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerPostsGet))

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)

	// Mount the version 1 API router to the main router
	router.Mount("/v1", v1Router)

	// Configure the HTTP server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Constants for the scraping process
	const collectionConcurrency = 10
	const collectionInterval = time.Minute

	// Start the scraping process in a separate goroutine
	go startScraping(dbQueries, collectionConcurrency, collectionInterval)

	// Log the serving port and start the HTTP server
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
