package binance

import (
	"context"
	"net/http"
)

type ListingService struct {
	c *Client
}

type GetListingResponse struct {
	OpenTime int64    `json:"openTime"`
	Symbols  []string `json:"symbols"`
}

func (s *ListingService) Do(ctx context.Context) (res []GetListingResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/spot/open-symbol-list",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &res)

	return res, err
}
