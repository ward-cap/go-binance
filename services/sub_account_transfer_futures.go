package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewSubAccountTransferFuturesService(
	fromId *string,
	toId *string,
	futuresType int64,
	clientTranId *string,
	asset string,
	amount float64,
) *SubAccountTransferFuturesService {
	return &SubAccountTransferFuturesService{
		c:            c,
		fromId:       fromId,
		toId:         toId,
		futuresType:  futuresType,
		clientTranId: clientTranId,
		asset:        asset,
		amount:       amount,
	}
}

type SubAccountTransferFuturesService struct {
	c            *Client
	fromId       *string
	toId         *string
	futuresType  int64
	clientTranId *string
	asset        string
	amount       float64
}

func (s *SubAccountTransferFuturesService) Do(ctx context.Context, opts ...RequestOption) (res AccountTransferFuturesResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/transfer/futures",
		secType:  secTypeSigned,
	}
	m := params{
		"fromId":       s.fromId,
		"toId":         s.toId,
		"futuresType":  s.futuresType,
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

type AccountTransferFuturesResponse struct {
	Success      bool   `json:"success"`
	TxnId        string `json:"txnId"`
	ClientTranId string `json:"clientTranId"`
}
