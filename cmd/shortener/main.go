package main

import (
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/transport/handlers"
)


func main() {

	mux := http.NewServeMux()
	mux.HandleFunc(`/`, handlers.Handler)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}

}
