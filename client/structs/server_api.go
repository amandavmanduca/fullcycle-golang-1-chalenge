package structs

import "encoding/json"

type ServerApiResponse struct {
	Status  string          `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type ExchangeRateResponse struct {
	ID  string      `json:"id"`
	Bid json.Number `json:"bid"`
}
