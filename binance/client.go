package binance

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ward-cap/go-binance/common"
	"io"
	"net/http"
	"net/url"
	"reflect"
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
}

func (r *request) setParam(key string, value any) *request {
	if r.query == nil {
		r.query = url.Values{}
	}

	if reflect.TypeOf(value).Kind() == reflect.Slice {
		v, err := json.Marshal(value)
		if err == nil {
			value = string(v)
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
	err = c.parseRequest(r)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(common.APIError)
		_ = json.Unmarshal(data, apiErr)
		return nil, apiErr
	}
	return data, nil
}

//// SetApiEndpoint set api Endpoint
//func (c *Client) SetApiEndpoint(url string) *Client {
//	c.BaseURL = url
//	return c
//}

func (c *Client) NewGetAllAssetsService() *GetAllAssetsService {
	return &GetAllAssetsService{c: c}
}

func (c *Client) NewGetAllAnnouncementsService() *GetAllAnnouncementsService {
	return &GetAllAnnouncementsService{c: c}
}
