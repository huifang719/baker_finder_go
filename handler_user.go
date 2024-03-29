package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/huifang719/baker_finder_go/internal/database"
)

func (app *application)  handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
			UserName       string `json:"user_name"`
			UserType       string	`json:"user_type"`
			Email          string	`json:"email"`
			PasswordDigest string	`json:"password_digest"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramters{}
	err := decoder.Decode(&params)
	if err != nil {
		app.errorLog.Print(err)
		respondWithError(w, 400, "Invalid request")
		return
	}
	// Create a new user
	user, err := app.config.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		UserName:       params.UserName,
		UserType:       params.UserType,
		Email:          params.Email,
		PasswordDigest: params.PasswordDigest,
	})
	if err != nil {
		app.errorLog.Println(err)
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" {
			respondWithError(w, 400, "Email already exists")
			return
		}
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_user_name_key\"" {
			respondWithError(w, 400, "User name already exists")
			return 
		}
		respondWithError(w, 500, "Failed to create user")
		return
	}
	repondWithJSON(w, 200, user)
}