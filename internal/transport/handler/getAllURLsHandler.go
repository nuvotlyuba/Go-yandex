package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/configs"
	"github.com/nuvotlyuba/Go-yandex/internal/jwt"
	"github.com/nuvotlyuba/Go-yandex/internal/utils"

	"github.com/nuvotlyuba/Go-yandex/internal/app/apiserver/logger"
	"github.com/nuvotlyuba/Go-yandex/internal/store"
	"go.uber.org/zap"
)

func (h Handler) GetAllURLsHandler(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	token, err := r.Cookie("token")
	// если куки нет, то создаем ее
	if errors.Is(err, http.ErrNoCookie) {

		cookie, err := utils.PrepareCookie()
		if err != nil {
			http.Error(w, error.Error(err), http.StatusBadRequest)
			return
		}
		http.SetCookie(w, cookie)
		status = http.StatusUnauthorized
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		return
	} else {
		// если кука есть, то проводим проверку
		userID := jwt.GetUserID(token.Value)
		if userID != configs.UserID {
			// если кука есть, но она не проходит проверку подлинности
			cookie, err := utils.PrepareCookie()
			if err != nil {
				http.Error(w, error.Error(err), http.StatusBadRequest)
				return
			}
			http.SetCookie(w, cookie)
			status = http.StatusUnauthorized

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			return
		}
		// если нет созданных коротких ссылок
		if errors.Is(err, store.ErrNoContent) {
			status = http.StatusNoContent
		}
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

// ErrNoCookie is returned by Request's Cookie method when a cookie is not found.
// var ErrNoCookie = errors.New("http: named cookie not present")
// http.ErrNoCookie
