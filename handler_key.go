package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func (app *application) handlerFetchKey(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()

	apiKey := os.Getenv("BING_API_KEY")
	if apiKey == "" {
		respondWithError(w, 500, "API key not found")
		return
	}
	app.infoLog.Println(apiKey)
	repondWithJSON(w, 200, struct {
		ApiKey string `json:"api_key"`
	}{ApiKey: apiKey})
}