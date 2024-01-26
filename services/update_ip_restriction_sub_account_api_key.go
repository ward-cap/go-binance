package binance

import (
	"context"
	"net/http"
	"strings"
)

func (c *Client) NewUpdateIPRestrictionSubAccountAPIKeyService(
	email, subAccountApiKey, status string,
	ipAddress []string,
) *UpdateIPRestrictionSubAccountAPIKeyService {
	return &UpdateIPRestrictionSubAccountAPIKeyService{
		c:                c,
		email:            email,
		subAccountApiKey: subAccountApiKey,
		ipAddress:        ipAddress,
		status:           status,
	}
}

type UpdateIPRestrictionSubAccountAPIKeyService struct {
	c                *Client
	email            string
	subAccountApiKey string
	ipAddress        []string
	status           string
}

func (s *UpdateIPRestrictionSubAccountAPIKeyService) Do(ctx context.Context, opts ...RequestOption) (def UpdateIPRestrictionSubAccountAPIKeyResponse, _ error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v2/sub-account/subAccountApi/ipRestriction",
		secType:  secTypeSigned,
	}
	m := params{
		"email":            s.email,
		"subAccountApiKey": s.subAccountApiKey,

		"status":    s.status,
		"ipAddress": strings.Join(s.ipAddress, ","),
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return def, err
	}
	err = json.Unmarshal(data, &def)

	return def, err
}

type UpdateIPRestrictionSubAccountAPIKeyResponse struct {
	Status     string   `json:"status"`
	IpList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
	ApiKey     string   `json:"apiKey"`
}
