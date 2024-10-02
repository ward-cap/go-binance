package binance

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetAllAssetsService struct {
	c *Client
}

type GetAllAssetsResponse struct {
	Data []struct {
		AssetCode  string `json:"assetCode"`
		LogoUrl    string `json:"logoUrl"`
		AssetDigit int64  `json:"assetDigit"`
		Trading    bool   `json:"trading"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (s *GetAllAssetsService) Do(ctx context.Context) (res GetAllAssetsResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/bapi/asset/v2/public/asset/asset/get-all-asset",
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	//res = make([]*Deposit, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}

	if !res.Success {
		err = errors.New("failed to get the assets")
	}
	return res, nil
}
