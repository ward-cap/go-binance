package binance

import (
	"context"
	"net/http"
)

func (c *Client) NewGetReferralRebateRecord() *GetReferralRebateRecord {
	return &GetReferralRebateRecord{
		c: c,
	}
}

type GetReferralRebateRecord struct {
	c         *Client
	startTime int64 // required
	endTime   int64 // required
	limit     int   // max 500
}

func (s *GetReferralRebateRecord) StartTime(startTime int64) *GetReferralRebateRecord {
	s.startTime = startTime
	return s
}

func (s *GetReferralRebateRecord) EndTime(endTime int64) *GetReferralRebateRecord {
	s.endTime = endTime
	return s
}

func (s *GetReferralRebateRecord) Limit(limit int) *GetReferralRebateRecord {
	s.limit = limit
	return s
}

func (s *GetReferralRebateRecord) Do(ctx context.Context, opts ...RequestOption) ([]byte, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/apiReferral/rebate/recentRecord",
		secType:  secTypeSigned,
	}

	m := params{
		//"subAccountId": s.subAccountID,
		//"futures":      true,
	}

	r.setParams(m)
	return s.c.callAPI(ctx, r, opts...)

	//return err
}
