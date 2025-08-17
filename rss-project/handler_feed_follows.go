package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Sudhanshu-Taneja/rsagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		errorResponse(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})

	if err != nil {
		errorResponse(w, 500, fmt.Sprintf("Error creating feed follows: %v", err))
		return
	}

	jsonResponse(w, 200, databaseFeedFollowToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		errorResponse(w, 500, fmt.Sprintf("Error fetching feed follows: %v", err))
		return
	}

	jsonResponse(w, 200, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		errorResponse(w, 400, fmt.Sprintf("Invalid feed follow ID: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		errorResponse(w, 500, fmt.Sprintf("Error deleting feed follows: %v", err))
		return
	}

	jsonResponse(w, 204, nil)
}
