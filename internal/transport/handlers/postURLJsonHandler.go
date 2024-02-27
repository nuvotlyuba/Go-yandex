package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/services"
	"go.uber.org/zap"
)

func (s Store) PostURLJsonHandler(w http.ResponseWriter, r *http.Request) {

	var req models.RequestBody
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		logger.Log.Debug("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	service := new(services.Service)
	data, err := service.CreateNewURL(req.URL)
	if err != nil {
		http.Error(w, error.Error(err), http.StatusBadRequest)
		return
	}

	resp := models.Response{
		Result: data.ShortURL,
	}
	fmt.Println(resp, "resp!!!")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		fmt.Println(err, "errrrrrrrrrrrr")
	logger.Log.Debug("error encoding response", zap.Error(err))
		return
	}
}
