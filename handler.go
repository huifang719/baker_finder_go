package main

import "net/http"


func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from baker finder"))
}