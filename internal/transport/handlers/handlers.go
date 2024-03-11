package handlers

import (
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/store"
)

type Handlers struct {
	store *store.Store
}


func New(store *store.Store) *Handlers {
	return &Handlers{
		store: store,
	}
}

type Handler interface {
	GetURLHandler(w http.ResponseWriter, r *http.Request)
	PostURLHandler(w http.ResponseWriter, r *http.Request)
	PostURLJsonHandler(w http.ResponseWriter, r *http.Request)
	GetConnDBHandler(w http.ResponseWriter, r *http.Request)
}
