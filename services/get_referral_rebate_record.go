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

func (s *GetReferralRebateRecord) StartTime(startTime int64) *GetReferralRebateRecord {
	s.startTime = &startTime
	return s
}

func (s *GetReferralRebateRecord) EndTime(endTime int64) *GetReferralRebateRecord {
	s.endTime = &endTime
	return s
}

func (s *GetReferralRebateRecord) Limit(limit int) *GetReferralRebateRecord {
	s.limit = &limit
	return s
}

func (s *GetReferralRebateRecord) Do(ctx context.Context, opts ...RequestOption) (res []*ReferralRebateRecordResponse, err error) {
	r := &request{
		service:  "GetReferralRebateRecord",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/apiReferral/rebate/recentRecord",
		secType:  secTypeSigned,
	}

	m := params{}
	if s.limit != nil {
		m["limit"] = *s.limit
	}
	if s.startTime != nil {
		m["startTime"] = *s.startTime
	}
	if s.endTime != nil {
		m["endTime"] = *s.endTime
	}

	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = jsonCodec.Unmarshal(data, &res)
	return
}
