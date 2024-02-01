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
	amount float64,
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
	amount       float64
}

func (s *SubAccountTransferSpotService) Do(ctx context.Context, opts ...RequestOption) (res AccountTransferSpotResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/transfer",
		secType:  secTypeSigned,
	}
	m := params{
		"fromId":       s.fromId,
		"toId":         s.toId,
		"clientTranId": s.clientTranId,
		"asset":        s.asset,
		"amount":       s.amount,
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
	TxnId        string `json:"txnId"`
	ClientTranId string `json:"clientTranId"`
}
