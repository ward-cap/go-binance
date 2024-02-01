package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewSubAccountTransferHistoryFuturesService(
	subAccountId string,
	futuresType int64,
	clientTranId *string,
	startTime *int64,
	endTime *int64,
	page *int,
	limit *int,
) *SubAccountTransferHistoryFuturesService {
	return &SubAccountTransferHistoryFuturesService{
		c:            c,
		subAccountId: &subAccountId,
		futuresType:  futuresType,
		clientTranId: clientTranId,
		startTime:    startTime,
		endTime:      endTime,
		page:         page,
		limit:        limit,
	}
}

type SubAccountTransferHistoryFuturesService struct {
	c            *Client
	subAccountId *string
	futuresType  int64
	clientTranId *string
	startTime    *int64
	endTime      *int64
	page         *int
	limit        *int
}

type Transfer struct {
	From         string `json:"from"`
	To           string `json:"to"`
	Asset        string `json:"asset"`
	Qty          string `json:"qty"`
	TranId       string `json:"tranId"`
	ClientTranId string `json:"clientTranId"`
	Time         int64  `json:"time"`
}

type AccountTransferHistoryFuturesResponse struct {
	Success     bool       `json:"success"`
	FuturesType int64      `json:"futuresType"`
	Transfers   []Transfer `json:"transfers"`
}

func (s *SubAccountTransferHistoryFuturesService) Do(ctx context.Context, opts ...RequestOption) (res AccountTransferHistoryFuturesResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/transfer/futures",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": *s.subAccountId,
		"futuresType":  s.futuresType,
	}

	if s.clientTranId != nil {
		m["clientTranId"] = *s.clientTranId
	}
	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}
	if s.page != nil {
		m["page"] = *s.page
	}
	if s.limit != nil {
		m["limit"] = *s.limit
	}

	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}
