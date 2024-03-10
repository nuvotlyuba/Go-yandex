package handlers

import (
	"context"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/services"
)

func (h Handlers) GetConnDBHandler(w http.ResponseWriter, r *http.Request) {
	//пингануть базу
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	s := new(services.Service)
	if err := s.PingDB(ctx);err != nil {
		http.Error(w, error.Error(err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
