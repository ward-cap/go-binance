package binance

import (
	"context"
	"github.com/ward-cap/go-binance/common"
	"net/http"
)

// ExchangeInfoService exchange info service
type ExchangeInfoService struct {
	c           *Client
	symbol      string
	symbols     []string
	permissions []string
}

// Symbol set symbol
func (s *ExchangeInfoService) Symbol(symbol string) *ExchangeInfoService {
	s.symbol = symbol
	return s
}

// Symbols set symbol
func (s *ExchangeInfoService) Symbols(symbols ...string) *ExchangeInfoService {
	s.symbols = symbols

	return s
}

// Permissions set permission
func (s *ExchangeInfoService) Permissions(permissions ...string) *ExchangeInfoService {
	s.permissions = permissions

	return s
}

// Do send request
func (s *ExchangeInfoService) Do(ctx context.Context, opts ...RequestOption) (res *ExchangeInfo, err error) {
	r := &request{
		service:  "ExchangeInfoService",
		method:   http.MethodGet,
		endpoint: "/api/v3/exchangeInfo",
		secType:  secTypeNone,
	}
	m := params{}
	if s.symbol != "" {
		m["symbol"] = s.symbol
	}
	if len(s.symbols) != 0 {
		m["symbols"] = s.symbols
	}
	if len(s.permissions) != 0 {
		m["permissions"] = s.permissions
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(ExchangeInfo)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// LotSizeFilter return lot size filter of symbol
func (s *Symbol) LotSizeFilter() *LotSizeFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeLotSize) {
			f := &LotSizeFilter{}
			if i, ok := filter["maxQty"]; ok {
				f.MaxQuantity = i.(string)
			}
			if i, ok := filter["minQty"]; ok {
				f.MinQuantity = i.(string)
			}
			if i, ok := filter["stepSize"]; ok {
				f.StepSize = i.(string)
			}
			return f
		}
	}
	return nil
}

// PriceFilter return price filter of symbol
func (s *Symbol) PriceFilter() *PriceFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypePriceFilter) {
			f := &PriceFilter{}
			if i, ok := filter["maxPrice"]; ok {
				f.MaxPrice = i.(string)
			}
			if i, ok := filter["minPrice"]; ok {
				f.MinPrice = i.(string)
			}
			if i, ok := filter["tickSize"]; ok {
				f.TickSize = i.(string)
			}
			return f
		}
	}
	return nil
}

// PercentPriceBySideFilter return percent price filter of symbol
func (s *Symbol) PercentPriceBySideFilter() *PercentPriceBySideFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypePercentPriceBySide) {
			f := &PercentPriceBySideFilter{}
			if i, ok := filter["avgPriceMins"]; ok {
				if apm, okk := common.ToInt(i); okk == nil {
					f.AveragePriceMins = apm
				}
			}
			if i, ok := filter["bidMultiplierUp"]; ok {
				f.BidMultiplierUp = i.(string)
			}
			if i, ok := filter["bidMultiplierDown"]; ok {
				f.BidMultiplierDown = i.(string)
			}
			if i, ok := filter["askMultiplierUp"]; ok {
				f.AskMultiplierUp = i.(string)
			}
			if i, ok := filter["askMultiplierDown"]; ok {
				f.AskMultiplierDown = i.(string)
			}
			return f
		}
	}
	return nil
}

// NotionalFilter return notional filter of symbol
func (s *Symbol) NotionalFilter() *NotionalFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeNotional) {
			f := &NotionalFilter{}
			if i, ok := filter["minNotional"]; ok {
				f.MinNotional = i.(string)
			}
			if i, ok := filter["applyMinToMarket"]; ok {
				f.ApplyMinToMarket = i.(bool)
			}
			if i, ok := filter["maxNotional"]; ok {
				f.MaxNotional = i.(string)
			}
			if i, ok := filter["applyMaxToMarket"]; ok {
				f.ApplyMaxToMarket = i.(bool)
			}
			if i, ok := filter["avgPriceMins"]; ok {
				if apm, okk := common.ToInt(i); okk == nil {
					f.AvgPriceMins = apm
				}
			}
			return f
		}
	}
	return nil
}

// IcebergPartsFilter return iceberg part filter of symbol
func (s *Symbol) IcebergPartsFilter() *IcebergPartsFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeIcebergParts) {
			f := &IcebergPartsFilter{}
			if i, ok := filter["limit"]; ok {
				if limit, okk := common.ToInt(i); okk == nil {
					f.Limit = limit
				}
			}
			return f
		}
	}
	return nil
}

// MarketLotSizeFilter return market lot size filter of symbol
func (s *Symbol) MarketLotSizeFilter() *MarketLotSizeFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeMarketLotSize) {
			f := &MarketLotSizeFilter{}
			if i, ok := filter["maxQty"]; ok {
				f.MaxQuantity = i.(string)
			}
			if i, ok := filter["minQty"]; ok {
				f.MinQuantity = i.(string)
			}
			if i, ok := filter["stepSize"]; ok {
				f.StepSize = i.(string)
			}
			return f
		}
	}
	return nil
}

// For specific meanings, please refer to the type definition MaxNumOrders
func (s *Symbol) MaxNumOrdersFilter() *MaxNumOrdersFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeMaxNumOrders) {
			f := &MaxNumOrdersFilter{}
			if i, ok := filter["maxNumOrders"]; ok {
				if mno, okk := common.ToInt(i); okk == nil {
					f.MaxNumOrders = mno
				}
			}
			return f
		}
	}
	return nil
}

// MaxNumAlgoOrdersFilter return max num algo orders filter of symbol
func (s *Symbol) MaxNumAlgoOrdersFilter() *MaxNumAlgoOrdersFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeMaxNumAlgoOrders) {
			f := &MaxNumAlgoOrdersFilter{}
			if i, ok := filter["maxNumAlgoOrders"]; ok {
				if mnao, okk := common.ToInt(i); okk == nil {
					f.MaxNumAlgoOrders = mnao
				}
			}
			return f
		}
	}
	return nil
}

// For specific meanings, please refer to the type definition TrailingDeltaFilter
func (s *Symbol) TrailingDeltaFilter() *TrailingDeltaFilter {
	for _, filter := range s.Filters {
		if filter["filterType"].(string) == string(SymbolFilterTypeTrailingDelta) {
			f := &TrailingDeltaFilter{}
			if i, ok := filter["minTrailingAboveDelta"]; ok {
				if mtad, okk := common.ToInt(i); okk == nil {
					f.MinTrailingAboveDelta = mtad
				}
			}
			if i, ok := filter["maxTrailingAboveDelta"]; ok {
				if mtad, okk := common.ToInt(i); okk == nil {
					f.MaxTrailingAboveDelta = mtad
				}
			}
			if i, ok := filter["minTrailingBelowDelta"]; ok {
				if mtbd, okk := common.ToInt(i); okk == nil {
					f.MinTrailingBelowDelta = mtbd
				}
			}
			if i, ok := filter["maxTrailingBelowDelta"]; ok {
				if mtbd, okk := common.ToInt(i); okk == nil {
					f.MaxTrailingBelowDelta = mtbd
				}
			}
			return f
		}
	}
	return nil
}
