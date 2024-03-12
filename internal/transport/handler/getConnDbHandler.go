package handler

import (
	"context"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"go.uber.org/zap"
)

func (h Handler) GetConnDBHandler(w http.ResponseWriter, r *http.Request) {

	logger.Info("GetConnDBHandler", zap.String("start", "work"))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := h.service.PingDB(ctx); err != nil {
		http.Error(w, error.Error(err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
