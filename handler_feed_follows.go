package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bootdotdev/projects/posts/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// handlerFeedFollowsGet handles the request to get all feed follows for the authenticated user.
func (cfg *apiConfig) handlerFeedFollowsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	// Fetch all feed follows for the user from the database.
	feedFollows, err := cfg.DB.GetFeedFollowsForUser(r.Context(), user.ID)
	if err != nil {
		// If there's an error, respond with an internal server error.
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feed follows")
		return
	}

	// Respond with the fetched feed follows in JSON format.
	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}

// handlerFeedFollowCreate handles the request to create a new feed follow for the authenticated user.
func (cfg *apiConfig) handlerFeedFollowCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	// Define a struct to hold the parameters from the request body.
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"` // ID of the feed to follow.
	}

	// Decode the JSON request body into the parameters struct.
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		// If there's an error decoding the request body, respond with an internal server error.
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	// Create a new feed follow in the database.
	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),       // Generate a new UUID for the feed follow.
		CreatedAt: time.Now().UTC(), // Set the creation time to the current UTC time.
		UpdatedAt: time.Now().UTC(), // Set the update time to the current UTC time.
		UserID:    user.ID,          // Set the user ID to the authenticated user's ID.
		FeedID:    params.FeedID,    // Set the feed ID from the request parameters.
	})
	if err != nil {
		// If there's an error creating the feed follow, respond with an internal server error.
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	// Respond with the created feed follow in JSON format.
	respondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}

// handlerFeedFollowDelete handles the request to delete a feed follow for the authenticated user.
func (cfg *apiConfig) handlerFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	// Extract the feed follow ID from the URL parameters.
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		// If the feed follow ID is invalid, respond with a bad request error.
		respondWithError(w, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}

	// Delete the feed follow from the database.
	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,      // Ensure the feed follow belongs to the authenticated user.
		ID:     feedFollowID, // Set the feed follow ID to be deleted.
	})
	if err != nil {
		// If there's an error deleting the feed follow, respond with an internal server error.
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}

	// Respond with an empty JSON object indicating successful deletion.
	respondWithJSON(w, http.StatusOK, struct{}{})
}
