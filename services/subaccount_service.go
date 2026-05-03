package binance

import (
	"context"
	"net/http"
)

// TransferToSubAccountService transfer to subaccount
type TransferToSubAccountService struct {
	c       *Client
	toEmail string
	asset   string
	amount  string
}

// ToEmail set toEmail
func (s *TransferToSubAccountService) ToEmail(toEmail string) *TransferToSubAccountService {
	s.toEmail = toEmail
	return s
}

// Asset set asset
func (s *TransferToSubAccountService) Asset(asset string) *TransferToSubAccountService {
	s.asset = asset
	return s
}

// Amount set amount
func (s *TransferToSubAccountService) Amount(amount string) *TransferToSubAccountService {
	s.amount = amount
	return s
}

func (s *TransferToSubAccountService) transferToSubaccount(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		service:  "TransferToSubAccountService",
		method:   "POST",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"toEmail": s.toEmail,
		"asset":   s.asset,
		"amount":  s.amount,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *TransferToSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *TransferToSubAccountResponse, err error) {
	data, err := s.transferToSubaccount(ctx, "/sapi/v1/sub-account/transfer/subToSub", opts...)
	if err != nil {
		return nil, err
	}
	res = &TransferToSubAccountResponse{}
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubaccountDepositAddressService struct {
	c       *Client
	email   string
	coin    string
	network string
}

// Email set email
func (s *SubaccountDepositAddressService) Email(email string) *SubaccountDepositAddressService {
	s.email = email
	return s
}

// Coin set coin
func (s *SubaccountDepositAddressService) Coin(coin string) *SubaccountDepositAddressService {
	s.coin = coin
	return s
}

// Network set network
func (s *SubaccountDepositAddressService) Network(network string) *SubaccountDepositAddressService {
	s.network = network
	return s
}

func (s *SubaccountDepositAddressService) subaccountDepositAddress(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		service:  "SubaccountDepositAddressService",
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"email":   s.email,
		"coin":    s.coin,
		"network": s.network,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountDepositAddressService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountDepositAddressResponse, err error) {
	data, err := s.subaccountDepositAddress(ctx, "/sapi/v1/capital/deposit/subAddress", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountDepositAddressResponse{}
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubaccountAssetsService struct {
	c     *Client
	email string
}

// Email set email
func (s *SubaccountAssetsService) Email(email string) *SubaccountAssetsService {
	s.email = email
	return s
}

func (s *SubaccountAssetsService) subaccountAssets(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		service:  "SubaccountAssetsService",
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"email": s.email,
	}
	r.setParams(m)
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountAssetsService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountAssetsResponse, err error) {
	data, err := s.subaccountAssets(ctx, "/sapi/v3/sub-account/assets", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountAssetsResponse{}
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubaccountSpotSummaryService struct {
	c     *Client
	email *string
	page  *int32
	size  *int32
}

// Email set email
func (s *SubaccountSpotSummaryService) Email(email string) *SubaccountSpotSummaryService {
	s.email = &email
	return s
}

func (s *SubaccountSpotSummaryService) Page(page int32) *SubaccountSpotSummaryService {
	s.page = &page
	return s
}

func (s *SubaccountSpotSummaryService) Size(size int32) *SubaccountSpotSummaryService {
	s.size = &size
	return s
}

func (s *SubaccountSpotSummaryService) subaccountSpotSummary(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, err error) {
	r := &request{
		service:  "SubaccountSpotSummaryService",
		method:   "GET",
		endpoint: endpoint,
		secType:  secTypeSigned,
	}

	if s.size != nil {
		r.setParam("size", *s.size)
	}

	if s.page != nil {
		r.setParam("page", *s.page)
	}

	if s.email != nil {
		r.setParam("email", *s.email)
	}
	data, err = s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

// Do send request
func (s *SubaccountSpotSummaryService) Do(ctx context.Context, opts ...RequestOption) (res *SubaccountSpotSummaryResponse, err error) {
	data, err := s.subaccountSpotSummary(ctx, "/sapi/v1/sub-account/spotSummary", opts...)
	if err != nil {
		return nil, err
	}
	res = &SubaccountSpotSummaryResponse{}
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubAccountListService Query Sub-account List (For Master Account)
// https://binance-docs.github.io/apidocs/spot/en/#query-sub-account-list-for-master-account
type SubAccountListService struct {
	c           *Client
	email       *string
	isFreeze    bool
	page, limit int
}

func (s *SubAccountListService) Email(v string) *SubAccountListService {
	s.email = &v
	return s
}

func (s *SubAccountListService) IsFreeze(v bool) *SubAccountListService {
	s.isFreeze = v
	return s
}

func (s *SubAccountListService) Page(v int) *SubAccountListService {
	s.page = v
	return s
}

func (s *SubAccountListService) Limit(v int) *SubAccountListService {
	s.limit = v
	return s
}

func (s *SubAccountListService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountList, err error) {
	r := &request{
		service:  "SubAccountListService",
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/list",
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}
	if s.isFreeze {
		r.setParam("isFreeze", "true")
	} else {
		r.setParam("isFreeze", "false")
	}
	if s.page > 0 {
		r.setParam("page", s.page)
	}
	if s.limit > 200 {
		r.setParam("limit", 200)
	} else if s.limit <= 0 {
		r.setParam("limit", 10)
	} else {
		r.setParam("limit", s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountList)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// ManagedSubAccountDepositService
// Deposit Assets Into The Managed Sub-account（For Investor Master Account）
// https://binance-docs.github.io/apidocs/spot/en/#deposit-assets-into-the-managed-sub-account-for-investor-master-account
type ManagedSubAccountDepositService struct {
	c       *Client
	toEmail string
	asset   string
	amount  float64
}

func (s *ManagedSubAccountDepositService) ToEmail(email string) *ManagedSubAccountDepositService {
	s.toEmail = email
	return s
}

func (s *ManagedSubAccountDepositService) Asset(asset string) *ManagedSubAccountDepositService {
	s.asset = asset
	return s
}

func (s *ManagedSubAccountDepositService) Amount(amount float64) *ManagedSubAccountDepositService {
	s.amount = amount
	return s
}

// Do send request
func (s *ManagedSubAccountDepositService) Do(ctx context.Context, opts ...RequestOption) (*ManagedSubAccountDepositResponse, error) {
	r := &request{
		service:  "ManagedSubAccountDepositService",
		method:   "POST",
		endpoint: "/sapi/v1/managed-subaccount/deposit",
		secType:  secTypeSigned,
	}

	r.setParam("toEmail", s.toEmail)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := &ManagedSubAccountDepositResponse{}
	if err := jsonCodec.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ManagedSubAccountWithdrawalService
// Withdrawal Assets From The Managed Sub-account（For Investor Master Account）
// https://binance-docs.github.io/apidocs/spot/en/#withdrawl-assets-from-the-managed-sub-account-for-investor-master-account
type ManagedSubAccountWithdrawalService struct {
	c            *Client
	fromEmail    string
	asset        string
	amount       float64
	transferDate int64 // Withdrawals is automatically occur on the transfer date(UTC0). If a date is not selected, the withdrawal occurs right now
}

func (s *ManagedSubAccountWithdrawalService) FromEmail(email string) *ManagedSubAccountWithdrawalService {
	s.fromEmail = email
	return s
}

func (s *ManagedSubAccountWithdrawalService) Asset(asset string) *ManagedSubAccountWithdrawalService {
	s.asset = asset
	return s
}

func (s *ManagedSubAccountWithdrawalService) Amount(amount float64) *ManagedSubAccountWithdrawalService {
	s.amount = amount
	return s
}

func (s *ManagedSubAccountWithdrawalService) TransferDate(val int64) *ManagedSubAccountWithdrawalService {
	s.transferDate = val
	return s
}

// Do send request
func (s *ManagedSubAccountWithdrawalService) Do(ctx context.Context, opts ...RequestOption) (*ManagedSubAccountWithdrawalResponse, error) {
	r := &request{
		service:  "ManagedSubAccountWithdrawalService",
		method:   "POST",
		endpoint: "/sapi/v1/managed-subaccount/withdraw",
		secType:  secTypeSigned,
	}

	r.setParam("fromEmail", s.fromEmail)
	r.setParam("asset", s.asset)
	r.setParam("amount", s.amount)
	if s.transferDate > 0 {
		r.setParam("transferDate", s.transferDate)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := &ManagedSubAccountWithdrawalResponse{}
	if err := jsonCodec.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// ManagedSubAccountAssetsService
// Query Managed Sub-account Asset Details（For Investor Master Account）
// https://binance-docs.github.io/apidocs/spot/en/#query-managed-sub-account-asset-details-for-investor-master-account
type ManagedSubAccountAssetsService struct {
	c     *Client
	email string
}

func (s *ManagedSubAccountAssetsService) Email(email string) *ManagedSubAccountAssetsService {
	s.email = email
	return s
}

func (s *ManagedSubAccountAssetsService) Do(ctx context.Context, opts ...RequestOption) ([]*ManagedSubAccountAsset, error) {
	r := &request{
		service:  "ManagedSubAccountAssetsService",
		method:   "GET",
		endpoint: "/sapi/v1/managed-subaccount/asset",
		secType:  secTypeSigned,
	}

	r.setParam("email", s.email)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*ManagedSubAccountAsset, 0)
	if err := jsonCodec.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// SubAccountFuturesAccountService Get Detail on Sub-account's Futures Account (For Master Account)
// https://binance-docs.github.io/apidocs/spot/en/#get-detail-on-sub-account-39-s-futures-account-for-master-account
type SubAccountFuturesAccountService struct {
	c     *Client
	email *string
}

func (s *SubAccountFuturesAccountService) Email(v string) *SubAccountFuturesAccountService {
	s.email = &v
	return s
}

func (s *SubAccountFuturesAccountService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesAccount, err error) {
	r := &request{
		service:  "SubAccountFuturesAccountService",
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/futures/account",
		secType:  secTypeSigned,
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesAccount)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubaccountFuturesSummaryV1Service Get Summary of Sub-account's Futures Account (For Master Account)
// https://binance-docs.github.io/apidocs/spot/en/#get-summary-of-sub-account-39-s-futures-account-for-master-account
type SubAccountFuturesSummaryV1Service struct {
	c *Client
}

func (s *SubAccountFuturesSummaryV1Service) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountFuturesSummaryV1, err error) {
	r := &request{
		service:  "SubAccountFuturesSummaryV1Service",
		method:   "GET",
		endpoint: "/sapi/v1/sub-account/futures/accountSummary",
		secType:  secTypeSigned,
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SubAccountFuturesSummaryV1)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CreateSubAccountService struct {
	c *Client
}

type QuerySubAccountService struct {
	c    *Client
	size uint64
	page uint64
}

func (s *QuerySubAccountService) Size(size uint64) *QuerySubAccountService {
	s.size = size
	return s
}

func (s *QuerySubAccountService) Page(page uint64) *QuerySubAccountService {
	s.page = page
	return s
}

// Do send request
func (s *QuerySubAccountService) Do(ctx context.Context, opts ...RequestOption) (res []*QuerySubAccountResponse, err error) {
	r := &request{
		service:  "QuerySubAccountService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/broker/subAccount",
		secType:  secTypeSigned,
	}
	m := params{}
	if s.size != 0 {
		m["size"] = s.size
	}
	if s.page != 0 {
		m["page"] = s.page
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = make([]*QuerySubAccountResponse, 0)
	err = jsonCodec.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Do send request
func (s *CreateSubAccountService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountResponse, err error) {
	r := &request{
		service:  "CreateSubAccountService",
		method:   "POST",
		endpoint: "/sapi/v1/broker/subAccount",
		secType:  secTypeSigned,
	}
	m := params{}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(SubAccountResponse)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type SubAccountApiKeyService struct {
	c            *Client
	canTrade     bool
	futuresTrade bool
	marginTrade  bool
	subAccountID string
	publicKey    string
}

func (s *SubAccountApiKeyService) CanTrade(canTrade bool) *SubAccountApiKeyService {
	s.canTrade = canTrade
	return s
}

func (s *SubAccountApiKeyService) PublicKey(publicKey string) *SubAccountApiKeyService {
	s.publicKey = publicKey
	return s
}

func (s *SubAccountApiKeyService) FuturesTrade(futuresTrade bool) *SubAccountApiKeyService {
	s.futuresTrade = futuresTrade
	return s
}

func (s *SubAccountApiKeyService) MarginTrade(marginTrade bool) *SubAccountApiKeyService {
	s.marginTrade = marginTrade
	return s
}

func (s *SubAccountApiKeyService) SubAccountID(subAccountID string) *SubAccountApiKeyService {
	s.subAccountID = subAccountID
	return s
}

// Do send request
func (s *SubAccountApiKeyService) Do(ctx context.Context, opts ...RequestOption) (res *SubAccountAPIKeyResponse, err error) {
	r := &request{
		service:  "SubAccountApiKeyService",
		method:   http.MethodPost,
		endpoint: "/sapi/v1/broker/subAccountApi",
		secType:  secTypeSigned,
	}
	m := params{
		"subAccountId": s.subAccountID,
		"canTrade":     s.canTrade,
		"futuresTrade": s.futuresTrade,
		"marginTrade":  s.marginTrade,
	}
	if s.publicKey != "" {
		m["publicKey"] = s.publicKey
	}

	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res = new(SubAccountAPIKeyResponse)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
