package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/store"
	"go.uber.org/zap"
)

func (h Handler) GetAllURLsHandler(w http.ResponseWriter, _ *http.Request) {


	data, err := h.service.GetAllURLs()

	status := http.StatusOK
	if errors.Is(err, store.ErrNoContent) {
		status = http.StatusNoContent
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err != nil {
		http.Error(w, error.Error(err), http.StatusBadRequest)
		return
	}

	// result := utils.ToResult(data)

	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		logger.Debug("error encoding response", zap.Error(err))
		return
	}
}
