package api

type RateResponse struct {
	Rate float64 `json:"rate"`
}

type EmailResponse struct {
	Email string `json:"email"`
}

type BinanceResponse struct {
	OpenPrice string `json:"openPrice"`
}
