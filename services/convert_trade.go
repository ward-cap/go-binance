package binance

import (
	"context"
	"net/http"
)

type ConvertTradeHistoryService struct {
	c         *Client
	startTime int64
	endTime   int64
	limit     *int32
}

// StartTime set startTime
func (s *ConvertTradeHistoryService) StartTime(startTime int64) *ConvertTradeHistoryService {
	s.startTime = startTime
	return s
}

// EndTime set endTime
func (s *ConvertTradeHistoryService) EndTime(endTime int64) *ConvertTradeHistoryService {
	s.endTime = endTime
	return s
}

// Limit set limit
func (s *ConvertTradeHistoryService) Limit(limit int32) *ConvertTradeHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ConvertTradeHistoryService) Do(ctx context.Context, opts ...RequestOption) (*ConvertTradeHistory, error) {
	r := &request{
		service:  "ConvertTradeHistoryService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/convert/tradeFlow",
		secType:  secTypeSigned,
	}
	r.setParam("startTime", s.startTime)
	r.setParam("endTime", s.endTime)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := ConvertTradeHistory{}
	if err = jsonCodec.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
