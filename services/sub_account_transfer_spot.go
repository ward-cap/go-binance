package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewSubAccountTransferSpotService(
	fromId *string,
	toId *string,
	clientTranId *string,
	asset string,
	amount string,
) *SubAccountTransferSpotService {
	return &SubAccountTransferSpotService{
		c:            c,
		fromId:       fromId,
		toId:         toId,
		clientTranId: clientTranId,
		asset:        asset,
		amount:       amount,
	}
}

type SubAccountTransferSpotService struct {
	c            *Client
	fromId       *string
	toId         *string
	clientTranId *string
	asset        string
	amount       string
}

func (s *SubAccountTransferSpotService) Do(ctx context.Context, opts ...RequestOption) (res AccountTransferSpotResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/transfer",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
	}

	if s.fromId != nil {
		m["fromId"] = *s.fromId
	}
	if s.toId != nil {
		m["toId"] = *s.toId
	}
	if s.clientTranId != nil {
		m["clientTranId"] = *s.clientTranId
	}

	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

type AccountTransferSpotResponse struct {
	TxnId     int64  `json:"txnId"`
	ErrorData string `json:"errorData"`
}
