package utils

import (
	"context"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func KeepAlive(ctx context.Context, c *websocket.Conn, timeout time.Duration, logger *zap.SugaredLogger) {
	_ = c.SetReadDeadline(time.Now().Add(timeout))
	c.SetPongHandler(func(string) error {
		logger.Info("pong")
		return c.SetReadDeadline(time.Now().Add(timeout))
	})

	t := time.NewTicker(timeout / 2)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			logger.Info("keepAlive: context done")
			_ = c.Close()
			return
		case <-t.C:
			dl := time.Now().Add(10 * time.Second)
			if err := c.WriteControl(websocket.PingMessage, nil, dl); err != nil {
				logger.Warnf("ping failed: %v", err)
				_ = c.Close()
				return
			}
		}
	}
}
