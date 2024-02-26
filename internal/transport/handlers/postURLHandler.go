package handlers

import (
	"io"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/services"
)

func (s Store) PostURLHandler(w http.ResponseWriter, r *http.Request) {
	responseData, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
		return
	}
	responseString := string(responseData)

	ss := new(services.Service)
	data, err := ss.CreateNewURL(responseString)
	if err != nil {
		http.Error(w, "Не удалось получить короткую ссылку", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, data.ShortURL)

}
