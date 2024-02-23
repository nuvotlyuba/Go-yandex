package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
	"go.uber.org/zap"
)

func (s Store) PostURLJsonHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		logger.Log.Debug("got request with bad content-type", zap.String("content-type", contentType))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req models.RequestBody
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		logger.Log.Debug("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := repo.CreateNewID(req.URL)
	shortenURL := utils.StringURL(configs.BaseURL, id)

	resp := models.Response{
		Result: shortenURL,
	}
	logger.Log.Info("Создание короткой ссылки")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
	logger.Log.Debug("error encoding response", zap.Error(err))
		return
	}
}
