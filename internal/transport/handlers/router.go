package handlers

import "github.com/go-chi/chi/v5"

func BasicRouter() chi.Router {
	r := chi.NewRouter()
	s := new(Store)

	r.Post("/", s.PostURLHandler)
	r.Get("/{id}", s.GetURLHandler)

	return r
}
