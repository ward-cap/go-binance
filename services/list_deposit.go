package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewListDepositsService(
	coin *string,
	status *int,
	startTime *int64,
	endTime *int64,
	offset *int,
	limit *int,
	txId *string,
) *ListDepositsService {
	return &ListDepositsService{
		c:         c,
		coin:      coin,
		status:    status,
		startTime: startTime,
		endTime:   endTime,
		offset:    offset,
		limit:     limit,
		txId:      txId,
	}
}

type ListDepositsService struct {
	c         *Client
	coin      *string
	status    *int
	startTime *int64
	endTime   *int64
	offset    *int
	limit     *int
	txId      *string
}

func (s *ListDepositsService) Do(ctx context.Context) (res []*Deposit, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/deposit/hisrec",
		secType:  secTypeSigned,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.txId != nil {
		r.setParam("txId", *s.txId)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res = make([]*Deposit, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}
