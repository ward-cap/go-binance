package futures

import (
	"context"
	"go.uber.org/zap"
	"golang.org/x/net/proxy"
	"net"
	"time"

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

type DialFunc func(network, addr string) (net.Conn, error)

var wsServe = func(
	ctx context.Context,
	cfg *WsConfig,
	handler WsHandler,
	errHandler ErrHandler,
	dialer DialFunc,
	logger *zap.SugaredLogger,
	sendMessageAfterConnect ...[]byte,
) (doneC, stopC chan struct{}, err error) {
	if dialer == nil {
		dialer = proxy.Direct.Dial
	}
	if ctx == nil {
		ctx = context.Background()
	}

	Dialer := websocket.Dialer{
		NetDial:           dialer,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}

	c, _, err := Dialer.DialContext(ctx, cfg.Endpoint, nil)
	if err != nil {
		return nil, nil, err
	}
	c.SetReadLimit(655350)
	doneC = make(chan struct{})
	stopC = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneC)
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout, logger)
		}
		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		silent := false
		go func() {
			select {
			case <-stopC:
				silent = true
			case <-doneC:
			}
			_ = c.Close()
		}()
		go func() {
			for _, msg := range sendMessageAfterConnect {
				err := c.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					logger.Error(err.Error())
				}
			}
		}()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if !silent {
					errHandler(err)
				}
				return
			}
			handler(message)
		}
	}()
	return
}

func keepAlive(c *websocket.Conn, timeout time.Duration, logger *zap.SugaredLogger) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		logger.Info("pong")
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > timeout {
				logger.Info("close")
				err = c.Close()
				if err != nil {
					logger.Info(err.Error())
				}
				return
			}
		}
	}()
}
