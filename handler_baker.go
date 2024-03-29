package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/huifang719/baker_finder_go/internal/database"
)

func (app *application)  handlerCreateBaker(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		ID int32 `json:"id"`
		Name string `json:"name"`
		Img   string `json:"img"`
		Address string `json:"address"`
		Suburb  string `json:"suburb"`
		Postcode  string `json:"postcode"`
		Contact   string `json:"contact"`
		Specialty string `json:"specialty"`
		Creator   string `json:"creator"`
	}
	decoder := json.NewDecoder(r.Body)	
	params := paramters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid request")
		return
	}

	// Create a new baker

	baker, err := app.config.DB.CreateBaker(r.Context(), database.CreateBakerParams{
		ID:        int32(uuid.New().ID()),
		Name:      sql.NullString{String: params.Name, Valid: true},
		Img:       sql.NullString{String: params.Img, Valid: true},
		Address:   sql.NullString{String: params.Address, Valid: true},
		Suburb:    sql.NullString{String: params.Suburb, Valid: true},
		Postcode:  sql.NullString{String: params.Postcode, Valid: true},
		Contact:   sql.NullString{String: params.Contact, Valid: true},
		Specialty: sql.NullString{String: params.Specialty, Valid: true},
		Creator:   sql.NullString{String: params.Creator, Valid: true},
	})
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to create baker")
		return
	}
	repondWithJSON(w, 200, baker)	
}

func (app *application) handlerDeleteBaker(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		BakerID int32 `json:"baker_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid request")
		return
	}

	// Delete a baker
	baker,err := app.config.DB.DeleteBaker(r.Context(), params.BakerID)
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to delete baker")
		return
	}
	repondWithJSON(w, 200, baker)	
}