package binance

import (
	"context"
	"github.com/ward-cap/go-binance/common"
	"net/http"
)

// ListBookTickersService list best price/qty on the order book for a symbol or symbols
type ListBookTickersService struct {
	c      *Client
	symbol *string
}

// Symbol set symbol
func (s *ListBookTickersService) Symbol(symbol string) *ListBookTickersService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *ListBookTickersService) Do(ctx context.Context, opts ...RequestOption) (res []*BookTicker, err error) {
	r := &request{
		service:  "ListBookTickersService",
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker/bookTicker",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	data = common.ToJSONList(data)
	if err != nil {
		return []*BookTicker{}, err
	}
	res = make([]*BookTicker, 0)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return []*BookTicker{}, err
	}
	return res, nil
}

// ListPricesService list latest price for a symbol or symbols
type ListPricesService struct {
	c       *Client
	symbol  *string
	symbols []string
}

// Symbol set symbol
func (s *ListPricesService) Symbol(symbol string) *ListPricesService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *ListPricesService) Do(ctx context.Context, opts ...RequestOption) (res []*SymbolPrice, err error) {
	r := &request{
		service:  "ListPricesService",
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker/price",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	} else if s.symbols != nil {
		s, _ := jsonCodec.Marshal(s.symbols)
		r.setParam("symbols", string(s))
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*SymbolPrice{}, err
	}
	data = common.ToJSONList(data)
	res = make([]*SymbolPrice, 0)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return []*SymbolPrice{}, err
	}
	return res, nil
}

// ListPriceChangeStatsService show stats of price change in last 24 hours for all symbols
type ListPriceChangeStatsService struct {
	c       *Client
	symbol  *string
	symbols []string
}

// Symbol set symbol
func (s *ListPriceChangeStatsService) Symbol(symbol string) *ListPriceChangeStatsService {
	s.symbol = &symbol
	return s
}

// Symbols set symbols
func (s *ListPriceChangeStatsService) Symbols(symbols []string) *ListPriceChangeStatsService {
	s.symbols = symbols
	return s
}

// Symbols set symbols
func (s *ListPricesService) Symbols(symbols []string) *ListPricesService {
	s.symbols = symbols
	return s
}

// Do send request
func (s *ListPriceChangeStatsService) Do(ctx context.Context, opts ...RequestOption) (res []*PriceChangeStats, err error) {
	r := &request{
		service:  "ListPriceChangeStatsService",
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker/24hr",
	}

	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	} else if s.symbols != nil {
		r.setParam("symbols", s.symbols)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	data = common.ToJSONList(data)
	res = make([]*PriceChangeStats, 0)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AveragePriceService show current average price for a symbol
type AveragePriceService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *AveragePriceService) Symbol(symbol string) *AveragePriceService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *AveragePriceService) Do(ctx context.Context, opts ...RequestOption) (res *AvgPrice, err error) {
	r := &request{
		service:  "AveragePriceService",
		method:   http.MethodGet,
		endpoint: "/api/v3/avgPrice",
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return res, err
	}
	res = new(AvgPrice)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type ListSymbolTickerService struct {
	c          *Client
	symbol     *string
	symbols    []string
	windowSize *string
}

func (s *ListSymbolTickerService) Symbol(symbol string) *ListSymbolTickerService {
	s.symbol = &symbol
	return s
}

func (s *ListSymbolTickerService) Symbols(symbols []string) *ListSymbolTickerService {
	s.symbols = symbols
	return s
}

// Defaults to 1d if no parameter provided
//
// Supported windowSize values:
//
// - 1m,2m....59m for minutes
//
// - 1h, 2h....23h - for hours
//
// - 1d...7d - for days
//
// Units cannot be combined (e.g. 1d2h is not allowed).
//
// Reference: https://binance-docs.github.io/apidocs/spot/en/#rolling-window-price-change-statistics
func (s *ListSymbolTickerService) WindowSize(windowSize string) *ListSymbolTickerService {
	s.windowSize = &windowSize
	return s
}

func (s *ListSymbolTickerService) Do(ctx context.Context, opts ...RequestOption) (res []*SymbolTicker, err error) {
	r := &request{
		service:  "ListSymbolTickerService",
		method:   http.MethodGet,
		endpoint: "/api/v3/ticker",
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	} else if s.symbols != nil {
		s, _ := jsonCodec.Marshal(s.symbols)
		r.setParam("symbols", string(s))
	}

	if s.windowSize != nil {
		r.setParam("windowSize", *s.windowSize)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	data = common.ToJSONList(data)
	if err != nil {
		return []*SymbolTicker{}, err
	}
	res = make([]*SymbolTicker, 0)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return []*SymbolTicker{}, err
	}
	return res, nil
}
