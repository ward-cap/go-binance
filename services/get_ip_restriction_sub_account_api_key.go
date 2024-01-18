package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewGetIPRestrictionSubAccountAPIKeyService(subAccountID, subAccountApiKey string) *GetIPRestrictionSubAccountAPIKeyService {
	return &GetIPRestrictionSubAccountAPIKeyService{
		c:                c,
		subAccountID:     subAccountID,
		subAccountApiKey: subAccountApiKey,
	}
}

type GetIPRestrictionSubAccountAPIKeyService struct {
	c                *Client
	subAccountID     string
	subAccountApiKey string
}

func (s *GetIPRestrictionSubAccountAPIKeyService) Do(ctx context.Context, opts ...RequestOption) (d GetIPRestrictionSubAccountAPIKeyResponse, _ error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/subAccountApi/ipRestriction",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":     s.subAccountID,
		"subAccountApiKey": s.subAccountApiKey,
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return d, err
	}
	err = json.Unmarshal(data, &d)

	return d, err
}

type GetIPRestrictionSubAccountAPIKeyResponse struct {
	SubAccountId string   `json:"subaccountId"`
	IpRestrict   string   `json:"ipRestrict"` // true/false
	Apikey       string   `json:"apikey"`
	IpList       []string `json:"ipList"`
	UpdateTime   int64    `json:"updateTime"`
}
