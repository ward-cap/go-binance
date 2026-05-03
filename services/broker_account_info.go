package binance

import (
	"context"
	"net/http"
)

type GetBrokerInfoService struct {
	c *Client
}

func (c *Client) NewGetBrokerInfoService() *GetBrokerInfoService {
	return &GetBrokerInfoService{c: c}
}

// Do send request
func (s *GetBrokerInfoService) Do(ctx context.Context, opts ...RequestOption) (res *GetBrokerInfoResponse, err error) {
	r := &request{
		service:  "GetBrokerInfoService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/info",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetBrokerInfoResponse)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
