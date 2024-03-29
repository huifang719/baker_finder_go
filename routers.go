package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"http://localhost:3000", "http://localhost:3001"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300, 
		}))
	v1Router := chi.NewRouter()
	// v1 router, in case we want to add more versions with breaking changes in the future while v1 is still in use
	router.Mount("/v1", v1Router)
	v1Router.Get("/healthz", app.HandlerReadiness)
	v1Router.Get("/err", app.HandlerError)
	v1Router.Post("/baker", app.handlerCreateBaker)
	return router
}