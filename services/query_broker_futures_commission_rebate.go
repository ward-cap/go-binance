package binance

import (
	"context"
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

func (s *QueryBrokerFuturesCommissionRebateService) Do(ctx context.Context, opts ...RequestOption) (d []BrokerCommissionRebateResponse, _ error) {
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
	r.setParams(m)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &d)
	return d, err
}
