package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h Handler) GetURLHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	res, err := h.service.FindURL(id)

	if err != nil {
		http.Error(w, error.Error(err), http.StatusBadRequest)
		return
	}

	if res == "" {
		http.Error(w, "Ссылка по ID не найдена", http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", res)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
