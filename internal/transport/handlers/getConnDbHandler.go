package handlers

import (
	"context"
	"net/http"
)

func (h Handlers) GetConnDBHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := h.store.DBRepo().Ping(ctx); err != nil {
		http.Error(w, error.Error(err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
