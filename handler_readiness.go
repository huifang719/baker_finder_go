package main

import (
	"net/http"
)


func (app *application) HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	repondWithJSON(w, 200, struct {
		Status string `json:"status"`
		Message string `json:"message"`
	}{Status: "ok", Message: "Hello from baker finder"	})
}