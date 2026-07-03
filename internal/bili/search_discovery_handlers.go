package bili

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
)

func HandleSearchAll(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "SearchAll"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, req.Params)
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params SearchAllRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	if params.Keyword == "" {
		return nil, fmt.Errorf("missing required parameter: keyword")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, params)
	res, err := client.SearchAll(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleSearchByType(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "SearchByType"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, req.Params)
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params SearchByTypeRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	if params.Keyword == "" {
		return nil, fmt.Errorf("missing required parameter: keyword")
	}
	if params.Type == "" {
		return nil, fmt.Errorf("missing required parameter: type")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, params)
	res, err := client.SearchByType(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetSearchHot(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetSearchHot"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, req.Params)
	res, err := client.GetSearchHot()
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetSearchSquare(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetSearchSquare"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, req.Params)
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params SearchSquareRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, params)
	res, err := client.GetSearchSquare(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetHotwords(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetHotwords"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, req.Params)
	res, err := client.GetHotwords()
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetPopular(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetPopular"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, req.Params)
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params PopularRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, params)
	res, err := client.GetPopular(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetTrendingRankings(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "GetTrendingRankings"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, req.Params)
	res, err := client.GetTrendingRankings()
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleSearchSuggestions(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := "SearchSuggestions"
	log.Printf("[MCP Handler] Tool: %s | Raw Params: %v", toolName, req.Params)
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params SearchSuggestionsRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	if params.Keyword == "" {
		return nil, fmt.Errorf("missing required parameter: keyword")
	}
	log.Printf("[MCP Handler] Tool: %s | Parsed Request: %+v", toolName, params)
	res, err := client.SearchSuggestions(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}
