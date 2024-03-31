package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{AllowedOrigins: []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:8081", "http://192.168.1.100:8081", "http://121.45.87.60:8081", "exp://192.168.1.100:8081s" },
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
	v1Router.Post("/getbaker", app.handlerGetBakerById)
	v1Router.Post("/review", app.handlerCreateReview)
	v1Router.Delete("/baker", app.handlerDeleteBaker)
	v1Router.Delete("/review", app.handlerDeleteReviews)
	v1Router.Post("/reviewsByBaker", app.handlerGetReviews)
	v1Router.Post("/getBakersByPostcode", app.handlerGetBakers)
	v1Router.Patch("/baker", app.handlerUpdateBaker)
	v1Router.Post("/user", app.handlerCreateUser)
	v1Router.Get("/listbakers", app.handlerGetAllBakers)
	v1Router.Post("/listreviews", app. handlerGetUserReviews)
	v1Router.Post("/getmybaker", app.handlerGetBakerByCreator)
	v1Router.Get("/fetchkey", app.handlerFetchKey)
	v1Router.Post("/getuser", app.handlerGetUser)
	return router
}