package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
	user, err := app.config.DB.GetUserById(r.Context(), params.UserID)
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to get the user")
	}
	var alreadyExistReview bool
	existReview, err := app.config.DB.CheckExistReviewForBaker(r.Context(), database.CheckExistReviewForBakerParams{
		BakerID: params.BakerID,
		UserID: params.UserID,
	})
	if err != nil {
		app.errorLog.Println(err)
		if err.Error() == "sql: no rows in result set" {
			alreadyExistReview = false
		} 	
	} else {
		app.infoLog.Println(existReview)
		alreadyExistReview = true
	}
	
	if alreadyExistReview {
		respondWithError(w, 400, "You have already reviewed this baker")
		return
	}
	// Create a new baker
	review, err := app.config.DB.CreateReview(r.Context(), database.CreateReviewParams{
		ID:       uuid.New(),
		BakerID:  params.BakerID,
		Review:   params.Review,
		Rating:   fmt.Sprint(params.Rating),
		UserID:  params.UserID,
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		app.errorLog.Println(err)
		if err.Error() == "pq: insert or update on table \"reviews\" violates foreign key constraint \"reviews_baker_id_fkey\"" {
			respondWithError(w, 400, "This baker does not exist")
			return
		}
		respondWithError(w, 500, "Failed to post the review")
		return
	}

	bakerName, err := app.config.DB.GetBakerByBakerId(r.Context(), review.BakerID)
		if err != nil {
			app.errorLog.Println(err)
			respondWithError(w, 500, "invalid baker id")
			return 
		}
	repondWithJSON(w, 200, databaseReviewtoReview(review, bakerName, user))
}

func (app *application) handlerDeleteReviews(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		ReviewID  uuid.UUID `json:"review_id"`
		UserID    uuid.UUID `json:"user_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramters{}
	err := decoder.Decode(&params)
	if err != nil {
		app.errorLog.Print(err)
		respondWithError(w, 400, "Invalid request")
		return
	}
	// Get the user
	user, err := app.config.DB.GetUserById(r.Context(), params.UserID)
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to get the user")
	}
	// Delete a review
	review,err := app.config.DB.DeleteReview(r.Context(), params.ReviewID)
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to delete the review")
		return
	}
	bakerName, err := app.config.DB.GetBakerByBakerId(r.Context(), review.BakerID)
		if err != nil {
			app.errorLog.Println(err)
			respondWithError(w, 500, "invalid baker id")
			return 
		}
	repondWithJSON(w, 200, databaseReviewtoReview(review, bakerName, user))
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
	bakerName, err := app.config.DB.GetBakerByBakerId(r.Context(), params.BakerID)
		if err != nil {
			app.errorLog.Println(err)
		}
	proceedReviews := []Review{}
	for _, review := range reviews {
		user, err := app.config.DB.GetUserById(r.Context(), review.UserID)
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to get the user")
	}
		proceedReviews = append(proceedReviews, databaseReviewtoReview(review, bakerName, user))
	}
	repondWithJSON(w, 200, proceedReviews)
}

// get all reviews for an user
func (app *application) handlerGetUserReviews(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		UserID  uuid.UUID `json:"user_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramters{}
	err := decoder.Decode(&params)
	if err != nil {
		app.errorLog.Print(err)
		respondWithError(w, 400, "Invalid request")
		return
	}

	// Get all reviews for a user
	reviews,err := app.config.DB.GetReviewsByUserId(r.Context(), params.UserID)
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to get the reviews")
		return
	}
	user, err := app.config.DB.GetUserById(r.Context(), params.UserID)
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to get the user")
	}
	proceedReviews := []Review{}
	for _, review := range reviews {
		bakerName, err := app.config.DB.GetBakerByBakerId(r.Context(), review.BakerID)
		if err != nil {
			app.errorLog.Println(err)
		}
		
		proceedReviews = append(proceedReviews, databaseReviewtoReview(review, bakerName, user))
	}
	repondWithJSON(w, 200, proceedReviews)
}