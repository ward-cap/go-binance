package futures

import (
	"context"
	"net/http"
)

// ContinuousKlinesService list klines
type ContinuousKlinesService struct {
	c            *Client
	pair         string
	contractType string
	interval     string
	limit        *int
	startTime    *int64
	endTime      *int64
}

// pair set pair
func (s *ContinuousKlinesService) Pair(pair string) *ContinuousKlinesService {
	s.pair = pair
	return s
}

// contractType set contractType
func (s *ContinuousKlinesService) ContractType(contractType string) *ContinuousKlinesService {
	s.contractType = contractType
	return s
}

// Interval set interval
func (s *ContinuousKlinesService) Interval(interval string) *ContinuousKlinesService {
	s.interval = interval
	return s
}

// Limit set limit
func (s *ContinuousKlinesService) Limit(limit int) *ContinuousKlinesService {
	s.limit = &limit
	return s
}

// StartTime set startTime
func (s *ContinuousKlinesService) StartTime(startTime int64) *ContinuousKlinesService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ContinuousKlinesService) EndTime(endTime int64) *ContinuousKlinesService {
	s.endTime = &endTime
	return s
}

// Do send request
func (s *ContinuousKlinesService) Do(ctx context.Context, opts ...RequestOption) (res []*ContinuousKline, err error) {
	r := &request{
		service:  "ContinuousKlinesService",
		method:   http.MethodGet,
		endpoint: "/fapi/v1/continuousKlines",
	}
	r.setParam("pair", s.pair)
	r.setParam("contractType", s.contractType)
	r.setParam("interval", s.interval)
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
		return []*ContinuousKline{}, err
	}
	res, err = parseContinuousKlines(data)
	if err != nil {
		return []*ContinuousKline{}, err
	}
	return res, nil
}
