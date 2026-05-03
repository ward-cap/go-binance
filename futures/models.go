package futures

import (
	"github.com/shopspring/decimal"
	"github.com/ward-cap/go-binance/common"
)

//go:generate easyjson -all models.go

// Ask is a type alias for PriceLevel.
type Ask = common.PriceLevel

// Bid is a type alias for PriceLevel.
type Bid = common.PriceLevel

// Account define account info
//
//easyjson:json
type Account struct {
	Assets                      []*AccountAsset    `json:"assets"`
	FeeTier                     int                `json:"feeTier"`
	CanTrade                    bool               `json:"canTrade"`
	CanDeposit                  bool               `json:"canDeposit"`
	CanWithdraw                 bool               `json:"canWithdraw"`
	UpdateTime                  int64              `json:"updateTime"`
	MultiAssetsMargin           bool               `json:"multiAssetsMargin"`
	TotalInitialMargin          string             `json:"totalInitialMargin"`
	TotalMaintMargin            string             `json:"totalMaintMargin"`
	TotalWalletBalance          string             `json:"totalWalletBalance"`
	TotalUnrealizedProfit       string             `json:"totalUnrealizedProfit"`
	TotalMarginBalance          string             `json:"totalMarginBalance"`
	TotalPositionInitialMargin  string             `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin string             `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     string             `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             string             `json:"totalCrossUnPnl"`
	AvailableBalance            string             `json:"availableBalance"`
	MaxWithdrawAmount           string             `json:"maxWithdrawAmount"`
	Positions                   []*AccountPosition `json:"positions"`
}

// AccountAsset define account asset
//
//easyjson:json
type AccountAsset struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintMargin            string `json:"maintMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
	CrossWalletBalance     string `json:"crossWalletBalance"`
	CrossUnPnl             string `json:"crossUnPnl"`
	AvailableBalance       string `json:"availableBalance"`
	MarginAvailable        bool   `json:"marginAvailable"`
	UpdateTime             int64  `json:"updateTime"`
}

// AccountPosition define account position
//
//easyjson:json
type AccountPosition struct {
	Isolated               bool             `json:"isolated"`
	Leverage               decimal.Decimal  `json:"leverage"`
	InitialMargin          string           `json:"initialMargin"`
	MaintMargin            string           `json:"maintMargin"`
	OpenOrderInitialMargin string           `json:"openOrderInitialMargin"`
	PositionInitialMargin  string           `json:"positionInitialMargin"`
	Symbol                 string           `json:"symbol"`
	UnrealizedProfit       string           `json:"unrealizedProfit"`
	EntryPrice             string           `json:"entryPrice"`
	MaxNotional            string           `json:"maxNotional"`
	PositionSide           PositionSideType `json:"positionSide"`
	PositionAmt            string           `json:"positionAmt"`
	Notional               string           `json:"notional"`
	BidNotional            string           `json:"bidNotional"`
	AskNotional            string           `json:"askNotional"`
	UpdateTime             int64            `json:"updateTime"`
}

// AccountTrade define account trade
//
//easyjson:json
type AccountTrade struct {
	Buyer           bool             `json:"buyer"`
	Commission      decimal.Decimal  `json:"commission"`
	CommissionAsset string           `json:"commissionAsset"`
	ID              int64            `json:"id"`
	Maker           bool             `json:"maker"`
	OrderID         int64            `json:"orderId"`
	Price           decimal.Decimal  `json:"price"`
	Quantity        decimal.Decimal  `json:"qty"`
	QuoteQuantity   string           `json:"quoteQty"`
	RealizedPnl     string           `json:"realizedPnl"`
	Side            SideType         `json:"side"`
	PositionSide    PositionSideType `json:"positionSide"`
	Symbol          string           `json:"symbol"`
	Time            int64            `json:"time"`
}

// AggTrade define aggregate trade info
//
//easyjson:json
type AggTrade struct {
	AggTradeID   int64  `json:"a"`
	Price        string `json:"p"`
	Quantity     string `json:"q"`
	FirstTradeID int64  `json:"f"`
	LastTradeID  int64  `json:"l"`
	Timestamp    int64  `json:"T"`
	IsBuyerMaker bool   `json:"m"`
}

//easyjson:json
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

// Balance define user balance of your account
//
//easyjson:json
type Balance struct {
	AccountAlias       string          `json:"accountAlias"`
	Asset              string          `json:"asset"`
	Balance            decimal.Decimal `json:"balance"`
	CrossWalletBalance decimal.Decimal `json:"crossWalletBalance"`
	CrossUnPnl         decimal.Decimal `json:"crossUnPnl"`
	AvailableBalance   decimal.Decimal `json:"availableBalance"`
	MaxWithdrawAmount  decimal.Decimal `json:"maxWithdrawAmount"`
}

// BookTicker define book ticker info
//
//easyjson:json
type BookTicker struct {
	Symbol      string `json:"symbol"`
	BidPrice    string `json:"bidPrice"`
	BidQuantity string `json:"bidQty"`
	AskPrice    string `json:"askPrice"`
	AskQuantity string `json:"askQty"`
}

// Bracket define the bracket
//
//easyjson:json
type Bracket struct {
	InitialLeverage decimal.Decimal `json:"initialLeverage"`
	NotionalCap     decimal.Decimal `json:"notionalCap"`

	//Bracket          int             `json:"bracket"`
	//NotionalFloor    float64         `json:"notionalFloor"`
	//MaintMarginRatio float64         `json:"maintMarginRatio"`
	//Cum              float64         `json:"cum"`
}

// CancelOrderResponse define response of canceling order
//
//easyjson:json
type CancelOrderResponse struct {
	ClientOrderID    string           `json:"clientOrderId"`
	CumQuantity      string           `json:"cumQty"`
	CumQuote         string           `json:"cumQuote"`
	ExecutedQuantity string           `json:"executedQty"`
	OrderID          int64            `json:"orderId"`
	OrigQuantity     string           `json:"origQty"`
	Price            string           `json:"price"`
	ReduceOnly       bool             `json:"reduceOnly"`
	Side             SideType         `json:"side"`
	Status           OrderStatusType  `json:"status"`
	StopPrice        string           `json:"stopPrice"`
	Symbol           string           `json:"symbol"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	UpdateTime       int64            `json:"updateTime"`
	WorkingType      WorkingType      `json:"workingType"`
	ActivatePrice    string           `json:"activatePrice"`
	PriceRate        string           `json:"priceRate"`
	OrigType         string           `json:"origType"`
	PositionSide     PositionSideType `json:"positionSide"`
	PriceProtect     bool             `json:"priceProtect"`
}

//easyjson:json
type CloseAlgoOrderResponse struct {
	AlgoId       int    `json:"algoId"`
	ClientAlgoId string `json:"clientAlgoId"`
	Code         string `json:"code"`
	Msg          string `json:"msg"`
}

// Commission Rate
//
//easyjson:json
type CommissionRate struct {
	Symbol              string `json:"symbol"`
	MakerCommissionRate string `json:"makerCommissionRate"`
	TakerCommissionRate string `json:"takerCommissionRate"`
}

// ContinuousKline define ContinuousKline info
//
//easyjson:json
type ContinuousKline struct {
	OpenTime                 int64  `json:"openTime"`
	Open                     string `json:"open"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	Close                    string `json:"close"`
	Volume                   string `json:"volume"`
	CloseTime                int64  `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	TradeNum                 int64  `json:"tradeNum"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
}

//easyjson:json
type CreateBatchOrdersResponse struct {
	Orders []*Order
}

// CreateOrderResponse define create order response
//
//easyjson:json
type CreateOrderResponse struct {
	Symbol            string           `json:"symbol"`
	OrderID           int64            `json:"orderId"`
	ClientOrderID     string           `json:"clientOrderId"`
	Price             string           `json:"price"`
	OrigQuantity      string           `json:"origQty"`
	ExecutedQuantity  string           `json:"executedQty"`
	CumQuote          string           `json:"cumQuote"`
	ReduceOnly        bool             `json:"reduceOnly"`
	Status            OrderStatusType  `json:"status"`
	StopPrice         string           `json:"stopPrice"`
	TimeInForce       TimeInForceType  `json:"timeInForce"`
	Type              OrderType        `json:"type"`
	Side              SideType         `json:"side"`
	UpdateTime        int64            `json:"updateTime"`
	WorkingType       WorkingType      `json:"workingType"`
	ActivatePrice     string           `json:"activatePrice"`
	PriceRate         string           `json:"priceRate"`
	AvgPrice          string           `json:"avgPrice"`
	PositionSide      PositionSideType `json:"positionSide"`
	ClosePosition     bool             `json:"closePosition"`
	PriceProtect      bool             `json:"priceProtect"`
	RateLimitOrder10s string           `json:"rateLimitOrder10s,omitempty"`
	RateLimitOrder1m  string           `json:"rateLimitOrder1m,omitempty"`
}

// DepthResponse define depth info with bids and asks
//
//easyjson:json
type DepthResponse struct {
	LastUpdateID int64 `json:"lastUpdateId"`
	Time         int64 `json:"E"`
	TradeTime    int64 `json:"T"`
	Bids         []Bid `json:"bids"`
	Asks         []Ask `json:"asks"`
}

// ExchangeInfo exchange info
//
//easyjson:json
type ExchangeInfo struct {
	Timezone        string        `json:"timezone"`
	ServerTime      int64         `json:"serverTime"`
	RateLimits      []RateLimit   `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Symbols         []Symbol      `json:"symbols"`
}

// FundingRate define funding rate of mark price
//
//easyjson:json
type FundingRate struct {
	Symbol      string `json:"symbol"`
	FundingRate string `json:"fundingRate"`
	FundingTime int64  `json:"fundingTime"`
	Time        int64  `json:"time"`
}

// IncomeHistory define position margin history info
//
//easyjson:json
type IncomeHistory struct {
	Asset      string `json:"asset"`
	Income     string `json:"income"`
	IncomeType string `json:"incomeType"`
	Info       string `json:"info"`
	Symbol     string `json:"symbol"`
	Time       int64  `json:"time"`
	TranID     int64  `json:"tranId"`
	TradeID    string `json:"tradeId"`
}

// Kline define kline info
//
//easyjson:json
type Kline struct {
	OpenTime                 int64  `json:"openTime"`
	Open                     string `json:"open"`
	High                     string `json:"high"`
	Low                      string `json:"low"`
	Close                    string `json:"close"`
	Volume                   string `json:"volume"`
	CloseTime                int64  `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	TradeNum                 int64  `json:"tradeNum"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
}

// LeverageBracket define the leverage bracket
//
//easyjson:json
type LeverageBracket struct {
	Symbol   string    `json:"symbol"`
	Brackets []Bracket `json:"brackets"`
}

// LiquidationOrder define liquidation order
//
//easyjson:json
type LiquidationOrder struct {
	Symbol           string          `json:"symbol"`
	Price            string          `json:"price"`
	OrigQuantity     string          `json:"origQty"`
	ExecutedQuantity string          `json:"executedQty"`
	AveragePrice     string          `json:"avragePrice"`
	Status           OrderStatusType `json:"status"`
	TimeInForce      TimeInForceType `json:"timeInForce"`
	Type             OrderType       `json:"type"`
	Side             SideType        `json:"side"`
	Time             int64           `json:"time"`
}

//easyjson:json
type LongShortRatio struct {
	Symbol         string `json:"symbol"`
	LongShortRatio string `json:"longShortRatio"`
	LongAccount    string `json:"longAccount"`
	ShortAccount   string `json:"shortAccount"`
	Timestamp      int64  `json:"timestamp"`
}

// LotSizeFilter define lot size filter of symbol
//
//easyjson:json
type LotSizeFilter struct {
	MaxQuantity string `json:"maxQty"`
	MinQuantity string `json:"minQty"`
	StepSize    string `json:"stepSize"`
}

// MarketLotSizeFilter define market lot size filter of symbol
//
//easyjson:json
type MarketLotSizeFilter struct {
	MaxQuantity string `json:"maxQty"`
	MinQuantity string `json:"minQty"`
	StepSize    string `json:"stepSize"`
}

// MaxNumAlgoOrdersFilter define max num algo orders filter of symbol
//
//easyjson:json
type MaxNumAlgoOrdersFilter struct {
	Limit int64 `json:"limit"`
}

// MaxNumOrdersFilter define max num orders filter of symbol
//
//easyjson:json
type MaxNumOrdersFilter struct {
	Limit int64 `json:"limit"`
}

// MinNotionalFilter define min notional filter of symbol
//
//easyjson:json
type MinNotionalFilter struct {
	Notional string `json:"notional"`
}

// Response of user's multi-asset mode
//
//easyjson:json
type MultiAssetMode struct {
	MultiAssetsMargin bool `json:"multiAssetsMargin"`
}

//easyjson:json
type OpenInterest struct {
	OpenInterest string `json:"openInterest"`
	Symbol       string `json:"symbol"`
	Time         int64  `json:"time"`
}

//easyjson:json
type OpenInterestStatistic struct {
	Symbol               string `json:"symbol"`
	SumOpenInterest      string `json:"sumOpenInterest"`
	SumOpenInterestValue string `json:"sumOpenInterestValue"`
	Timestamp            int64  `json:"timestamp"`
}

// Order define order info
//
//easyjson:json
type Order struct {
	Symbol           string              `json:"symbol"`
	OrderID          int64               `json:"orderId"`
	ClientOrderID    string              `json:"clientOrderId"`
	Price            decimal.Decimal     `json:"price"`
	ReduceOnly       bool                `json:"reduceOnly"`
	OrigQuantity     decimal.Decimal     `json:"origQty"`
	ExecutedQuantity decimal.NullDecimal `json:"executedQty"`
	CumQuantity      string              `json:"cumQty"`
	CumQuote         string              `json:"cumQuote"`
	Status           OrderStatusType     `json:"status"`
	TimeInForce      TimeInForceType     `json:"timeInForce"`
	Type             OrderType           `json:"type"`
	Side             SideType            `json:"side"`
	StopPrice        decimal.NullDecimal `json:"stopPrice"`
	Time             int64               `json:"time"`
	UpdateTime       int64               `json:"updateTime"`
	WorkingType      WorkingType         `json:"workingType"`
	ActivatePrice    string              `json:"activatePrice"`
	PriceRate        string              `json:"priceRate"`
	AvgPrice         string              `json:"avgPrice"`
	OrigType         string              `json:"origType"`
	PositionSide     PositionSideType    `json:"positionSide"`
	PriceProtect     bool                `json:"priceProtect"`
	ClosePosition    bool                `json:"closePosition"`
}

// PercentPriceFilter define percent price filter of symbol
//
//easyjson:json
type PercentPriceFilter struct {
	MultiplierDecimal string `json:"multiplierDecimal"`
	MultiplierUp      string `json:"multiplierUp"`
	MultiplierDown    string `json:"multiplierDown"`
}

// PositionMarginHistory define position margin history info
//
//easyjson:json
type PositionMarginHistory struct {
	Amount       string `json:"amount"`
	Asset        string `json:"asset"`
	Symbol       string `json:"symbol"`
	Time         int64  `json:"time"`
	Type         int    `json:"type"`
	PositionSide string `json:"positionSide"`
}

// Response of user's position mode
//
//easyjson:json
type PositionMode struct {
	DualSidePosition bool `json:"dualSidePosition"`
}

// PositionRisk define position risk info
//
//easyjson:json
type PositionRisk struct {
	EntryPrice       decimal.Decimal `json:"entryPrice"`
	BreakEvenPrice   string          `json:"breakEvenPrice"`
	MarginType       string          `json:"marginType"`
	IsAutoAddMargin  string          `json:"isAutoAddMargin"`
	IsolatedMargin   string          `json:"isolatedMargin"`
	Leverage         string          `json:"leverage"`
	LiquidationPrice string          `json:"liquidationPrice"`
	MarkPrice        string          `json:"markPrice"`
	MaxNotionalValue string          `json:"maxNotionalValue"`
	PositionAmt      decimal.Decimal `json:"positionAmt"`
	Symbol           string          `json:"symbol"`
	UnRealizedProfit string          `json:"unRealizedProfit"`
	PositionSide     string          `json:"positionSide"`
	Notional         string          `json:"notional"`
	IsolatedWallet   string          `json:"isolatedWallet"`
}

// PremiumIndex define premium index of mark price
//
//easyjson:json
type PremiumIndex struct {
	Symbol          string `json:"symbol"`
	MarkPrice       string `json:"markPrice"`
	LastFundingRate string `json:"lastFundingRate"`
	NextFundingTime int64  `json:"nextFundingTime"`
	Time            int64  `json:"time"`
}

// PriceChangeStats define price change stats
//
//easyjson:json
type PriceChangeStats struct {
	Symbol             string              `json:"symbol"`
	PriceChange        string              `json:"priceChange"`
	PriceChangePercent string              `json:"priceChangePercent"`
	WeightedAvgPrice   string              `json:"weightedAvgPrice"`
	PrevClosePrice     string              `json:"prevClosePrice"`
	LastPrice          string              `json:"lastPrice"`
	LastQuantity       string              `json:"lastQty"`
	OpenPrice          decimal.NullDecimal `json:"openPrice"`
	HighPrice          string              `json:"highPrice"`
	LowPrice           string              `json:"lowPrice"`
	Volume             string              `json:"volume"`
	QuoteVolume        string              `json:"quoteVolume"`
	OpenTime           int64               `json:"openTime"`
	CloseTime          int64               `json:"closeTime"`
	FirstID            int64               `json:"firstId"`
	LastID             int64               `json:"lastId"`
	Count              int64               `json:"count"`
}

// PriceFilter define price filter of symbol
//
//easyjson:json
type PriceFilter struct {
	MaxPrice string `json:"maxPrice"`
	MinPrice string `json:"minPrice"`
	TickSize string `json:"tickSize"`
}

// RateLimit struct
//
//easyjson:json
type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int64  `json:"intervalNum"`
	Limit         int64  `json:"limit"`
}

// PositionRisk define position risk info
//
//easyjson:json
type RebateNewUser struct {
	BrokerId      string `json:"brokerId"`
	RebateWorking bool   `json:"rebateWorking"`
	IfNewUser     bool   `json:"ifNewUser"`
}

//easyjson:json
type ReferralOverviewResponse struct {
	BrokerId                  string `json:"brokerId"`
	NewTraderRebateCommission string `json:"newTraderRebateCommission"`
	OldTraderRebateCommission string `json:"oldTraderRebateCommission"`
	TotalTradeUser            int    `json:"totalTradeUser"`
	Unit                      string `json:"unit"`
	TotalTradeVol             string `json:"totalTradeVol"`
	TotalRebateVol            string `json:"totalRebateVol"`
	Time                      int64  `json:"time"`
}

// Symbol market symbol
//
//easyjson:json
type Symbol struct {
	Symbol                string                   `json:"symbol"`
	Pair                  string                   `json:"pair"`
	ContractType          ContractType             `json:"contractType"`
	DeliveryDate          int64                    `json:"deliveryDate"`
	OnboardDate           int64                    `json:"onboardDate"`
	Status                string                   `json:"status"`
	MaintMarginPercent    string                   `json:"maintMarginPercent"`
	RequiredMarginPercent string                   `json:"requiredMarginPercent"`
	PricePrecision        int                      `json:"pricePrecision"`
	QuantityPrecision     int                      `json:"quantityPrecision"`
	BaseAssetPrecision    int                      `json:"baseAssetPrecision"`
	QuotePrecision        int                      `json:"quotePrecision"`
	UnderlyingType        string                   `json:"underlyingType"`
	UnderlyingSubType     []string                 `json:"underlyingSubType"`
	SettlePlan            int64                    `json:"settlePlan"`
	TriggerProtect        string                   `json:"triggerProtect"`
	OrderType             []OrderType              `json:"orderType"`
	TimeInForce           []TimeInForceType        `json:"timeInForce"`
	Filters               []map[string]interface{} `json:"filters"`
	QuoteAsset            string                   `json:"quoteAsset"`
	MarginAsset           string                   `json:"marginAsset"`
	BaseAsset             string                   `json:"baseAsset"`
	LiquidationFee        string                   `json:"liquidationFee"`
	MarketTakeBound       string                   `json:"marketTakeBound"`
}

// SymbolLeverage define leverage info of symbol
//
//easyjson:json
type SymbolLeverage struct {
	Leverage         int    `json:"leverage"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	Symbol           string `json:"symbol"`
}

// SymbolPrice define symbol and price pair
//
//easyjson:json
type SymbolPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

// Trade define trade info
//
//easyjson:json
type Trade struct {
	ID            int64  `json:"id"`
	Price         string `json:"price"`
	Quantity      string `json:"qty"`
	QuoteQuantity string `json:"quoteQty"`
	Time          int64  `json:"time"`
	IsBuyerMaker  bool   `json:"isBuyerMaker"`
}

// TradeV3 define v3 trade info
//
//easyjson:json
type TradeV3 struct {
	ID              int64  `json:"id"`
	Symbol          string `json:"symbol"`
	OrderID         int64  `json:"orderId"`
	Price           string `json:"price"`
	Quantity        string `json:"qty"`
	QuoteQuantity   string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}

//easyjson:json
type TraderSummaryResponse struct {
	CustomerId string `json:"customerId"`
	Unit       string `json:"unit"`
	TradeVol   string `json:"tradeVol"`
	RebateVol  string `json:"rebateVol"`
	Time       int64  `json:"time"`
}

// UserLiquidationOrder defines user's liquidation order
//
//easyjson:json
type UserLiquidationOrder struct {
	OrderId          int64            `json:"orderId"`
	Symbol           string           `json:"symbol"`
	Status           OrderStatusType  `json:"status"`
	ClientOrderId    string           `json:"clientOrderId"`
	Price            string           `json:"price"`
	AveragePrice     string           `json:"avgPrice"`
	OrigQuantity     string           `json:"origQty"`
	ExecutedQuantity string           `json:"executedQty"`
	CumQuote         string           `json:"cumQuote"`
	TimeInForce      TimeInForceType  `json:"timeInForce"`
	Type             OrderType        `json:"type"`
	ReduceOnly       bool             `json:"reduceOnly"`
	ClosePosition    bool             `json:"closePosition"`
	Side             SideType         `json:"side"`
	PositionSide     PositionSideType `json:"positionSide"`
	StopPrice        string           `json:"stopPrice"`
	WorkingType      WorkingType      `json:"workingType"`
	OrigType         string           `json:"origType"`
	Time             int64            `json:"time"`
	UpdateTime       int64            `json:"updateTime"`
}
