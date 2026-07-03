package bili

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
)

// --- Tool 1: get_video_info ---

type VideoInfoRequest struct {
	Bvid string `json:"bvid"`
}

type VideoInfoResponse struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
	Cover string `json:"cover"`
	Stat  struct {
		View     int64 `json:"view"`
		Like     int64 `json:"like"`
		Coin     int64 `json:"coin"`
		Favorite int64 `json:"favorite"`
	} `json:"stat"`
}

func (c *BiliClient) GetVideoInfo(req VideoInfoRequest) (*VideoInfoResponse, error) {
	params := map[string]string{"bvid": req.Bvid}
	data, err := c.doRequest("GET", "/x/web-interface/view", params, true)
	if err != nil {
		return nil, err
	}
	var res VideoInfoResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("failed to decode video info: %w", err)
	}
	return &res, nil
}

func HandleGetVideoInfo(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "HandleGetVideoInfo"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req VideoInfoRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.Bvid == "" {
		return nil, fmt.Errorf("missing required parameter: bvid")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetVideoInfo(req)
	if err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(data)), nil
}

// --- Tool 2: get_video_detail ---

type VideoDetailRequest struct {
	Bvid string `json:"bvid"`
}

type VideoDetailResponse struct {
	Title   string   `json:"title"`
	Desc    string   `json:"desc"`
	Tags    []string `json:"tags"`
	PubDate int64    `json:"pubdate"`
	Owner   struct {
		Mid  int64  `json:"mid"`
		Name string `json:"name"`
	} `json:"owner"`
}

func (c *BiliClient) GetVideoDetail(req VideoDetailRequest) (*VideoDetailResponse, error) {
	params := map[string]string{"bvid": req.Bvid}
	data, err := c.doRequest("GET", "/x/web-interface/view/detail", params, true)
	if err != nil {
		return nil, err
	}
	var res VideoDetailResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("failed to decode video detail: %w", err)
	}
	return &res, nil
}

func HandleGetVideoDetail(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "HandleGetVideoDetail"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req VideoDetailRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.Bvid == "" {
		return nil, fmt.Errorf("missing required parameter: bvid")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetVideoDetail(req)
	if err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(data)), nil
}

// --- Tool 3: get_video_stat ---

type VideoStatRequest struct {
	Bvid string `json:"bvid"`
}

type VideoStatResponse struct {
	View     int64 `json:"view"`
	Like     int64 `json:"like"`
	Coin     int64 `json:"coin"`
	Favorite int64 `json:"favorite"`
}

func (c *BiliClient) GetVideoStat(req VideoStatRequest) (*VideoStatResponse, error) {
	params := map[string]string{"bvid": req.Bvid}
	data, err := c.doRequest("GET", "/x/web-interface/archive/stat", params, true)
	if err != nil {
		return nil, err
	}
	var res VideoStatResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("failed to decode video stat: %w", err)
	}
	return &res, nil
}

func HandleGetVideoStat(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "HandleGetVideoStat"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req VideoStatRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.Bvid == "" {
		return nil, fmt.Errorf("missing required parameter: bvid")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetVideoStat(req)
	if err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(data)), nil
}

// --- Tool 4: get_video_tags ---

type VideoTagsRequest struct {
	Bvid string `json:"bvid"`
}

type VideoTagsResponse struct {
	Tags []string `json:"tags"`
}

func (c *BiliClient) GetVideoTags(req VideoTagsRequest) (*VideoTagsResponse, error) {
	params := map[string]string{"bvid": req.Bvid}
	data, err := c.doRequest("GET", "/x/web-interface/view/detail/tag", params, true)
	if err != nil {
		return nil, err
	}
	var res VideoTagsResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("failed to decode video tags: %w", err)
	}
	return &res, nil
}

func HandleGetVideoTags(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "HandleGetVideoTags"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req VideoTagsRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.Bvid == "" {
		return nil, fmt.Errorf("missing required parameter: bvid")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetVideoTags(req)
	if err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(data)), nil
}

// --- Tool 5: get_video_desc ---

type VideoDescRequest struct {
	Bvid string `json:"bvid"`
}

type VideoDescResponse struct {
	Desc string `json:"desc"`
}

func (c *BiliClient) GetVideoDesc(req VideoDescRequest) (*VideoDescResponse, error) {
	params := map[string]string{"bvid": req.Bvid}
	data, err := c.doRequest("GET", "/x/web-interface/archive/desc", params, true)
	if err != nil {
		return nil, err
	}
	var res VideoDescResponse
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("failed to decode video desc: %w", err)
	}
	return &res, nil
}

func HandleGetVideoDesc(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "HandleGetVideoDesc"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req VideoDescRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.Bvid == "" {
		return nil, fmt.Errorf("missing required parameter: bvid")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetVideoDesc(req)
	if err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(data)), nil
}
