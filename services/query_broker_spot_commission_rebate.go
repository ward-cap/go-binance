package binance

import (
	"context"
	"net/http"

	"github.com/shopspring/decimal"
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

type BrokerCommissionRebateResponse struct {
	SubAccountID string          `json:"subaccountId"`
	Income       decimal.Decimal `json:"income"`
	Asset        string          `json:"asset"`
	Symbol       string          `json:"symbol"`
	Time         int64           `json:"time"`
	TradeId      int             `json:"tradeId"`
	Status       int             `json:"status"`
}

func (s *QueryBrokerSpotCommissionRebateService) Do(ctx context.Context, opts ...RequestOption) (d []BrokerCommissionRebateResponse, _ error) {
	r := &request{
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
	//r.setFormParams(m)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &d)
	return d, err
}
