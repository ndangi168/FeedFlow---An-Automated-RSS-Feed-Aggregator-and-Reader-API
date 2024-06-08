package main

import "net/http"

// handlerReadiness handles readiness checks for the server.
// It responds with a JSON object indicating that the server is ready.
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// Respond with a status OK and a JSON payload {"status": "ok"}.
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// handlerErr handles error responses for the server.
// It simulates an internal server error response.
func handlerErr(w http.ResponseWriter, r *http.Request) {
	// Respond with an internal server error and a JSON payload {"error": "Internal Server Error"}.
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
