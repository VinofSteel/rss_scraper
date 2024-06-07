package handlers

import "net/http"

func (a *ApiConfig) Error(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
