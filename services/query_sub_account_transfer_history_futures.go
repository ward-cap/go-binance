package binance

import (
	"context"
	"net/http"
	"time"
)

func (c *Client) NewSubAccountTransferHistoryFuturesService(
	subAccountId string,
	futuresType int64,
	clientTranId *string,
	startTime *time.Time,
	endTime *time.Time,
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
	startTime    *time.Time
	endTime      *time.Time
	page         *int
	limit        *int
}

func (s *SubAccountTransferHistoryFuturesService) Do(ctx context.Context, opts ...RequestOption) (res AccountTransferHistoryFuturesResponse, err error) {
	r := &request{
		service:  "SubAccountTransferHistoryFuturesService",
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
		m["startTime"] = s.startTime.UnixMilli()
	}
	if s.endTime != nil {
		m["endTime"] = s.endTime.UnixMilli()
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

	err = jsonCodec.Unmarshal(data, &res)

	return res, err
}
