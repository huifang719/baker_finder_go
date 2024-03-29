package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/huifang719/baker_finder_go/internal/database"
)

func (app *application)  handlerCreateReview(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		BakerID  int32 `json:"baker_id"`
		Review   string `json:"review"`
		Rating   int32 `json:"rating"`
		UserName string `json:"user_name"`
	}
	decoder := json.NewDecoder(r.Body)	
	params := paramters{}
	err := decoder.Decode(&params)
	if err != nil {
		app.errorLog.Print(err)
		respondWithError(w, 400, "Invalid request")
		return
	}
	// Create a new baker

	review, err := app.config.DB.CreateReview(r.Context(), database.CreateReviewParams{
		ID:       int32(uuid.New().ID()),
		BakerID:  sql.NullString{String: fmt.Sprint(params.BakerID), Valid: true},
		Review:   sql.NullString{String: params.Review, Valid: true},
		Rating:   sql.NullString{String: fmt.Sprint(params.Rating), Valid: true},
		UserName: sql.NullString{String: params.UserName, Valid: true},
	})
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to post the review")
		return
	}
	repondWithJSON(w, 200, review)	
}

func (app *application) handlerDeleteReviews(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		ReviewID  int32 `json:"review_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramters{}
	err := decoder.Decode(&params)
	if err != nil {
		app.errorLog.Print(err)
		respondWithError(w, 400, "Invalid request")
		return
	}

	// Delete a review
	review,err := app.config.DB.DeleteReview(r.Context(), params.ReviewID)
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to delete the review")
		return
	}
	repondWithJSON(w, 200, review)
}