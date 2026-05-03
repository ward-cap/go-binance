package binance

import (
	"github.com/shopspring/decimal"
	"github.com/ward-cap/go-binance/common"
)

//go:generate easyjson -all models.go

// Ask is a type alias for PriceLevel.
type Ask = common.PriceLevel

// Bid is a type alias for PriceLevel.
type Bid = common.PriceLevel

// APIKeyPermission define API key permission
//
//easyjson:json
type APIKeyPermission struct {
	IPRestrict                     bool   `json:"ipRestrict"`
	CreateTime                     uint64 `json:"createTime"`
	EnableWithdrawals              bool   `json:"enableWithdrawals"`
	EnableInternalTransfer         bool   `json:"enableInternalTransfer"`
	PermitsUniversalTransfer       bool   `json:"permitsUniversalTransfer"`
	EnableVanillaOptions           bool   `json:"enableVanillaOptions"`
	EnableReading                  bool   `json:"enableReading"`
	EnableFutures                  bool   `json:"enableFutures"`
	EnableMargin                   bool   `json:"enableMargin"`
	EnableSpotAndMarginTrading     bool   `json:"enableSpotAndMarginTrading"`
	TradingAuthorityExpirationTime uint64 `json:"tradingAuthorityExpirationTime"`
}

// Account define account info
//
//easyjson:json
type Account struct {
	MakerCommission  int64           `json:"makerCommission"`
	TakerCommission  int64           `json:"takerCommission"`
	BuyerCommission  int64           `json:"buyerCommission"`
	SellerCommission int64           `json:"sellerCommission"`
	CommissionRates  CommissionRates `json:"commissionRates"`
	CanTrade         bool            `json:"canTrade"`
	CanWithdraw      bool            `json:"canWithdraw"`
	CanDeposit       bool            `json:"canDeposit"`
	UpdateTime       uint64          `json:"updateTime"`
	AccountType      string          `json:"accountType"`
	Balances         []Balance       `json:"balances"`
	Permissions      []string        `json:"permissions"`
}

//easyjson:json
type AccountTransferFuturesResponse struct {
	TxnId     int64  `json:"txnId"`
	ErrorData string `json:"errorData"`
}

//easyjson:json
type AccountTransferHistoryFuturesResponse struct {
	Success     bool       `json:"success"`
	FuturesType int64      `json:"futuresType"`
	Transfers   []Transfer `json:"transfers"`
}

//easyjson:json
type AccountTransferHistorySpotResponse struct {
	FromId string          `json:"fromId"`
	ToId   string          `json:"toId"`
	Asset  string          `json:"asset"`
	Qty    decimal.Decimal `json:"qty"`
	Time   int64           `json:"time"`
	TxnId  int64           `json:"txnId"`
}

//easyjson:json
type AccountTransferSpotResponse struct {
	TxnId     int64  `json:"txnId"`
	ErrorData string `json:"errorData"`
}

//easyjson:json
type AddLiquidityPreviewResponse struct {
	QuoteAsset string `json:"quoteAsset"`
	BaseAsset  string `json:"baseAsset"` // only existed when type is COMBINATION
	QuoteAmt   string `json:"quoteAmt"`
	BaseAmt    string `json:"baseAmt"` // only existed when type is COMBINATION
	Price      string `json:"price"`
	Share      string `json:"share"`
	Slippage   string `json:"slippage"`
	Fee        string `json:"fee"`
}

//easyjson:json
type AddLiquidityResponse struct {
	OperationId int64 `json:"operationId"`
}

// AggTrade define aggregate trade info
//
//easyjson:json
type AggTrade struct {
	AggTradeID       int64  `json:"a"`
	Price            string `json:"p"`
	Quantity         string `json:"q"`
	FirstTradeID     int64  `json:"f"`
	LastTradeID      int64  `json:"l"`
	Timestamp        int64  `json:"T"`
	IsBuyerMaker     bool   `json:"m"`
	IsBestPriceMatch bool   `json:"M"`
}

//easyjson:json
type AssetBalance struct {
	Asset  string  `json:"asset"`
	Free   float64 `json:"free"`
	Locked float64 `json:"locked"`
}

// AssetDetail represents the detail of an asset
//
//easyjson:json
type AssetDetail struct {
	MinWithdrawAmount string `json:"minWithdrawAmount"`
	DepositStatus     bool   `json:"depositStatus"`
	WithdrawFee       string `json:"withdrawFee"`
	WithdrawStatus    bool   `json:"withdrawStatus"`
	DepositTip        string `json:"depositTip"`
}

//easyjson:json
type AssetFundingResponse struct {
	Asset        string              `json:"asset"`
	Free         decimal.Decimal     `json:"free"`
	Locked       decimal.Decimal     `json:"locked"`
	Freeze       decimal.Decimal     `json:"freeze"`
	Withdrawing  decimal.NullDecimal `json:"withdrawing"`
	BtcValuation decimal.NullDecimal `json:"btcValuation"`
}

// AvgPrice define average price
//
//easyjson:json
type AvgPrice struct {
	Mins  int64  `json:"mins"`
	Price string `json:"price"`
}

// BNBBurn response
//
//easyjson:json
type BNBBurn struct {
	SpotBNBBurn     bool `json:"spotBNBBurn"`
	InterestBNBBurn bool `json:"interestBNBBurn"`
}

// Balance define user balance of your account
//
//easyjson:json
type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
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

//easyjson:json
type BrokerCommissionRebateResponse struct {
	SubAccountID string          `json:"subaccountId"`
	Income       decimal.Decimal `json:"income"`
	Asset        string          `json:"asset"`
	Symbol       string          `json:"symbol"`
	Time         int64           `json:"time"`
	TradeId      int             `json:"tradeId"`
	Status       int             `json:"status"`
}

// C2CRecord a record of c2c
//
//easyjson:json
type C2CRecord struct {
	OrderNumber         string `json:"orderNumber"`
	AdvNo               string `json:"advNo"`
	TradeType           string `json:"tradeType"`
	Asset               string `json:"asset"`
	Fiat                string `json:"fiat"`
	FiatSymbol          string `json:"fiatSymbol"`
	Amount              string `json:"amount"`
	TotalPrice          string `json:"totalPrice"`
	UnitPrice           string `json:"unitPrice"`
	OrderStatus         string `json:"orderStatus"`
	CreateTime          int64  `json:"createTime"`
	Commission          string `json:"commission"`
	CounterPartNickName string `json:"counterPartNickName"`
	AdvertisementRole   string `json:"advertisementRole"`
}

// C2CTradeHistory response
//
//easyjson:json
type C2CTradeHistory struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    []C2CRecord `json:"data"`
	Total   int64       `json:"total"`
	Success bool        `json:"success"`
}

// CancelMarginOCOResponse define create cancelled oco response.
//
//easyjson:json
type CancelMarginOCOResponse struct {
	OrderListID       int64                   `json:"orderListId"`
	ContingencyType   string                  `json:"contingencyType"`
	ListStatusType    string                  `json:"listStatusType"`
	ListOrderStatus   string                  `json:"listOrderStatus"`
	ListClientOrderID string                  `json:"listClientOrderId"`
	TransactionTime   int64                   `json:"transactionTime"`
	Symbol            string                  `json:"symbol"`
	IsIsolated        bool                    `json:"isIsolated"`
	Orders            []*MarginOCOOrder       `json:"orders"`
	OrderReports      []*MarginOCOOrderReport `json:"orderReports"`
}

// CancelMarginOrderResponse define response of canceling order
//
//easyjson:json
type CancelMarginOrderResponse struct {
	Symbol                   string          `json:"symbol"`
	OrigClientOrderID        string          `json:"origClientOrderId"`
	OrderID                  string          `json:"orderId"`
	ClientOrderID            string          `json:"clientOrderId"`
	TransactTime             int64           `json:"transactTime"`
	Price                    string          `json:"price"`
	OrigQuantity             string          `json:"origQty"`
	ExecutedQuantity         string          `json:"executedQty"`
	CummulativeQuoteQuantity string          `json:"cummulativeQuoteQty"`
	Status                   OrderStatusType `json:"status"`
	TimeInForce              TimeInForceType `json:"timeInForce"`
	Type                     OrderType       `json:"type"`
	Side                     SideType        `json:"side"`
}

// CancelOCOResponse may be returned included in a CancelOpenOrdersResponse.
//
//easyjson:json
type CancelOCOResponse struct {
	OrderListID       int64             `json:"orderListId"`
	ContingencyType   string            `json:"contingencyType"`
	ListStatusType    string            `json:"listStatusType"`
	ListOrderStatus   string            `json:"listOrderStatus"`
	ListClientOrderID string            `json:"listClientOrderId"`
	TransactionTime   int64             `json:"transactionTime"`
	Symbol            string            `json:"symbol"`
	Orders            []*OCOOrder       `json:"orders"`
	OrderReports      []*OCOOrderReport `json:"orderReports"`
}

// CancelOpenOrdersResponse defines cancel open orders response.
//
//easyjson:json
type CancelOpenOrdersResponse struct {
	Orders    []*CancelOrderResponse
	OCOOrders []*CancelOCOResponse
}

// CancelOrderResponse may be returned included in a CancelOpenOrdersResponse.
//
//easyjson:json
type CancelOrderResponse struct {
	Symbol                   string          `json:"symbol"`
	OrigClientOrderID        string          `json:"origClientOrderId"`
	OrderID                  int64           `json:"orderId"`
	OrderListID              int64           `json:"orderListId"`
	ClientOrderID            string          `json:"clientOrderId"`
	TransactTime             int64           `json:"transactTime"`
	Price                    string          `json:"price"`
	OrigQuantity             string          `json:"origQty"`
	ExecutedQuantity         string          `json:"executedQty"`
	CummulativeQuoteQuantity string          `json:"cummulativeQuoteQty"`
	Status                   OrderStatusType `json:"status"`
	TimeInForce              TimeInForceType `json:"timeInForce"`
	Type                     OrderType       `json:"type"`
	Side                     SideType        `json:"side"`
}

//easyjson:json
type ClaimRewardResponse struct {
	Success bool `json:"success"`
}

//easyjson:json
type ClaimedRewardHistory struct {
	PoolId        int               `json:"poolId"`
	PoolName      string            `json:"poolName"`
	AssetRewards  string            `json:"assetRewards"`
	ClaimedAt     int64             `json:"claimedTime"`
	ClaimedAmount string            `json:"claimAmount"`
	Status        RewardClaimStatus `json:"status"`
}

//easyjson:json
type CoinInfo struct {
	Coin              string    `json:"coin"`
	DepositAllEnable  bool      `json:"depositAllEnable"`
	Free              string    `json:"free"`
	Freeze            string    `json:"freeze"`
	Ipoable           string    `json:"ipoable"`
	Ipoing            string    `json:"ipoing"`
	IsLegalMoney      bool      `json:"isLegalMoney"`
	Locked            string    `json:"locked"`
	Name              string    `json:"name"`
	NetworkList       []Network `json:"networkList"`
	Storage           string    `json:"storage"`
	Trading           bool      `json:"trading"`
	WithdrawAllEnable bool      `json:"withdrawAllEnable"`
	Withdrawing       string    `json:"withdrawing"`
}

//easyjson:json
type CommissionRateSubAccount struct {
	SubAccountId    int64  `json:"subAccountId"`
	Symbol          string `json:"symbol"`
	MakerAdjustment int    `json:"makerAdjustment"`
	TakerAdjustment int    `json:"takerAdjustment"`
	MakerCommission int    `json:"makerCommission"`
	TakerCommission int    `json:"takerCommission"`
}

//easyjson:json
type CommissionRates struct {
	Maker  string `json:"maker"`
	Taker  string `json:"taker"`
	Buyer  string `json:"buyer"`
	Seller string `json:"seller"`
}

// ConvertTradeHistory define the convert trade history
//
//easyjson:json
type ConvertTradeHistory struct {
	List      []ConvertTradeHistoryItem `json:"list"`
	StartTime int64                     `json:"startTime"`
	EndTime   int64                     `json:"endTime"`
	Limit     int32                     `json:"limit"`
	MoreData  bool                      `json:"moreData"`
}

// ConvertTradeHistoryItem define a convert trade history item
//
//easyjson:json
type ConvertTradeHistoryItem struct {
	QuoteId      string `json:"quoteId"`
	OrderId      int64  `json:"orderId"`
	OrderStatus  string `json:"orderStatus"`
	FromAsset    string `json:"fromAsset"`
	FromAmount   string `json:"fromAmount"`
	ToAsset      string `json:"toAsset"`
	ToAmount     string `json:"toAmount"`
	Ratio        string `json:"ratio"`
	InverseRatio string `json:"inverseRatio"`
	CreateTime   int64  `json:"createTime"`
}

// CreateMarginOCOResponse define create order response
//
//easyjson:json
type CreateMarginOCOResponse struct {
	OrderListID           int64                   `json:"orderListId"`
	ContingencyType       string                  `json:"contingencyType"`
	ListStatusType        string                  `json:"listStatusType"`
	ListOrderStatus       string                  `json:"listOrderStatus"`
	ListClientOrderID     string                  `json:"listClientOrderId"`
	TransactionTime       int64                   `json:"transactionTime"`
	Symbol                string                  `json:"symbol"`
	MarginBuyBorrowAmount string                  `json:"marginBuyBorrowAmount"`
	MarginBuyBorrowAsset  string                  `json:"marginBuyBorrowAsset"`
	IsIsolated            bool                    `json:"isIsolated"`
	Orders                []*MarginOCOOrder       `json:"orders"`
	OrderReports          []*MarginOCOOrderReport `json:"orderReports"`
}

// CreateOCOResponse define create order response
//
//easyjson:json
type CreateOCOResponse struct {
	OrderListID       int64             `json:"orderListId"`
	ContingencyType   string            `json:"contingencyType"`
	ListStatusType    string            `json:"listStatusType"`
	ListOrderStatus   string            `json:"listOrderStatus"`
	ListClientOrderID string            `json:"listClientOrderId"`
	TransactionTime   int64             `json:"transactionTime"`
	Symbol            string            `json:"symbol"`
	Orders            []*OCOOrder       `json:"orders"`
	OrderReports      []*OCOOrderReport `json:"orderReports"`
}

// CreateOrderResponse define create order response
//
//easyjson:json
type CreateOrderResponse struct {
	Symbol                   string `json:"symbol"`
	OrderID                  int64  `json:"orderId"`
	ClientOrderID            string `json:"clientOrderId"`
	TransactTime             int64  `json:"transactTime"`
	Price                    string `json:"price"`
	OrigQuantity             string `json:"origQty"`
	ExecutedQuantity         string `json:"executedQty"`
	CummulativeQuoteQuantity string `json:"cummulativeQuoteQty"`
	IsIsolated               bool   `json:"isIsolated"` // for isolated margin

	Status      OrderStatusType `json:"status"`
	TimeInForce TimeInForceType `json:"timeInForce"`
	Type        OrderType       `json:"type"`
	Side        SideType        `json:"side"`

	// for order response is set to FULL
	Fills                 []*Fill `json:"fills"`
	MarginBuyBorrowAmount string  `json:"marginBuyBorrowAmount"` // for margin
	MarginBuyBorrowAsset  string  `json:"marginBuyBorrowAsset"`
}

// CreateUserUniversalTransferResponse represents a response from CreateUserUniversalTransferResponse.
//
//easyjson:json
type CreateUserUniversalTransferResponse struct {
	ID int64 `json:"tranId"`
}

// CreateWithdrawResponse represents a response from CreateWithdrawService.
//
//easyjson:json
type CreateWithdrawResponse struct {
	ID string `json:"id"`
}

// Deposit represents a single deposit entry.
//
//easyjson:json
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

// DepthResponse define depth info with bids and asks
//
//easyjson:json
type DepthResponse struct {
	LastUpdateID int64 `json:"lastUpdateId"`
	Bids         []Bid `json:"bids"`
	Asks         []Ask `json:"asks"`
}

// DividendResponse represents a response from AssetDividendService.
//
//easyjson:json
type DividendResponse struct {
	ID     int64  `json:"id"`
	Amount string `json:"amount"`
	Asset  string `json:"asset"`
	Info   string `json:"enInfo"`
	Time   int64  `json:"divTime"`
	TranID int64  `json:"tranId"`
}

// DividendResponseWrapper represents a wrapper around a AssetDividendService.
//
//easyjson:json
type DividendResponseWrapper struct {
	Rows  *[]DividendResponse `json:"rows"`
	Total int32               `json:"total"`
}

// DustResult represents the result of a DustLog API Call.
//
//easyjson:json
type DustResult struct {
	Total              uint8               `json:"total"` //Total counts of exchange
	UserAssetDribblets []UserAssetDribblet `json:"userAssetDribblets"`
}

// DustTransferResponse represents the response from DustTransferService.
//
//easyjson:json
type DustTransferResponse struct {
	TotalServiceCharge string                `json:"totalServiceCharge"`
	TotalTransfered    string                `json:"totalTransfered"`
	TransferResult     []*DustTransferResult `json:"transferResult"`
}

// DustTransferResult represents the result of a dust transfer.
//
//easyjson:json
type DustTransferResult struct {
	Amount              string `json:"amount"`
	FromAsset           string `json:"fromAsset"`
	OperateTime         int64  `json:"operateTime"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	TranID              int64  `json:"tranId"`
	TransferedAmount    string `json:"transferedAmount"`
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

// FiatDepositWithdrawHistory define the fiat deposit/withdraw history
//
//easyjson:json
type FiatDepositWithdrawHistory struct {
	Code    string                           `json:"code"`
	Message string                           `json:"message"`
	Data    []FiatDepositWithdrawHistoryItem `json:"data"`
	Total   int32                            `json:"total"`
	Success bool                             `json:"success"`
}

// FiatDepositWithdrawHistoryItem define a fiat deposit/withdraw history item
//
//easyjson:json
type FiatDepositWithdrawHistoryItem struct {
	OrderNo         string `json:"orderNo"`
	FiatCurrency    string `json:"fiatCurrency"`
	IndicatedAmount string `json:"indicatedAmount"`
	Amount          string `json:"amount"`
	TotalFee        string `json:"totalFee"`
	Method          string `json:"method"`
	Status          string `json:"status"`
	CreateTime      int64  `json:"createTime"`
	UpdateTime      int64  `json:"updateTime"`
}

// FiatPaymentsHistory define the fiat payments history
//
//easyjson:json
type FiatPaymentsHistory struct {
	Code    string                    `json:"code"`
	Message string                    `json:"message"`
	Data    []FiatPaymentsHistoryItem `json:"data"`
	Total   int32                     `json:"total"`
	Success bool                      `json:"success"`
}

// FiatPaymentsHistoryItem define a fiat payments history item
//
//easyjson:json
type FiatPaymentsHistoryItem struct {
	OrderNo        string `json:"orderNo"`
	SourceAmount   string `json:"sourceAmount"`
	FiatCurrency   string `json:"fiatCurrency"`
	ObtainAmount   string `json:"obtainAmount"`
	CryptoCurrency string `json:"cryptoCurrency"`
	TotalFee       string `json:"totalFee"`
	Price          string `json:"price"`
	Status         string `json:"status"`
	CreateTime     int64  `json:"createTime"`
	UpdateTime     int64  `json:"updateTime"`
}

// Fill may be returned in an array of fills in a CreateOrderResponse.
//
//easyjson:json
type Fill struct {
	TradeID         int64  `json:"tradeId"`
	Price           string `json:"price"`
	Quantity        string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
}

//easyjson:json
type FundsDetail struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

// FuturesTransfer define futures transfer history item
//
//easyjson:json
type FuturesTransfer struct {
	Asset     string                    `json:"asset"`
	TranID    int64                     `json:"tranId"`
	Amount    string                    `json:"amount"`
	Type      int64                     `json:"type"`
	Timestamp int64                     `json:"timestamp"`
	Status    FuturesTransferStatusType `json:"status"`
}

// FuturesTransferHistory define futures transfer history
//
//easyjson:json
type FuturesTransferHistory struct {
	Rows  []FuturesTransfer `json:"rows"`
	Total int64             `json:"total"`
}

// GetAPIKeyPermission get API Key permission info
//
//easyjson:json
type GetAPIKeyPermission struct {
	c *Client
}

//easyjson:json
type GetBrokerInfoResponse struct {
	MaxMakerCommission decimal.Decimal `json:"maxMakerCommission"`
	MinMakerCommission decimal.Decimal `json:"minMakerCommission"`
	MaxTakerCommission decimal.Decimal `json:"maxTakerCommission"`
	MinTakerCommission decimal.Decimal `json:"minTakerCommission"`

	SubAccountQty    int64 `json:"subAccountQty"`
	MaxSubAccountQty int64 `json:"maxSubAccountQty"`
}

// GetDepositAddressResponse represents a response from GetDepositsAddressService.
//
//easyjson:json
type GetDepositAddressResponse struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
	Coin    string `json:"coin"`
	URL     string `json:"url"`
}

//easyjson:json
type GetIPRestrictionSubAccountAPIKeyResponse struct {
	SubAccountId string   `json:"subaccountId"`
	IpRestrict   string   `json:"ipRestrict"` // true/false
	Apikey       string   `json:"apikey"`
	IpList       []string `json:"ipList"`
	UpdateTime   int64    `json:"updateTime"`
}

//easyjson:json
type GetListingResponse struct {
	OpenTime int64    `json:"openTime"`
	Symbols  []string `json:"symbols"`
}

//easyjson:json
type GetReferralRebateRecord struct {
	c         *Client
	startTime *int64 // required
	endTime   *int64 // required
	limit     *int   // max 500
}

//easyjson:json
type GetSwapQuoteResponse struct {
	QuoteAsset string `json:"quoteAsset"`
	BaseAsset  string `json:"baseAsset"`
	QuoteQty   string `json:"quoteQty"`
	BaseQty    string `json:"baseQty"`
	Price      string `json:"price"`
	Slippage   string `json:"slippage"`
	Fee        string `json:"fee"`
}

// IcebergPartsFilter define iceberg part filter of symbol
//
//easyjson:json
type IcebergPartsFilter struct {
	Limit int `json:"limit"`
}

// InterestHistory represents a response from InterestHistoryService.
//
//easyjson:json
type InterestHistory []InterestHistoryElement

//easyjson:json
type InterestHistoryElement struct {
	Asset       string      `json:"asset"`
	Interest    string      `json:"interest"`
	LendingType LendingType `json:"lendingType"`
	ProductName string      `json:"productName"`
	Time        int64       `json:"time"`
}

//easyjson:json
type InternalUniversalTransfer struct {
	TranId          int64  `json:"tranId"`
	ClientTranId    string `json:"clientTranId"`
	FromEmail       string `json:"fromEmail"`
	ToEmail         string `json:"toEmail"`
	Asset           string `json:"asset"`
	Amount          string `json:"amount"`
	FromAccountType string `json:"fromAccountType"`
	ToAccountType   string `json:"toAccountType"`
	Status          string `json:"status"`
	CreateTimeStamp uint64 `json:"createTimeStamp"`
}

//easyjson:json
type InternalUniversalTransferHistoryResponse struct {
	Result     []*InternalUniversalTransfer `json:"result"`
	TotalCount int                          `json:"totalCount"`
}

//easyjson:json
type InternalUniversalTransferResponse struct {
	ID           int64  `json:"tranId"`
	ClientTranID string `json:"clientTranId"`
}

// IsolatedMarginAccount defines isolated user assets of margin account
//
//easyjson:json
type IsolatedMarginAccount struct {
	TotalAssetOfBTC     string                `json:"totalAssetOfBtc"`
	TotalLiabilityOfBTC string                `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBTC  string                `json:"totalNetAssetOfBtc"`
	Assets              []IsolatedMarginAsset `json:"assets"`
}

// IsolatedMarginAllPair define isolated margin pair info
//
//easyjson:json
type IsolatedMarginAllPair struct {
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// IsolatedMarginAsset defines isolated margin asset information, like margin level, liquidation price... etc
//
//easyjson:json
type IsolatedMarginAsset struct {
	Symbol     string            `json:"symbol"`
	QuoteAsset IsolatedUserAsset `json:"quoteAsset"`
	BaseAsset  IsolatedUserAsset `json:"baseAsset"`

	IsolatedCreated   bool   `json:"isolatedCreated"`
	Enabled           bool   `json:"enabled"`
	MarginLevel       string `json:"marginLevel"`
	MarginLevelStatus string `json:"marginLevelStatus"`
	MarginRatio       string `json:"marginRatio"`
	IndexPrice        string `json:"indexPrice"`
	LiquidatePrice    string `json:"liquidatePrice"`
	LiquidateRate     string `json:"liquidateRate"`
	TradeEnabled      bool   `json:"tradeEnabled"`
}

// IsolatedUserAsset defines isolated user assets of the margin account
//
//easyjson:json
type IsolatedUserAsset struct {
	Asset         string `json:"asset"`
	Borrowed      string `json:"borrowed"`
	Free          string `json:"free"`
	Interest      string `json:"interest"`
	Locked        string `json:"locked"`
	NetAsset      string `json:"netAsset"`
	NetAssetOfBtc string `json:"netAssetOfBtc"`

	BorrowEnabled bool   `json:"borrowEnabled"`
	RepayEnabled  bool   `json:"repayEnabled"`
	TotalAsset    string `json:"totalAsset"`
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

//easyjson:json
type LiquidityPool struct {
	PoolId   int64    `json:"poolId"`
	PoolName string   `json:"poolName"`
	Assets   []string `json:"assets"`
}

//easyjson:json
type LiquidityPoolDetail struct {
	PoolId     int64                 `json:"poolId"`
	PoolName   string                `json:"poolName"`
	UpdateTime int64                 `json:"updateTime"`
	Liquidity  map[string]string     `json:"liquidity"`
	Share      *PoolShareInformation `json:"share"`
}

//easyjson:json
type ListDustDetail struct {
	Asset            string `json:"asset"`
	AssetFullName    string `json:"assetFullName"`
	AmountFree       string `json:"amountFree"`
	ToBTC            string `json:"toBTC"`
	ToBNB            string `json:"toBNB"`
	ToBNBOffExchange string `json:"toBNBOffExchange"`
	Exchange         string `json:"exchange"`
}

//easyjson:json
type ListDustResponse struct {
	Details            []ListDustDetail `json:"details"`
	TotalTransferBtc   string           `json:"totalTransferBtc"`
	TotalTransferBNB   string           `json:"totalTransferBNB"`
	DribbletPercentage string           `json:"dribbletPercentage"`
}

// LotSizeFilter define lot size filter of symbol
//
//easyjson:json
type LotSizeFilter struct {
	MaxQuantity string `json:"maxQty"`
	MinQuantity string `json:"minQty"`
	StepSize    string `json:"stepSize"`
}

//easyjson:json
type ManagedSubAccountAsset struct {
	Coin             string `json:"coin"`
	Name             string `json:"name"`
	TotalBalance     string `json:"totalBalance"`
	AvailableBalance string `json:"availableBalance"`
	InOrder          string `json:"inOrder"`
	BtcValue         string `json:"btcValue"`
}

//easyjson:json
type ManagedSubAccountDepositResponse struct {
	ID int64 `json:"tranId"`
}

//easyjson:json
type ManagedSubAccountWithdrawalResponse struct {
	ID int64 `json:"tranId"`
}

// MarginAccount define margin account info
//
//easyjson:json
type MarginAccount struct {
	BorrowEnabled       bool        `json:"borrowEnabled"`
	MarginLevel         string      `json:"marginLevel"`
	TotalAssetOfBTC     string      `json:"totalAssetOfBtc"`
	TotalLiabilityOfBTC string      `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBTC  string      `json:"totalNetAssetOfBtc"`
	TradeEnabled        bool        `json:"tradeEnabled"`
	TransferEnabled     bool        `json:"transferEnabled"`
	UserAssets          []UserAsset `json:"userAssets"`
}

// MarginAllPair define margin pair info
//
//easyjson:json
type MarginAllPair struct {
	ID            int64  `json:"id"`
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// MarginAsset define margin asset info
//
//easyjson:json
type MarginAsset struct {
	FullName      string `json:"assetFullName"`
	Name          string `json:"assetName"`
	Borrowable    bool   `json:"isBorrowable"`
	Mortgageable  bool   `json:"isMortgageable"`
	UserMinBorrow string `json:"userMinBorrow"`
	UserMinRepay  string `json:"userMinRepay"`
}

// MarginLoan define margin loan
//
//easyjson:json
type MarginLoan struct {
	Asset     string               `json:"asset"`
	Principal string               `json:"principal"`
	Timestamp int64                `json:"timestamp"`
	Status    MarginLoanStatusType `json:"status"`
}

// MarginLoanResponse define margin loan response
//
//easyjson:json
type MarginLoanResponse struct {
	Rows  []MarginLoan `json:"rows"`
	Total int64        `json:"total"`
}

// MarginOCOOrder may be returned in an array of MarginOCOOrder in a CreateMarginOCOResponse
//
//easyjson:json
type MarginOCOOrder struct {
	Symbol        string `json:"symbol"`
	OrderID       int64  `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
}

// MarginOCOOrderReport may be returned in an array of MarginOCOOrderReport in a CreateMarginOCOResponse
//
//easyjson:json
type MarginOCOOrderReport struct {
	Symbol                   string          `json:"symbol"`
	OrderID                  int64           `json:"orderId"`
	OrderListID              int64           `json:"orderListId"`
	ClientOrderID            string          `json:"clientOrderId"`
	TransactionTime          int64           `json:"transactionTime"`
	Price                    string          `json:"price"`
	OrigQuantity             string          `json:"origQty"`
	ExecutedQuantity         string          `json:"executedQty"`
	CummulativeQuoteQuantity string          `json:"cummulativeQuoteQty"`
	Status                   OrderStatusType `json:"status"`
	TimeInForce              TimeInForceType `json:"timeInForce"`
	Type                     OrderType       `json:"type"`
	Side                     SideType        `json:"side"`
	StopPrice                string          `json:"stopPrice"`
}

// MarginPair define margin pair info
//
//easyjson:json
type MarginPair struct {
	ID            int64  `json:"id"`
	Symbol        string `json:"symbol"`
	Base          string `json:"base"`
	Quote         string `json:"quote"`
	IsMarginTrade bool   `json:"isMarginTrade"`
	IsBuyAllowed  bool   `json:"isBuyAllowed"`
	IsSellAllowed bool   `json:"isSellAllowed"`
}

// MarginPriceIndex define margin price index
//
//easyjson:json
type MarginPriceIndex struct {
	CalcTime int64  `json:"calcTime"`
	Price    string `json:"price"`
	Symbol   string `json:"symbol"`
}

// MarginRepay define margin repay
//
//easyjson:json
type MarginRepay struct {
	Asset     string                `json:"asset"`
	Amount    string                `json:"amount"`
	Interest  string                `json:"interest"`
	Principal string                `json:"principal"`
	Timestamp int64                 `json:"timestamp"`
	Status    MarginRepayStatusType `json:"status"`
	TxID      int64                 `json:"txId"`
}

// MarginRepayResponse define margin repay response
//
//easyjson:json
type MarginRepayResponse struct {
	Rows  []MarginRepay `json:"rows"`
	Total int64         `json:"total"`
}

// MarketLotSizeFilter define market lot size filter of symbol
//
//easyjson:json
type MarketLotSizeFilter struct {
	MaxQuantity string `json:"maxQty"`
	MinQuantity string `json:"minQty"`
	StepSize    string `json:"stepSize"`
}

// MaxBorrowable define max borrowable response
//
//easyjson:json
type MaxBorrowable struct {
	Amount string `json:"amount"`
}

// MaxNumAlgoOrdersFilter define max num algo orders filter of symbol
//
//easyjson:json
type MaxNumAlgoOrdersFilter struct {
	MaxNumAlgoOrders int `json:"maxNumAlgoOrders"`
}

// The "Algo" order is STOP_ LOSS, STOP_ LOS_ LIMITED, TAKE_ PROFIT and TAKE_ PROFIT_ Limit Stop Loss Order.
// Therefore, orders other than the above types are non conditional(Algo) orders, and MaxNumOrders defines the maximum
// number of orders placed for these types of orders
//
//easyjson:json
type MaxNumOrdersFilter struct {
	MaxNumOrders int `json:"maxNumOrders"`
}

// MaxTransferable define max transferable response
//
//easyjson:json
type MaxTransferable struct {
	Amount string `json:"amount"`
}

//easyjson:json
type Network struct {
	AddressRegex            string `json:"addressRegex"`
	Coin                    string `json:"coin"`
	DepositDesc             string `json:"depositDesc,omitempty"` // 仅在充值关闭时返回
	DepositEnable           bool   `json:"depositEnable"`
	IsDefault               bool   `json:"isDefault"`
	MemoRegex               string `json:"memoRegex"`
	MinConfirm              int    `json:"minConfirm"` // 上账所需的最小确认数
	Name                    string `json:"name"`
	Network                 string `json:"network"`
	ResetAddressStatus      bool   `json:"resetAddressStatus"`
	EstimatedArrivalTime    int64  `json:"estimatedArrivalTime"`
	SpecialTips             string `json:"specialTips"`
	UnLockConfirm           int    `json:"unLockConfirm"`          // 解锁需要的确认数
	WithdrawDesc            string `json:"withdrawDesc,omitempty"` // 仅在提现关闭时返回
	WithdrawEnable          bool   `json:"withdrawEnable"`
	WithdrawFee             string `json:"withdrawFee"`
	WithdrawIntegerMultiple string `json:"withdrawIntegerMultiple"`
	WithdrawMax             string `json:"withdrawMax"`
	WithdrawMin             string `json:"withdrawMin"`
	SameAddress             bool   `json:"sameAddress"` // 是否需要memo
}

// NotionalFilter define notional filter of symbol
//
//easyjson:json
type NotionalFilter struct {
	MinNotional      string `json:"minNotional"`
	ApplyMinToMarket bool   `json:"applyMinToMarket"`
	MaxNotional      string `json:"maxNotional"`
	ApplyMaxToMarket bool   `json:"applyMaxToMarket"`
	AvgPriceMins     int    `json:"avgPriceMins"`
}

// OCOOrder may be returned in an array of OCOOrder in a CreateOCOResponse.
//
//easyjson:json
type OCOOrder struct {
	Symbol        string `json:"symbol"`
	OrderID       int64  `json:"orderId"`
	ClientOrderID string `json:"clientOrderId"`
}

// OCOOrderReport may be returned in an array of OCOOrderReport in a CreateOCOResponse.
//
//easyjson:json
type OCOOrderReport struct {
	Symbol                   string          `json:"symbol"`
	OrderID                  int64           `json:"orderId"`
	OrderListID              int64           `json:"orderListId"`
	ClientOrderID            string          `json:"clientOrderId"`
	OrigClientOrderID        string          `json:"origClientOrderId"`
	TransactionTime          int64           `json:"transactionTime"`
	Price                    string          `json:"price"`
	OrigQuantity             string          `json:"origQty"`
	ExecutedQuantity         string          `json:"executedQty"`
	CummulativeQuoteQuantity string          `json:"cummulativeQuoteQty"`
	Status                   OrderStatusType `json:"status"`
	TimeInForce              TimeInForceType `json:"timeInForce"`
	Type                     OrderType       `json:"type"`
	Side                     SideType        `json:"side"`
	StopPrice                string          `json:"stopPrice"`
	IcebergQuantity          string          `json:"icebergQty"`
}

// oco define oco info
//
//easyjson:json
type Oco struct {
	Symbol            string   `json:"symbol"`
	OrderListId       int64    `json:"orderListId"`
	ContingencyType   string   `json:"contingencyType"`
	ListStatusType    string   `json:"listStatusType"`
	ListOrderStatus   string   `json:"listOrderStatus"`
	ListClientOrderID string   `json:"listClientOrderId"`
	TransactionTime   int64    `json:"transactionTime"`
	Orders            []*Order `json:"orders"`
}

// Order define order info
//
//easyjson:json
type Order struct {
	Symbol                   string              `json:"symbol"`
	OrderID                  int64               `json:"orderId"`
	OrderListId              int64               `json:"orderListId"`
	ClientOrderID            string              `json:"clientOrderId"`
	Price                    decimal.Decimal     `json:"price"`
	OrigQuantity             decimal.Decimal     `json:"origQty"`
	ExecutedQuantity         decimal.NullDecimal `json:"executedQty"`
	CummulativeQuoteQuantity string              `json:"cummulativeQuoteQty"`
	Status                   OrderStatusType     `json:"status"`
	TimeInForce              TimeInForceType     `json:"timeInForce"`
	Type                     OrderType           `json:"type"`
	Side                     SideType            `json:"side"`
	StopPrice                decimal.NullDecimal `json:"stopPrice"`
	IcebergQuantity          string              `json:"icebergQty"`
	Time                     int64               `json:"time"`
	UpdateTime               int64               `json:"updateTime"`
	IsWorking                bool                `json:"isWorking"`
	IsIsolated               bool                `json:"isIsolated"`
	OrigQuoteOrderQuantity   string              `json:"origQuoteOrderQty"`
}

//easyjson:json
type PayTradeHistory struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    []PayTradeItem `json:"data"`
	Success bool           `json:"success"`
}

//easyjson:json
type PayTradeItem struct {
	OrderType       string        `json:"orderType"`
	TransactionID   string        `json:"transactionId"`
	TransactionTime int64         `json:"transactionTime"`
	Amount          string        `json:"amount"`
	Currency        string        `json:"currency"`
	PayerInfo       *PayerInfo    `json:"payerInfo"`
	ReceiverInfo    *ReceiverInfo `json:"receiverInfo"`
	FundsDetail     []FundsDetail `json:"fundsDetail"`
}

//easyjson:json
type PayerInfo struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Email       string `json:"email"`
	BinanceId   int    `json:"binanceId"`
	AccountId   int    `json:"accountId"`
	CountryCode int    `json:"countryCode"`
	PhoneNumber string `json:"phoneNumber"`
	MobileCode  string `json:"mobileCode"`
	UnmaskData  bool   `json:"unmaskData"`
}

// PERCENT_PRICE_BY_SIDE define percent price filter of symbol by side
//
//easyjson:json
type PercentPriceBySideFilter struct {
	AveragePriceMins  int    `json:"avgPriceMins"`
	BidMultiplierUp   string `json:"bidMultiplierUp"`
	BidMultiplierDown string `json:"bidMultiplierDown"`
	AskMultiplierUp   string `json:"askMultiplierUp"`
	AskMultiplierDown string `json:"askMultiplierDown"`
}

//easyjson:json
type PoolShareInformation struct {
	ShareAmount     string            `json:"shareAmount"`
	SharePercentage string            `json:"sharePercentage"`
	Assets          map[string]string `json:"asset"`
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
	LastQty            string              `json:"lastQty"`
	BidPrice           string              `json:"bidPrice"`
	BidQty             string              `json:"bidQty"`
	AskPrice           string              `json:"askPrice"`
	AskQty             string              `json:"askQty"`
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

//easyjson:json
type PurchaseSavingsFlexibleProductResponse struct {
	PurchaseId uint64 `json:"purchaseId"`
}

//easyjson:json
type QuerySubAccountResponse struct {
	SubaccountId          string          `json:"subaccountId"`
	Email                 string          `json:"email"`
	Tag                   string          `json:"tag"`
	MakerCommission       decimal.Decimal `json:"makerCommission"`
	TakerCommission       decimal.Decimal `json:"takerCommission"`
	MarginMakerCommission decimal.Decimal `json:"marginMakerCommission"`
	MarginTakerCommission decimal.Decimal `json:"marginTakerCommission"`
	CreateTime            int64           `json:"createTime"`
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

//easyjson:json
type RateLimitFull struct {
	RateLimitType RateLimitType     `json:"rateLimitType"`
	Interval      RateLimitInterval `json:"interval"`
	IntervalNum   int               `json:"intervalNum"`
	Limit         int               `json:"limit"`
	Count         int               `json:"count"`
}

//easyjson:json
type ReceiverInfo struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Email       string `json:"email,omitempty"`
	BinanceId   int    `json:"binanceId,omitempty"`
	AccountId   int    `json:"accountId"`
	CountryCode int    `json:"countryCode,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	MobileCode  string `json:"mobileCode,omitempty"`
	UnmaskData  bool   `json:"unmaskData"`
	Extend      struct {
		PhoneOrEmailChanged bool `json:"phoneOrEmailChanged"`
	} `json:"extend,omitempty"`
}

//easyjson:json
type ReferralRebateRecordResponse struct {
	CustomerId      string `json:"customerId"`
	Email           string `json:"email"`
	Income          string `json:"income"`
	Asset           string `json:"asset"`
	Symbol          string `json:"symbol"`
	Time            int64  `json:"time"`
	OrderId         int    `json:"orderId"`
	TradeId         int    `json:"tradeId"`
	DistributeTime  int64  `json:"distributeTime"`
	CommissionAsset string `json:"commissionAsset"`
	Commission      string `json:"commission"`
	ConvertPrice    string `json:"convertPrice"`
}

//easyjson:json
type RemoveLiquidityResponse struct {
	OperationId int64 `json:"operationId"`
}

// SavingFixedProjectPosition represents a saving flexible product position.
//
//easyjson:json
type SavingFixedProjectPosition struct {
	Asset           string `json:"asset"`
	CanTransfer     bool   `json:"canTransfer"`
	CreateTimestamp int64  `json:"createTimestamp"`
	Duration        int64  `json:"duration"`
	StartTime       int64  `json:"startTime"`
	EndTime         int64  `json:"endTime"`
	PurchaseTime    int64  `json:"purchaseTime"`
	RedeemDate      string `json:"redeemDate"`
	Interest        string `json:"interest"`
	InterestRate    string `json:"interestRate"`
	Lot             int32  `json:"lot"`
	PositionId      int64  `json:"positionId"`
	Principal       string `json:"principal"`
	ProjectId       string `json:"projectId"`
	ProjectName     string `json:"projectName"`
	Status          string `json:"status"`
	ProjectType     string `json:"type"`
}

// SavingFlexibleProductPosition represents a saving flexible product position.
//
//easyjson:json
type SavingFlexibleProductPosition struct {
	Asset                 string `json:"asset"`
	ProductId             string `json:"productId"`
	ProductName           string `json:"productName"`
	AvgAnnualInterestRate string `json:"avgAnnualInterestRate"`
	AnnualInterestRate    string `json:"annualInterestRate"`
	DailyInterestRate     string `json:"dailyInterestRate"`
	TotalInterest         string `json:"totalInterest"`
	TotalAmount           string `json:"totalAmount"`
	TotalPurchasedAmount  string `json:"todayPurchasedAmount"`
	RedeemingAmount       string `json:"redeemingAmount"`
	FreeAmount            string `json:"freeAmount"`
	FreezeAmount          string `json:"freezeAmount,omitempty"`
	LockedAmount          string `json:"lockedAmount,omitempty"`
	CanRedeem             bool   `json:"canRedeem"`
}

// SavingsFixedProduct define a fixed product (Savings)
//
//easyjson:json
type SavingsFixedProduct struct {
	Asset              string `json:"asset"`
	DisplayPriority    int    `json:"displayPriority"`
	Duration           int    `json:"duration"`
	InterestPerLot     string `json:"interestPerLot"`
	InterestRate       string `json:"interestRate"`
	LotSize            string `json:"lotSize"`
	LotsLowLimit       int    `json:"lotsLowLimit"`
	LotsPurchased      int    `json:"lotsPurchased"`
	LotsUpLimit        int    `json:"lotsUpLimit"`
	MaxLotsPerUser     int    `json:"maxLotsPerUser"`
	NeedKyc            bool   `json:"needKyc"`
	ProjectId          string `json:"projectId"`
	ProjectName        string `json:"projectName"`
	Status             string `json:"status"`
	Type               string `json:"type"`
	WithAreaLimitation bool   `json:"withAreaLimitation"`
}

// SavingsFlexibleProduct define a flexible product (Savings)
//
//easyjson:json
type SavingsFlexibleProduct struct {
	Asset                    string `json:"asset"`
	AvgAnnualInterestRate    string `json:"avgAnnualInterestRate"`
	CanPurchase              bool   `json:"canPurchase"`
	CanRedeem                bool   `json:"canRedeem"`
	DailyInterestPerThousand string `json:"dailyInterestPerThousand"`
	Featured                 bool   `json:"featured"`
	MinPurchaseAmount        string `json:"minPurchaseAmount"`
	ProductId                string `json:"productId"`
	PurchasedAmount          string `json:"purchasedAmount"`
	Status                   string `json:"status"`
	UpLimit                  string `json:"upLimit"`
	UpLimitPerUser           string `json:"upLimitPerUser"`
}

// Snapshot define snapshot
//
//easyjson:json
type Snapshot struct {
	Code     int            `json:"code"`
	Msg      string         `json:"msg"`
	Snapshot []*SnapshotVos `json:"snapshotVos"`
}

// SnapshotAssets define snapshot assets
//
//easyjson:json
type SnapshotAssets struct {
	Asset         string `json:"asset"`
	MarginBalance string `json:"marginBalance"`
	WalletBalance string `json:"walletBalance"`
}

// SnapshotBalances define snapshot balances
//
//easyjson:json
type SnapshotBalances struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}

// SnapshotData define content of a snapshot
//
//easyjson:json
type SnapshotData struct {
	MarginLevel         string `json:"marginLevel"`
	TotalAssetOfBtc     string `json:"totalAssetOfBtc"`
	TotalLiabilityOfBtc string `json:"totalLiabilityOfBtc"`
	TotalNetAssetOfBtc  string `json:"totalNetAssetOfBtc"`

	Balances   []*SnapshotBalances   `json:"balances"`
	UserAssets []*SnapshotUserAssets `json:"userAssets"`
	Assets     []*SnapshotAssets     `json:"assets"`
	Positions  []*SnapshotPositions  `json:"position"`
}

// SnapshotPositions define snapshot positions
//
//easyjson:json
type SnapshotPositions struct {
	EntryPrice       string `json:"entryPrice"`
	MarkPrice        string `json:"markPrice"`
	PositionAmt      string `json:"positionAmt"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
}

// SnapshotUserAssets define snapshot user assets
//
//easyjson:json
type SnapshotUserAssets struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

// SnapshotVos define content of a snapshot
//
//easyjson:json
type SnapshotVos struct {
	Data       *SnapshotData `json:"data"`
	Type       string        `json:"type"`
	UpdateTime int64         `json:"updateTime"`
}

// SpotRebateHistory define the spot rebate history
//
//easyjson:json
type SpotRebateHistory struct {
	Status string                `json:"status"`
	Type   string                `json:"type"`
	Code   string                `json:"code"`
	Data   SpotRebateHistoryData `json:"data"`
}

// SpotRebateHistoryData define the data part of the spot rebate history
//
//easyjson:json
type SpotRebateHistoryData struct {
	Page         int32                       `json:"page"`
	TotalRecords int32                       `json:"totalRecords"`
	TotalPageNum int32                       `json:"totalPageNum"`
	Data         []SpotRebateHistoryDataItem `json:"data"`
}

// SpotRebateHistoryDataItem define a spot rebate history data item
//
//easyjson:json
type SpotRebateHistoryDataItem struct {
	Asset      string `json:"asset"`
	Type       int32  `json:"type"`
	Amount     string `json:"amount"`
	UpdateTime int64  `json:"updateTime"`
}

//easyjson:json
type SpotSubUserAssetBtcVoList struct {
	Email      string `json:"email"`
	TotalAsset string `json:"totalAsset"`
}

// StakingHistory represents a list of staking history transactions.
//
//easyjson:json
type StakingHistory []StakingHistoryTransaction

// StakingHistoryTransaction represents a staking history transaction.
//
//easyjson:json
type StakingHistoryTransaction struct {
	PositionId  int64  `json:"positionId"`
	Time        int64  `json:"time"`
	Asset       string `json:"asset"`
	Project     string `json:"project"`
	Amount      string `json:"amount"`
	LockPeriod  int64  `json:"lockPeriod"`
	DeliverDate int64  `json:"deliverDate"`
	Type        string `json:"type"`
	Status      string `json:"status"`
}

// StakingProductPosition represents a staking product position.
//
//easyjson:json
type StakingProductPosition struct {
	PositionId                 int64  `json:"positionId"`
	ProductId                  string `json:"productId"`
	Asset                      string `json:"asset"`
	Amount                     string `json:"amount"`
	PurchaseTime               int64  `json:"purchaseTime"`
	Duration                   int64  `json:"duration"`
	AccrualDays                int64  `json:"accrualDays"`
	RewardAsset                string `json:"rewardAsset"`
	APY                        string `json:"apy"`
	RewardAmount               string `json:"rewardAmt"`
	ExtraRewardAsset           string `json:"extraRewardAsset"`
	ExtraRewardAPY             string `json:"extraRewardAPY"`
	EstimatedExtraRewardAmount string `json:"estExtraRewardAmt"`
	NextInterestPay            string `json:"nextInterestPay"`
	NextInterestPayDate        int64  `json:"nextInterestPayDate"`
	PayInterestPeriod          int64  `json:"payInterestPeriod"`
	RedeemAmountEarly          string `json:"redeemAmountEarly"`
	InterestEndDate            int64  `json:"interestEndDate"`
	DeliverDate                int64  `json:"deliverDate"`
	RedeemPeriod               int64  `json:"redeemPeriod"`
	RedeemingAmount            string `json:"redeemingAmt"`
	PartialAmountDeliverDate   int64  `json:"partialAmtDeliverDate"`
	CanRedeemEarly             bool   `json:"canRedeemEarly"`
	Renewable                  bool   `json:"renewable"`
	Type                       string `json:"type"`
	Status                     string `json:"status"`
}

// StakingProductPositions represents a list of staking product positions.
//
//easyjson:json
type StakingProductPositions []StakingProductPosition

//easyjson:json
type SubAccount struct {
	Email                       string `json:"email"`
	IsFreeze                    bool   `json:"isFreeze"`
	CreateTime                  uint64 `json:"createTime"`
	IsManagedSubAccount         bool   `json:"isManagedSubAccount"`
	IsAssetManagementSubAccount bool   `json:"isAssetManagementSubAccount"`
}

//easyjson:json
type SubAccountAPIKeyResponse struct {
	APIKey    string `json:"apiKey"`
	SecretKey string `json:"secretKey"`
}

//easyjson:json
type SubAccountFuturesAccount struct {
	Email                       string                          `json:"email"`
	Asset                       string                          `json:"asset"`
	Assets                      []SubAccountFuturesAccountAsset `json:"assets"`
	CanDeposit                  bool                            `json:"canDeposit"`
	CanTrade                    bool                            `json:"canTrade"`
	CanWithdraw                 bool                            `json:"canWithdraw"`
	FeeTier                     int                             `json:"feeTier"`
	MaxWithdrawAmount           string                          `json:"maxWithdrawAmount"`
	TotalInitialMargin          string                          `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string                          `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string                          `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string                          `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string                          `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string                          `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string                          `json:"totalWalletBalance"`
	UpdateTime                  int64                           `json:"updateTime"`
}

//easyjson:json
type SubAccountFuturesAccountAsset struct {
	Asset                  string `json:"asset"`
	InitialMargin          string `json:"initialMargin"`
	MaintenanceMargin      string `json:"maintenanceMargin"`
	MarginBalance          string `json:"marginBalance"`
	MaxWithdrawAmount      string `json:"maxWithdrawAmount"`
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"`
	PositionInitialMargin  string `json:"positionInitialMargin"`
	UnrealizedProfit       string `json:"unrealizedProfit"`
	WalletBalance          string `json:"walletBalance"`
}

//easyjson:json
type SubAccountFuturesSummaryCommon struct {
	Asset                       string `json:"asset"`
	TotalInitialMargin          string `json:"totalInitialMargin"`
	TotalMaintenanceMargin      string `json:"totalMaintenanceMargin"`
	TotalMarginBalance          string `json:"totalMarginBalance"`
	TotalOpenOrderInitialMargin string `json:"totalOpenOrderInitialMargin"`
	TotalPositionInitialMargin  string `json:"totalPositionInitialMargin"`
	TotalUnrealizedProfit       string `json:"totalUnrealizedProfit"`
	TotalWalletBalance          string `json:"totalWalletBalance"`
}

//easyjson:json
type SubAccountFuturesSummaryV1 struct {
	SubAccountFuturesSummaryCommon
	SubAccountList []SubAccountFuturesSummaryV1SubAccountList `json:"subAccountList"`
}

//easyjson:json
type SubAccountFuturesSummaryV1SubAccountList struct {
	Email string `json:"email"`
	SubAccountFuturesSummaryCommon
}

//easyjson:json
type SubAccountList struct {
	SubAccounts []SubAccount `json:"subAccounts"`
}

//easyjson:json
type SubAccountResponse struct {
	Email        string `json:"email"`
	SubAccountID string `json:"subaccountId"`
	Tag          string `json:"tag"`
}

// SubaccountAssetsResponse Query Sub-account Assets response
//
//easyjson:json
type SubaccountAssetsResponse struct {
	Balances []AssetBalance `json:"balances"`
}

//easyjson:json
type SubaccountDepositAddressResponse struct {
	Address string `json:"address"`
	Coin    string `json:"coin"`
	Tag     string `json:"tag"`
	URL     string `json:"url"`
}

// SubaccountSpotSummaryResponse Query Sub-account Spot Assets Summary response
//
//easyjson:json
type SubaccountSpotSummaryResponse struct {
	TotalCount                int64                       `json:"totalCount"`
	MasterAccountTotalAsset   string                      `json:"masterAccountTotalAsset"`
	SpotSubUserAssetBtcVoList []SpotSubUserAssetBtcVoList `json:"spotSubUserAssetBtcVoList"`
}

//easyjson:json
type SwapRecord struct {
	SwapId     int64          `json:"swapId"`
	SwapTime   int64          `json:"swapTime"`
	Status     SwappingStatus `json:"status"`
	QuoteAsset string         `json:"quoteAsset"`
	BaseAsset  string         `json:"baseAsset"`
	QuoteQty   string         `json:"quoteQty"`
	BaseQty    string         `json:"baseQty"`
	Price      string         `json:"price"`
	Fee        string         `json:"fee"`
}

//easyjson:json
type SwapResponse struct {
	SwapId int64 `json:"swapId"`
}

// Symbol market symbol
//
//easyjson:json
type Symbol struct {
	Symbol                     string                   `json:"symbol"`
	Status                     string                   `json:"status"`
	BaseAsset                  string                   `json:"baseAsset"`
	BaseAssetPrecision         int                      `json:"baseAssetPrecision"`
	QuoteAsset                 string                   `json:"quoteAsset"`
	QuotePrecision             int                      `json:"quotePrecision"`
	QuoteAssetPrecision        int                      `json:"quoteAssetPrecision"`
	BaseCommissionPrecision    int32                    `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision   int32                    `json:"quoteCommissionPrecision"`
	OrderTypes                 []string                 `json:"orderTypes"`
	IcebergAllowed             bool                     `json:"icebergAllowed"`
	OcoAllowed                 bool                     `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed bool                     `json:"quoteOrderQtyMarketAllowed"`
	IsSpotTradingAllowed       bool                     `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed     bool                     `json:"isMarginTradingAllowed"`
	Filters                    []map[string]interface{} `json:"filters"`
	Permissions                []string                 `json:"permissions"`
}

// SymbolPrice define symbol and price pair
//
//easyjson:json
type SymbolPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

//easyjson:json
type SymbolTicker struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	LastPrice          string `json:"lastPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstId            int64  `json:"firstId"`
	LastId             int64  `json:"lastId"`
	Count              int64  `json:"count"`
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
	IsBestMatch   bool   `json:"isBestMatch"`
	IsIsolated    bool   `json:"isIsolated"`
}

// TradeFeeDetails represents details about fees
//
//easyjson:json
type TradeFeeDetails struct {
	Symbol          string `json:"symbol"`
	MakerCommission string `json:"makerCommission"`
	TakerCommission string `json:"takerCommission"`
}

// TradeV3 define v3 trade info
//
//easyjson:json
type TradeV3 struct {
	ID              int64           `json:"id"`
	Symbol          string          `json:"symbol"`
	OrderID         int64           `json:"orderId"`
	OrderListId     int64           `json:"orderListId"`
	Price           decimal.Decimal `json:"price"`
	Quantity        decimal.Decimal `json:"qty"`
	QuoteQuantity   string          `json:"quoteQty"`
	Commission      decimal.Decimal `json:"commission"`
	CommissionAsset string          `json:"commissionAsset"`
	Time            int64           `json:"time"`
	IsBuyer         bool            `json:"isBuyer"`
	IsMaker         bool            `json:"isMaker"`
	IsBestMatch     bool            `json:"isBestMatch"`
	IsIsolated      bool            `json:"isIsolated"`
}

// Spot trading supports tracking stop orders
// Tracking stop loss sets an automatic trigger price based on market price using a new parameter trailingDelta
//
//easyjson:json
type TrailingDeltaFilter struct {
	MinTrailingAboveDelta int `json:"minTrailingAboveDelta"`
	MaxTrailingAboveDelta int `json:"maxTrailingAboveDelta"`
	MinTrailingBelowDelta int `json:"minTrailingBelowDelta"`
	MaxTrailingBelowDelta int `json:"maxTrailingBelowDelta"`
}

// TransactionResponse define transaction response
//
//easyjson:json
type TransactionResponse struct {
	TranID int64 `json:"tranId"`
}

//easyjson:json
type Transfer struct {
	From   string          `json:"from"`
	To     string          `json:"to"`
	Asset  string          `json:"asset"`
	Qty    decimal.Decimal `json:"qty"`
	TranId string          `json:"tranId"`
	//ClientTranId string          `json:"clientTranId"`
	Time int64 `json:"time"`
}

// TransferToSubAccountResponse define transfer to subaccount response
//
//easyjson:json
type TransferToSubAccountResponse struct {
	TxnID int64 `json:"txnId"`
}

//easyjson:json
type UniversalTransferServiceResponse struct {
	TxnId int64 `json:"txnId"`
}

//easyjson:json
type UpdateIPRestrictionSubAccountAPIKeyResponse struct {
	Status     string   `json:"status"`
	IpList     []string `json:"ipList"`
	UpdateTime int64    `json:"updateTime"`
	ApiKey     string   `json:"apiKey"`
}

// UserAsset define user assets of margin account
//
//easyjson:json
type UserAsset struct {
	Asset    string `json:"asset"`
	Borrowed string `json:"borrowed"`
	Free     string `json:"free"`
	Interest string `json:"interest"`
	Locked   string `json:"locked"`
	NetAsset string `json:"netAsset"`
}

// UserAssetDribblet represents one dust log row
//
//easyjson:json
type UserAssetDribblet struct {
	OperateTime              int64                     `json:"operateTime"`
	TotalTransferedAmount    string                    `json:"totalTransferedAmount"`    //Total transfered BNB amount for this exchange.
	TotalServiceChargeAmount string                    `json:"totalServiceChargeAmount"` //Total service charge amount for this exchange.
	TransID                  int64                     `json:"transId"`
	UserAssetDribbletDetails []UserAssetDribbletDetail `json:"userAssetDribbletDetails"` //Details of this exchange.
}

// DustLog represents one dust log informations
//
//easyjson:json
type UserAssetDribbletDetail struct {
	TransID             int    `json:"transId"`
	ServiceChargeAmount string `json:"serviceChargeAmount"`
	Amount              string `json:"amount"`
	OperateTime         int64  `json:"operateTime"` //The time of this exchange.
	TransferedAmount    string `json:"transferedAmount"`
	FromAsset           string `json:"fromAsset"`
}

//easyjson:json
type UserAssetRecord struct {
	Asset        string          `json:"asset"`
	Free         decimal.Decimal `json:"free"`
	Locked       decimal.Decimal `json:"locked"`
	Freeze       decimal.Decimal `json:"freeze"`
	Withdrawing  string          `json:"withdrawing"`
	Ipoable      string          `json:"ipoable"`
	BtcValuation string          `json:"btcValuation"`
}

// Withdraw represents a single withdraw entry.
//
//easyjson:json
type Withdraw struct {
	Address         string `json:"address"`
	Amount          string `json:"amount"`
	ApplyTime       string `json:"applyTime"`
	Coin            string `json:"coin"`
	ID              string `json:"id"`
	WithdrawOrderID string `json:"withdrawOrderId"`
	Network         string `json:"network"`
	TransferType    int    `json:"transferType"`
	Status          int    `json:"status"`
	TransactionFee  string `json:"transactionFee"`
	ConfirmNo       int32  `json:"confirmNo"`
	Info            string `json:"info"`
	TxID            string `json:"txId"`
}
