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
	amount string,
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
	amount       string
}

func (s *SubAccountTransferFuturesService) Do(ctx context.Context, opts ...RequestOption) (res AccountTransferFuturesResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/transfer/futures",
		secType:  secTypeSigned,
	}
	m := params{
		"futuresType": s.futuresType,
		"asset":       s.asset,
		"amount":      s.amount,
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

type AccountTransferFuturesResponse struct {
	TxnId     int64  `json:"txnId"`
	ErrorData string `json:"errorData"`
}
