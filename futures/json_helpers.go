package futures

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type listenKeyResponse struct {
	ListenKey string `json:"listenKey"`
}

type serverTimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}

type depthPayload struct {
	Time         int64             `json:"E"`
	TradeTime    int64             `json:"T"`
	LastUpdateID int64             `json:"lastUpdateId"`
	Bids         []priceLevelTuple `json:"bids"`
	Asks         []priceLevelTuple `json:"asks"`
}

type priceLevelTuple [2]string

type klineTuple []json.RawMessage

func parseListenKey(data []byte) (string, error) {
	var res listenKeyResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return "", err
	}
	return res.ListenKey, nil
}

func parseServerTime(data []byte) (int64, error) {
	var res serverTimeResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return 0, err
	}
	return res.ServerTime, nil
}

func parseDepth(data []byte) (*DepthResponse, error) {
	var payload depthPayload
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, err
	}
	res := &DepthResponse{
		Time:         payload.Time,
		TradeTime:    payload.TradeTime,
		LastUpdateID: payload.LastUpdateID,
	}
	res.Bids = make([]Bid, len(payload.Bids))
	for i, bid := range payload.Bids {
		res.Bids[i] = Bid{Price: bid[0], Quantity: bid[1]}
	}
	res.Asks = make([]Ask, len(payload.Asks))
	for i, ask := range payload.Asks {
		res.Asks[i] = Ask{Price: ask[0], Quantity: ask[1]}
	}
	return res, nil
}

func parseKlines(data []byte) ([]*Kline, error) {
	var rows []klineTuple
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, err
	}
	res := make([]*Kline, len(rows))
	for i, row := range rows {
		if len(row) < 11 {
			return nil, fmt.Errorf("invalid kline response")
		}
		kline, err := parseFullKline(row)
		if err != nil {
			return nil, err
		}
		res[i] = kline
	}
	return res, nil
}

func parsePriceKlines(data []byte) ([]*Kline, error) {
	var rows []klineTuple
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, err
	}
	res := make([]*Kline, len(rows))
	for i, row := range rows {
		if len(row) < 11 {
			return nil, fmt.Errorf("invalid kline response")
		}
		kline, err := parseCompactKline(row)
		if err != nil {
			return nil, err
		}
		res[i] = kline
	}
	return res, nil
}

func parseContinuousKlines(data []byte) ([]*ContinuousKline, error) {
	var rows []klineTuple
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, err
	}
	res := make([]*ContinuousKline, len(rows))
	for i, row := range rows {
		if len(row) < 11 {
			return nil, fmt.Errorf("invalid kline response")
		}
		kline, err := parseContinuousKline(row)
		if err != nil {
			return nil, err
		}
		res[i] = kline
	}
	return res, nil
}

func parseFullKline(row klineTuple) (*Kline, error) {
	openTime, closeTime, tradeNum, err := parseKlineTimes(row)
	if err != nil {
		return nil, err
	}
	return &Kline{
		OpenTime:                 openTime,
		Open:                     rawString(row[1]),
		High:                     rawString(row[2]),
		Low:                      rawString(row[3]),
		Close:                    rawString(row[4]),
		Volume:                   rawString(row[5]),
		CloseTime:                closeTime,
		QuoteAssetVolume:         rawString(row[7]),
		TradeNum:                 tradeNum,
		TakerBuyBaseAssetVolume:  rawString(row[9]),
		TakerBuyQuoteAssetVolume: rawString(row[10]),
	}, nil
}

func parseCompactKline(row klineTuple) (*Kline, error) {
	openTime, err := rawInt64(row[0])
	if err != nil {
		return nil, err
	}
	closeTime, err := rawInt64(row[6])
	if err != nil {
		return nil, err
	}
	return &Kline{
		OpenTime:  openTime,
		Open:      rawString(row[1]),
		High:      rawString(row[2]),
		Low:       rawString(row[3]),
		Close:     rawString(row[4]),
		CloseTime: closeTime,
	}, nil
}

func parseContinuousKline(row klineTuple) (*ContinuousKline, error) {
	openTime, closeTime, tradeNum, err := parseKlineTimes(row)
	if err != nil {
		return nil, err
	}
	return &ContinuousKline{
		OpenTime:                 openTime,
		Open:                     rawString(row[1]),
		High:                     rawString(row[2]),
		Low:                      rawString(row[3]),
		Close:                    rawString(row[4]),
		Volume:                   rawString(row[5]),
		CloseTime:                closeTime,
		QuoteAssetVolume:         rawString(row[7]),
		TradeNum:                 tradeNum,
		TakerBuyBaseAssetVolume:  rawString(row[9]),
		TakerBuyQuoteAssetVolume: rawString(row[10]),
	}, nil
}

func parseKlineTimes(row klineTuple) (int64, int64, int64, error) {
	openTime, err := rawInt64(row[0])
	if err != nil {
		return 0, 0, 0, err
	}
	closeTime, err := rawInt64(row[6])
	if err != nil {
		return 0, 0, 0, err
	}
	tradeNum, err := rawInt64(row[8])
	if err != nil {
		return 0, 0, 0, err
	}
	return openTime, closeTime, tradeNum, nil
}

func rawString(data json.RawMessage) string {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		return s
	}
	return string(data)
}

func rawInt64(data json.RawMessage) (int64, error) {
	var n int64
	if err := json.Unmarshal(data, &n); err == nil {
		return n, nil
	}
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return 0, err
	}
	return strconv.ParseInt(s, 10, 64)
}
