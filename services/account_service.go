package binance

import (
	"context"
	"net/http"
)

// GetAccountService get account info
type GetAccountService struct {
	c *Client
}

// Do send request
func (s *GetAccountService) Do(ctx context.Context, opts ...RequestOption) (res *Account, err error) {
	r := &request{
		service:  "GetAccountService",
		method:   http.MethodGet,
		endpoint: "/api/v3/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(Account)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetAccountSnapshotService all account orders; active, canceled, or filled
type GetAccountSnapshotService struct {
	c           *Client
	accountType string
	startTime   *int64
	endTime     *int64
	limit       *int
}

// Type set account type ("SPOT", "MARGIN", "FUTURES")
func (s *GetAccountSnapshotService) Type(accountType string) *GetAccountSnapshotService {
	s.accountType = accountType
	return s
}

// StartTime set starttime
func (s *GetAccountSnapshotService) StartTime(startTime int64) *GetAccountSnapshotService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *GetAccountSnapshotService) EndTime(endTime int64) *GetAccountSnapshotService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *GetAccountSnapshotService) Limit(limit int) *GetAccountSnapshotService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetAccountSnapshotService) Do(ctx context.Context, opts ...RequestOption) (res *Snapshot, err error) {
	r := &request{
		service:  "GetAccountSnapshotService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/accountSnapshot",
		secType:  secTypeSigned,
	}
	r.setParam("type", s.accountType)

	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &Snapshot{}, err
	}
	res = new(Snapshot)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return &Snapshot{}, err
	}
	return res, nil
}

// Do send request
func (s *GetAPIKeyPermission) Do(ctx context.Context, opts ...RequestOption) (res *APIKeyPermission, err error) {
	r := &request{
		service:  "GetAPIKeyPermission",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/account/apiRestrictions",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(APIKeyPermission)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
