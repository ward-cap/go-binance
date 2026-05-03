package binance

import (
	"context"
	"net/http"
)

// StakingProductPositionService fetches the staking product positions
type StakingProductPositionService struct {
	c         *Client
	product   StakingProduct
	productId *string
	asset     *string
	current   *int32
	size      *int32
}

// Product sets the product parameter.
func (s *StakingProductPositionService) Product(product StakingProduct) *StakingProductPositionService {
	s.product = product
	return s
}

// ProductId sets the productId parameter.
func (s *StakingProductPositionService) ProductId(productId string) *StakingProductPositionService {
	s.productId = &productId
	return s
}

// Asset sets the asset parameter.
func (s *StakingProductPositionService) Asset(asset string) *StakingProductPositionService {
	s.asset = &asset
	return s
}

// Current sets the current parameter.
func (s *StakingProductPositionService) Current(current int32) *StakingProductPositionService {
	s.current = &current
	return s
}

// Size sets the size parameter.
func (s *StakingProductPositionService) Size(size int32) *StakingProductPositionService {
	s.size = &size
	return s
}

// Do sends the request.
func (s *StakingProductPositionService) Do(ctx context.Context) (*StakingProductPositions, error) {
	r := &request{
		service:  "StakingProductPositionService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/position",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	if s.productId != nil {
		r.setParam("productId", *s.productId)
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(StakingProductPositions)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// StakingHistoryService fetches the staking history
type StakingHistoryService struct {
	c               *Client
	product         StakingProduct
	transactionType StakingTransactionType
	asset           *string
	startTime       *int64
	endTime         *int64
	current         *int32
	size            *int32
}

// Product sets the product parameter.
func (s *StakingHistoryService) Product(product StakingProduct) *StakingHistoryService {
	s.product = product
	return s
}

// TransactionType sets the txnType parameter.
func (s *StakingHistoryService) TransactionType(transactionType StakingTransactionType) *StakingHistoryService {
	s.transactionType = transactionType
	return s
}

// Asset sets the asset parameter.
func (s *StakingHistoryService) Asset(asset string) *StakingHistoryService {
	s.asset = &asset
	return s
}

// StartTime sets the startTime parameter.
func (s *StakingHistoryService) StartTime(startTime int64) *StakingHistoryService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
func (s *StakingHistoryService) EndTime(endTime int64) *StakingHistoryService {
	s.endTime = &endTime
	return s
}

// Current sets the current parameter.
func (s *StakingHistoryService) Current(current int32) *StakingHistoryService {
	s.current = &current
	return s
}

// Size sets the size parameter.
func (s *StakingHistoryService) Size(size int32) *StakingHistoryService {
	s.size = &size
	return s
}

// Do sends the request.
func (s *StakingHistoryService) Do(ctx context.Context) (*StakingHistory, error) {
	r := &request{
		service:  "StakingHistoryService",
		method:   http.MethodGet,
		endpoint: "/sapi/v1/staking/stakingRecord",
		secType:  secTypeSigned,
	}
	r.setParam("product", s.product)
	r.setParam("txnType", s.transactionType)
	if s.asset != nil {
		r.setParam("asset", *s.asset)
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
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(StakingHistory)
	err = jsonCodec.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
