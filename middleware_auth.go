package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aaryan003/GoFeed/internal/auth"
	"github.com/aaryan003/GoFeed/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Missing or invalid API key")
			return
		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			log.Printf("get user failed: %v", err)
			respondWithError(w, http.StatusNotFound, fmt.Sprintf("User not found: %v", err))
			return
		}
		handler(w, r, user)
	}
}
