package common

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"strings"
	"sync/atomic"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type httpConnTraceContextKey struct{}

type httpConnTraceState struct {
	observed atomic.Bool
	reused   atomic.Bool
}

func AppendContextField(ctx context.Context, fields []any) []any {
	if ctx == nil {
		return fields
	}

	return append(fields, "ctx", ctx)
}

func WithHTTPConnTrace(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	state := &httpConnTraceState{}
	ctx = context.WithValue(ctx, httpConnTraceContextKey{}, state)

	trace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			state.observed.Store(true)
			state.reused.Store(info.Reused)
		},
	}

	return httptrace.WithClientTrace(ctx, trace)
}

func HTTPConnReused(ctx context.Context) (reused bool, ok bool) {
	if ctx == nil {
		return false, false
	}

	state, _ := ctx.Value(httpConnTraceContextKey{}).(*httpConnTraceState)
	if state == nil || !state.observed.Load() {
		return false, false
	}

	return state.reused.Load(), true
}

func StartRequestSpan(ctx context.Context, instrumentationName, spanName string) (context.Context, trace.Span) {
	if ctx == nil {
		ctx = context.Background()
	}

	if spanName == "" {
		spanName = "binance.api"
	}

	return otel.Tracer(instrumentationName).Start(ctx, spanName)
}

func SanitizeURL(raw string, hiddenParams ...string) string {
	if raw == "" {
		return ""
	}

	parsed, err := url.Parse(raw)
	if err != nil {
		return raw
	}

	query := parsed.Query()
	for _, key := range hiddenParams {
		for queryKey := range query {
			if strings.EqualFold(queryKey, key) {
				query.Set(queryKey, "***")
			}
		}
	}
	parsed.RawQuery = query.Encode()

	return parsed.String()
}

func SanitizeHeaders(header http.Header) http.Header {
	if header == nil {
		return nil
	}

	cloned := header.Clone()
	for key, values := range cloned {
		if strings.EqualFold(key, "X-MBX-APIKEY") || strings.EqualFold(key, "Authorization") {
			masked := make([]string, len(values))
			for i, value := range values {
				masked[i] = MaskAPIKey(value)
			}
			cloned[key] = masked
		}
	}

	return cloned
}

func MaskAPIKey(value string) string {
	if value == "" {
		return ""
	}

	if len(value) <= 8 {
		return "****"
	}

	return value[:4] + "****" + value[len(value)-4:]
}

func ReadBodyForLog(body io.Reader) string {
	if body == nil {
		return ""
	}

	var payload string
	switch v := body.(type) {
	case *bytes.Buffer:
		payload = v.String()
	case *bytes.Reader:
		cloned := *v
		if data, err := io.ReadAll(&cloned); err == nil {
			payload = string(data)
		}
	case *strings.Reader:
		cloned := *v
		if data, err := io.ReadAll(&cloned); err == nil {
			payload = string(data)
		}
	}

	return payload
}
