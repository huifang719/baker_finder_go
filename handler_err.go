package main

import (
	"net/http"
)

func (app *application) HandlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong!")
}