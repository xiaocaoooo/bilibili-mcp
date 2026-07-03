package bili

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
)

func HandleGetLiveRoomInfo(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetLiveRoomInfo"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, req.Params)
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params RoomInfoRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	if params.RoomID == "" {
		return nil, fmt.Errorf("missing required parameter: room_id")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, params)
	res, err := client.GetLiveRoomInfo(params)
	if err != nil {
		return nil, err
	}

	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetLivePlayURL(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetLivePlayURL"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, req.Params)
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params PlayURLRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	if params.RoomID == "" {
		return nil, fmt.Errorf("missing required parameter: room_id")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, params)
	res, err := client.GetLivePlayURL(params)
	if err != nil {
		return nil, err
	}

	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetAnchorInfo(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetAnchorInfo"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, req.Params)
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params AnchorInfoRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	if params.UID == "" {
		return nil, fmt.Errorf("missing required parameter: uid")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, params)
	res, err := client.GetAnchorInfo(params)
	if err != nil {
		return nil, err
	}

	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetLiveRecommend(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetLiveRecommend"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, req.Params)
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params LiveRecommendRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, params)
	res, err := client.GetLiveRecommend(params)
	if err != nil {
		return nil, err
	}

	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetRoomStatusBatch(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetRoomStatusBatch"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, req.Params)
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params RoomStatusBatchRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	if len(params.UIDs) == 0 {
		return nil, fmt.Errorf("missing required parameter: uids")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, params)
	res, err := client.GetRoomStatusBatch(params)
	if err != nil {
		return nil, err
	}

	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}
