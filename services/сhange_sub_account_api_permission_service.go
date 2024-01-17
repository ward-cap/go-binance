package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewChangeSubAccountApiPermissionService() *ChangeSubAccountApiPermissionService {
	return &ChangeSubAccountApiPermissionService{c: c}
}

type ChangeSubAccountApiPermissionService struct {
	c                *Client
	subAccountID     string
	subAccountApiKey string
	canTrade         bool
	futuresTrade     bool
	marginTrade      bool
}

func (s *ChangeSubAccountApiPermissionService) SubAccountID(subAccountID string) *ChangeSubAccountApiPermissionService {
	s.subAccountID = subAccountID
	return s
}

func (s *ChangeSubAccountApiPermissionService) CanTrade(canTrade bool) *ChangeSubAccountApiPermissionService {
	s.canTrade = canTrade
	return s
}

func (s *ChangeSubAccountApiPermissionService) FuturesTrade(futuresTrade bool) *ChangeSubAccountApiPermissionService {
	s.futuresTrade = futuresTrade
	return s
}

func (s *ChangeSubAccountApiPermissionService) MarginTrade(marginTrade bool) *ChangeSubAccountApiPermissionService {
	s.marginTrade = marginTrade
	return s
}

func (s *ChangeSubAccountApiPermissionService) SubAccountApiKey(subAccountApiKey string) *ChangeSubAccountApiPermissionService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *ChangeSubAccountApiPermissionService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/broker/subAccountApi/permission",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountID,
		"subAccountApiKey": s.subAccountApiKey,
		"canTrade":         s.canTrade,
		"futuresTrade":     s.futuresTrade,
		"marginTrade":      s.marginTrade,
	}
	r.setParams(m)
	_, err = s.c.callAPI(ctx, r, opts...)

	return err
}
