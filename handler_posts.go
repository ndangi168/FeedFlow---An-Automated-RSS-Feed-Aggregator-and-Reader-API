package main

import (
	"net/http"
	"strconv"

	"github.com/bootdotdev/projects/posts/internal/database"
)

// handlerPostsGet handles the request to get posts for a specific user.
// It retrieves the limit parameter from the URL query string and fetches posts accordingly.
func (cfg *apiConfig) handlerPostsGet(w http.ResponseWriter, r *http.Request, user database.User) {
	// Get the "limit" parameter from the URL query string.
	limitStr := r.URL.Query().Get("limit")
	limit := 10 // Default limit if not specified.

	// If a limit is specified and can be converted to an integer, use it.
	if specifiedLimit, err := strconv.Atoi(limitStr); err == nil {
		limit = specifiedLimit
	}

	// Fetch the posts for the user from the database with the specified limit.
	posts, err := cfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		// If there is an error fetching posts, respond with an internal server error.
		respondWithError(w, http.StatusInternalServerError, "Couldn't get posts for user")
		return
	}

	// Respond with the fetched posts in JSON format.
	respondWithJSON(w, http.StatusOK, databasePostsToPosts(posts))
}
