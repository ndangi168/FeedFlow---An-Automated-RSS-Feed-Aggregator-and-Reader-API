package main

import (
	"net/http"

	"github.com/bootdotdev/projects/posts/internal/auth"
	"github.com/bootdotdev/projects/posts/internal/database"
)

// authedHandler is a custom type for handlers that require authentication.
// It takes an http.ResponseWriter, an *http.Request, and a database.User as arguments.
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// middlewareAuth is a method of apiConfig that returns an http.HandlerFunc.
// It wraps a handler with authentication logic.
func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the API key from the request headers using the auth package.
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			// If the API key is not found, respond with a 401 Unauthorized error.
			respondWithError(w, http.StatusUnauthorized, "Couldn't find api key")
			return
		}

		// Fetch the user associated with the API key from the database.
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			// If the user is not found, respond with a 404 Not Found error.
			respondWithError(w, http.StatusNotFound, "Couldn't get user")
			return
		}

		// Call the original handler with the authenticated user.
		handler(w, r, user)
	}
}
