package futures

import (
	"context"
	"encoding/json"
	"net/http"
)

func (c *Client) NewQueryBrokerFuturesCommissionRebateService(
	futuresType uint,
	startTime uint,
	endTime uint,
	page uint,
	size uint,
) *QueryBrokerFuturesCommissionRebateService {
	return &QueryBrokerFuturesCommissionRebateService{
		c:           c,
		futuresType: futuresType,
		startTime:   startTime,
		endTime:     endTime,
		page:        page,
		size:        size,
	}
}

type QueryBrokerFuturesCommissionRebateService struct {
	c           *Client
	futuresType uint
	startTime   uint
	endTime     uint
	page        uint
	size        uint
}

type BrokerFuturesCommissionRebateResponse struct {
	SubaccountId string `json:"subaccountId"`
	Income       string `json:"income"`
	Asset        string `json:"asset"`
	Symbol       string `json:"symbol"`
	TradeId      int    `json:"tradeId"`
	Time         int64  `json:"time"`
	Status       int    `json:"status"`
}

func (s *QueryBrokerFuturesCommissionRebateService) Do(ctx context.Context, opts ...RequestOption) (d []BrokerFuturesCommissionRebateResponse, _ error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/rebate/futures/recentRecord",
		secType:  secTypeSigned,
	}
	m := params{
		"futuresType": s.futuresType,
		"startTime":   s.startTime,
		"endTime":     s.endTime,
		"page":        s.page,
		"size":        s.size,
	}
	r.setFormParams(m)

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &d)
	return d, err
}
