package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewCommissionSubAccountRateService() *CommissionSubAccountRateService {
	return &CommissionSubAccountRateService{c: c}
}

type CommissionSubAccountRateService struct {
	c      *Client
	symbol string
	subAcc string
}

// Symbol set symbol
func (s *CommissionSubAccountRateService) Symbol(symbol string) *CommissionSubAccountRateService {
	s.symbol = symbol
	return s
}

func (s *CommissionSubAccountRateService) SubAcc(acc string) *CommissionSubAccountRateService {
	s.subAcc = acc
	return s
}

// Do send request
func (s *CommissionSubAccountRateService) Do(
	ctx context.Context,
	opts ...RequestOption,
) (res []CommissionRateSubAccount, err error) {
	r := &request{
		service:  "CommissionSubAccountRateService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/subAccountApi/commission/futures",
		secType:  secTypeSigned,
	}
	r.setParam("subAccountId", s.subAcc)
	if s.symbol != "" {
		r.setParam("symbol", s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
