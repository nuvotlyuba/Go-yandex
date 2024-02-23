package models

type URLItem struct {
	ID      string `json:"id"`
	URL     string `json:"Url"`
}

type URLData []URLItem

type RequestBody struct {
	URL string     `json:"url"`
}
type Response struct {
	Result string `json:"result"`
}
