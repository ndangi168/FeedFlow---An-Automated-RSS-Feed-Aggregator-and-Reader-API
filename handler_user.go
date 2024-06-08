package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bootdotdev/projects/posts/internal/database"
	"github.com/google/uuid"
)

// handlerUsersCreate handles the creation of a new user.
func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	// Define a struct to hold the expected parameters from the request body.
	type parameters struct {
		Name string
	}

	// Create a new JSON decoder for the request body.
	decoder := json.NewDecoder(r.Body)

	// Initialize a parameters instance to hold the decoded values.
	params := parameters{}

	// Decode the request body into the parameters struct.
	err := decoder.Decode(&params)
	if err != nil {
		// If decoding fails, respond with an internal server error.
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	// Create a new user in the database with the provided parameters.
	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),       // Generate a new UUID for the user.
		CreatedAt: time.Now().UTC(), // Set the creation timestamp.
		UpdatedAt: time.Now().UTC(), // Set the update timestamp.
		Name:      params.Name,      // Use the provided name.
	})
	if err != nil {
		// Log the error and respond with an internal server error if user creation fails.
		log.Println(err)
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	// Respond with the created user in JSON format.
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}

// handlerUsersGet handles the retrieval of user information.
func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request, user database.User) {
	// Respond with the user information in JSON format.
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
