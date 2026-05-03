package futures

import (
	"context"
	"net/http"
)

// StartUserStreamService create listen key for user stream service
type StartUserStreamService struct {
	c *Client
}

// Do send request
func (s *StartUserStreamService) Do(ctx context.Context, opts ...RequestOption) (listenKey string, err error) {
	r := &request{
		service:  "StartUserStreamService",
		method:   http.MethodPost,
		endpoint: "/fapi/v1/listenKey",
		secType:  secTypeSigned,
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}
	return parseListenKey(data)
}

// KeepaliveUserStreamService update listen key
type KeepaliveUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *KeepaliveUserStreamService) ListenKey(listenKey string) *KeepaliveUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *KeepaliveUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		service:  "KeepaliveUserStreamService",
		method:   http.MethodPut,
		endpoint: "/fapi/v1/listenKey",
		secType:  secTypeSigned,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, _, err = s.c.callAPI(ctx, r, opts...)
	return err
}

// CloseUserStreamService delete listen key
type CloseUserStreamService struct {
	c         *Client
	listenKey string
}

// ListenKey set listen key
func (s *CloseUserStreamService) ListenKey(listenKey string) *CloseUserStreamService {
	s.listenKey = listenKey
	return s
}

// Do send request
func (s *CloseUserStreamService) Do(ctx context.Context, opts ...RequestOption) (err error) {
	r := &request{
		service:  "CloseUserStreamService",
		method:   http.MethodDelete,
		endpoint: "/fapi/v1/listenKey",
		secType:  secTypeSigned,
	}
	r.setFormParam("listenKey", s.listenKey)
	_, _, err = s.c.callAPI(ctx, r, opts...)
	return err
}
