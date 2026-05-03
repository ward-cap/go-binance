package futures

import (
	"context"
	"net/http"
)

// PremiumIndexKlinesService list klines
type PremiumIndexKlinesService struct {
	c         *Client
	symbol    string
	interval  string
	limit     *int
	startTime *int64
	endTime   *int64
}

// Symbol sets symbol
func (piks *PremiumIndexKlinesService) Symbol(symbol string) *PremiumIndexKlinesService {
	piks.symbol = symbol
	return piks
}

// Interval set interval
func (piks *PremiumIndexKlinesService) Interval(interval string) *PremiumIndexKlinesService {
	piks.interval = interval
	return piks
}

// Limit set limit
func (piks *PremiumIndexKlinesService) Limit(limit int) *PremiumIndexKlinesService {
	piks.limit = &limit
	return piks
}

// StartTime set startTime
func (piks *PremiumIndexKlinesService) StartTime(startTime int64) *PremiumIndexKlinesService {
	piks.startTime = &startTime
	return piks
}

// EndTime set endTime
func (piks *PremiumIndexKlinesService) EndTime(endTime int64) *PremiumIndexKlinesService {
	piks.endTime = &endTime
	return piks
}

// Do send request
func (piks *PremiumIndexKlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*Kline, err error) {
	r := &request{
		service:  "PremiumIndexKlinesService",
		method:   http.MethodGet,
		endpoint: "/fapi/v1/premiumIndexKlines",
	}
	r.setParam("symbol", piks.symbol)
	r.setParam("interval", piks.interval)
	if piks.limit != nil {
		r.setParam("limit", *piks.limit)
	}
	if piks.startTime != nil {
		r.setParam("startTime", *piks.startTime)
	}
	if piks.endTime != nil {
		r.setParam("endTime", *piks.endTime)
	}
	data, _, err := piks.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*Kline{}, err
	}
	res, err = parsePriceKlines(data)
	if err != nil {
		return []*Kline{}, err
	}
	return res, nil
}
