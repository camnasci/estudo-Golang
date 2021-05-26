package model

type CurrencyResponse map[string]ContentResponse

type ContentResponse struct {
	Usd float64 `json:"usd"`
}

type Response struct {
	ID      string   `json:"id"`
	Content *Content `json:"content,omitempty"`
	Partial bool     `json:"partial"`
}

type Content struct {
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
}
