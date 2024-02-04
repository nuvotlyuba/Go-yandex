package main

import (
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/transport/handlers"
)


func main() {



	err := http.ListenAndServe(`:8080`, handlers.BasicRouter())
	if err != nil {
		panic(err)
	}

}
