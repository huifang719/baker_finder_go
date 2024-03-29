package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Printf("responding with 5XX error: %s", message)
	} 
	type errorResponse struct {
		Error string `json:"error"`
	}
	repondWithJSON(w, code, errorResponse{Error: message})
}
func repondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	log.Printf("responding: %v", payload)

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling struct to json: %v", payload)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("responding with %d: %s", code, data)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}