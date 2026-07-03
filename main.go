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

func wrapHandler(name string, client *bili.BiliClient, handler func(context.Context, *bili.BiliClient, mcp.CallToolRequest) (*mcp.CallToolResult, error)) func(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("[MCP Request] Tool: %s | Raw Params: %v", name, request.Params)
		return handler(ctx, client, request)
	}
}

func registerTools(s *server.MCPServer, client *bili.BiliClient) {
	// Video Tools
	s.AddTool(mcp.NewTool("get_video_info", mcp.WithDescription("Get basic information of a video by its bvid"), mcp.WithInputSchema[bili.VideoInfoRequest]()), wrapHandler("get_video_info", client, bili.HandleGetVideoInfo))
	s.AddTool(mcp.NewTool("get_video_detail", mcp.WithDescription("Get detailed information of a video"), mcp.WithInputSchema[bili.VideoDetailRequest]()), wrapHandler("get_video_detail", client, bili.HandleGetVideoDetail))
	s.AddTool(mcp.NewTool("get_video_stat", mcp.WithDescription("Get statistics (views, likes, etc.) of a video"), mcp.WithInputSchema[bili.VideoStatRequest]()), wrapHandler("get_video_stat", client, bili.HandleGetVideoStat))
	s.AddTool(mcp.NewTool("get_video_tags", mcp.WithDescription("Get tags of a video"), mcp.WithInputSchema[bili.VideoTagsRequest]()), wrapHandler("get_video_tags", client, bili.HandleGetVideoTags))
	s.AddTool(mcp.NewTool("get_video_desc", mcp.WithDescription("Get description of a video"), mcp.WithInputSchema[bili.VideoDescRequest]()), wrapHandler("get_video_desc", client, bili.HandleGetVideoDesc))
	s.AddTool(mcp.NewTool("get_video_related", mcp.WithDescription("Get related videos"), mcp.WithInputSchema[bili.RelatedRequest]()), wrapHandler("get_video_related", client, bili.HandleGetVideoRelated))
	s.AddTool(mcp.NewTool("get_video_conclusion", mcp.WithDescription("Get video conclusion/summary"), mcp.WithInputSchema[bili.ConclusionRequest]()), wrapHandler("get_video_conclusion", client, bili.HandleGetVideoConclusion))
	s.AddTool(mcp.NewTool("get_video_online_count", mcp.WithDescription("Get current online viewers of a video"), mcp.WithInputSchema[bili.OnlineCountRequest]()), wrapHandler("get_video_online_count", client, bili.HandleGetVideoOnlineCount))
	s.AddTool(mcp.NewTool("get_danmaku_snapshot", mcp.WithDescription("Get danmaku snapshot"), mcp.WithInputSchema[bili.DanmakuSnapshotRequest]()), wrapHandler("get_danmaku_snapshot", client, bili.HandleGetDanmakuSnapshot))
	s.AddTool(mcp.NewTool("get_danmaku_list", mcp.WithDescription("Get danmaku list"), mcp.WithInputSchema[bili.DanmakuListRequest]()), wrapHandler("get_danmaku_list", client, bili.HandleGetDanmakuList))
	s.AddTool(mcp.NewTool("get_danmaku_config", mcp.WithDescription("Get danmaku configuration"), mcp.WithInputSchema[bili.DanmakuConfigRequest]()), wrapHandler("get_danmaku_config", client, bili.HandleGetDanmakuConfig))
	s.AddTool(mcp.NewTool("get_danmaku_buzzword", mcp.WithDescription("Get danmaku buzzwords"), mcp.WithInputSchema[bili.DanmakuBuzzwordRequest]()), wrapHandler("get_danmaku_buzzword", client, bili.HandleGetDanmakuBuzzword))

	// User Tools
	s.AddTool(mcp.NewTool("get_user_profile", mcp.WithDescription("Get user profile information"), mcp.WithInputSchema[bili.UserProfileRequest]()), wrapHandler("get_user_profile", client, bili.HandleGetUserProfile))
	s.AddTool(mcp.NewTool("get_user_card", mcp.WithDescription("Get user card information"), mcp.WithInputSchema[bili.UserCardRequest]()), wrapHandler("get_user_card", client, bili.HandleGetUserCard))
	s.AddTool(mcp.NewTool("get_user_relation_stat", mcp.WithDescription("Get user relation statistics"), mcp.WithInputSchema[bili.UserRelationStatRequest]()), wrapHandler("get_user_relation_stat", client, bili.HandleGetUserRelationStat))
	s.AddTool(mcp.NewTool("get_user_public_dynamics", mcp.WithDescription("Get user public dynamics"), mcp.WithInputSchema[bili.UserPublicDynamicsRequest]()), wrapHandler("get_user_public_dynamics", client, bili.HandleGetUserPublicDynamics))
	s.AddTool(mcp.NewTool("get_dynamic_detail", mcp.WithDescription("Get detailed information of a dynamic"), mcp.WithInputSchema[bili.DynamicDetailRequest]()), wrapHandler("get_dynamic_detail", client, bili.HandleGetDynamicDetail))

	// Search Tools
	s.AddTool(mcp.NewTool("search_all", mcp.WithDescription("Search across all categories"), mcp.WithInputSchema[bili.SearchAllRequest]()), wrapHandler("search_all", client, bili.HandleSearchAll))
	s.AddTool(mcp.NewTool("search_by_type", mcp.WithDescription("Search by specific type"), mcp.WithInputSchema[bili.SearchByTypeRequest]()), wrapHandler("search_by_type", client, bili.HandleSearchByType))
	s.AddTool(mcp.NewTool("get_search_hot", mcp.WithDescription("Get current hot search keywords")), wrapHandler("get_search_hot", client, bili.HandleGetSearchHot))
	s.AddTool(mcp.NewTool("get_search_square", mcp.WithDescription("Get search square content"), mcp.WithInputSchema[bili.SearchSquareRequest]()), wrapHandler("get_search_square", client, bili.HandleGetSearchSquare))
	s.AddTool(mcp.NewTool("get_hotwords", mcp.WithDescription("Get trending hotwords")), wrapHandler("get_hotwords", client, bili.HandleGetHotwords))
	s.AddTool(mcp.NewTool("get_popular", mcp.WithDescription("Get popular content"), mcp.WithInputSchema[bili.PopularRequest]()), wrapHandler("get_popular", client, bili.HandleGetPopular))
	s.AddTool(mcp.NewTool("get_trending_rankings", mcp.WithDescription("Get trending rankings")), wrapHandler("get_trending_rankings", client, bili.HandleGetTrendingRankings))
	s.AddTool(mcp.NewTool("search_suggestions", mcp.WithDescription("Get search suggestions"), mcp.WithInputSchema[bili.SearchSuggestionsRequest]()), wrapHandler("search_suggestions", client, bili.HandleSearchSuggestions))

	// Social Tools
	s.AddTool(mcp.NewTool("get_column_info", mcp.WithDescription("Get column information"), mcp.WithInputSchema[bili.ColumnInfoRequest]()), wrapHandler("get_column_info", client, bili.HandleGetColumnInfo))
	s.AddTool(mcp.NewTool("get_column_list", mcp.WithDescription("Get column list"), mcp.WithInputSchema[bili.ColumnListRequest]()), wrapHandler("get_column_list", client, bili.HandleGetColumnList))
	s.AddTool(mcp.NewTool("get_topic_details", mcp.WithDescription("Get topic details"), mcp.WithInputSchema[bili.TopicDetailsRequest]()), wrapHandler("get_topic_details", client, bili.HandleGetTopicDetails))
	s.AddTool(mcp.NewTool("get_video_comments", mcp.WithDescription("Get comments for a video"), mcp.WithInputSchema[bili.VideoCommentsRequest]()), wrapHandler("get_video_comments", client, bili.HandleGetVideoComments))

	// Live Tools
	s.AddTool(mcp.NewTool("get_live_room_info", mcp.WithDescription("Get live room information"), mcp.WithInputSchema[bili.RoomInfoRequest]()), wrapHandler("get_live_room_info", client, bili.HandleGetLiveRoomInfo))
	s.AddTool(mcp.NewTool("get_live_play_url", mcp.WithDescription("Get live play URL"), mcp.WithInputSchema[bili.PlayURLRequest]()), wrapHandler("get_live_play_url", client, bili.HandleGetLivePlayURL))
	s.AddTool(mcp.NewTool("get_anchor_info", mcp.WithDescription("Get anchor information"), mcp.WithInputSchema[bili.AnchorInfoRequest]()), wrapHandler("get_anchor_info", client, bili.HandleGetAnchorInfo))
	s.AddTool(mcp.NewTool("get_live_recommend", mcp.WithDescription("Get live recommendations"), mcp.WithInputSchema[bili.LiveRecommendRequest]()), wrapHandler("get_live_recommend", client, bili.HandleGetLiveRecommend))
	s.AddTool(mcp.NewTool("get_room_status_batch", mcp.WithDescription("Get status of multiple live rooms"), mcp.WithInputSchema[bili.RoomStatusBatchRequest]()), wrapHandler("get_room_status_batch", client, bili.HandleGetRoomStatusBatch))

	// Ranking Tools
	s.AddTool(mcp.NewTool("get_region_ranking", mcp.WithDescription("Get regional rankings"), mcp.WithInputSchema[bili.RegionRankingRequest]()), wrapHandler("get_region_ranking", client, bili.HandleGetRegionRanking))
	s.AddTool(mcp.NewTool("get_weekly_must_watch", mcp.WithDescription("Get weekly must-watch list")), wrapHandler("get_weekly_must_watch", client, bili.HandleGetWeeklyMustWatch))
	s.AddTool(mcp.NewTool("get_weekly_must_watch_detail", mcp.WithDescription("Get weekly must-watch detail"), mcp.WithInputSchema[bili.WeeklyMustWatchDetailRequest]()), wrapHandler("get_weekly_must_watch_detail", client, bili.HandleGetWeeklyMustWatchDetail))
	s.AddTool(mcp.NewTool("get_popular_precious", mcp.WithDescription("Get popular precious videos")), wrapHandler("get_popular_precious", client, bili.HandleGetPopularPrecious))
}
