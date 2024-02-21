package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewUniversalTransferService(
	fromID, toID,
	fromAccountType, toAccountType,
	asset string,
	amount float64,
) *UniversalTransferService {
	return &UniversalTransferService{
		c:               c,
		fromID:          fromID,
		toID:            toID,
		fromAccountType: fromAccountType,
		toAccountType:   toAccountType,
		asset:           asset,
		amount:          amount,
	}
}

type UniversalTransferService struct {
	c                              *Client
	fromID, toID                   string
	fromAccountType, toAccountType string
	asset                          string
	amount                         float64
}

func (s *UniversalTransferService) Do(ctx context.Context, opts ...RequestOption) (def UpdateIPRestrictionSubAccountAPIKeyResponse, _ error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/universalTransfer",
		secType:  secTypeSigned,
	}
	m := params{
		"fromId":          s.fromID,
		"toId":            s.toID,
		"fromAccountType": s.fromAccountType,
		"toAccountType":   s.toAccountType,
		"asset":           s.asset,
		"amount":          s.amount,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return def, err
	}
	err = json.Unmarshal(data, &def)

	return def, err
}

type UniversalTransferServiceResd struct {
	Status     string   `json:"status"`
	IpList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
	ApiKey     string   `json:"apiKey"`
}
