package main

import (
	"io"
	"math/rand"
	"net/http"
	"strings"
)

type UrlItem struct {
	id      string
	longUrl string
}

var urlData []UrlItem

func GetItemById(data []UrlItem, id string) (item UrlItem, isFind bool) {
	for _, item := range data {
		if item.id == id {
			return item, true
		}
	}
	return UrlItem{}, false
}

func GenerateToken(length int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
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
	if err != nil {
		http.Error(w, "Не удалось прочитать тело запроса", http.StatusBadRequest)
		return
	}
	responseString := string(responseData)
	//генерируем уникальный токен
	token := GenerateToken(8)
	// записываем в структуру
	item := UrlItem{id: token, longUrl: responseString}
	urlData = append(urlData, item)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "http://localhost:8080/"+token)

}

func GetUrlById(w http.ResponseWriter, r *http.Request) {

	id := r.URL.String()[1:]
	url, isFind := GetItemById(urlData, id)
	if !isFind {
		http.Error(w, "Ссылка по ID не найдена", http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", url.longUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc(`/`, Handler)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}

}
