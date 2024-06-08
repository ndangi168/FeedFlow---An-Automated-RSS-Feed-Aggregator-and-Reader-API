package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// respondWithError sends an error response with the specified HTTP status code and message.
// It logs the error if it's a server error (5XX status code).
func respondWithError(w http.ResponseWriter, code int, msg string) {
	// Log the error message if it's a server-side error.
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	// Define the structure of the error response.
	type errorResponse struct {
		Error string `json:"error"`
	}

	// Send the error response as JSON.
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

// respondWithJSON sends a response with the specified HTTP status code and payload as JSON.
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Set the Content-Type header to indicate the response is in JSON format.
	w.Header().Set("Content-Type", "application/json")

	// Marshal the payload into JSON format.
	dat, err := json.Marshal(payload)
	if err != nil {
		// Log an error message if JSON marshaling fails and respond with a 500 Internal Server Error.
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}

	// Set the HTTP status code and write the JSON response.
	w.WriteHeader(code)
	w.Write(dat)
}
