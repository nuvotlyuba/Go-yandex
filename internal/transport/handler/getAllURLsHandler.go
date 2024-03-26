package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/jwt"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/store"
	"go.uber.org/zap"
)

func (h Handler) GetAllURLsHandler(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("token")
	if err != nil {
		http.Error(w, error.Error(err), http.StatusBadRequest)
		return
	}

	status := http.StatusOK
	userID := jwt.GetUserID(token.Value)
	if userID != configs.UserID {
		// http.SetCookie(w, cookie)
		status = http.StatusUnauthorized
	}
	if errors.Is(err, store.ErrNoContent) {
		status = http.StatusNoContent
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	data, err := h.service.GetAllURLs()

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
