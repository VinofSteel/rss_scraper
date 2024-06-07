package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/vinofsteel/rssscraper/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/healthz", handlers.Readiness)
	mux.HandleFunc("GET /v1/err", handlers.Error)

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
