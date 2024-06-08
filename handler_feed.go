package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bootdotdev/projects/posts/internal/database"
	"github.com/google/uuid"
)

// handlerFeedCreate handles the request to create a new feed and follow it for the user.
func (cfg *apiConfig) handlerFeedCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	// Define a struct to hold the parameters from the request body.
	type parameters struct {
		Name string `json:"name"` // Name of the feed.
		URL  string `json:"url"`  // URL of the feed.
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

	// Create a new feed in the database.
	feed, err := cfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),       // Generate a new UUID for the feed.
		CreatedAt: time.Now().UTC(), // Set the creation time to the current UTC time.
		UpdatedAt: time.Now().UTC(), // Set the update time to the current UTC time.
		UserID:    user.ID,          // Set the user ID to the authenticated user's ID.
		Name:      params.Name,      // Set the feed name from the request parameters.
		Url:       params.URL,       // Set the feed URL from the request parameters.
	})
	if err != nil {
		// If there's an error creating the feed, respond with an internal server error.
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed")
		return
	}

	// Create a new feed follow in the database, linking the user to the newly created feed.
	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),       // Generate a new UUID for the feed follow.
		CreatedAt: time.Now().UTC(), // Set the creation time to the current UTC time.
		UpdatedAt: time.Now().UTC(), // Set the update time to the current UTC time.
		UserID:    user.ID,          // Set the user ID to the authenticated user's ID.
		FeedID:    feed.ID,          // Set the feed ID to the newly created feed's ID.
	})
	if err != nil {
		// If there's an error creating the feed follow, respond with an internal server error.
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed follow")
		return
	}

	// Respond with the created feed and feed follow in JSON format.
	respondWithJSON(w, http.StatusOK, struct {
		feed       Feed
		feedFollow FeedFollow
	}{
		feed:       databaseFeedToFeed(feed),
		feedFollow: databaseFeedFollowToFeedFollow(feedFollow),
	})
}

// handlerGetFeeds handles the request to get all feeds.
func (cfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	// Fetch all feeds from the database.
	feeds, err := cfg.DB.GetFeeds(r.Context())
	if err != nil {
		// If there's an error fetching the feeds, respond with an internal server error.
		respondWithError(w, http.StatusInternalServerError, "Couldn't get feeds")
		return
	}

	// Respond with the fetched feeds in JSON format.
	respondWithJSON(w, http.StatusOK, databaseFeedsToFeeds(feeds))
}
