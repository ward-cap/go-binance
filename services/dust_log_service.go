/*******************************************************************************
** @Author:					Thomas Bouder <Tbouder>
** @Email:					Tbouder@protonmail.com
** @Date:					Friday 19 June 2020 - 18:51:45
** @Filename:				0_dust_service.go
**
** @Last modified by:		Tbouder
*******************************************************************************/

package binance

import (
	"context"
	"net/http"
)

// ListDustLogService fetch small amounts of assets exchanged versus BNB
// See https://binance-docs.github.io/apidocs/spot/en/#dustlog-user_data
type ListDustLogService struct {
	c         *Client
	startTime *int64
	endTime   *int64
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListDustLogService) StartTime(startTime int64) *ListDustLogService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListDustLogService) EndTime(endTime int64) *ListDustLogService {
	s.endTime = &endTime
	return s
}

// Do sends the request.
func (s *ListDustLogService) Do(ctx context.Context) (withdraws *DustResult, err error) {
	r := &request{
		service:  "ListDustLogService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/asset/dribblet",
		secType:  secTypeSigned,
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res := new(DustResult)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res, nil
}

// DustTransferService convert dust assets to BNB.
// See https://binance-docs.github.io/apidocs/spot/en/#dust-transfer-user_data
type DustTransferService struct {
	c     *Client
	asset []string
}

// Asset set asset.
func (s *DustTransferService) Asset(asset []string) *DustTransferService {
	s.asset = asset
	return s
}

// Do sends the request.
func (s *DustTransferService) Do(ctx context.Context) (withdraws *DustTransferResponse, err error) {
	r := &request{
		service:  "DustTransferService",
		method:   http.MethodPost,
		endpoint: "/sapi/v1/asset/dust",
		secType:  secTypeSigned,
	}
	for _, a := range s.asset {
		r.addParam("asset", a)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res := new(DustTransferResponse)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res, nil
}

// ListDustService get list of dust to BNB.
// See https://binance-docs.github.io/apidocs/spot/en/#get-assets-that-can-be-converted-into-bnb-user_data
type ListDustService struct {
	c *Client
}

// Do sends the request.
func (s *ListDustService) Do(ctx context.Context) (res *ListDustResponse, err error) {
	r := &request{
		service:  "ListDustService",
		method:   http.MethodPost,
		endpoint: "/sapi/v1/asset/dust-btc",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	res = new(ListDustResponse)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return
	}
	return res, nil
}
