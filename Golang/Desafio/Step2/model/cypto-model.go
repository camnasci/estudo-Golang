package model

type Cryptoresponse struct {
	Bitcoin struct {
		Usd int `json:"usd"`
	} `json:"bitcoin"`
}

type Response struct {
	ID      string  `json:"id"`
	Content Content `json:"content"`
	Partial bool    `json:"partial"`
}

type Content struct {
	Price    int    `json:"price"`
	Currency string `json:"currency"`
}
