package binance

import (
	"context"
	"net/http"
)

// C2CTradeHistoryService retrieve c2c trade history
type C2CTradeHistoryService struct {
	c              *Client
	tradeType      SideType
	startTimestamp *int64
	endTimestamp   *int64
	page           *int32
	rows           *int32
}

// TransactionType set transaction type
func (s *C2CTradeHistoryService) TradeType(tradeType SideType) *C2CTradeHistoryService {
	s.tradeType = tradeType
	return s
}

// BeginTime set beginTime
func (s *C2CTradeHistoryService) StartTimestamp(startTimestamp int64) *C2CTradeHistoryService {
	s.startTimestamp = &startTimestamp
	return s
}

// EndTime set endTime
func (s *C2CTradeHistoryService) EndTime(endTimestamp int64) *C2CTradeHistoryService {
	s.endTimestamp = &endTimestamp
	return s
}

// Page set page
func (s *C2CTradeHistoryService) Page(page int32) *C2CTradeHistoryService {
	s.page = &page
	return s
}

// Rows set rows
func (s *C2CTradeHistoryService) Rows(rows int32) *C2CTradeHistoryService {
	s.rows = &rows
	return s
}

// Do send request
func (s *C2CTradeHistoryService) Do(ctx context.Context, opts ...RequestOption) (*C2CTradeHistory, error) {
	r := &request{
		service:  "C2CTradeHistoryService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/c2c/orderMatch/listUserOrderHistory",
		secType:  secTypeSigned,
	}
	r.setParam("tradeType", s.tradeType)
	if s.startTimestamp != nil {
		r.setParam("startTimestamp", *s.startTimestamp)
	}
	if s.endTimestamp != nil {
		r.setParam("endTime", *s.endTimestamp)
	}
	if s.page != nil {
		r.setParam("page", *s.page)
	}
	if s.rows != nil {
		r.setParam("rows", *s.rows)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := C2CTradeHistory{}
	if err = jsonCodec.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
