package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/nuvotlyuba/Go-yandex/internal/storage"
)


func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		PostUrl(w, r)
	case http.MethodGet:
		GetUrlById(w, r)
	default:
		http.Error(w, "Неверный метод", http.StatusBadRequest)
		return
	}
}

func PostUrl(w http.ResponseWriter, r *http.Request) {
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

	shortUrl := storage.CreateNewShotUrl(responseString)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, shortUrl)

}

func GetUrlById(w http.ResponseWriter, r *http.Request) {

	id := r.URL.String()[1:]
	url, isFind := storage.GetItemById(id)
	if !isFind {
		http.Error(w, "Ссылка по ID не найдена", http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", url.LongUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
