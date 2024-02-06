package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewChangeSubAccountApiUniversalTransferService(
	subAccountID string,
	subAccountApiKey string,
	canUniversalTransfer bool,
) *ChangeSubAccountApiUniversalTransferService {
	return &ChangeSubAccountApiUniversalTransferService{
		c:                    c,
		subAccountID:         subAccountID,
		subAccountApiKey:     subAccountApiKey,
		canUniversalTransfer: canUniversalTransfer,
	}
}

type ChangeSubAccountApiUniversalTransferService struct {
	c                    *Client
	subAccountID         string
	subAccountApiKey     string
	canUniversalTransfer bool
}

func (s *ChangeSubAccountApiUniversalTransferService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/subAccountApi/permission/universalTransfer",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId":         s.subAccountID,
		"subAccountApiKey":     s.subAccountApiKey,
		"canUniversalTransfer": s.canUniversalTransfer,
	}
	r.setParams(m)
	_, err = s.c.callAPI(ctx, r, opts...)

	return err
}
