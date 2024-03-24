package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"github.com/nuvotlyuba/Go-yandex/internal/store"
	"go.uber.org/zap"
)

func (h Handler) PostURLJsonHandler(w http.ResponseWriter, r *http.Request) {
	var req models.RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Debug("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := h.service.CreateURL(req.URL)

	resp := models.Response{
		Result: data.ShortURL,
	}

	status := http.StatusCreated
	if errors.Is(err, store.ErrConflict) {
		status = http.StatusConflict
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err != nil && err != store.ErrConflict {
		http.Error(w, error.Error(err), http.StatusBadRequest)
		return
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		logger.Debug("error encoding response", zap.Error(err))
		return
	}

}
