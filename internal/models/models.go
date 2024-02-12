package models

type URLItem struct {
	ID      string `json:"id"`
	LongURL string `json:"longUrl"`
}

type URLData []URLItem
