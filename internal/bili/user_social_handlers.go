package bili

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func HandleGetUserProfile(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params UserProfileRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.GetUserProfile(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetUserCard(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params UserCardRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.GetUserCard(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetUserRelationStat(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params UserRelationStatRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.GetUserRelationStat(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetUserPublicDynamics(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params UserPublicDynamicsRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.GetUserPublicDynamics(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetDynamicDetail(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params DynamicDetailRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.GetDynamicDetail(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetColumnInfo(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params ColumnInfoRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.GetColumnInfo(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetColumnList(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params ColumnListRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.GetColumnList(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetTopicDetails(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params TopicDetailsRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.GetTopicDetails(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}

func HandleGetVideoComments(ctx context.Context, client *BiliClient, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	paramsBytes, err := json.Marshal(req.Params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}
	var params VideoCommentsRequest
	if err := json.Unmarshal(paramsBytes, &params); err != nil {
		return nil, fmt.Errorf("invalid params: %w", err)
	}
	res, err := client.GetVideoComments(params)
	if err != nil {
		return nil, err
	}
	out, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(out)), nil
}
