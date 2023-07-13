package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	if statusCode > 499 {
		log.Printf("Responding with 5xx Error: %s", message)
	}

	type ErrorResponse struct {
		Error string `json:"error"`
	}

	RespondWithJSON(w, r, statusCode, ErrorResponse{Error: message})
}

func RespondWithJSON(w http.ResponseWriter, r *http.Request, statusCode int, payload interface{}) {

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marsharl JSON response : %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}
