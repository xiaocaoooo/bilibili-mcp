package bili

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
)

// --- Models for get_region_ranking ---

type RegionRankingRequest struct {
	RegionID string `json:"region_id"`
}

type RankingVideo struct {
	AvID    int64  `json:"aid"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Rank    int    `json:"rank"`
	View    int64  `json:"view"`
	Cover   string `json:"cover"`
	URL     string `json:"url"`
}

type RegionRankingResponse struct {
	List []RankingVideo `json:"list"`
}

// --- Models for get_weekly_must_watch ---

type WeeklyMustWatchResponse struct {
	List []WeeklyIssue `json:"list"`
}

type WeeklyIssue struct {
	Num     int    `json:"num"`
	Title   string `json:"title"`
	Cover   string `json:"cover"`
	Update  string `json:"update_time"`
}

// --- Models for get_weekly_must_watch_detail ---

type WeeklyMustWatchDetailRequest struct {
	Num string `json:"num"`
}

type WeeklyMustWatchDetailResponse struct {
	Title   string          `json:"title"`
	Issues  []WeeklyVideo   `json:"issues"`
	Update  string          `json:"update_time"`
}

type WeeklyVideo struct {
	AvID    int64  `json:"aid"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Cover   string `json:"cover"`
	URL     string `json:"url"`
	Reason  string `json:"reason"`
}

// --- Models for get_popular_precious ---

type PopularPreciousResponse struct {
	List []PreciousVideo `json:"list"`
}

type PreciousVideo struct {
	AvID    int64  `json:"aid"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Cover   string `json:"cover"`
	URL     string `json:"url"`
	Reason  string `json:"reason"`
}

// --- Client Methods ---

func (c *BiliClient) GetRegionRanking(req RegionRankingRequest) (*RegionRankingResponse, error) {
	params := map[string]string{
		"region_id": req.RegionID,
	}
	// This API typically uses WBI signing
	data, err := c.doRequest("GET", "/x/web-interface/ranking/v2", params, true)
	if err != nil {
		return nil, err
	}

	var resp RegionRankingResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode ranking response: %w", err)
	}
	return &resp, nil
}

func (c *BiliClient) GetWeeklyMustWatchList() (*WeeklyMustWatchResponse, error) {
	params := map[string]string{}
	data, err := c.doRequest("GET", "/x/web-interface/popular/series/list", params, false)
	if err != nil {
		return nil, err
	}

	var resp WeeklyMustWatchResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode weekly list response: %w", err)
	}
	return &resp, nil
}

func (c *BiliClient) GetWeeklyMustWatchDetail(req WeeklyMustWatchDetailRequest) (*WeeklyMustWatchDetailResponse, error) {
	params := map[string]string{
		"num": req.Num,
	}
	data, err := c.doRequest("GET", "/x/web-interface/popular/series/one", params, false)
	if err != nil {
		return nil, err
	}

	var resp WeeklyMustWatchDetailResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode weekly detail response: %w", err)
	}
	return &resp, nil
}

func (c *BiliClient) GetPopularPrecious() (*PopularPreciousResponse, error) {
	params := map[string]string{}
	data, err := c.doRequest("GET", "/x/web-interface/popular/precious", params, false)
	if err != nil {
		return nil, err
	}

	var resp PopularPreciousResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode precious response: %w", err)
	}
	return &resp, nil
}

// --- MCP Handlers ---

func HandleGetRegionRanking(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetRegionRanking"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req RegionRankingRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.RegionID == "" {
		return nil, fmt.Errorf("missing required parameter: region_id")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetRegionRanking(req)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("API error: %v", err)), nil
	}

	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetWeeklyMustWatch(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetWeeklyMustWatch"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	res, err := client.GetWeeklyMustWatchList()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("API error: %v", err)), nil
	}

	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetWeeklyMustWatchDetail(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetWeeklyMustWatchDetail"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req WeeklyMustWatchDetailRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.Num == "" {
		return nil, fmt.Errorf("missing required parameter: num")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetWeeklyMustWatchDetail(req)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("API error: %v", err)), nil
	}

	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetPopularPrecious(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetPopularPrecious"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	res, err := client.GetPopularPrecious()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("API error: %v", err)), nil
	}

	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}
