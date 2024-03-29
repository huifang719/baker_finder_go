package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/huifang719/baker_finder_go/internal/database"
)

func (app *application)  handlerCreateReview(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		BakerID  uuid.UUID `json:"baker_id"`
		Review   string `json:"review"`
		Rating   int32 `json:"rating"`
		UserID uuid.UUID `json:"user_id"`
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
		ID:       uuid.New(),
		BakerID:  params.BakerID,
		Review:   params.Review,
		Rating:   fmt.Sprint(params.Rating),
		UserID:  params.UserID,
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
		ReviewID  uuid.UUID `json:"review_id"`
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

// get all reviews for a baker
func (app *application) handlerGetReviews(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		BakerID  uuid.UUID `json:"baker_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramters{}
	err := decoder.Decode(&params)
	if err != nil {
		app.errorLog.Print(err)
		respondWithError(w, 400, "Invalid request")
		return
	}

	// Get all reviews for a baker
	reviews,err := app.config.DB.GetAllReviews(r.Context(), params.BakerID)
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to get the reviews")
		return
	}
	repondWithJSON(w, 200, reviews)
}
