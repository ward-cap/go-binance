package binance

import (
	"context"
	"errors"
	"time"

	"github.com/ward-cap/go-binance/futures"
	"github.com/ward-cap/go-binance/utils"
	"go.uber.org/zap"
	"golang.org/x/net/proxy"

	"github.com/gorilla/websocket"
)

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
	}
}

var wsServe = func(
	ctx context.Context,
	cfg *WsConfig,
	handler WsHandler,
	errHandler ErrHandler,
	dialer futures.DialFunc,
	logger *zap.SugaredLogger,
) (doneC chan struct{}, err error) {
	if dialer == nil {
		dialer = proxy.Direct.Dial
	}
	if ctx == nil {
		return nil, errors.New("context is required")
	}

	ctx, cancel := context.WithCancel(ctx)

	Dialer := websocket.Dialer{
		NetDial:           dialer,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}

	c, _, err := Dialer.DialContext(ctx, cfg.Endpoint, nil)
	if err != nil {
		cancel()
		return nil, err
	}
	c.SetReadLimit(655350)
	doneC = make(chan struct{})
	go func() {
		defer cancel()
		defer close(doneC)
		go utils.KeepAlive(ctx, c, WebsocketTimeout, logger)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				cancel()
				errHandler(err)
				return
			}
			handler(message)
		}
	}()
	return
}
