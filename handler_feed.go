package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aaryan003/GoFeed/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if params.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Name is required")
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		log.Printf("create feed failed: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to create feed")
		return
	}
	respondWithJSON(w, http.StatusCreated, databaseFeedToAPIFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeed(w http.ResponseWriter, r *http.Request) {
	feed, err := apiCfg.DB.GetFeed(r.Context())
	if err != nil {
		log.Printf("get feed failed: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to get feed")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedsToAPIFeeds(feed))
}
