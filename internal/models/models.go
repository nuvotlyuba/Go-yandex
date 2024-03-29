package models

type BatchURL []*URL

type URL struct {
	ID          string `json:"id"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type RequestBody struct {
	URL string `json:"url"`
}
type Response struct {
	Result string `json:"result"`
}

type RequestBatch []RequestItem
type ResponseBatch []ResponseItem

type RequestItem struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
}
type ResponseItem struct {
	CorrelationID string `json:"correlation_id"`
	ShortURL      string `json:"short_url"`
}
