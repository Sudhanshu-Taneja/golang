package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sudhanshu-Taneja/rsagg/internal/auth"
	"github.com/Sudhanshu-Taneja/rsagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		errorResponse(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		errorResponse(w, 500, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	jsonResponse(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {

	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		errorResponse(w, 401, fmt.Sprintf("Unauthorized: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		errorResponse(w, 404, fmt.Sprintf("User not found: %v", err))
		return
	}

	jsonResponse(w, 200, databaseUserToUser(user))
}
