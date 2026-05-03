package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewUniversalTransferService(
	fromID, toID *string,
	fromAccountType, toAccountType,
	asset string,
	amount string,
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
	fromID, toID                   *string
	fromAccountType, toAccountType string
	asset                          string
	amount                         string
}

func (s *UniversalTransferService) Do(ctx context.Context, opts ...RequestOption) (def UniversalTransferServiceResponse, _ error) {
	r := &request{
		service:  "UniversalTransferService",
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/universalTransfer",
		secType:  secTypeSigned,
	}
	m := params{
		"fromAccountType": s.fromAccountType,
		"toAccountType":   s.toAccountType,
		"asset":           s.asset,
		"amount":          s.amount,
	}
	if s.fromID != nil {
		m["fromId"] = *s.fromID
	}
	if s.toID != nil {
		m["toId"] = *s.toID
	}

	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return def, err
	}
	err = jsonCodec.Unmarshal(data, &def)

	return def, err
}
