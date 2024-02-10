package handlers

import (
	"flag"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/nuvotlyuba/Go-yandex/internal/repository"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
)

var baseURL string

func init() {
	flag.StringVar(&baseURL, "b", "", "Base url for short links")
}

func BasicRouter() chi.Router {
	r := chi.NewRouter()

	r.Post("/", PostUrlHandler)
	r.Get("/{id}", GetUrlHandler)

	return r
}

func PostUrlHandler(w http.ResponseWriter, r *http.Request) {
	flag.Parse()
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
	baseURL = parseBaseUrl(baseURL)

	id := repository.CreateNewID(responseString)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, utils.StringURL(baseURL, id))

}

func GetUrlHandler(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	url, isFind := repository.GetItemByID(id)
	if !isFind {
		http.Error(w, "Ссылка по ID не найдена", http.StatusBadRequest)
		return
	}

	w.Header().Set("Location", url.LongURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
