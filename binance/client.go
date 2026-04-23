package binance

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/ward-cap/go-binance/common"
	"go.uber.org/zap"
)

func NewClient(client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	return &Client{

		BaseURL:    "https://www.binance.com",
		UserAgent:  "Binance/golang",
		HTTPClient: client,
	}
}

type Client struct {
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	TimeOffset int64

	Logger *zap.SugaredLogger
}

type request struct {
	method   string
	endpoint string
	query    url.Values
	form     url.Values
	//recvWindow int64
	//secType    secType
	header  http.Header
	body    io.Reader
	fullURL string
	service string
}

func (r *request) setParam(key string, value any) *request {
	if r.query == nil {
		r.query = url.Values{}
	}

	if value != nil {
		if valueType := reflect.TypeOf(value); valueType != nil && valueType.Kind() == reflect.Slice {
			v, err := json.Marshal(value)
			if err == nil {
				value = string(v)
			}
		}
	}

	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

func (c *Client) parseRequest(r *request) (err error) {

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)

	//if r.secType == secTypeSigned {
	//	r.setParam(timestampKey, currentTimestamp()-c.TimeOffset)
	//}
	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	}
	//if r.secType == secTypeAPIKey || r.secType == secTypeSigned {
	//	header.Set("X-MBX-APIKEY", c.APIKey)
	//}

	//if r.secType == secTypeSigned {
	//	raw := fmt.Sprintf("%s%s", queryString, bodyString)
	//	mac := hmac.New(sha256.New, []byte(c.SecretKey))
	//	_, err = mac.Write([]byte(raw))
	//	if err != nil {
	//		return err
	//	}
	//	v := url.Values{}
	//	v.Set(signatureKey, fmt.Sprintf("%x", mac.Sum(nil)))
	//	if queryString == "" {
	//		queryString = v.Encode()
	//	} else {
	//		queryString = fmt.Sprintf("%s&%s", queryString, v.Encode())
	//	}
	//}
	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}

	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request) (data []byte, err error) {
	startedAt := time.Now()
	service := r.service
	ctx, span := common.StartRequestSpan(ctx, "go-binance/binance", service)
	defer span.End()

	err = c.parseRequest(r)
	if err != nil {
		c.logAPIError(ctx, service, r, nil, startedAt, err)
		return []byte{}, err
	}
	req, err := http.NewRequestWithContext(ctx, r.method, r.fullURL, r.body)
	if err != nil {
		c.logAPIError(ctx, service, r, nil, startedAt, err)
		return []byte{}, err
	}
	req = req.WithContext(common.WithHTTPConnTrace(req.Context()))
	req.Header = r.header

	c.logAPIRequest(ctx, service, r, req)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		c.logAPIError(ctx, service, r, req, startedAt, err)
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		c.logAPIError(ctx, service, r, req, startedAt, err)
		return []byte{}, err
	}

	c.logAPIResponse(ctx, service, r, req, res, data, startedAt)

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(common.APIError)
		_ = json.Unmarshal(data, apiErr)
		c.logAPIError(ctx, service, r, req, startedAt, apiErr)
		return nil, apiErr
	}
	return data, nil
}

func (c *Client) logAPIRequest(ctx context.Context, service string, r *request, req *http.Request) {
	if c == nil || c.Logger == nil || req == nil {
		return
	}

	fields := []any{
		"binance.package", "binance",
		"binance.service", service,
		"http.method", r.method,
		"url.full", common.SanitizeURL(req.URL.String()),
		"server.address", req.URL.Host,
		"http.request.header", common.SanitizeHeaders(req.Header),
		"http.request.body", common.ReadBodyForLog(r.body),
	}
	fields = common.AppendContextField(ctx, fields)

	c.Logger.Debugw("binance api request", fields...)
}

func (c *Client) logAPIResponse(ctx context.Context, service string, r *request, req *http.Request, res *http.Response, data []byte, startedAt time.Time) {
	if c == nil || c.Logger == nil || req == nil || res == nil {
		return
	}

	fields := []any{
		"binance.package", "binance",
		"binance.service", service,
		"http.method", r.method,
		"url.full", common.SanitizeURL(req.URL.String()),
		"server.address", req.URL.Host,
		"http.response.status_code", res.StatusCode,
		"http.response.header", common.SanitizeHeaders(res.Header),
		"http.response.body", string(data),
		"event.duration", time.Since(startedAt),
	}
	if reused, ok := common.HTTPConnReused(req.Context()); ok {
		fields = append(fields, "http.connection.reused", reused)
	}
	fields = common.AppendContextField(ctx, fields)

	c.Logger.Debugw("binance api response", fields...)
}

func (c *Client) logAPIError(ctx context.Context, service string, r *request, req *http.Request, startedAt time.Time, err error) {
	if c == nil || c.Logger == nil || err == nil {
		return
	}

	serverAddress := ""
	if req != nil && req.URL != nil {
		serverAddress = req.URL.Host
	}

	fields := []any{
		"binance.package", "binance",
		"binance.service", service,
		"http.method", r.method,
		"url.full", common.SanitizeURL(r.fullURL),
		"server.address", serverAddress,
		"event.duration", time.Since(startedAt),
		"error", err,
	}
	fields = common.AppendContextField(ctx, fields)

	c.Logger.Errorw("binance api error", fields...)
}

func (c *Client) NewGetAllAssetsService() *GetAllAssetsService {
	return &GetAllAssetsService{c: c}
}
