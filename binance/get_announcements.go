package binance

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type GetAllAnnouncementsService struct {
	c *Client
}

type GetAllAnnouncementsResponse struct {
	Message *string `json:"message"`
	Data    struct {
		Catalogs []struct {
			CatalogId   int    `json:"catalogId"`
			Icon        string `json:"icon"`
			CatalogName string `json:"catalogName"`
			CatalogType int    `json:"catalogType"`
			Total       int    `json:"total"`
			Articles    []struct {
				Id          int    `json:"id"`
				Code        string `json:"code"`
				Title       string `json:"title"`
				Type        int    `json:"type"`
				ReleaseDate int64  `json:"releaseDate"`
			} `json:"articles"`
			//ParentCatalogId interface{}   `json:"parentCatalogId"`
			//Description     interface{}   `json:"description"`
			//Catalogs        []interface{} `json:"catalogs"`
		} `json:"catalogs"`
	} `json:"data"`
	Success bool `json:"success"`
}

func (s *GetAllAnnouncementsService) Do(ctx context.Context) (res GetAllAnnouncementsResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/bapi/apex/v1/public/apex/cms/article/list/query?type=1&pageNo=1&pageSize=10&catalogId=48", // todo: move params to rq
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return
	}
	//res = make([]*Deposit, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}

	if !res.Success {
		text := "failed to get announcements"
		if res.Message != nil {
			text += ": " + *res.Message
		}
		err = errors.New(text)
	}
	return res, err
}
