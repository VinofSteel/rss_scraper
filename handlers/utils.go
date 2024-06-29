package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/vinofsteel/rssscraper/internal/database"
	"github.com/vinofsteel/rssscraper/internal/validation"
)

type ApiConfig struct {
	DB *database.Queries
	Validator *validation.Validator
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithErrorMap(w http.ResponseWriter, code int, errors map[string]string) {
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", errors)
	}

	type errorResponse struct {
		FieldErrors map[string]string `json:"field_errors"`
	}
	respondWithJSON(w, code, errorResponse{
		FieldErrors: errors,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	if code == http.StatusNoContent {
		w.WriteHeader(code)
		return
	}

	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(code)
	w.Write(dat)
}

func nullTimeToTimePtr(t sql.NullTime) *time.Time {
	if t.Valid {
		return &t.Time
	}

	return nil
}

func nullStringToStringPtr(s sql.NullString) *string {
	if s.Valid {
		return &s.String
	}

	return nil
}
