package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewQueryBrokerSpotCommissionRebateService(
	startTime uint,
	endTime uint,
	page uint,
	size uint,
) *QueryBrokerSpotCommissionRebateService {
	return &QueryBrokerSpotCommissionRebateService{
		c:         c,
		startTime: startTime,
		endTime:   endTime,
		page:      page,
		size:      size,
	}
}

type QueryBrokerSpotCommissionRebateService struct {
	c         *Client
	startTime uint
	endTime   uint
	page      uint
	size      uint
}

func (s *QueryBrokerSpotCommissionRebateService) Do(ctx context.Context, opts ...RequestOption) (d []BrokerCommissionRebateResponse, _ error) {
	r := &request{
		service:  "QueryBrokerSpotCommissionRebateService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/rebate/recentRecord",
		secType:  secTypeSigned,
	}
	m := params{
		"startTime": s.startTime,
		"endTime":   s.endTime,
		"page":      s.page,
		"size":      s.size,
	}
	r.setParams(m)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	err = jsonCodec.Unmarshal(data, &d)
	return d, err
}
