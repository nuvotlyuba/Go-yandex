package handlers

import (
	"io"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
)

func (s Store) PostURLHandler(w http.ResponseWriter, r *http.Request) {
	// contentType := r.Header.Get("Content-Type")
	// if !strings.Contains(contentType, "text/plain") {
	// 	http.Error(w, "", http.StatusUnsupportedMediaType)
	// 	return
	// }
	responseData, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
		return
	}
	responseString := string(responseData)


	id := repo.CreateNewID(responseString)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, utils.StringURL(configs.BaseURL, id))

}
