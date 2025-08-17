package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sudhanshu-Taneja/rsagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"Url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		errorResponse(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		errorResponse(w, 500, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	jsonResponse(w, 200, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		errorResponse(w, 500, fmt.Sprintf("Error fetching feeds: %v", err))
		return
	}

	jsonResponse(w, 200, databaseFeedsToFeeds(feeds))
}
