package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/nuvotlyuba/Go-yandex/config"
	"github.com/nuvotlyuba/Go-yandex/internal/repository"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
)


func PostURLHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/plain") {
		http.Error(w, "", http.StatusUnsupportedMediaType)
		return
	}
	responseData, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
		return
	}
	responseString := string(responseData)

	id := repository.CreateNewID(responseString)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, utils.StringURL(config.BaseURL, id))

}

func GetURLHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	url, isFind := repository.GetItemByID(id)
	if !isFind {
		http.Error(w, "Ссылка по ID не найдена", http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", url.LongURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
