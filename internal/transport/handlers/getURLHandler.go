package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nuvotlyuba/Go-yandex/internal/services"
)

func (s Store) GetURLHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	service := new(services.Service)
	data, err := service.FindUrl(id)
	if err != nil {
		http.Error(w, error.Error(err), http.StatusBadRequest)
		return
	}

	if data == nil {
		http.Error(w, "Ссылка по ID не найдена", http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", data.OriginalUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
