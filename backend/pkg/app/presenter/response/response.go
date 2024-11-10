package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// OK - 200
func OK(w http.ResponseWriter, body any) {
	writeResponseJSON(w, http.StatusOK, body)
}

// Created - 201
func Created(w http.ResponseWriter, body any) {
	writeResponseJSON(w, http.StatusCreated, body)
}

// NoContent - 204
func NoContent(w http.ResponseWriter, body any) {
	writeResponseJSON(w, http.StatusNoContent, body)
}

func writeResponseJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, fmt.Errorf("failed to encode json. %w", err).Error(), http.StatusInternalServerError)
	}
}
