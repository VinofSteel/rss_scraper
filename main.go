package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vinofsteel/rssscraper/handlers"
	"github.com/vinofsteel/rssscraper/internal/database"
)

func main() {
	// Loading env variables
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// Opening a connection to a database
	var (
		user     string = os.Getenv("PGUSER")
		password string = os.Getenv("PGPASSWORD")
		host     string = os.Getenv("PGHOST")
		port     string = os.Getenv("PGPORT")
		dbName   string = os.Getenv("PGDATABASE")
	)
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)

	log.Printf("Opening connection with database %s on port %s...\n", dbName, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening db connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging db: %v", err)
	}
	log.Println("Connection opened succesfully!")

	// Setting up our api configuration
	apiConfig := handlers.ApiConfig{
		DB: database.New(db),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthz", apiConfig.Readiness)
	mux.HandleFunc("GET /v1/err", apiConfig.Error)

	// Users
	mux.HandleFunc("POST /v1/users", apiConfig.UsersCreate)

	server := http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	log.Printf("Server running in port :%s\n", os.Getenv("PORT"))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
