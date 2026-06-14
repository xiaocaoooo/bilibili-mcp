package bili

import (
	"encoding/json"
	"fmt"
)

// --- Request/Response Structures ---

type RelatedRequest struct {
	Aid string `json:"aid"`
}

type RelatedResponse struct {
	List []struct {
		Title string `json:"title"`
		Bvid  string `json:"bvid"`
		Author struct {
			Name string `json:"name"`
		} `json:"author"`
	} `json:"list"`
}

type ConclusionRequest struct {
	Bvid string `json:"bvid"`
}

type ConclusionResponse struct {
	Conclusion string `json:"conclusion"`
}

type OnlineCountRequest struct {
	Bvid string `json:"bvid"`
}

type OnlineCountResponse struct {
	Total int `json:"total"`
}

type DanmakuSnapshotRequest struct {
	Bvid string `json:"bvid"`
}

type DanmakuSnapshotResponse struct {
	List []struct {
		Text string `json:"text"`
		Time int    `json:"time"`
	} `json:"list"`
}

type DanmakuListRequest struct {
	Aid string `json:"aid"`
}

type DanmakuConfigRequest struct {
	Bvid string `json:"bvid"`
}

type DanmakuConfigResponse struct {
	Config map[string]interface{} `json:"config"`
}

type DanmakuBuzzwordRequest struct {
	Bvid string `json:"bvid"`
}

type DanmakuBuzzwordResponse struct {
	Buzzwords []string `json:"buzzwords"`
}

// --- Client Methods ---

func (c *BiliClient) GetVideoRelated(req RelatedRequest) (*RelatedResponse, error) {
	params := map[string]string{"aid": req.Aid}
	data, err := c.doRequest("GET", "/x/web-interface/archive/related", params, false)
	if err != nil {
		return nil, err
	}

	var resp RelatedResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode related response: %w", err)
	}
	return &resp, nil
}

func (c *BiliClient) GetVideoConclusion(req ConclusionRequest) (*ConclusionResponse, error) {
	params := map[string]string{"bvid": req.Bvid}
	data, err := c.doRequest("GET", "/x/web-interface/view/conclusion/get", params, false)
	if err != nil {
		return nil, err
	}

	var resp ConclusionResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode conclusion response: %w", err)
	}
	return &resp, nil
}

func (c *BiliClient) GetVideoOnlineCount(req OnlineCountRequest) (*OnlineCountResponse, error) {
	params := map[string]string{"bvid": req.Bvid}
	data, err := c.doRequest("GET", "/x/player/online/total", params, false)
	if err != nil {
		return nil, err
	}

	var resp OnlineCountResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode online count response: %w", err)
	}
	return &resp, nil
}

func (c *BiliClient) GetDanmakuSnapshot(req DanmakuSnapshotRequest) (*DanmakuSnapshotResponse, error) {
	params := map[string]string{"bvid": req.Bvid}
	data, err := c.doRequest("GET", "/x/v2/dm/ajax", params, false)
	if err != nil {
		return nil, err
	}

	var resp DanmakuSnapshotResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode snapshot response: %w", err)
	}
	return &resp, nil
}

func (c *BiliClient) GetDanmakuList(req DanmakuListRequest) ([]byte, error) {
	params := map[string]string{"aid": req.Aid}
	// This returns XML, use doRawRequest
	return c.doRawRequest("GET", "/x/v1/dm/list.so", params, false)
}

func (c *BiliClient) GetDanmakuConfig(req DanmakuConfigRequest) (*DanmakuConfigResponse, error) {
	params := map[string]string{"bvid": req.Bvid}
	data, err := c.doRequest("GET", "/x/v2/dm/web/config", params, false)
	if err != nil {
		return nil, err
	}

	var resp DanmakuConfigResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode config response: %w", err)
	}
	return &resp, nil
}

func (c *BiliClient) GetDanmakuBuzzword(req DanmakuBuzzwordRequest) (*DanmakuBuzzwordResponse, error) {
	params := map[string]string{"bvid": req.Bvid}
	data, err := c.doRequest("GET", "/x/v2/dm/buzzword/list", params, false)
	if err != nil {
		return nil, err
	}

	var resp DanmakuBuzzwordResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode buzzword response: %w", err)
	}
	return &resp, nil
}
