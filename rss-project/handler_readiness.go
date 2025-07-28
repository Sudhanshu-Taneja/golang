package main

import "net/http"

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, struct{}{})
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	errorResponse(w, http.StatusInternalServerError, "An error occurred")
}
