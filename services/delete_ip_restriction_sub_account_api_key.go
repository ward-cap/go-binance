package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewDeleteIPRestrictionSubAccountAPIKeyService(subAccountID, subAccountApiKey, ipAddress string) *DeleteIPRestrictionSubAccountAPIKeyService {
	return &DeleteIPRestrictionSubAccountAPIKeyService{
		c:                c,
		subAccountID:     subAccountID,
		subAccountApiKey: subAccountApiKey,
		ipAddress:        ipAddress,
	}
}

type DeleteIPRestrictionSubAccountAPIKeyService struct {
	c                *Client
	subAccountID     string
	subAccountApiKey string
	ipAddress        string
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
