package futures

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ward-cap/go-binance/common"
)

type SignTradeFiService struct {
	c *Client
}

func (s *SignTradeFiService) Do(ctx context.Context, opts ...RequestOption) (res common.APIError, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/fapi/v1/stock/contract",
		secType:  secTypeSigned,
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return common.APIError{}, err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return common.APIError{}, err
	}
	return res, nil
}
