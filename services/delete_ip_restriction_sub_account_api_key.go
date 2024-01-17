package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewDeleteIPRestrictionSubAccountAPIKeyService() *DeleteIPRestrictionSubAccountAPIKeyService {
	return &DeleteIPRestrictionSubAccountAPIKeyService{c: c}
}

type DeleteIPRestrictionSubAccountAPIKeyService struct {
	c                *Client
	subAccountID     string
	subAccountApiKey string
}

func (s *DeleteIPRestrictionSubAccountAPIKeyService) SubAccountID(subAccountID string) *DeleteIPRestrictionSubAccountAPIKeyService {
	s.subAccountID = subAccountID
	return s
}

func (s *DeleteIPRestrictionSubAccountAPIKeyService) SubAccountApiKey(subAccountApiKey string) *DeleteIPRestrictionSubAccountAPIKeyService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *DeleteIPRestrictionSubAccountAPIKeyService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/broker/subAccountApi/ipRestriction/ipList",
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
