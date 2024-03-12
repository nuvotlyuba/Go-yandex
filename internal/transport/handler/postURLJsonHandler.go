package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"go.uber.org/zap"
)

func (h Handler) PostURLJsonHandler(w http.ResponseWriter, r *http.Request) {
	var req models.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Debug("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := h.service.CreateNewURL(req.URL)
	if err != nil {
		http.Error(w, error.Error(err), http.StatusBadRequest)
		return
	}

	resp := models.Response{
		Result: data.ShortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		logger.Debug("error encoding response", zap.Error(err))
		return
	}
}
