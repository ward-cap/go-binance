package binance

import (
	"context"
	"encoding/json"
	"net/http"
)

type GetAllAnnouncementsService struct {
	c *Client
}

type GetAllAnnouncementsResponse struct {
	OpenTime int64    `json:"openTime"`
	Symbols  []string `json:"symbols"`
}

func (s *GetAllAnnouncementsService) Do(ctx context.Context) (res []GetAllAnnouncementsResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/spot/open-symbol-list",
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &res)
	//if err != nil {
	//	return
	//}

	//if !res.Success {
	//	text := "failed to get announcements"
	//	if res.Message != nil {
	//		text += ": " + *res.Message
	//	}
	//	err = errors.New(text)
	//}
	return res, err
}
