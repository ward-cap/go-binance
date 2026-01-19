package futures

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type OpenAlgoOrdersService struct {
	c *Client
}

// Do send request
func (s *OpenAlgoOrdersService) Do(ctx context.Context, opts ...RequestOption) (res []*AlgoOrders, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/fapi/v1/openAlgoOrders",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &res)
	return
}

type AlgoOrders struct {
	AlgoId                  int         `json:"algoId"`
	ClientAlgoId            string      `json:"clientAlgoId"`
	AlgoType                string      `json:"algoType"`
	OrderType               string      `json:"orderType"`
	Symbol                  string      `json:"symbol"`
	Side                    string      `json:"side"`
	PositionSide            string      `json:"positionSide"`
	TimeInForce             string      `json:"timeInForce"`
	Quantity                string      `json:"quantity"`
	AlgoStatus              string      `json:"algoStatus"`
	ActualOrderId           string      `json:"actualOrderId"`
	ActualPrice             string      `json:"actualPrice"`
	TriggerPrice            string      `json:"triggerPrice"`
	Price                   string      `json:"price"`
	IcebergQuantity         interface{} `json:"icebergQuantity"`
	TpTriggerPrice          string      `json:"tpTriggerPrice"`
	TpPrice                 string      `json:"tpPrice"`
	SlTriggerPrice          string      `json:"slTriggerPrice"`
	SlPrice                 string      `json:"slPrice"`
	TpOrderType             string      `json:"tpOrderType"`
	SelfTradePreventionMode string      `json:"selfTradePreventionMode"`
	WorkingType             string      `json:"workingType"`
	PriceMatch              string      `json:"priceMatch"`
	ClosePosition           bool        `json:"closePosition"`
	PriceProtect            bool        `json:"priceProtect"`
	ReduceOnly              bool        `json:"reduceOnly"`
	CreateTime              int64       `json:"createTime"`
	UpdateTime              int64       `json:"updateTime"`
	TriggerTime             int         `json:"triggerTime"`
	GoodTillDate            int         `json:"goodTillDate"`
}

type CloseAlgoOrderResponse struct {
	AlgoId       int    `json:"algoId"`
	ClientAlgoId string `json:"clientAlgoId"`
	Code         string `json:"code"`
	Msg          string `json:"msg"`
}

type CloseAlgoOrdersService struct {
	c *Client

	// only 1 param is required
	algoID int64  // real type is LONG
	symbol string // closes all by symbol
}

func (s *CloseAlgoOrdersService) SetAlgoID(algoId int64) *CloseAlgoOrdersService {
	s.algoID = algoId
	return s
}

func (s *CloseAlgoOrdersService) SetSymbol(symbol string) *CloseAlgoOrdersService {
	s.symbol = symbol
	return s
}

func (s *CloseAlgoOrdersService) Do(ctx context.Context, opts ...RequestOption) (res CloseAlgoOrderResponse, err error) {
	r := &request{
		method:  http.MethodDelete,
		secType: secTypeSigned,
	}
	if (s.algoID == 0) == (s.symbol == "") {
		return res, fmt.Errorf("either algoID or symbol must be set, but not both")
	}

	if s.algoID != 0 {
		r.setFormParam("algoId", s.algoID)
		r.endpoint = "/fapi/v1/algoOrder"
	}
	if s.symbol != "" {
		r.setFormParam("symbol", s.symbol)
		r.endpoint = "fapi/v1/algoOpenOrders"
	}

	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &res)
	return
}
