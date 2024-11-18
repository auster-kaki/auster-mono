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
func NoContent(w http.ResponseWriter) {
	writeResponseJSON(w, http.StatusNoContent, nil)
}

// BadRequest - 400
func BadRequest(w http.ResponseWriter, err error) {
	writeResponseJSON(w, http.StatusBadRequest, map[string]any{
		"message": err.Error(),
	})
}

// Unauthorized - 401
func Unauthorized(w http.ResponseWriter, err error) {
	writeResponseJSON(w, http.StatusUnauthorized, map[string]any{
		"message": err.Error(),
	})
}

// Forbidden - 403
func Forbidden(w http.ResponseWriter, err error) {
	writeResponseJSON(w, http.StatusForbidden, map[string]any{
		"message": err.Error(),
	})
}

// NotFound - 404
func NotFound(w http.ResponseWriter, err error) {
	writeResponseJSON(w, http.StatusNotFound, map[string]any{
		"message": err.Error(),
	})
}

// InternalError - 500
func InternalError(w http.ResponseWriter, err error) {
	writeResponseJSON(w, http.StatusInternalServerError, map[string]any{
		"message": err.Error(),
	})
}

func writeResponseJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, fmt.Errorf("failed to encode json. %w", err).Error(), http.StatusInternalServerError)
	}
}
