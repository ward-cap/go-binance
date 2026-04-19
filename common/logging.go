package common

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

//const maxLoggedBodySize = 8 * 1024

func AppendTraceFields(ctx context.Context, fields []any) []any {
	if ctx == nil {
		return fields
	}

	sc := trace.SpanContextFromContext(ctx)
	if !sc.IsValid() {
		return fields
	}

	return append(fields,
		"trace_id", sc.TraceID().String(),
		"span_id", sc.SpanID().String(),
	)
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
