package handler

import (
	"net/http"

	"github.com/nuvotlyuba/Go-yandex/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

type Handlers interface {
	GetURLHandler(w http.ResponseWriter, r *http.Request)
	PostURLHandler(w http.ResponseWriter, r *http.Request)
	PostURLJsonHandler(w http.ResponseWriter, r *http.Request)
	GetConnDBHandler(w http.ResponseWriter, r *http.Request)
	PostURLBatchHandler(w http.ResponseWriter, r *http.Request)
	GetAllUrlsHandler(w http.ResponseWriter, r *http.Request)
}
