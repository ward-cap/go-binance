package binance

import (
	"context"
	"net/http"
)

// GetAssetDetailService fetches all asset detail.
//
// See https://binance-docs.github.io/apidocs/spot/en/#asset-detail-user_data
type GetAssetDetailService struct {
	c     *Client
	asset *string
}

// Asset sets the asset parameter.
func (s *GetAssetDetailService) Asset(asset string) *GetAssetDetailService {
	s.asset = &asset
	return s
}

// Do sends the request.
func (s *GetAssetDetailService) Do(ctx context.Context) (res map[string]AssetDetail, err error) {
	r := &request{
		service:  "GetAssetDetailService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/asset/assetDetail",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res = make(map[string]AssetDetail)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

type GetAllCoinsInfoService struct {
	c     *Client
	asset *string
}

// Do send request
func (s *GetAllCoinsInfoService) Do(ctx context.Context) (res []*CoinInfo, err error) {
	r := &request{
		service:  "GetAllCoinsInfoService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/config/getall",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return []*CoinInfo{}, err
	}
	res = make([]*CoinInfo, 0)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return []*CoinInfo{}, err
	}
	return res, nil
}

// GetUserAssetService Get user assets
// See https://binance-docs.github.io/apidocs/spot/en/#user-asset-user_data
type GetUserAssetService struct {
	c                *Client
	asset            *string
	needBtcValuation bool
}

func (s *GetUserAssetService) Asset(asset string) *GetUserAssetService {
	s.asset = &asset
	return s
}

func (s *GetUserAssetService) NeedBtcValuation(val bool) *GetUserAssetService {
	s.needBtcValuation = val
	return s
}

func (s *GetUserAssetService) Do(ctx context.Context) (res []UserAssetRecord, err error) {
	r := &request{
		service:  "GetUserAssetService",
		method:   http.MethodPost,
		endpoint: "/sapi/v3/asset/getUserAsset",
		secType:  secTypeSigned,
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.needBtcValuation {
		r.setParam("needBtcValuation", s.needBtcValuation)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	err = jsonCodec.Unmarshal(data, &res)
	return
}
