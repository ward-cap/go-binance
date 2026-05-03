package binance

import (
	"context"
	"net/http"
	"strings"
)

// MarginTransferService transfer between spot account and margin account
type MarginTransferService struct {
	c            *Client
	asset        string
	amount       string
	transferType int
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginTransferService) Asset(asset string) *MarginTransferService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginTransferService) Amount(amount string) *MarginTransferService {
	s.amount = amount
	return s
}

// Type 1: transfer from main account to margin account 2: transfer from margin account to main account
func (s *MarginTransferService) Type(transferType MarginTransferType) *MarginTransferService {
	s.transferType = int(transferType)
	return s
}

// Do send request
func (s *MarginTransferService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		service:  "MarginTransferService",
		method:   http.MethodPost,
		endpoint: "/sapi/v1/margin/transfer",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
		"type":   s.transferType,
	}
	r.setFormParams(m)
	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginLoanService apply for a loan
type MarginLoanService struct {
	c          *Client
	asset      string
	amount     string
	isIsolated bool
	symbol     *string
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginLoanService) Asset(asset string) *MarginLoanService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginLoanService) Amount(amount string) *MarginLoanService {
	s.amount = amount
	return s
}

// IsIsolated is for isolated margin or not, "TRUE", "FALSE"，default "FALSE"
func (s *MarginLoanService) IsIsolated(isIsolated bool) *MarginLoanService {
	s.isIsolated = isIsolated
	return s
}

// Symbol set isolated symbol
func (s *MarginLoanService) Symbol(symbol string) *MarginLoanService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *MarginLoanService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		service:  "MarginLoanService",
		method:   http.MethodPost,
		endpoint: "/sapi/v1/margin/loan",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
	}
	r.setFormParams(m)
	if s.isIsolated {
		r.setParam("isIsolated", "TRUE")
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// MarginRepayService repay loan for margin account
type MarginRepayService struct {
	c          *Client
	asset      string
	amount     string
	isIsolated bool
	symbol     *string
}

// Asset set asset being transferred, e.g., BTC
func (s *MarginRepayService) Asset(asset string) *MarginRepayService {
	s.asset = asset
	return s
}

// Amount the amount to be transferred
func (s *MarginRepayService) Amount(amount string) *MarginRepayService {
	s.amount = amount
	return s
}

// IsIsolated is for isolated margin or not, "TRUE", "FALSE"，default "FALSE"
func (s *MarginRepayService) IsIsolated(isIsolated bool) *MarginRepayService {
	s.isIsolated = isIsolated
	return s
}

// Symbol set isolated symbol
func (s *MarginRepayService) Symbol(symbol string) *MarginRepayService {
	s.symbol = &symbol
	return s
}

// Do send request
func (s *MarginRepayService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		service:  "MarginRepayService",
		method:   http.MethodPost,
		endpoint: "/sapi/v1/margin/repay",
		secType:  secTypeSigned,
	}
	m := params{
		"asset":  s.asset,
		"amount": s.amount,
	}
	r.setFormParams(m)
	if s.isIsolated {
		r.setParam("isIsolated", "TRUE")
	}
	if s.symbol != nil {
		r.setParam("symbol", *s.symbol)
	}

	res = new(TransactionResponse)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListMarginLoansService list loan record
type ListMarginLoansService struct {
	c         *Client
	asset     string
	txID      *int64
	startTime *int64
	endTime   *int64
	current   *int64
	size      *int64
}

// Asset set asset
func (s *ListMarginLoansService) Asset(asset string) *ListMarginLoansService {
	s.asset = asset
	return s
}

// TxID set transaction id
func (s *ListMarginLoansService) TxID(txID int64) *ListMarginLoansService {
	s.txID = &txID
	return s
}

// StartTime set start time
func (s *ListMarginLoansService) StartTime(startTime int64) *ListMarginLoansService {
	s.startTime = &startTime
	return s
}

// EndTime set end time
func (s *ListMarginLoansService) EndTime(endTime int64) *ListMarginLoansService {
	s.endTime = &endTime
	return s
}

// Current currently querying page. Start from 1. Default:1
func (s *ListMarginLoansService) Current(current int64) *ListMarginLoansService {
	s.current = &current
	return s
}

// Size default:10 max:100
func (s *ListMarginLoansService) Size(size int64) *ListMarginLoansService {
	s.size = &size
	return s
}

// Do send request
func (s *ListMarginLoansService) Do(ctx context.Context, opts ...RequestOption) (res *MarginLoanResponse, err error) {
	r := &request{
		service:  "ListMarginLoansService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/loan",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	if s.txID != nil {
		r.setParam("txId", *s.txID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginLoanResponse)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListMarginRepaysService list repay record
type ListMarginRepaysService struct {
	c         *Client
	asset     string
	txID      *int64
	startTime *int64
	endTime   *int64
	current   *int64
	size      *int64
}

// Asset set asset
func (s *ListMarginRepaysService) Asset(asset string) *ListMarginRepaysService {
	s.asset = asset
	return s
}

// TxID set transaction id
func (s *ListMarginRepaysService) TxID(txID int64) *ListMarginRepaysService {
	s.txID = &txID
	return s
}

// StartTime set start time
func (s *ListMarginRepaysService) StartTime(startTime int64) *ListMarginRepaysService {
	s.startTime = &startTime
	return s
}

// EndTime set end time
func (s *ListMarginRepaysService) EndTime(endTime int64) *ListMarginRepaysService {
	s.endTime = &endTime
	return s
}

// Current currently querying page. Start from 1. Default:1
func (s *ListMarginRepaysService) Current(current int64) *ListMarginRepaysService {
	s.current = &current
	return s
}

// Size default:10 max:100
func (s *ListMarginRepaysService) Size(size int64) *ListMarginRepaysService {
	s.size = &size
	return s
}

// Do send request
func (s *ListMarginRepaysService) Do(ctx context.Context, opts ...RequestOption) (res *MarginRepayResponse, err error) {
	r := &request{
		service:  "ListMarginRepaysService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/repay",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	if s.txID != nil {
		r.setParam("txId", *s.txID)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginRepayResponse)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetIsolatedMarginAccountService gets isolated margin account info
type GetIsolatedMarginAccountService struct {
	c *Client

	symbols []string
}

// Symbols set symbols to the isolated margin account
func (s *GetIsolatedMarginAccountService) Symbols(symbols ...string) *GetIsolatedMarginAccountService {
	s.symbols = symbols
	return s
}

// Do send request
func (s *GetIsolatedMarginAccountService) Do(ctx context.Context, opts ...RequestOption) (res *IsolatedMarginAccount, err error) {
	r := &request{
		service:  "GetIsolatedMarginAccountService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/isolated/account",
		secType:  secTypeSigned,
	}

	if len(s.symbols) > 0 {
		r.setParam("symbols", strings.Join(s.symbols, ","))
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(IsolatedMarginAccount)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMarginAccountService get margin account info
type GetMarginAccountService struct {
	c *Client
}

// Do send request
func (s *GetMarginAccountService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAccount, err error) {
	r := &request{
		service:  "GetMarginAccountService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginAccount)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMarginAssetService get margin asset info
type GetMarginAssetService struct {
	c     *Client
	asset string
}

// Asset set asset
func (s *GetMarginAssetService) Asset(asset string) *GetMarginAssetService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetMarginAssetService) Do(ctx context.Context, opts ...RequestOption) (res *MarginAsset, err error) {
	r := &request{
		service:  "GetMarginAssetService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/asset",
		secType:  secTypeAPIKey,
	}
	r.setParam("asset", s.asset)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginAsset)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMarginPairService get margin pair info
type GetMarginPairService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetMarginPairService) Symbol(symbol string) *GetMarginPairService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetMarginPairService) Do(ctx context.Context, opts ...RequestOption) (res *MarginPair, err error) {
	r := &request{
		service:  "GetMarginPairService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/pair",
		secType:  secTypeAPIKey,
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginPair)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMarginAllPairsService get margin pair info
type GetMarginAllPairsService struct {
	c *Client
}

// Do send request
func (s *GetMarginAllPairsService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginAllPair, err error) {
	r := &request{
		service:  "GetMarginAllPairsService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/allPairs",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginAllPair{}, err
	}
	res = make([]*MarginAllPair, 0)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return []*MarginAllPair{}, err
	}
	return res, nil
}

// GetMarginPriceIndexService get margin price index
type GetMarginPriceIndexService struct {
	c      *Client
	symbol string
}

// Symbol set symbol
func (s *GetMarginPriceIndexService) Symbol(symbol string) *GetMarginPriceIndexService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetMarginPriceIndexService) Do(ctx context.Context, opts ...RequestOption) (res *MarginPriceIndex, err error) {
	r := &request{
		service:  "GetMarginPriceIndexService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/priceIndex",
		secType:  secTypeAPIKey,
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MarginPriceIndex)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ListMarginTradesService list trades
type ListMarginTradesService struct {
	c          *Client
	symbol     string
	startTime  *int64
	endTime    *int64
	limit      *int
	fromID     *int64
	isIsolated bool
}

// Symbol set symbol
func (s *ListMarginTradesService) Symbol(symbol string) *ListMarginTradesService {
	s.symbol = symbol
	return s
}

// IsIsolated set isIsolated
func (s *ListMarginTradesService) IsIsolated(isIsolated bool) *ListMarginTradesService {
	s.isIsolated = isIsolated
	return s
}

// StartTime set starttime
func (s *ListMarginTradesService) StartTime(startTime int64) *ListMarginTradesService {
	s.startTime = &startTime
	return s
}

// EndTime set endtime
func (s *ListMarginTradesService) EndTime(endTime int64) *ListMarginTradesService {
	s.endTime = &endTime
	return s
}

// Limit set limit
func (s *ListMarginTradesService) Limit(limit int) *ListMarginTradesService {
	s.limit = &limit
	return s
}

// FromID set fromID
func (s *ListMarginTradesService) FromID(fromID int64) *ListMarginTradesService {
	s.fromID = &fromID
	return s
}

// Do send request
func (s *ListMarginTradesService) Do(ctx context.Context, opts ...RequestOption) (res []*TradeV3, err error) {
	r := &request{
		service:  "ListMarginTradesService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/myTrades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}
	if s.isIsolated {
		r.setParam("isIsolated", "TRUE")
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*TradeV3{}, err
	}
	res = make([]*TradeV3, 0)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return []*TradeV3{}, err
	}
	return res, nil
}

// GetMaxBorrowableService get max borrowable of asset
type GetMaxBorrowableService struct {
	c              *Client
	asset          string
	isolatedSymbol string
}

// Asset set asset
func (s *GetMaxBorrowableService) Asset(asset string) *GetMaxBorrowableService {
	s.asset = asset
	return s
}

// IsolatedSymbol set isolatedSymbol
func (s *GetMaxBorrowableService) IsolatedSymbol(isolatedSymbol string) *GetMaxBorrowableService {
	s.isolatedSymbol = isolatedSymbol
	return s
}

// Do send request
func (s *GetMaxBorrowableService) Do(ctx context.Context, opts ...RequestOption) (res *MaxBorrowable, err error) {
	r := &request{
		service:  "GetMaxBorrowableService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/maxBorrowable",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	if s.isolatedSymbol != "" {
		r.setParam("isolatedSymbol", s.isolatedSymbol)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MaxBorrowable)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMaxTransferableService get max transferable of asset
type GetMaxTransferableService struct {
	c     *Client
	asset string
}

// Asset set asset
func (s *GetMaxTransferableService) Asset(asset string) *GetMaxTransferableService {
	s.asset = asset
	return s
}

// Do send request
func (s *GetMaxTransferableService) Do(ctx context.Context, opts ...RequestOption) (res *MaxTransferable, err error) {
	r := &request{
		service:  "GetMaxTransferableService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/maxTransferable",
		secType:  secTypeSigned,
	}
	r.setParam("asset", s.asset)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(MaxTransferable)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// StartIsolatedMarginUserStreamService create listen key for margin user stream service
type StartIsolatedMarginUserStreamService struct {
	c      *Client
	symbol string
}

// Symbol sets the user stream to isolated margin user stream
func (s *StartIsolatedMarginUserStreamService) Symbol(symbol string) *StartIsolatedMarginUserStreamService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *StartIsolatedMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (listenKey string, err error) {
	r := &request{
		service:  "StartIsolatedMarginUserStreamService",
		method:   http.MethodPost,
		endpoint: "/sapi/v1/userDataStream/isolated",
		secType:  secTypeAPIKey,
	}

	r.setFormParam("symbol", s.symbol)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	return parseListenKey(data)
}

// KeepaliveIsolatedMarginUserStreamService updates listen key for isolated margin user data stream
type KeepaliveIsolatedMarginUserStreamService struct {
	c         *Client
	listenKey string
	symbol    string
}

// Symbol set symbol to the isolated margin keepalive request
func (s *KeepaliveIsolatedMarginUserStreamService) Symbol(symbol string) *KeepaliveIsolatedMarginUserStreamService {
	s.symbol = symbol
	return s
}

// ListenKey set listen key
func (s *KeepaliveIsolatedMarginUserStreamService) ListenKey(listenKey string) *KeepaliveIsolatedMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveIsolatedMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		service:  "KeepaliveIsolatedMarginUserStreamService",
		method:   http.MethodPut,
		endpoint: "/sapi/v1/userDataStream/isolated",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	r.setFormParam("symbol", s.symbol)

	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// CloseIsolatedMarginUserStreamService delete listen key
type CloseIsolatedMarginUserStreamService struct {
	c         *Client
	listenKey string

	symbol string
}

// ListenKey set listen key
func (s *CloseIsolatedMarginUserStreamService) ListenKey(listenKey string) *CloseIsolatedMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Symbol set symbol to the isolated margin user stream close request
func (s *CloseIsolatedMarginUserStreamService) Symbol(symbol string) *CloseIsolatedMarginUserStreamService {
	s.symbol = symbol
	return s
}

// Do send request
func (s *CloseIsolatedMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		service:  "CloseIsolatedMarginUserStreamService",
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/userDataStream/isolated",
		secType:  secTypeAPIKey,
	}

	r.setFormParam("listenKey", s.listenKey)
	r.setFormParam("symbol", s.symbol)

	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// StartMarginUserStreamService create listen key for margin user stream service
type StartMarginUserStreamService struct {
	c *Client
}

// Do send request
func (s *StartMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (listenKey string, err error) {
	r := &request{
		service:  "StartMarginUserStreamService",
		method:   http.MethodPost,
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	return parseListenKey(data)
}

// KeepaliveMarginUserStreamService update listen key
type KeepaliveMarginUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *KeepaliveMarginUserStreamService) ListenKey(listenKey string) *KeepaliveMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		service:  "KeepaliveMarginUserStreamService",
		method:   http.MethodPut,
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// CloseMarginUserStreamService delete listen key
type CloseMarginUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseMarginUserStreamService) ListenKey(listenKey string) *CloseMarginUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseMarginUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		service:  "CloseMarginUserStreamService",
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/userDataStream",
		secType:  secTypeAPIKey,
	}

	r.setFormParam("listenKey", s.listenKey)

	_, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// GetAllMarginAssetsService get margin pair info
type GetAllMarginAssetsService struct {
	c *Client
}

// Do send request
func (s *GetAllMarginAssetsService) Do(ctx context.Context, opts ...RequestOption) (res []*MarginAsset, err error) {
	r := &request{
		service:  "GetAllMarginAssetsService",
		method:   "GET",
		endpoint: "/sapi/v1/margin/allAssets",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*MarginAsset{}, err
	}
	res = make([]*MarginAsset, 0)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return []*MarginAsset{}, err
	}
	return res, nil
}

// GetIsolatedMarginAllPairsService get isolated margin pair info
type GetIsolatedMarginAllPairsService struct {
	c *Client
}

// Do send request
func (s *GetIsolatedMarginAllPairsService) Do(ctx context.Context, opts ...RequestOption) (res []*IsolatedMarginAllPair, err error) {
	r := &request{
		service:  "GetIsolatedMarginAllPairsService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/margin/isolated/allPairs",
		secType:  secTypeAPIKey,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*IsolatedMarginAllPair{}, err
	}
	res = make([]*IsolatedMarginAllPair, 0)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return []*IsolatedMarginAllPair{}, err
	}
	return res, nil
}

// IsolatedMarginTransferService transfer assets between spot and isolated margin.
type IsolatedMarginTransferService struct {
	c *Client

	asset     string
	symbol    string
	transFrom AccountType
	transTo   AccountType
	amount    string
}

func (s *IsolatedMarginTransferService) Symbol(symbol string) *IsolatedMarginTransferService {
	s.symbol = symbol
	return s
}

func (s *IsolatedMarginTransferService) Asset(asset string) *IsolatedMarginTransferService {
	s.asset = asset
	return s
}

// TransFrom supports account types: "SPOT", "ISOLATED_MARGIN"
func (s *IsolatedMarginTransferService) TransFrom(transFrom AccountType) *IsolatedMarginTransferService {
	s.transFrom = transFrom
	return s
}

// TransTo supports account types: "SPOT", "ISOLATED_MARGIN"
func (s *IsolatedMarginTransferService) TransTo(transTo AccountType) *IsolatedMarginTransferService {
	s.transTo = transTo
	return s
}

func (s *IsolatedMarginTransferService) Amount(amount string) *IsolatedMarginTransferService {
	s.amount = amount
	return s
}

// Do send request
func (s *IsolatedMarginTransferService) Do(ctx context.Context, opts ...RequestOption) (res *TransactionResponse, err error) {
	r := &request{
		service:  "IsolatedMarginTransferService",
		method:   http.MethodPost,
		endpoint: "/sapi/v1/margin/isolated/transfer",
		secType:  secTypeSigned,
	}
	r.setFormParam("asset", s.asset)
	r.setFormParam("symbol", s.symbol)
	r.setFormParam("transFrom", s.transFrom)
	r.setFormParam("transTo", s.transTo)
	r.setFormParam("amount", s.amount)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(TransactionResponse)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
