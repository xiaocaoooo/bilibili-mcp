package bili

import (
	"encoding/json"
	"fmt"
)

type SearchAllRequest struct {
	Keyword string `json:"keyword"`
	Page    int    `json:"page"`
}

type SearchByTypeRequest struct {
	Keyword string `json:"keyword"`
	Type    string `json:"type"`
	Page    int    `json:"page"`
}

type SearchHotRequest struct{}

type SearchSquareRequest struct {
	Page int `json:"page"`
}

type HotwordsRequest struct{}

type PopularRequest struct {
	Page int `json:"page"`
}

type TrendingRankingsRequest struct{}

type SearchSuggestionsRequest struct {
	Keyword string `json:"keyword"`
}

func (c *BiliClient) SearchAll(req SearchAllRequest) (json.RawMessage, error) {
	params := map[string]string{
		"keyword": req.Keyword,
		"page":    fmt.Sprintf("%d", req.Page),
	}
	return c.doRequest("GET", "/x/web-interface/search/all", params, true)
}

func (c *BiliClient) SearchByType(req SearchByTypeRequest) (json.RawMessage, error) {
	params := map[string]string{
		"keyword": req.Keyword,
		"type":    req.Type,
		"page":    fmt.Sprintf("%d", req.Page),
	}
	return c.doRequest("GET", "/x/web-interface/search/type", params, true)
}

func (c *BiliClient) GetSearchHot() (json.RawMessage, error) {
	return c.doRequest("GET", "/x/web-interface/search/hot", nil, true)
}

func (c *BiliClient) GetSearchSquare(req SearchSquareRequest) (json.RawMessage, error) {
	params := map[string]string{
		"page": fmt.Sprintf("%d", req.Page),
	}
	return c.doRequest("GET", "/x/web-interface/search/square", params, true)
}

func (c *BiliClient) GetHotwords() (json.RawMessage, error) {
	return c.doRequest("GET", "/x/web-interface/search/hotwords", nil, true)
}

func (c *BiliClient) GetPopular(req PopularRequest) (json.RawMessage, error) {
	params := map[string]string{
		"page": fmt.Sprintf("%d", req.Page),
	}
	return c.doRequest("GET", "/x/web-interface/popular", params, true)
}

func (c *BiliClient) GetTrendingRankings() (json.RawMessage, error) {
	return c.doRequest("GET", "/x/web-interface/ranking/trending", nil, true)
}

func (c *BiliClient) SearchSuggestions(req SearchSuggestionsRequest) (json.RawMessage, error) {
	params := map[string]string{
		"keyword": req.Keyword,
	}
	return c.doRequest("GET", "/x/web-interface/search/suggest", params, true)
}
