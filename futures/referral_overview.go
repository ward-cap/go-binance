package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

type ReferralOverview struct {
	c     *Client
	_type *int // 1:USDT Margined Futures, 2:COIN Margined Futures Default：USDT Margined Futures
}

func (s *ReferralOverview) Type(_type int) *ReferralOverview {
	s._type = &_type
	return s
}

type ReferralOverviewResponse struct {
	BrokerId                  string `json:"brokerId"`
	NewTraderRebateCommission string `json:"newTraderRebateCommission"`
	OldTraderRebateCommission string `json:"oldTraderRebateCommission"`
	TotalTradeUser            int    `json:"totalTradeUser"`
	Unit                      string `json:"unit"`
	TotalTradeVol             string `json:"totalTradeVol"`
	TotalRebateVol            string `json:"totalRebateVol"`
	Time                      int64  `json:"time"`
}

// Do send request
func (s *ReferralOverview) Do(ctx context.Context, opts ...RequestOption) (res ReferralOverviewResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/apiReferral/overview",
		secType:  secTypeSigned,
	}

	if s._type != nil {
		r.setParam("type", s._type)
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return
}
