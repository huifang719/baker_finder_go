package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/huifang719/baker_finder_go/internal/database"
)

type BakerParamters struct {
	BakerID uuid.UUID `json:"baker_id"`
	Name string `json:"name"`
	Img   string `json:"img"`
	Address string `json:"address"`
	Suburb  string `json:"suburb"`
	Postcode  string `json:"postcode"`
	Contact   string `json:"contact"`
	Specialty string `json:"specialty"`
	Creator   uuid.UUID `json:"creator"`
}
func (app *application)  handlerCreateBaker(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)	
	params := BakerParamters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid request")
		return
	}

	// Create a new baker

	baker, err := app.config.DB.CreateBaker(r.Context(), database.CreateBakerParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Img:       sql.NullString{String: params.Img, Valid: true},
		Address:   sql.NullString{String: params.Address, Valid: true},
		Suburb:    sql.NullString{String: params.Suburb, Valid: true},
		Postcode:  sql.NullString{String: params.Postcode, Valid: true},
		Contact:   sql.NullString{String: params.Contact, Valid: true},
		Specialty: sql.NullString{String: params.Specialty, Valid: true},
		Creator:   params.Creator,
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
		BakerID uuid.UUID `json:"baker_id"`
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

// update baker
func (app *application) handlerUpdateBaker(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := BakerParamters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid request")
		return
	}

	bakerToUpdate, err := app.config.DB.GetBakerById(r.Context(), params.BakerID)
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to get baker")
		return
	}
	app.infoLog.Println(bakerToUpdate)
	updatedBaker, err := app.config.DB.UpdateBaker(r.Context(), database.UpdateBakerParams{
		ID:        params.BakerID,
		Name:      params.Name,
		Img:       sql.NullString{String: params.Img, Valid: true},
		Address:   sql.NullString{String: params.Address, Valid: true},
		Suburb:    sql.NullString{String: params.Suburb, Valid: true},
		Postcode:  sql.NullString{String: params.Postcode, Valid: true},
		Contact:   sql.NullString{String: params.Contact, Valid: true},
		Specialty: sql.NullString{String: params.Specialty, Valid: true},
		Creator:   bakerToUpdate.Creator,
	})
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to update baker")
		return
	}
	repondWithJSON(w, 200, updatedBaker)
}
// get all bakers from the same postcode
func (app *application) handlerGetBakers(w http.ResponseWriter, r *http.Request) {
	type paramters struct {
		Postcode string `json:"postcode"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, "Invalid request")
		return
	}

	// Get all bakers
	bakers, err := app.config.DB.GetBakersByPostcode(r.Context(), sql.NullString{String: params.Postcode, Valid: true})
	if err != nil {
		app.errorLog.Println(err)
		respondWithError(w, 500, "Failed to get bakers")
		return
	}
 
	repondWithJSON(w, 200, bakers)
}