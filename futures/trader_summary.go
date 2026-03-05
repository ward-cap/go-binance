package futures

import (
	"context"
	"net/http"
)

type TraderSummaryService struct {
	c          *Client
	customerId *string
	startTime  *int64
	endTime    *int64
	_type      *int // 1:USDT Margined Futures, 2:COIN Margined Futures Default： USDT Margined Futures
	limit      *int // default 500, max 1000
}

func (s *TraderSummaryService) CustomerId(customerId string) *TraderSummaryService {
	s.customerId = &customerId
	return s
}

func (s *TraderSummaryService) Type(_type int) *TraderSummaryService {
	s._type = &_type
	return s
}

func (s *TraderSummaryService) Limit(limit int) *TraderSummaryService {
	s.limit = &limit
	return s
}

func (s *TraderSummaryService) StartTime(startTime int64) *TraderSummaryService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *TraderSummaryService) EndTime(endTime int64) *TraderSummaryService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *TraderSummaryService) Do(ctx context.Context, opts ...RequestOption) (res []byte, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/markPriceKlines",
	}
	if s.customerId != nil {
		r.setParam("customerId", *s.customerId)
	}
	if s._type != nil {
		r.setParam("type", s._type)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	return data, nil
}
