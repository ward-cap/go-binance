package binance

import (
	"context"
	"github.com/shopspring/decimal"
	"net/http"
)

// Deposit represents a single deposit entry.
type Deposit struct {
	Amount        decimal.Decimal `json:"amount"`
	Coin          string          `json:"coin"`
	Network       string          `json:"network"`
	Status        int             `json:"status"`
	Address       string          `json:"address"`
	AddressTag    string          `json:"addressTag"`
	TxID          string          `json:"txId"`
	InsertTime    int64           `json:"insertTime"`
	TransferType  int64           `json:"transferType"`
	UnlockConfirm int64           `json:"unlockConfirm"`
	ConfirmTimes  string          `json:"confirmTimes"`
}

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
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetDepositAddressResponse represents a response from GetDepositsAddressService.
type GetDepositAddressResponse struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
	Coin    string `json:"coin"`
	URL     string `json:"url"`
}
