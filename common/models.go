package common

//go:generate easyjson -all models.go

// APIError define API error when response status is 4xx or 5xx
//
//easyjson:json
type APIError struct {
	Code    int64  `json:"code"`
	Message string `json:"msg"`
}

// PriceLevel is a common structure for bids and asks in the
// order book.
//
//easyjson:json
type PriceLevel struct {
	Price    string
	Quantity string
}
