package main

import (
	"net/http"
)


func (app *application) HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	repondWithJSON(w, 200, struct {message string}{
		message: "hello from baker finder"})
}