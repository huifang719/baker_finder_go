package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/huifang719/baker_finder_go/internal/database"
)

func (app *application)  handlerCreateReview(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		BakerID  string `json:"baker_id"`
		Review   string `json:"review"`
		Rating   int32 `json:"rating"`
		UserName int32 `json:"user_name"`
	}
	decoder := json.NewDecoder(r.Body)	
	params := paramters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid request")
		return
	}

	// Create a new baker

	review, err := app.config.DB.CreateReview(r.Context(), database.CreateReviewParams{
		ID:       int32(uuid.New().ID()),
		BakerID:  sql.NullString{String: params.BakerID, Valid: true},
		Review:   sql.NullString{String: params.Review, Valid: true},
		Rating:   sql.NullString{String: string(rune(params.Rating)), Valid: true},
	})
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to post the review")
		return
	}
	repondWithJSON(w, 200, review)	
}