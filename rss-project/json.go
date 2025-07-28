package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func errorResponse(w http.ResponseWriter, code int, message string) {

	if code > 499 {
		log.Println("Error response:", message)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	jsonResponse(w, code, errResponse{
		Error: message,
	})

}

func jsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
