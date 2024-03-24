package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"
	"go.uber.org/zap"
)

func (h Handler) GetAllUrlsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAllUrs()

	if err != nil {
		http.Error(w, error.Error(err), http.StatusBadRequest)
		return
	}

	if data == nil {
		http.Error(w, error.Error(err), http.StatusNoContent)
		return
	}


	result := utils.ToResult(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	enc := json.NewEncoder(w)
	if err := enc.Encode(result); err != nil {
		logger.Debug("error encoding response", zap.Error(err))
		return
	}
}
