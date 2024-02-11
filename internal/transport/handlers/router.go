package handlers

import "github.com/go-chi/chi/v5"


func BasicRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", PostURLHandler)
	r.Get("/{id}", GetURLHandler)

	return r
}
