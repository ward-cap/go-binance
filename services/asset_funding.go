package binance

import (
	"context"
	"net/http"
)

type GetAssetFundingDetailService struct {
	c                *Client
	asset            *string
	needBtcValuation *bool
}

func (s *GetAssetFundingDetailService) Asset(asset string) *GetAssetFundingDetailService {
	s.asset = &asset
	return s
}

func (s *GetAssetFundingDetailService) NeedBtcValuation(needBtcValuation bool) *GetAssetFundingDetailService {
	s.needBtcValuation = &needBtcValuation
	return s
}

// Do sends the request.
func (s *GetAssetFundingDetailService) Do(ctx context.Context) (res []AssetFundingResponse, err error) {
	r := &request{
		service:  "GetAssetFundingDetailService",
		method:   http.MethodPost,
		endpoint: "/sapi/v1/asset/get-funding-asset",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setFormParam("asset", *s.asset)
	}
	if s.needBtcValuation != nil {
		r.setFormParam("needBtcValuation", *s.needBtcValuation)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	//res = make(map[string]AssetDetail)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}
