package bili

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"github.com/mark3labs/mcp-go/mcp"
)

// --- Tool 1: get_video_related ---

func HandleGetVideoRelated(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "HandleGetVideoRelated"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req RelatedRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.Aid == "" {
		return nil, fmt.Errorf("missing required parameter: aid")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetVideoRelated(req)
	if err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(data)), nil
}

// --- Tool 2: get_video_conclusion ---

func HandleGetVideoConclusion(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "HandleGetVideoConclusion"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req ConclusionRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.Bvid == "" {
		return nil, fmt.Errorf("missing required parameter: bvid")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetVideoConclusion(req)
	if err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(data)), nil
}

// --- Tool 3: get_video_online_count ---

func HandleGetVideoOnlineCount(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "HandleGetVideoOnlineCount"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req OnlineCountRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.Bvid == "" {
		return nil, fmt.Errorf("missing required parameter: bvid")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetVideoOnlineCount(req)
	if err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(data)), nil
}

// --- Tool 4: get_danmaku_snapshot ---

func HandleGetDanmakuSnapshot(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "HandleGetDanmakuSnapshot"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req DanmakuSnapshotRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.Bvid == "" {
		return nil, fmt.Errorf("missing required parameter: bvid")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetDanmakuSnapshot(req)
	if err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(data)), nil
}

// --- Tool 5: get_danmaku_list ---

func HandleGetDanmakuList(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "HandleGetDanmakuList"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req DanmakuListRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.Aid == "" {
		return nil, fmt.Errorf("missing required parameter: aid")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetDanmakuList(req)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(string(res)), nil
}

// --- Tool 6: get_danmaku_config ---

func HandleGetDanmakuConfig(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "HandleGetDanmakuConfig"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req DanmakuConfigRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.Bvid == "" {
		return nil, fmt.Errorf("missing required parameter: bvid")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetDanmakuConfig(req)
	if err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(data)), nil
}

// --- Tool 7: get_danmaku_buzzword ---

func HandleGetDanmakuBuzzword(ctx context.Context, client *BiliClient, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "HandleGetDanmakuBuzzword"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, request.Params)
	paramsBytes, err := json.Marshal(request.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var req DanmakuBuzzwordRequest
	if err := json.Unmarshal(paramsBytes, &req); err != nil {
		return nil, fmt.Errorf("invalid parameters: %w", err)
	}
	if req.Bvid == "" {
		return nil, fmt.Errorf("missing required parameter: bvid")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, req)
	res, err := client.GetDanmakuBuzzword(req)
	if err != nil {
		return nil, err
	}
	data, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal response: %w", err)
	}
	return mcp.NewToolResultText(string(data)), nil
}
