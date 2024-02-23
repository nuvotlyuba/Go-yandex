package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
)

func (s Store) GetURLHandler(w http.ResponseWriter, r *http.Request) {


	id := chi.URLParam(r, "id")

	data, isFind := repo.GetItemByID(id)
	if !isFind {
		http.Error(w, "Ссылка по ID не найдена", http.StatusBadRequest)
		return
	}

	logger.Log.Info("Получение оригинальной ссылки")
	w.Header().Set("Location", data.URL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
