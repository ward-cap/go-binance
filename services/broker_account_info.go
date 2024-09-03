package binance

import (
	"context"
	"github.com/shopspring/decimal"
	"net/http"
)

type GetBrokerInfoService struct {
	c *Client
}

func (c *Client) NewGetBrokerInfoService() *GetBrokerInfoService {
	return &GetBrokerInfoService{c: c}
}

// Do send request
func (s *GetBrokerInfoService) Do(ctx context.Context, opts ...RequestOption) (res *GetBrokerInfoResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/info",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetBrokerInfoResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetBrokerInfoResponse struct {
	MaxMakerCommission decimal.Decimal `json:"maxMakerCommission"`
	MinMakerCommission decimal.Decimal `json:"minMakerCommission"`
	MaxTakerCommission decimal.Decimal `json:"maxTakerCommission"`
	MinTakerCommission decimal.Decimal `json:"minTakerCommission"`

	SubAccountQty    int64 `json:"subAccountQty"`
	MaxSubAccountQty int64 `json:"maxSubAccountQty"`
}
