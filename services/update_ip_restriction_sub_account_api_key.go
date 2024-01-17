package binance

import (
	"context"
	"net/http"
	"strings"
)

func (c *Client) NewUpdateIPRestrictionSubAccountAPIKeyService() *UpdateIPRestrictionSubAccountAPIKeyService {
	return &UpdateIPRestrictionSubAccountAPIKeyService{c: c}
}

type UpdateIPRestrictionSubAccountAPIKeyService struct {
	c                *Client
	subAccountID     string
	subAccountApiKey string
	ipAddress        []string
	status           string
}

func (s *UpdateIPRestrictionSubAccountAPIKeyService) SubAccountID(subAccountID string) *UpdateIPRestrictionSubAccountAPIKeyService {
	s.subAccountID = subAccountID
	return s
}

func (s *UpdateIPRestrictionSubAccountAPIKeyService) SetStatus(status string) *UpdateIPRestrictionSubAccountAPIKeyService {
	s.status = status
	return s
}

func (s *UpdateIPRestrictionSubAccountAPIKeyService) SetIP(ips []string) *UpdateIPRestrictionSubAccountAPIKeyService {
	s.ipAddress = ips
	return s
}

func (s *UpdateIPRestrictionSubAccountAPIKeyService) SubAccountApiKey(subAccountApiKey string) *UpdateIPRestrictionSubAccountAPIKeyService {
	s.subAccountApiKey = subAccountApiKey
	return s
}

func (s *UpdateIPRestrictionSubAccountAPIKeyService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v2/broker/subAccountApi/ipRestriction",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountID,
		"subAccountApiKey": s.subAccountApiKey,

		"status":    s.status,
		"ipAddress": strings.Join(s.ipAddress, ","),
	}
	r.setParams(m)
	_, err = s.c.callAPI(ctx, r, opts...)

	return err
}
