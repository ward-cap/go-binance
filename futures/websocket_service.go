package futures

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

// Endpoints
const (
	baseWsMainUrl       = "wss://fstream.binance.com/ws"
	baseCombinedMainURL = "wss://fstream.binance.com/stream?streams="
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 60
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = false
	// UseTestnet switch all the WS streams from production to the testnet
	UseTestnet = false
)

// getWsEndpoint return the base endpoint of the WS according the UseTestnet flag
func getWsEndpoint() string {
	return baseWsMainUrl
}

// getCombinedEndpoint return the base endpoint of the combined stream according the UseTestnet flag
func getCombinedEndpoint() string {
	return baseCombinedMainURL
}

// WsAggTradeEvent define websocket aggTrde event.
type WsAggTradeEvent struct {
	Event            string `json:"e"`
	Time             int64  `json:"E"`
	Symbol           string `json:"s"`
	AggregateTradeID int64  `json:"a"`
	Price            string `json:"p"`
	Quantity         string `json:"q"`
	FirstTradeID     int64  `json:"f"`
	LastTradeID      int64  `json:"l"`
	TradeTime        int64  `json:"T"`
	Maker            bool   `json:"m"`
}

// WsAggTradeHandler handle websocket that push trade information that is aggregated for a single taker order.
type WsAggTradeHandler func(event *WsAggTradeEvent)

// WsAggTradeServe serve websocket that push trade information that is aggregated for a single taker order.
func WsAggTradeServe(ctx context.Context, symbol string, handler WsAggTradeHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@aggTrade", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsAggTradeEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsCombinedAggTradeServe is similar to WsAggTradeServe, but it handles multiple symbols
func WsCombinedAggTradeServe(ctx context.Context, symbols []string, handler WsAggTradeHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@aggTrade", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		stream := j.Get("stream").MustString()
		data := j.Get("data").MustMap()

		symbol := strings.Split(stream, "@")[0]

		jsonData, _ := json.Marshal(data)

		event := new(WsAggTradeEvent)
		err = json.Unmarshal(jsonData, event)
		if err != nil {
			errHandler(err)
			return
		}
		event.Symbol = strings.ToUpper(symbol)

		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsMarkPriceEvent define websocket markPriceUpdate event.
type WsMarkPriceEvent struct {
	Event                string          `json:"e"`
	Time                 int64           `json:"E"`
	Symbol               string          `json:"s"`
	MarkPrice            decimal.Decimal `json:"p"`
	IndexPrice           string          `json:"i"`
	EstimatedSettlePrice string          `json:"P"`
	FundingRate          string          `json:"r"`
	NextFundingTime      int64           `json:"T"`
}

// WsMarkPriceHandler handle websocket that pushes price and funding rate for a single symbol.
type WsMarkPriceHandler func(event *WsMarkPriceEvent)

func wsMarkPriceServe(
	ctx context.Context,
	endpoint string,
	handler WsMarkPriceHandler,
	errHandler ErrHandler,
	logger *zap.SugaredLogger,
) (doneC chan struct{}, err error) {
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsMarkPriceEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsMarkPriceServe serve websocket that pushes price and funding rate for a single symbol.
func WsMarkPriceServe(
	ctx context.Context,
	symbol string,
	handler WsMarkPriceHandler,
	errHandler ErrHandler,
	logger *zap.SugaredLogger,
) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@markPrice", getWsEndpoint(), strings.ToLower(symbol))
	return wsMarkPriceServe(ctx, endpoint, handler, errHandler, logger)
}

// WsMarkPriceServeWithRate serve websocket that pushes price and funding rate for a single symbol and rate.
func WsMarkPriceServeWithRate(ctx context.Context, symbol string, rate time.Duration, handler WsMarkPriceHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	var rateStr string
	switch rate {
	case 3 * time.Second:
		rateStr = ""
	case 1 * time.Second:
		rateStr = "@1s"
	default:
		return nil, errors.New("Invalid rate")
	}
	endpoint := fmt.Sprintf("%s/%s@markPrice%s", getWsEndpoint(), strings.ToLower(symbol), rateStr)
	return wsMarkPriceServe(ctx, endpoint, handler, errHandler, logger)
}

func wsCombinedMarkPriceServe(ctx context.Context, endpoint string, handler WsMarkPriceHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		data := j.Get("data").MustMap()
		jsonData, _ := json.Marshal(data)

		event := new(WsMarkPriceEvent)
		err = json.Unmarshal(jsonData, event)
		if err != nil {
			errHandler(err)
			return
		}

		handler(event)
	}

	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsCombinedMarkPriceServe is similar to WsMarkPriceServe, but it handles multiple symbols
func WsCombinedMarkPriceServe(
	ctx context.Context,
	symbols []string,
	handler WsMarkPriceHandler,
	errHandler ErrHandler,
	logger *zap.SugaredLogger,
) (doneC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@markPrice", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]

	return wsCombinedMarkPriceServe(ctx, endpoint, handler, errHandler, logger)
}

// WsCombinedMarkPriceServeWithRate is similar to WsMarkPriceServeWithRate, but it for multiple symbols
func WsCombinedMarkPriceServeWithRate(
	ctx context.Context,
	symbolLevels map[string]time.Duration,
	handler WsMarkPriceHandler,
	errHandler ErrHandler,
	logger *zap.SugaredLogger,
) (doneC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for symbol, rate := range symbolLevels {
		var rateStr string
		switch rate {
		case 3 * time.Second:
			rateStr = ""
		case 1 * time.Second:
			rateStr = "@1s"
		default:
			return nil, fmt.Errorf("invalid rate. Symbol %s (rate %d)", symbol, rate)
		}

		endpoint += fmt.Sprintf("%s@markPrice%s", strings.ToLower(symbol), rateStr) + "/"
	}

	endpoint = endpoint[:len(endpoint)-1]

	return wsCombinedMarkPriceServe(ctx, endpoint, handler, errHandler, logger)
}

// WsAllMarkPriceEvent defines an array of websocket markPriceUpdate events.
type WsAllMarkPriceEvent []*WsMarkPriceEvent

// WsAllMarkPriceHandler handle websocket that pushes price and funding rate for all symbol.
type WsAllMarkPriceHandler func(event WsAllMarkPriceEvent)

func wsAllMarkPriceServe(ctx context.Context, endpoint string, handler WsAllMarkPriceHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMarkPriceEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsAllMarkPriceServe serve websocket that pushes price and funding rate for all symbol.
func WsAllMarkPriceServe(
	ctx context.Context,
	handler WsAllMarkPriceHandler,
	errHandler ErrHandler,
	logger *zap.SugaredLogger,
) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!markPrice@arr", getWsEndpoint())
	return wsAllMarkPriceServe(ctx, endpoint, handler, errHandler, logger)
}

// WsAllMarkPriceServeWithRate serve websocket that pushes price and funding rate for all symbol and rate.
func WsAllMarkPriceServeWithRate(
	ctx context.Context,
	rate time.Duration,
	handler WsAllMarkPriceHandler,
	errHandler ErrHandler,
	logger *zap.SugaredLogger,
) (doneC chan struct{}, err error) {
	var rateStr string
	switch rate {
	case 3 * time.Second:
		rateStr = ""
	case 1 * time.Second:
		rateStr = "@1s"
	default:
		return nil, errors.New("Invalid rate")
	}
	endpoint := fmt.Sprintf("%s/!markPrice@arr%s", getWsEndpoint(), rateStr)
	return wsAllMarkPriceServe(ctx, endpoint, handler, errHandler, logger)
}

// WsKlineEvent define websocket kline event
type WsKlineEvent struct {
	Event  string  `json:"e"`
	Time   int64   `json:"E"`
	Symbol string  `json:"s"`
	Kline  WsKline `json:"k"`
}

// WsKline define websocket kline
type WsKline struct {
	StartTime            int64  `json:"t"`
	EndTime              int64  `json:"T"`
	Symbol               string `json:"s"`
	Interval             string `json:"i"`
	FirstTradeID         int64  `json:"f"`
	LastTradeID          int64  `json:"L"`
	Open                 string `json:"o"`
	Close                string `json:"c"`
	High                 string `json:"h"`
	Low                  string `json:"l"`
	Volume               string `json:"v"`
	TradeNum             int64  `json:"n"`
	IsFinal              bool   `json:"x"`
	QuoteVolume          string `json:"q"`
	ActiveBuyVolume      string `json:"V"`
	ActiveBuyQuoteVolume string `json:"Q"`
}

// WsKlineHandler handle websocket kline event
type WsKlineHandler func(event *WsKlineEvent)

// WsKlineServe serve websocket kline handler with a symbol and interval like 15m, 30s
func WsKlineServe(ctx context.Context, symbol string, interval string, handler WsKlineHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@kline_%s", getWsEndpoint(), strings.ToLower(symbol), interval)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsKlineEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsCombinedKlineServe is similar to WsKlineServe, but it handles multiple symbols with it interval
func WsCombinedKlineServe(
	ctx context.Context,
	symbolIntervalPair map[string]string,
	handler WsKlineHandler,
	errHandler ErrHandler,
	logger *zap.SugaredLogger,
) (doneC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for symbol, interval := range symbolIntervalPair {
		endpoint += fmt.Sprintf("%s@kline_%s", strings.ToLower(symbol), interval) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		stream := j.Get("stream").MustString()
		data := j.Get("data").MustMap()

		symbol := strings.Split(stream, "@")[0]

		jsonData, _ := json.Marshal(data)

		event := new(WsKlineEvent)
		err = json.Unmarshal(jsonData, event)
		if err != nil {
			errHandler(err)
			return
		}
		event.Symbol = strings.ToUpper(symbol)

		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsContinuousKlineEvent define websocket continuous kline event
type WsContinuousKlineEvent struct {
	Event        string            `json:"e"`
	Time         int64             `json:"E"`
	PairSymbol   string            `json:"ps"`
	ContractType string            `json:"ct"`
	Kline        WsContinuousKline `json:"k"`
}

// WsContinuousKline define websocket continuous kline
type WsContinuousKline struct {
	StartTime            int64  `json:"t"`
	EndTime              int64  `json:"T"`
	Interval             string `json:"i"`
	FirstTradeID         int64  `json:"f"`
	LastTradeID          int64  `json:"L"`
	Open                 string `json:"o"`
	Close                string `json:"c"`
	High                 string `json:"h"`
	Low                  string `json:"l"`
	Volume               string `json:"v"`
	TradeNum             int64  `json:"n"`
	IsFinal              bool   `json:"x"`
	QuoteVolume          string `json:"q"`
	ActiveBuyVolume      string `json:"V"`
	ActiveBuyQuoteVolume string `json:"Q"`
}

// WsContinuousKlineSubcribeArgs used with WsContinuousKlineServe or WsCombinedContinuousKlineServe
type WsContinuousKlineSubcribeArgs struct {
	Pair         string
	ContractType string
	Interval     string
}

// WsContinuousKlineHandler handle websocket continuous kline event
type WsContinuousKlineHandler func(event *WsContinuousKlineEvent)

// WsContinuousKlineServe serve websocket continuous kline handler with a pair and contractType and interval like 15m, 30s
func WsContinuousKlineServe(
	ctx context.Context,
	subscribeArgs *WsContinuousKlineSubcribeArgs,
	handler WsContinuousKlineHandler,
	errHandler ErrHandler,
	logger *zap.SugaredLogger,
) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s_%s@continuousKline_%s", getWsEndpoint(), strings.ToLower(subscribeArgs.Pair),
		strings.ToLower(subscribeArgs.ContractType), subscribeArgs.Interval)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsContinuousKlineEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

type WsRPC struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
	ID     int      `json:"id"`
}

// WsCombinedContinuousKlineServe is similar to WsContinuousKlineServe, but it handles multiple pairs of different contractType with its interval
func WsCombinedContinuousKlineServe(
	ctx context.Context,
	subscribeArgsList []*WsContinuousKlineSubcribeArgs,
	handler WsContinuousKlineHandler,
	errHandler ErrHandler,
	logger *zap.SugaredLogger,
) (doneC chan struct{}, err error) {
	const endpoint = "wss://fstream.binance.com/stream"

	rpc := WsRPC{
		Method: "SUBSCRIBE",
		ID:     1,
	}

	for _, val := range subscribeArgsList {
		subscribe := fmt.Sprintf("%s_%s@continuousKline_%s",
			strings.ToLower(val.Pair),
			strings.ToLower(val.ContractType),
			val.Interval,
		)
		rpc.Params = append(rpc.Params, subscribe)
	}

	rpcMsg, err := json.Marshal(rpc)
	if err != nil {
		return nil, err
	}

	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}

		data := j.Get("data").MustMap()

		jsonData, _ := json.Marshal(data)

		event := new(WsContinuousKlineEvent)
		err = json.Unmarshal(jsonData, event)
		if err != nil {
			errHandler(err)
			return
		}

		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger, rpcMsg)
}

// WsMiniMarketTickerEvent define websocket mini market ticker event.
type WsMiniMarketTickerEvent struct {
	Event       string `json:"e"`
	Time        int64  `json:"E"`
	Symbol      string `json:"s"`
	ClosePrice  string `json:"c"`
	OpenPrice   string `json:"o"`
	HighPrice   string `json:"h"`
	LowPrice    string `json:"l"`
	Volume      string `json:"v"`
	QuoteVolume string `json:"q"`
}

// WsMiniMarketTickerHandler handle websocket that pushes 24hr rolling window mini-ticker statistics for a single symbol.
type WsMiniMarketTickerHandler func(event *WsMiniMarketTickerEvent)

// WsMiniMarketTickerServe serve websocket that pushes 24hr rolling window mini-ticker statistics for a single symbol.
func WsMiniMarketTickerServe(ctx context.Context, symbol string, handler WsMiniMarketTickerHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@miniTicker", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsMiniMarketTickerEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsAllMiniMarketTickerEvent define an array of websocket mini market ticker events.
type WsAllMiniMarketTickerEvent []*WsMiniMarketTickerEvent

// WsAllMiniMarketTickerHandler handle websocket that pushes price and funding rate for all markets.
type WsAllMiniMarketTickerHandler func(event WsAllMiniMarketTickerEvent)

// WsAllMiniMarketTickerServe serve websocket that pushes price and funding rate for all markets.
func WsAllMiniMarketTickerServe(ctx context.Context, handler WsAllMiniMarketTickerHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!miniTicker@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMiniMarketTickerEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsMarketTickerEvent define websocket market ticker event.
type WsMarketTickerEvent struct {
	Event              string `json:"e"`
	Time               int64  `json:"E"`
	Symbol             string `json:"s"`
	PriceChange        string `json:"p"`
	PriceChangePercent string `json:"P"`
	WeightedAvgPrice   string `json:"w"`
	ClosePrice         string `json:"c"`
	CloseQty           string `json:"Q"`
	OpenPrice          string `json:"o"`
	HighPrice          string `json:"h"`
	LowPrice           string `json:"l"`
	BaseVolume         string `json:"v"`
	QuoteVolume        string `json:"q"`
	OpenTime           int64  `json:"O"`
	CloseTime          int64  `json:"C"`
	FirstID            int64  `json:"F"`
	LastID             int64  `json:"L"`
	TradeCount         int64  `json:"n"`
}

// WsMarketTickerHandler handle websocket that pushes 24hr rolling window mini-ticker statistics for a single symbol.
type WsMarketTickerHandler func(event *WsMarketTickerEvent)

// WsMarketTickerServe serve websocket that pushes 24hr rolling window mini-ticker statistics for a single symbol.
func WsMarketTickerServe(ctx context.Context, symbol string, handler WsMarketTickerHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@ticker", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsMarketTickerEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsAllMarketTickerEvent define an array of websocket mini ticker events.
type WsAllMarketTickerEvent []*WsMarketTickerEvent

// WsAllMarketTickerHandler handle websocket that pushes price and funding rate for all markets.
type WsAllMarketTickerHandler func(event WsAllMarketTickerEvent)

// WsAllMarketTickerServe serve websocket that pushes price and funding rate for all markets.
func WsAllMarketTickerServe(ctx context.Context, handler WsAllMarketTickerHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!ticker@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMarketTickerEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsBookTickerEvent define websocket best book ticker event.
type WsBookTickerEvent struct {
	Event           string `json:"e"`
	UpdateID        int64  `json:"u"`
	Time            int64  `json:"E"`
	TransactionTime int64  `json:"T"`
	Symbol          string `json:"s"`
	BestBidPrice    string `json:"b"`
	BestBidQty      string `json:"B"`
	BestAskPrice    string `json:"a"`
	BestAskQty      string `json:"A"`
}

// WsBookTickerHandler handle websocket that pushes updates to the best bid or ask price or quantity in real-time for a specified symbol.
type WsBookTickerHandler func(event *WsBookTickerEvent)

// WsBookTickerServe serve websocket that pushes updates to the best bid or ask price or quantity in real-time for a specified symbol.
func WsBookTickerServe(ctx context.Context, symbol string, handler WsBookTickerHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@bookTicker", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsBookTickerEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsAllBookTickerServe serve websocket that pushes updates to the best bid or ask price or quantity in real-time for all symbols.
func WsAllBookTickerServe(ctx context.Context, handler WsBookTickerHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!bookTicker", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsBookTickerEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsLiquidationOrderEvent define websocket liquidation order event.
type WsLiquidationOrderEvent struct {
	Event            string             `json:"e"`
	Time             int64              `json:"E"`
	LiquidationOrder WsLiquidationOrder `json:"o"`
}

// WsLiquidationOrder define websocket liquidation order.
type WsLiquidationOrder struct {
	Symbol               string          `json:"s"`
	Side                 SideType        `json:"S"`
	OrderType            OrderType       `json:"o"`
	TimeInForce          TimeInForceType `json:"f"`
	OrigQuantity         string          `json:"q"`
	Price                string          `json:"p"`
	AvgPrice             string          `json:"ap"`
	OrderStatus          OrderStatusType `json:"X"`
	LastFilledQty        string          `json:"l"`
	AccumulatedFilledQty string          `json:"z"`
	TradeTime            int64           `json:"T"`
}

// WsLiquidationOrderHandler handle websocket that pushes force liquidation order information for specific symbol.
type WsLiquidationOrderHandler func(event *WsLiquidationOrderEvent)

// WsLiquidationOrderServe serve websocket that pushes force liquidation order information for specific symbol.
func WsLiquidationOrderServe(ctx context.Context, symbol string, handler WsLiquidationOrderHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@forceOrder", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsLiquidationOrderEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsAllLiquidationOrderServe serve websocket that pushes force liquidation order information for all symbols.
func WsAllLiquidationOrderServe(ctx context.Context, handler WsLiquidationOrderHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!forceOrder@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsLiquidationOrderEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsDepthEvent define websocket depth book event
type WsDepthEvent struct {
	Event            string `json:"e"`
	Time             int64  `json:"E"`
	TransactionTime  int64  `json:"T"`
	Symbol           string `json:"s"`
	FirstUpdateID    int64  `json:"U"`
	LastUpdateID     int64  `json:"u"`
	PrevLastUpdateID int64  `json:"pu"`
	Bids             []Bid  `json:"b"`
	Asks             []Ask  `json:"a"`
}

// WsDepthHandler handle websocket depth event
type WsDepthHandler func(event *WsDepthEvent)

func wsPartialDepthServe(ctx context.Context, symbol string, levels int, rate *time.Duration, handler WsDepthHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	if levels != 5 && levels != 10 && levels != 20 {
		return nil, errors.New("Invalid levels")
	}
	levelsStr := fmt.Sprintf("%d", levels)
	return wsDepthServe(ctx, symbol, levelsStr, rate, handler, errHandler, logger)
}

// WsPartialDepthServe serve websocket partial depth handler.
func WsPartialDepthServe(ctx context.Context, symbol string, levels int, handler WsDepthHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	return wsPartialDepthServe(ctx, symbol, levels, nil, handler, errHandler, logger)
}

// WsPartialDepthServeWithRate serve websocket partial depth handler with rate.
func WsPartialDepthServeWithRate(ctx context.Context, symbol string, levels int, rate time.Duration, handler WsDepthHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	return wsPartialDepthServe(ctx, symbol, levels, &rate, handler, errHandler, logger)
}

// WsDiffDepthServe serve websocket diff. depth handler.
func WsDiffDepthServe(ctx context.Context, symbol string, handler WsDepthHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	return wsDepthServe(ctx, symbol, "", nil, handler, errHandler, logger)
}

// WsCombinedDepthServe is similar to WsPartialDepthServe, but it for multiple symbols
func WsCombinedDepthServe(
	ctx context.Context,
	symbolLevels map[string]string,
	handler WsDepthHandler,
	errHandler ErrHandler,
	logger *zap.SugaredLogger,
) (doneC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for s, l := range symbolLevels {
		endpoint += fmt.Sprintf("%s@depth%s", strings.ToLower(s), l) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsDepthEvent)
		data := j.Get("data").MustMap()
		event.Event = data["e"].(string)
		event.Time, _ = data["E"].(json.Number).Int64()
		event.TransactionTime, _ = data["T"].(json.Number).Int64()
		event.Symbol = data["s"].(string)
		event.FirstUpdateID, _ = data["U"].(json.Number).Int64()
		event.LastUpdateID, _ = data["u"].(json.Number).Int64()
		event.PrevLastUpdateID, _ = data["pu"].(json.Number).Int64()
		bidsLen := len(data["b"].([]interface{}))
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := data["b"].([]interface{})[i].([]interface{})
			event.Bids[i] = Bid{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		asksLen := len(data["a"].([]interface{}))
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {
			item := data["a"].([]interface{})[i].([]interface{})
			event.Asks[i] = Ask{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsCombinedDiffDepthServe is similar to WsDiffDepthServe, but it for multiple symbols
func WsCombinedDiffDepthServe(
	ctx context.Context,
	symbols []string,
	handler WsDepthHandler,
	errHandler ErrHandler,
	logger *zap.SugaredLogger,
) (doneC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@depth", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsDepthEvent)
		data := j.Get("data").MustMap()
		event.Event = data["e"].(string)
		event.Time, _ = data["E"].(json.Number).Int64()
		event.TransactionTime, _ = data["T"].(json.Number).Int64()
		event.Symbol = data["s"].(string)
		event.FirstUpdateID, _ = data["U"].(json.Number).Int64()
		event.LastUpdateID, _ = data["u"].(json.Number).Int64()
		event.PrevLastUpdateID, _ = data["pu"].(json.Number).Int64()
		bidsLen := len(data["b"].([]interface{}))
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := data["b"].([]interface{})[i].([]interface{})
			event.Bids[i] = Bid{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		asksLen := len(data["a"].([]interface{}))
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {
			item := data["a"].([]interface{})[i].([]interface{})
			event.Asks[i] = Ask{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsDiffDepthServeWithRate serve websocket diff. depth handler with rate.
func WsDiffDepthServeWithRate(ctx context.Context, symbol string, rate time.Duration, handler WsDepthHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	return wsDepthServe(ctx, symbol, "", &rate, handler, errHandler, logger)
}

func wsDepthServe(ctx context.Context, symbol string, levels string, rate *time.Duration, handler WsDepthHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	var rateStr string
	if rate != nil {
		switch *rate {
		case 250 * time.Millisecond:
			rateStr = ""
		case 500 * time.Millisecond:
			rateStr = "@500ms"
		case 100 * time.Millisecond:
			rateStr = "@100ms"
		default:
			return nil, errors.New("Invalid rate")
		}
	}
	endpoint := fmt.Sprintf("%s/%s@depth%s%s", getWsEndpoint(), strings.ToLower(symbol), levels, rateStr)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsDepthEvent)
		event.Event = j.Get("e").MustString()
		event.Time = j.Get("E").MustInt64()
		event.TransactionTime = j.Get("T").MustInt64()
		event.Symbol = j.Get("s").MustString()
		event.FirstUpdateID = j.Get("U").MustInt64()
		event.LastUpdateID = j.Get("u").MustInt64()
		event.PrevLastUpdateID = j.Get("pu").MustInt64()
		bidsLen := len(j.Get("b").MustArray())
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := j.Get("b").GetIndex(i)
			event.Bids[i] = Bid{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		asksLen := len(j.Get("a").MustArray())
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {
			item := j.Get("a").GetIndex(i)
			event.Asks[i] = Ask{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsBLVTInfoEvent define websocket BLVT info event
type WsBLVTInfoEvent struct {
	Event          string         `json:"e"`
	Time           int64          `json:"E"`
	Symbol         string         `json:"s"`
	Issued         float64        `json:"m"`
	Baskets        []WsBLVTBasket `json:"b"`
	Nav            float64        `json:"n"`
	Leverage       float64        `json:"l"`
	TargetLeverage int64          `json:"t"`
	FundingRate    float64        `json:"f"`
}

// WsBLVTBasket define websocket BLVT basket
type WsBLVTBasket struct {
	Symbol   string `json:"s"`
	Position int64  `json:"n"`
}

// WsBLVTInfoHandler handle websocket BLVT event
type WsBLVTInfoHandler func(event *WsBLVTInfoEvent)

// WsBLVTInfoServe serve BLVT info stream
func WsBLVTInfoServe(ctx context.Context, name string, handler WsBLVTInfoHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@tokenNav", getWsEndpoint(), strings.ToUpper(name))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsBLVTInfoEvent)
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsBLVTKlineEvent define BLVT kline event
type WsBLVTKlineEvent struct {
	Event  string      `json:"e"`
	Time   int64       `json:"E"`
	Symbol string      `json:"s"`
	Kline  WsBLVTKline `json:"k"`
}

// WsBLVTKline BLVT kline
type WsBLVTKline struct {
	StartTime       int64  `json:"t"`
	CloseTime       int64  `json:"T"`
	Symbol          string `json:"s"`
	Interval        string `json:"i"`
	FirstUpdateTime int64  `json:"f"`
	LastUpdateTime  int64  `json:"L"`
	OpenPrice       string `json:"o"`
	ClosePrice      string `json:"c"`
	HighPrice       string `json:"h"`
	LowPrice        string `json:"l"`
	Leverage        string `json:"v"`
	Count           int64  `json:"n"`
}

// WsBLVTKlineHandler BLVT kline handler
type WsBLVTKlineHandler func(event *WsBLVTKlineEvent)

// WsBLVTKlineServe serve BLVT kline stream
func WsBLVTKlineServe(ctx context.Context, name string, interval string, handler WsBLVTKlineHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@nav_Kline_%s", getWsEndpoint(), strings.ToUpper(name), interval)
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsBLVTKlineEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

// WsCompositeIndexEvent websocket composite index event
type WsCompositeIndexEvent struct {
	Event       string          `json:"e"`
	Time        int64           `json:"E"`
	Symbol      string          `json:"s"`
	Price       string          `json:"p"`
	Composition []WsComposition `json:"c"`
}

// WsComposition websocket composite index event composition
type WsComposition struct {
	BaseAsset    string `json:"b"`
	WeightQty    string `json:"w"`
	WeighPercent string `json:"W"`
}

// WsCompositeIndexHandler websocket composite index handler
type WsCompositeIndexHandler func(event *WsCompositeIndexEvent)

// WsCompositiveIndexServe serve composite index information for index symbols
func WsCompositiveIndexServe(ctx context.Context, symbol string, handler WsCompositeIndexHandler, errHandler ErrHandler, logger *zap.SugaredLogger) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@compositeIndex", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsCompositeIndexEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil, logger)
}

type WsUserDataEventType struct {
	Event UserDataEventType `json:"e"`
	Time  int64             `json:"E"`
}

// WsUserDataEvent define user data event
type WsUserDataEvent struct {
	WsUserDataEventType

	CrossWalletBalance  string                 `json:"cw"`
	MarginCallPositions []WsPosition           `json:"p,omitempty"`
	TransactionTime     int64                  `json:"T"`
	AccountUpdate       *WsAccountUpdate       `json:"a,omitempty"`
	OrderTradeUpdate    *WsOrderTradeUpdate    `json:"o,omitempty"`
	AccountConfigUpdate *WsAccountConfigUpdate `json:"ac,omitempty"`
}

// WsAccountUpdate define account update
type WsAccountUpdate struct {
	Reason    UserDataEventReasonType `json:"m"`
	Balances  []WsBalance             `json:"B"`
	Positions []WsPosition            `json:"P"`
}

// WsBalance define balance
type WsBalance struct {
	Asset              string              `json:"a"`
	Balance            decimal.Decimal     `json:"wb"`
	CrossWalletBalance string              `json:"cw"`
	ChangeBalance      decimal.NullDecimal `json:"bc"`
}

// WsPosition define position
type WsPosition struct {
	Symbol                    string           `json:"s"`
	Side                      PositionSideType `json:"ps"`
	Amount                    decimal.Decimal  `json:"pa"`
	MarginType                MarginType       `json:"mt"`
	IsolatedWallet            string           `json:"iw"`
	EntryPrice                decimal.Decimal  `json:"ep"`
	MarkPrice                 string           `json:"mp"`
	UnrealizedPnL             string           `json:"up"`
	AccumulatedRealized       string           `json:"cr"`
	MaintenanceMarginRequired string           `json:"mm"`
}

// WsOrderTradeUpdate define order trade update
type WsOrderTradeUpdate struct {
	Symbol               string              `json:"s"`
	ClientOrderID        string              `json:"c"`
	Side                 SideType            `json:"S"`
	Type                 OrderType           `json:"o"`
	TimeInForce          TimeInForceType     `json:"f"`
	OriginalQty          decimal.Decimal     `json:"q"`
	OriginalPrice        decimal.Decimal     `json:"p"`
	AveragePrice         string              `json:"ap"`
	StopPrice            decimal.NullDecimal `json:"sp"`
	ExecutionType        OrderExecutionType  `json:"x"`
	Status               OrderStatusType     `json:"X"`
	ID                   int64               `json:"i"`
	AlgoID               int64               `json:"aid"`
	LastFilledQty        decimal.Decimal     `json:"l"`
	AccumulatedFilledQty decimal.NullDecimal `json:"z"`
	LastFilledPrice      decimal.Decimal     `json:"L"`
	CommissionAsset      string              `json:"N"`
	Commission           decimal.Decimal     `json:"n"`
	TradeTime            int64               `json:"T"`
	TradeID              int64               `json:"t"`
	BidsNotional         string              `json:"b"`
	AsksNotional         string              `json:"a"`
	IsMaker              bool                `json:"m"`
	IsReduceOnly         bool                `json:"R"`
	WorkingType          WorkingType         `json:"wt"`
	OriginalType         OrderType           `json:"ot"`
	PositionSide         PositionSideType    `json:"ps"`
	IsClosingPosition    bool                `json:"cp"`
	ActivationPrice      decimal.NullDecimal `json:"AP"`
	CallbackRate         string              `json:"cr"`
	RealizedPnL          string              `json:"rp"`
	AlgoType             string              `json:"at"`
}

// WsAccountConfigUpdate define account config update
type WsAccountConfigUpdate struct {
	Symbol   string `json:"s"`
	Leverage int64  `json:"l"`
}

// WsUserDataHandler handle WsUserDataEvent
type WsUserDataHandler func([]byte)

// WsUserDataServe serve user data handler with listen key
func WsUserDataServe(
	ctx context.Context,
	listenKey string,
	handler func([]byte), // WsUserDataEvent
	errHandler ErrHandler,
	proxyFunc DialFunc,
	logger *zap.SugaredLogger,
) (doneC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s", getWsEndpoint(), listenKey)
	cfg := newWsConfig(endpoint)
	//wsHandler := func(message []byte) {
	//	//event := new(WsUserDataEvent)
	//	//err := json.Unmarshal(message, event)
	//	//if err != nil {
	//	//	errHandler(fmt.Errorf("error: %v. raw: %s", err, message))
	//	//	return
	//	//}
	//	handler(message)
	//}
	return wsServe(ctx, cfg, handler, errHandler, proxyFunc, logger)
}
