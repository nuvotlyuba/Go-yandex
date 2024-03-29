package handler

import (
	"io"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/store"
)

func (h Handler) PostURLHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
		return
	}
	url := string(body)

	data, err := h.service.CreateURL(url)

	if err == store.ErrConflict {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusConflict)
		io.WriteString(w, data.ShortURL)
		return
	}
	if err != nil {
		http.Error(w, "Не удалось получить короткую ссылку", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, data.ShortURL)

}
