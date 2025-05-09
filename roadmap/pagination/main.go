package main

import (
	"database/sql"
	"log"
	"net/http"
	"os" // For environment variables
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/vldcreation/go-ressources/roadmap/pagination/handler"
	"github.com/vldcreation/go-ressources/roadmap/pagination/repository"
	"github.com/vldcreation/go-ressources/roadmap/pagination/service"
)

var (
	DB_DSN = os.Getenv("DB_DSN")
)

func init() {
	if DB_DSN == "" {
		log.Fatal("DB_DSN environment variable not set")
	}
}

func main() {
	db, err := sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	// Configure connection pool settings.
	db.SetMaxOpenConns(25)                 // Max number of open connections to the database.
	db.SetMaxIdleConns(25)                 // Max number of connections in the idle connection pool.
	db.SetConnMaxLifetime(5 * time.Minute) // Max amount of time a connection may be reused.

	// Ping the database to verify the connection.
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Successfully connected to the database!")

	// --- Dependency Injection ---
	videoRepo := repository.NewVideoRepository(db)
	videoSvc := service.NewVideoService(videoRepo)
	videoHdlr := handler.NewVideoHandler(videoSvc)

	// --- HTTP Router Setup ---
	mux := http.NewServeMux()
	videoHdlr.SetupRoutes(mux) // Register video routes

	// --- Start HTTP Server ---
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	serverAddr := ":" + port
	log.Printf("Starting server on %s\n", serverAddr)

	server := &http.Server{
		Addr:         serverAddr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", serverAddr, err)
	}
}
