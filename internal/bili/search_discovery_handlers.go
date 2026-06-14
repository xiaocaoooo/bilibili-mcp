package bili

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func HandleSearchAll(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params SearchAllRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.SearchAll(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleSearchByType(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params SearchByTypeRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.SearchByType(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetSearchHot(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	res, err := client.GetSearchHot()
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetSearchSquare(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params SearchSquareRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.GetSearchSquare(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetHotwords(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	res, err := client.GetHotwords()
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetPopular(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params PopularRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.GetPopular(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetTrendingRankings(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	res, err := client.GetTrendingRankings()
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleSearchSuggestions(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params SearchSuggestionsRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.SearchSuggestions(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}
