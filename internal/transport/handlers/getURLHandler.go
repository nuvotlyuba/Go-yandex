package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (s Store) GetURLHandler(w http.ResponseWriter, r *http.Request) {


	id := chi.URLParam(r, "id")

	data, isFind := repo.GetItemByID(id)
	if !isFind {
		http.Error(w, "Ссылка по ID не найдена", http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", data.URL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
