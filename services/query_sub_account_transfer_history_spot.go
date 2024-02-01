package binance

import (
	"context"
	"net/http"
	"strconv"
)

func (c *Client) NewSubAccountTransferHistorySpotService() *SubAccountTransferHistorySpotService {
	return &SubAccountTransferHistorySpotService{c: c}
}

type SubAccountTransferHistorySpotService struct {
	c             *Client
	fromId        *string
	toId          *string
	clientTranId  *string
	showAllStatus *bool
	startTime     *int64
	endTime       *int64
	page          *int
	limit         *int
}

func (s *SubAccountTransferHistorySpotService) Do(ctx context.Context, opts ...RequestOption) (res []AccountTransferHistorySpotResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/transfer",
		secType:  secTypeSigned,
	}
	m := params{}

	if s.fromId != nil {
		m["fromId"] = *s.fromId
	}
	if s.toId != nil {
		m["toId"] = *s.toId
	}
	if s.clientTranId != nil {
		m["clientTranId"] = *s.clientTranId
	}
	if s.showAllStatus != nil {
		m["showAllStatus"] = strconv.FormatBool(*s.showAllStatus)
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
		return nil, err
	}

	err = json.Unmarshal(data, &res)

	return res, err
}

type AccountTransferHistorySpotResponse struct {
	FromId       string `json:"fromId"`
	ToId         string `json:"toId"`
	Asset        string `json:"asset"`
	Qty          string `json:"qty"`
	Time         int64  `json:"time"`
	TxnId        string `json:"txnId"`
	ClientTranId string `json:"clientTranId"`
	Status       string `json:"status"`
}
