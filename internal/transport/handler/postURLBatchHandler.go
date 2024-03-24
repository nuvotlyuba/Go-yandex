package handler

import (
	"encoding/json"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/models"
	"go.uber.org/zap"
)

func (h *Handler) PostURLBatchHandler(w http.ResponseWriter, r *http.Request) {
	req := make([]models.RequestItem, 0)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.Debug("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := h.service.CreateBatchURL(req)
	if err != nil {
		http.Error(w, error.Error(err), http.StatusBadRequest)
	}

	result := make([]models.ResponseItem, 0)
	for _, v := range data {
		item := models.ResponseItem{
			CorrelationID: v.ID,
			ShortURL:      v.ShortURL,
		}
		result = append(result, item)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	enc := json.NewEncoder(w)
	if err := enc.Encode(result); err != nil {
		logger.Debug("error encoding response", zap.Error(err))
		return
	}
}
