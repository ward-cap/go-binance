package binance

import (
	"context"
	"net/http"
)

// FiatDepositWithdrawHistoryService retrieve the fiat deposit/withdraw history
type FiatDepositWithdrawHistoryService struct {
	c               *Client
	transactionType TransactionType
	beginTime       *int64
	endTime         *int64
	page            *int32
	rows            *int32
}

// TransactionType set transactionType
func (s *FiatDepositWithdrawHistoryService) TransactionType(transactionType TransactionType) *FiatDepositWithdrawHistoryService {
	s.transactionType = transactionType
	return s
}

// BeginTime set beginTime
func (s *FiatDepositWithdrawHistoryService) BeginTime(beginTime int64) *FiatDepositWithdrawHistoryService {
	s.beginTime = &beginTime
	return s
}

// EndTime set endTime
func (s *FiatDepositWithdrawHistoryService) EndTime(endTime int64) *FiatDepositWithdrawHistoryService {
	s.endTime = &endTime
	return s
}

// Page set page
func (s *FiatDepositWithdrawHistoryService) Page(page int32) *FiatDepositWithdrawHistoryService {
	s.page = &page
	return s
}

// Rows set rows
func (s *FiatDepositWithdrawHistoryService) Rows(rows int32) *FiatDepositWithdrawHistoryService {
	s.rows = &rows
	return s
}

// Do send request
func (s *FiatDepositWithdrawHistoryService) Do(ctx context.Context, opts ...RequestOption) (*FiatDepositWithdrawHistory, error) {
	r := &request{
		service:  "FiatDepositWithdrawHistoryService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/fiat/orders",
		secType:  secTypeSigned,
	}
	r.setParam("transactionType", s.transactionType)
	if s.beginTime != nil {
		r.setParam("beginTime", *s.beginTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
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
	res := FiatDepositWithdrawHistory{}
	if err = jsonCodec.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// FiatPaymentsHistoryService retrieve the fiat payments history
type FiatPaymentsHistoryService struct {
	c               *Client
	transactionType TransactionType
	beginTime       *int64
	endTime         *int64
	page            *int32
	rows            *int32
}

// TransactionType set transactionType
func (s *FiatPaymentsHistoryService) TransactionType(transactionType TransactionType) *FiatPaymentsHistoryService {
	s.transactionType = transactionType
	return s
}

// BeginTime set beginTime
func (s *FiatPaymentsHistoryService) BeginTime(beginTime int64) *FiatPaymentsHistoryService {
	s.beginTime = &beginTime
	return s
}

// EndTime set endTime
func (s *FiatPaymentsHistoryService) EndTime(endTime int64) *FiatPaymentsHistoryService {
	s.endTime = &endTime
	return s
}

// Page set page
func (s *FiatPaymentsHistoryService) Page(page int32) *FiatPaymentsHistoryService {
	s.page = &page
	return s
}

// Rows set rows
func (s *FiatPaymentsHistoryService) Rows(rows int32) *FiatPaymentsHistoryService {
	s.rows = &rows
	return s
}

// Do send request
func (s *FiatPaymentsHistoryService) Do(ctx context.Context, opts ...RequestOption) (*FiatPaymentsHistory, error) {
	r := &request{
		service:  "FiatPaymentsHistoryService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/fiat/payments",
		secType:  secTypeSigned,
	}
	r.setParam("transactionType", s.transactionType)
	if s.beginTime != nil {
		r.setParam("beginTime", *s.beginTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
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
	res := FiatPaymentsHistory{}
	if err = jsonCodec.Unmarshal(data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
