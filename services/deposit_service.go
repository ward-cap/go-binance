package binance

import (
	"context"
	"net/http"
)

// GetDepositsAddressService retrieves the details of a deposit address.
//
// See https://binance-docs.github.io/apidocs/spot/en/#deposit-address-supporting-network-user_data
type GetDepositsAddressService struct {
	c       *Client
	coin    string
	network *string
}

// Coin sets the coin parameter (MANDATORY).
func (s *GetDepositsAddressService) Coin(coin string) *GetDepositsAddressService {
	s.coin = coin
	return s
}

// Network sets the network parameter.
func (s *GetDepositsAddressService) Network(network string) *GetDepositsAddressService {
	s.network = &network
	return s
}

// Do sends the request.
func (s *GetDepositsAddressService) Do(ctx context.Context) (*GetDepositAddressResponse, error) {
	r := &request{
		service:  "GetDepositsAddressService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/deposit/address",
		secType:  secTypeSigned,
	}
	r.setParam("coin", s.coin)
	if s.network != nil {
		r.setParam("network", *s.network)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &GetDepositAddressResponse{}
	if err := jsonCodec.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}
