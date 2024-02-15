package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func BasicRouter(r *chi.Mux) chi.Router {
	s := new(Store)
	r.Post("/", s.PostURLHandler)
	r.Get("/{id}", s.GetURLHandler)

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		route = strings.Replace(route, "/*/", "/", -1)
		fmt.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	return r
}
