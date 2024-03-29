package main

import (
	"net/http"
)


func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	type jsonResponse struct {
		Message string `json:"message"`
	}
	repondWithJSON(w, 200, jsonResponse  {Message: "hello from baker finder"})
}