package models

import "github.com/google/uuid"

type URLItem struct {
	ID      string `json:"id"`
	URL     string `json:"Url"`
}

type URL struct {
	UUID        uuid.UUID `json:"uuid"`
	ShortURL    string    `json:"short_url"`
	OriginalURL string    `json:"original_url"`
}

type URLData []URLItem

type RequestBody struct {
	URL string     `json:"url"`
}
type Response struct {
	Result string `json:"result"`
}
