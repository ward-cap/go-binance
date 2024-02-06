package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewEnableFuturesForSubAccountService(subAccountID string) *EnableFuturesForSubAccountService {
	return &EnableFuturesForSubAccountService{
		c:            c,
		subAccountID: subAccountID,
	}
}

type EnableFuturesForSubAccountService struct {
	c            *Client
	subAccountID string
}

func (s *EnableFuturesForSubAccountService) Do(ctx context.Context, opts ...RequestOption) error {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/subAccount/futures",
		secType:  secTypeSigned,
	}

	m := params{
		"subAccountID": s.subAccountID,
		"futures":      true,
	}

	r.setParams(m)
	_, err := s.c.callAPI(ctx, r, opts...)

	return err
}
