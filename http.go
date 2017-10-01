package main

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON writes an object to the http response as JSON and sets the
// correct headers
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// RespondWithError writes an error to the http response as JSON and sets the
// correct headers
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}
