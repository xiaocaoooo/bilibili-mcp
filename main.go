package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/xiaocaoooo/bilibili-mcp/internal/bili"
)

func main() {
	// 1. 初始化 BiliClient
	client := bili.NewBiliClient()

	// 2. 初始化 MCP 服务器
	s := server.NewMCPServer("Bilibili MCP Server", "1.0.0")

	// 3. 注册所有工具
	registerTools(s, client)

	// 4. 传输层配置
	port := os.Getenv("MCP_PORT")
	if port == "" {
		port = "8080"
	}

	// 实现优雅停机
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	httpServer := server.NewStreamableHTTPServer(s)

	go func() {
		log.Printf("Starting Bilibili MCP Server (streamable-http) on port %s...", port)
		// 使用 streamable-http 传输协议监听
		if err := httpServer.Start(":" + port); err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down Bilibili MCP Server...")
	if err := httpServer.Shutdown(context.Background()); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
}

func registerTools(s *server.MCPServer, client *bili.BiliClient) {
	tools := []struct {
		name        string
		description string
		handler     func(context.Context, *bili.BiliClient, mcp.CallToolRequest) (*mcp.CallToolResult, error)
	}{
		// Video Tools
		{"get_video_info", "Get basic information of a video by its bvid", bili.HandleGetVideoInfo},
		{"get_video_detail", "Get detailed information of a video", bili.HandleGetVideoDetail},
		{"get_video_stat", "Get statistics (views, likes, etc.) of a video", bili.HandleGetVideoStat},
		{"get_video_tags", "Get tags of a video", bili.HandleGetVideoTags},
		{"get_video_desc", "Get description of a video", bili.HandleGetVideoDesc},
		{"get_video_related", "Get related videos", bili.HandleGetVideoRelated},
		{"get_video_conclusion", "Get video conclusion/summary", bili.HandleGetVideoConclusion},
		{"get_video_online_count", "Get current online viewers of a video", bili.HandleGetVideoOnlineCount},
		{"get_danmaku_snapshot", "Get danmaku snapshot", bili.HandleGetDanmakuSnapshot},
		{"get_danmaku_list", "Get danmaku list", bili.HandleGetDanmakuList},
		{"get_danmaku_config", "Get danmaku configuration", bili.HandleGetDanmakuConfig},
		{"get_danmaku_buzzword", "Get danmaku buzzwords", bili.HandleGetDanmakuBuzzword},

		// User Tools
		{"get_user_profile", "Get user profile information", bili.HandleGetUserProfile},
		{"get_user_card", "Get user card information", bili.HandleGetUserCard},
		{"get_user_relation_stat", "Get user relation statistics", bili.HandleGetUserRelationStat},
		{"get_user_public_dynamics", "Get user public dynamics", bili.HandleGetUserPublicDynamics},
		{"get_dynamic_detail", "Get detailed information of a dynamic", bili.HandleGetDynamicDetail},

		// Search Tools
		{"search_all", "Search across all categories", bili.HandleSearchAll},
		{"search_by_type", "Search by specific type", bili.HandleSearchByType},
		{"get_search_hot", "Get current hot search keywords", bili.HandleGetSearchHot},
		{"get_search_square", "Get search square content", bili.HandleGetSearchSquare},
		{"get_hotwords", "Get trending hotwords", bili.HandleGetHotwords},
		{"get_popular", "Get popular content", bili.HandleGetPopular},
		{"get_trending_rankings", "Get trending rankings", bili.HandleGetTrendingRankings},
		{"search_suggestions", "Get search suggestions", bili.HandleSearchSuggestions},

		// Social Tools
		{"get_column_info", "Get column information", bili.HandleGetColumnInfo},
		{"get_column_list", "Get column list", bili.HandleGetColumnList},
		{"get_topic_details", "Get topic details", bili.HandleGetTopicDetails},
		{"get_video_comments", "Get comments for a video", bili.HandleGetVideoComments},

		// Live Tools
		{"get_live_room_info", "Get live room information", bili.HandleGetLiveRoomInfo},
		{"get_live_play_url", "Get live play URL", bili.HandleGetLivePlayURL},
		{"get_anchor_info", "Get anchor information", bili.HandleGetAnchorInfo},
		{"get_live_recommend", "Get live recommendations", bili.HandleGetLiveRecommend},
		{"get_room_status_batch", "Get status of multiple live rooms", bili.HandleGetRoomStatusBatch},

		// Ranking Tools
		{"get_region_ranking", "Get regional rankings", bili.HandleGetRegionRanking},
		{"get_weekly_must_watch", "Get weekly must-watch list", bili.HandleGetWeeklyMustWatch},
		{"get_weekly_must_watch_detail", "Get detailed info of a weekly must-watch issue", bili.HandleGetWeeklyMustWatchDetail},
		{"get_popular_precious", "Get popular precious videos", bili.HandleGetPopularPrecious},
	}

	for _, t := range tools {
		s.AddTool(mcp.NewTool(t.name,
			mcp.WithDescription(t.description),
		), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			return t.handler(ctx, client, request)
		})
	}
}
