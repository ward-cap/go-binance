package binance

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/ward-cap/go-binance/futures"
	"strings"
	"time"

	stdjson "encoding/json"
)

// Endpoints
const (
	baseWsMainURL       = "wss://stream.binance.com:9443/ws"
	baseCombinedMainURL = "wss://stream.binance.com:9443/stream?streams="
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 60
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = false
)

// getWsEndpoint return the base endpoint of the WS according the UseTestnet flag
func getWsEndpoint() string {
	return baseWsMainURL
}

// getCombinedEndpoint return the base endpoint of the combined stream according the UseTestnet flag
func getCombinedEndpoint() string {
	return baseCombinedMainURL
}

// WsPartialDepthEvent define websocket partial depth book event
type WsPartialDepthEvent struct {
	Symbol       string
	LastUpdateID int64 `json:"lastUpdateId"`
	Bids         []Bid `json:"bids"`
	Asks         []Ask `json:"asks"`
}

// WsPartialDepthHandler handle websocket partial depth event
type WsPartialDepthHandler func(event *WsPartialDepthEvent)

// WsPartialDepthServe serve websocket partial depth handler with a symbol, using 1sec updates
func WsPartialDepthServe(ctx context.Context, symbol string, levels string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth%s", getWsEndpoint(), strings.ToLower(symbol), levels)
	return wsPartialDepthServe(ctx, endpoint, symbol, handler, errHandler)
}

// WsPartialDepthServe100Ms serve websocket partial depth handler with a symbol, using 100msec updates
func WsPartialDepthServe100Ms(ctx context.Context, symbol string, levels string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth%s@100ms", getWsEndpoint(), strings.ToLower(symbol), levels)
	return wsPartialDepthServe(ctx, endpoint, symbol, handler, errHandler)
}

// WsPartialDepthServe serve websocket partial depth handler with a symbol
func wsPartialDepthServe(ctx context.Context, endpoint string, symbol string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsPartialDepthEvent)
		event.Symbol = symbol
		event.LastUpdateID = j.Get("lastUpdateId").MustInt64()
		bidsLen := len(j.Get("bids").MustArray())
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := j.Get("bids").GetIndex(i)
			event.Bids[i] = Bid{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		asksLen := len(j.Get("asks").MustArray())
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {
			item := j.Get("asks").GetIndex(i)
			event.Asks[i] = Ask{
				Price:    item.GetIndex(0).MustString(),
				Quantity: item.GetIndex(1).MustString(),
			}
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsCombinedPartialDepthServe is similar to WsPartialDepthServe, but it for multiple symbols
func WsCombinedPartialDepthServe(ctx context.Context, symbolLevels map[string]string, handler WsPartialDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
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
		event := new(WsPartialDepthEvent)
		stream := j.Get("stream").MustString()
		symbol := strings.Split(stream, "@")[0]
		event.Symbol = strings.ToUpper(symbol)
		data := j.Get("data").MustMap()
		event.LastUpdateID, _ = data["lastUpdateId"].(stdjson.Number).Int64()
		bidsLen := len(data["bids"].([]interface{}))
		event.Bids = make([]Bid, bidsLen)
		for i := 0; i < bidsLen; i++ {
			item := data["bids"].([]interface{})[i].([]interface{})
			event.Bids[i] = Bid{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		asksLen := len(data["asks"].([]interface{}))
		event.Asks = make([]Ask, asksLen)
		for i := 0; i < asksLen; i++ {

			item := data["asks"].([]interface{})[i].([]interface{})
			event.Asks[i] = Ask{
				Price:    item[0].(string),
				Quantity: item[1].(string),
			}
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsDepthHandler handle websocket depth event
type WsDepthHandler func(event *WsDepthEvent)

// WsDepthServe serve websocket depth handler with a symbol, using 1sec updates
func WsDepthServe(ctx context.Context, symbol string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth", getWsEndpoint(), strings.ToLower(symbol))
	return wsDepthServe(ctx, endpoint, handler, errHandler)
}

// WsDepthServe100Ms serve websocket depth handler with a symbol, using 100msec updates
func WsDepthServe100Ms(ctx context.Context, symbol string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@depth@100ms", getWsEndpoint(), strings.ToLower(symbol))
	return wsDepthServe(ctx, endpoint, handler, errHandler)
}

// WsDepthServe serve websocket depth handler with an arbitrary endpoint address
func wsDepthServe(ctx context.Context, endpoint string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
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
		event.Symbol = j.Get("s").MustString()
		event.LastUpdateID = j.Get("u").MustInt64()
		event.FirstUpdateID = j.Get("U").MustInt64()
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
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsDepthEvent define websocket depth event
type WsDepthEvent struct {
	Event         string `json:"e"`
	Time          int64  `json:"E"`
	Symbol        string `json:"s"`
	LastUpdateID  int64  `json:"u"`
	FirstUpdateID int64  `json:"U"`
	Bids          []Bid  `json:"b"`
	Asks          []Ask  `json:"a"`
}

// WsCombinedDepthServe is similar to WsDepthServe, but it for multiple symbols
func WsCombinedDepthServe(ctx context.Context, symbols []string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@depth", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	return wsCombinedDepthServe(ctx, endpoint, handler, errHandler)
}

func WsCombinedDepthServe100Ms(ctx context.Context, symbols []string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@depth@100ms", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	return wsCombinedDepthServe(ctx, endpoint, handler, errHandler)
}

func wsCombinedDepthServe(ctx context.Context, endpoint string, handler WsDepthHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		j, err := newJSON(message)
		if err != nil {
			errHandler(err)
			return
		}
		event := new(WsDepthEvent)
		stream := j.Get("stream").MustString()
		symbol := strings.Split(stream, "@")[0]
		event.Symbol = strings.ToUpper(symbol)
		data := j.Get("data").MustMap()
		event.Time, _ = data["E"].(stdjson.Number).Int64()
		event.LastUpdateID, _ = data["u"].(stdjson.Number).Int64()
		event.FirstUpdateID, _ = data["U"].(stdjson.Number).Int64()
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
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsKlineHandler handle websocket kline event
type WsKlineHandler func(event *WsKlineEvent)

// WsCombinedKlineServe is similar to WsKlineServe, but it handles multiple symbols with it interval
func WsCombinedKlineServe(ctx context.Context, symbolIntervalPair map[string]string, handler WsKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
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
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsKlineServe serve websocket kline handler with a symbol and interval like 15m, 30s
func WsKlineServe(ctx context.Context, symbol string, interval string, handler WsKlineHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
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
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
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

// WsAggTradeHandler handle websocket aggregate trade event
type WsAggTradeHandler func(event *WsAggTradeEvent)

// WsAggTradeServe serve websocket aggregate handler with a symbol
func WsAggTradeServe(ctx context.Context, symbol string, handler WsAggTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@aggTrade", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsAggTradeEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsCombinedAggTradeServe is similar to WsAggTradeServe, but it handles multiple symbolx
func WsCombinedAggTradeServe(ctx context.Context, symbols []string, handler WsAggTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for s := range symbols {
		endpoint += fmt.Sprintf("%s@aggTrade", strings.ToLower(symbols[s])) + "/"
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
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsAggTradeEvent define websocket aggregate trade event
type WsAggTradeEvent struct {
	Event                 string `json:"e"`
	Time                  int64  `json:"E"`
	Symbol                string `json:"s"`
	AggTradeID            int64  `json:"a"`
	Price                 string `json:"p"`
	Quantity              string `json:"q"`
	FirstBreakdownTradeID int64  `json:"f"`
	LastBreakdownTradeID  int64  `json:"l"`
	TradeTime             int64  `json:"T"`
	IsBuyerMaker          bool   `json:"m"`
	Placeholder           bool   `json:"M"` // add this field to avoid case insensitive unmarshaling
}

// WsTradeHandler handle websocket trade event
type WsTradeHandler func(event *WsTradeEvent)
type WsCombinedTradeHandler func(event *WsCombinedTradeEvent)

// WsTradeServe serve websocket handler with a symbol
func WsTradeServe(ctx context.Context, symbol string, handler WsTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@trade", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsTradeEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

func WsCombinedTradeServe(ctx context.Context, symbols []string, handler WsCombinedTradeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@trade/", strings.ToLower(s))
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsCombinedTradeEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsTradeEvent define websocket trade event
type WsTradeEvent struct {
	Event         string          `json:"e"`
	Time          int64           `json:"E"`
	Symbol        string          `json:"s"`
	TradeID       int64           `json:"t"`
	Price         decimal.Decimal `json:"p"`
	Quantity      string          `json:"q"`
	BuyerOrderID  int64           `json:"b"`
	SellerOrderID int64           `json:"a"`
	TradeTime     int64           `json:"T"`
	IsBuyerMaker  bool            `json:"m"`
	Placeholder   bool            `json:"M"` // add this field to avoid case insensitive unmarshaling
}

type WsCombinedTradeEvent struct {
	Stream string       `json:"stream"`
	Data   WsTradeEvent `json:"data"`
}

type WsUserDataEventType struct {
	Event UserDataEventType `json:"e"`
	Time  int64             `json:"E"`
}

// WsUserDataEvent define user data event
type WsUserDataEvent struct {
	WsUserDataEventType

	TransactionTime   int64 `json:"T"`
	AccountUpdateTime int64 `json:"u"`
	AccountUpdate     WsAccountUpdateList
	BalanceUpdate     WsBalanceUpdate
	OrderUpdate       WsOrderUpdate
	OCOUpdate         WsOCOUpdate
}

func (w *WsUserDataEvent) UnmarshalJSON(message []byte) error {
	j, err := newJSON(message)
	if err != nil {
		return err
	}

	err = json.Unmarshal(message, &w.WsUserDataEventType)
	if err != nil {
		return err
	}

	switch w.WsUserDataEventType.Event {
	case UserDataEventTypeOutboundAccountPosition:
		err = json.Unmarshal(message, &w.AccountUpdate)
		if err != nil {
			return err
		}

	case UserDataEventTypeBalanceUpdate:
		err = json.Unmarshal(message, &w.BalanceUpdate)
		if err != nil {
			return err
		}

	case UserDataEventTypeExecutionReport:
		err = json.Unmarshal(message, &w.OrderUpdate)
		if err != nil {
			return err
		}

		w.TransactionTime = j.Get("T").MustInt64()
		w.OrderUpdate.TransactionTime = j.Get("T").MustInt64()
		w.OrderUpdate.Id = j.Get("i").MustInt64()
		w.OrderUpdate.TradeId = j.Get("t").MustInt64()
		w.OrderUpdate.FeeAsset = j.Get("N").MustString()

	case UserDataEventTypeListStatus:
		err = json.Unmarshal(message, &w.OCOUpdate)
		if err != nil {
			return err
		}
	}

	return nil
}

type WsAccountUpdateList struct {
	WsAccountUpdates []WsAccountUpdate `json:"B"`
}

// WsAccountUpdate define account update
type WsAccountUpdate struct {
	Asset  string          `json:"a"`
	Free   decimal.Decimal `json:"f"`
	Locked decimal.Decimal `json:"l"`
}

type WsBalanceUpdate struct {
	Asset  string          `json:"a"`
	Change decimal.Decimal `json:"d"`
}

type WsOrderUpdate struct {
	Symbol                  string              `json:"s"`
	ClientOrderId           string              `json:"c"`
	Side                    string              `json:"S"`
	Type                    string              `json:"o"`
	TimeInForce             TimeInForceType     `json:"f"`
	Volume                  decimal.Decimal     `json:"q"`
	Price                   decimal.Decimal     `json:"p"`
	StopPrice               decimal.NullDecimal `json:"P"`
	TrailingDelta           int64               `json:"d"` // Trailing Delta
	IceBergVolume           string              `json:"F"`
	OrderListId             int64               `json:"g"` // for OCO
	OrigCustomOrderId       string              `json:"C"` // customized order ID for the original order
	ExecutionType           string              `json:"x"` // execution type for this event NEW/TRADE...
	Status                  string              `json:"X"` // order status
	RejectReason            string              `json:"r"`
	Id                      int64               `json:"i"` // order id
	LatestVolume            decimal.Decimal     `json:"l"` // quantity for the latest trade
	FilledVolume            decimal.NullDecimal `json:"z"`
	LatestPrice             decimal.Decimal     `json:"L"` // price for the latest trade
	FeeAsset                string              `json:"N"`
	FeeCost                 decimal.Decimal     `json:"n"`
	TransactionTime         int64               `json:"T"`
	TradeId                 int64               `json:"t"`
	IsInOrderBook           bool                `json:"w"` // is the order in the order book?
	IsMaker                 bool                `json:"m"` // is this order maker?
	CreateTime              int64               `json:"O"`
	FilledQuoteVolume       string              `json:"Z"` // the quote volume that already filled
	LatestQuoteVolume       string              `json:"Y"` // the quote volume for the latest trade
	QuoteVolume             string              `json:"Q"`
	TrailingTime            int64               `json:"D"` // Trailing Time
	StrategyId              int64               `json:"j"` // Strategy ID
	StrategyType            int64               `json:"J"` // Strategy Type
	WorkingTime             int64               `json:"W"` // Working Time
	SelfTradePreventionMode string              `json:"V"`
}

type WsOCOUpdate struct {
	Symbol          string `json:"s"`
	OrderListId     int64  `json:"g"`
	ContingencyType string `json:"c"`
	ListStatusType  string `json:"l"`
	ListOrderStatus string `json:"L"`
	RejectReason    string `json:"r"`
	ClientOrderId   string `json:"C"` // List Client Order ID
	Orders          WsOCOOrderList
}

type WsOCOOrderList struct {
	WsOCOOrders []WsOCOOrder `json:"O"`
}

type WsOCOOrder struct {
	Symbol        string `json:"s"`
	OrderId       int64  `json:"i"`
	ClientOrderId string `json:"c"`
}

// WsUserDataHandler handle WsUserDataEvent
//type WsUserDataHandler func(event *WsUserDataEvent)

// WsUserDataServe serve user data handler with listen key
func WsUserDataServe(
	ctx context.Context,
	listenKey string,
	handler func([]byte),
	errHandler ErrHandler,
	proxyFunc futures.DialFunc,
) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s", getWsEndpoint(), listenKey)
	cfg := newWsConfig(endpoint)

	return wsServe(ctx, cfg, handler, errHandler, proxyFunc)
}

// WsMarketStatHandler handle websocket that push single market statistics for 24hr
type WsMarketStatHandler func(event *WsMarketStatEvent)

// WsCombinedMarketStatServe is similar to WsMarketStatServe, but it handles multiple symbolx
func WsCombinedMarketStatServe(ctx context.Context, symbols []string, handler WsMarketStatHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := getCombinedEndpoint()
	for s := range symbols {
		endpoint += fmt.Sprintf("%s@ticker", strings.ToLower(symbols[s])) + "/"
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

		event := new(WsMarketStatEvent)
		err = json.Unmarshal(jsonData, event)
		if err != nil {
			errHandler(err)
			return
		}

		event.Symbol = strings.ToUpper(symbol)

		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsMarketStatServe serve websocket that push 24hr statistics for single market every second
func WsMarketStatServe(ctx context.Context, symbol string, handler WsMarketStatHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/%s@ticker", getWsEndpoint(), strings.ToLower(symbol))
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsMarketStatEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(&event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsAllMarketsStatHandler handle websocket that push all markets statistics for 24hr
type WsAllMarketsStatHandler func(event WsAllMarketsStatEvent)

// WsAllMarketsStatServe serve websocket that push 24hr statistics for all market every second
func WsAllMarketsStatServe(ctx context.Context, handler WsAllMarketsStatHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!ticker@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMarketsStatEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsAllMarketsStatEvent define array of websocket market statistics events
type WsAllMarketsStatEvent []*WsMarketStatEvent

// WsMarketStatEvent define websocket market statistics event
type WsMarketStatEvent struct {
	Event              string `json:"e"`
	Time               int64  `json:"E"`
	Symbol             string `json:"s"`
	PriceChange        string `json:"p"`
	PriceChangePercent string `json:"P"`
	WeightedAvgPrice   string `json:"w"`
	PrevClosePrice     string `json:"x"`
	LastPrice          string `json:"c"`
	CloseQty           string `json:"Q"`
	BidPrice           string `json:"b"`
	BidQty             string `json:"B"`
	AskPrice           string `json:"a"`
	AskQty             string `json:"A"`
	OpenPrice          string `json:"o"`
	HighPrice          string `json:"h"`
	LowPrice           string `json:"l"`
	BaseVolume         string `json:"v"`
	QuoteVolume        string `json:"q"`
	OpenTime           int64  `json:"O"`
	CloseTime          int64  `json:"C"`
	FirstID            int64  `json:"F"`
	LastID             int64  `json:"L"`
	Count              int64  `json:"n"`
}

// WsAllMiniMarketsStatServeHandler handle websocket that push all mini-ticker market statistics for 24hr
type WsAllMiniMarketsStatServeHandler func(event WsAllMiniMarketsStatEvent)

// WsAllMiniMarketsStatServe serve websocket that push mini version of 24hr statistics for all market every second
func WsAllMiniMarketsStatServe(ctx context.Context, handler WsAllMiniMarketsStatServeHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := fmt.Sprintf("%s/!miniTicker@arr", getWsEndpoint())
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		var event WsAllMiniMarketsStatEvent
		err := json.Unmarshal(message, &event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsAllMiniMarketsStatEvent define array of websocket market mini-ticker statistics events
type WsAllMiniMarketsStatEvent []*WsMiniMarketsStatEvent

// WsMiniMarketsStatEvent define websocket market mini-ticker statistics event
type WsMiniMarketsStatEvent struct {
	Event       string `json:"e"`
	Time        int64  `json:"E"`
	Symbol      string `json:"s"`
	LastPrice   string `json:"c"`
	OpenPrice   string `json:"o"`
	HighPrice   string `json:"h"`
	LowPrice    string `json:"l"`
	BaseVolume  string `json:"v"`
	QuoteVolume string `json:"q"`
}

// WsBookTickerEvent define websocket best book ticker event.
type WsBookTickerEvent struct {
	UpdateID     int64  `json:"u"`
	Symbol       string `json:"s"`
	BestBidPrice string `json:"b"`
	BestBidQty   string `json:"B"`
	BestAskPrice string `json:"a"`
	BestAskQty   string `json:"A"`
}

type WsCombinedBookTickerEvent struct {
	Data   *WsBookTickerEvent `json:"data"`
	Stream string             `json:"stream"`
}

// WsBookTickerHandler handle websocket that pushes updates to the best bid or ask price or quantity in real-time for a specified symbol.
type WsBookTickerHandler func(event *WsBookTickerEvent)

// WsBookTickerServe serve websocket that pushes updates to the best bid or ask price or quantity in real-time for a specified symbol.
func WsBookTickerServe(ctx context.Context, symbol string, handler WsBookTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
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
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsCombinedBookTickerServe is similar to WsBookTickerServe, but it is for multiple symbols
func WsCombinedBookTickerServe(ctx context.Context, symbols []string, handler WsBookTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	endpoint := baseCombinedMainURL
	for _, s := range symbols {
		endpoint += fmt.Sprintf("%s@bookTicker", strings.ToLower(s)) + "/"
	}
	endpoint = endpoint[:len(endpoint)-1]
	cfg := newWsConfig(endpoint)
	wsHandler := func(message []byte) {
		event := new(WsCombinedBookTickerEvent)
		err := json.Unmarshal(message, event)
		if err != nil {
			errHandler(err)
			return
		}
		handler(event.Data)
	}
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}

// WsAllBookTickerServe serve websocket that pushes updates to the best bid or ask price or quantity in real-time for all symbols.
func WsAllBookTickerServe(ctx context.Context, handler WsBookTickerHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
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
	return wsServe(ctx, cfg, wsHandler, errHandler, nil)
}
