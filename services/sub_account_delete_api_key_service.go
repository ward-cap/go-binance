package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewSubAccountDeleteApiKeyService() *SubAccountDeleteApiKeyService {
	return &SubAccountDeleteApiKeyService{c: c}
}

type SubAccountDeleteApiKeyService struct {
	c                *Client
	subAccountID     string
	subAccountApiKey string
}

func (s *SubAccountDeleteApiKeyService) SubAccountID(subAccountID string) *SubAccountDeleteApiKeyService {
	s.subAccountID = subAccountID
	return s
}

func (s *SubAccountDeleteApiKeyService) SubAccountApiKey(subAccountApiKey string) *SubAccountDeleteApiKeyService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *SubAccountDeleteApiKeyService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/broker/subAccountApi",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountID,
		"subAccountApiKey": s.subAccountApiKey,
	}
	r.setParams(m)
	_, err = s.c.callAPI(ctx, r, opts...)

	return err
}
