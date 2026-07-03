# bilibili-mcp

[![Go Version](https://img.shields.io/badge/Go->=1.25-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![MCP](https://img.shields.io/badge/Protocol-MCP-orange.svg)](https://modelcontextprotocol.io)

`bilibili-mcp` 是一个基于 **MCP (Model Context Protocol)** 协议的 Bilibili 服务端实现，采用 `streamable-http` 传输协议。它将 Bilibili 的丰富 API 封装为标准的 MCP Tools，使得 AI 模型能够直接调用这些工具来获取 Bilibili 的视频、用户、搜索、直播及排行榜等实时数据。

## ✨ 核心特性

本项目将 Bilibili 的能力划分为六大功能模块，为 AI 提供全方位的内容获取能力：

- **📹 视频工具**：深度获取视频元数据、统计信息、标签及弹幕数据。
- **👤 用户工具**：查询用户画像、名片、关系统计及公开动态。
- **🔍 搜索工具**：支持全站搜索、热词追踪、趋势排行榜及搜索建议。
- **💬 社交工具**：获取专栏内容、话题详情以及视频评论。
- **📺 直播工具**：实时查询直播间状态、主播信息及推荐直播。
- **🏆 排行榜工具**：追踪分区排名、每周必看及精品内容。

## 🚀 快速开始

### 环境要求
- **Go**: $\ge 1.25$
- **Docker** (可选): 用于容器化部署

### 本地编译运行
如果你想在本地快速启动服务，可以按照以下步骤操作：

```bash
# 克隆仓库
git clone https://github.com/your-username/bilibili-mcp.git
cd bilibili-mcp

# 编译项目
go build -o bilibili-mcp main.go

# 运行服务
./bilibili-mcp
```

### 使用 Docker 部署
通过 Docker Compose 实现一键快速部署：

```bash
docker-compose up -d
```

## 🤖 AI 交互示例

你可以尝试向 AI 提出以下问题，它将自动调用相应的工具来为你提供答案：

- **查询视频数据**：
  - 问：“这个视频 `BVxxxxxx` 的播放量是多少？” $\rightarrow$ 调用 `get_video_stat`
- **搜索热门内容**：
  - 问：“帮我搜一下 B 站上关于 'Go 语言' 的热门视频” $\rightarrow$ 调用 `search_all` 或 `get_popular`
- **关注 UP 主动态**：
  - 问：“这个 UP 主最近发了什么动态？” $\rightarrow$ 调用 `get_user_public_dynamics`

## ⚙️ 配置指南

服务通过环境变量进行配置：

| 变量名 | 描述 | 默认值 | 备注 |
| :--- | :--- | :--- | :--- |
| `MCP_PORT` | MCP 服务监听端口 | `8080` | 确保该端口在防火墙中已开放 |

## 🛠️ 工具详单 (Tools)

AI 模型可以通过以下工具与 Bilibili 交互：

### 1. 视频工具 (Video Tools)
| 工具名称 | 功能描述 |
| :--- | :--- |
| `get_video_info` | 获取视频的基本概览信息 |
| `get_video_detail` | 获取视频的详细内容信息 |
| `get_video_stat` | 获取视频的播放量、点赞、投币等统计数据 |
| `get_video_tags` | 获取视频关联的标签列表 |
| `get_video_desc` | 获取视频的详细简介描述 |
| `get_video_related` | 获取与当前视频相关的推荐视频 |
| `get_video_conclusion` | 获取视频的内容总结/结论 |
| `get_video_online_count` | 获取视频当前的实时在线观看人数 |
| `get_danmaku_snapshot` | 获取弹幕快照数据 |
| `get_danmaku_list` | 获取指定时间段的弹幕列表 |
| `get_danmaku_config` | 获取弹幕相关的配置参数 |
| `get_danmaku_buzzword` | 获取弹幕中的热门词汇/梗 |

### 2. 用户工具 (User Tools)
| 工具名称 | 功能描述 |
| :--- | :--- |
| `get_user_profile` | 获取用户的详细个人资料 |
| `get_user_card` | 获取用户的名片信息 |
| `get_user_relation_stat` | 获取用户的关注、粉丝等关系统计 |
| `get_user_public_dynamics` | 获取用户的公开动态列表 |
| `get_dynamic_detail` | 获取单条动态的详细内容 |

### 3. 搜索工具 (Search Tools)
| 工具名称 | 功能描述 |
| :--- | :--- |
| `search_all` | 执行全站综合搜索 |
| `search_by_type` | 根据指定类型（视频/用户/专栏等）进行搜索 |
| `get_search_hot` | 获取当前的搜索热词 |
| `get_search_square` | 获取搜索广场的推荐内容 |
| `get_hotwords` | 获取当前平台的热门关键词 |
| `get_popular` | 获取当前最热门的内容列表 |
| `get_trending_rankings` | 获取实时趋势排行榜 |
| `search_suggestions` | 获取搜索输入时的自动补全建议 |

### 4. 社交工具 (Social Tools)
| 工具名称 | 功能描述 |
| :--- | :--- |
| `get_column_info` | 获取专栏文章的详细信息 |
| `get_column_list` | 获取特定用户的专栏文章列表 |
| `get_topic_details` | 获取特定话题的详情及相关内容 |
| `get_video_comments` | 获取视频的评论区内容 |

### 5. 直播工具 (Live Tools)
| 工具名称 | 功能描述 |
| :--- | :--- |
| `get_live_room_info` | 获取直播间的实时状态和信息 |
| `get_live_play_url` | 获取直播间的播放地址 |
| `get_anchor_info` | 获取直播主播的详细资料 |
| `get_live_recommend` | 获取推荐的直播间列表 |
| `get_room_status_batch` | 批量获取多个直播间的状态 |

### 6. 排行榜工具 (Ranking Tools)
| 工具名称 | 功能描述 |
| :--- | :--- |
| `get_region_ranking` | 获取指定分区的排行榜 |
| `get_weekly_must_watch` | 获取本周必看视频列表 |
| `get_weekly_must_watch_detail` | 获取本周必看视频的详细信息 |
| `get_popular_precious` | 获取热门精品内容推荐 |

## 💻 技术栈

- **语言**: [Go](https://golang.org/)
- **协议**: [Model Context Protocol (MCP)](https://modelcontextprotocol.io)
- **框架**: `github.com/mark3labs/mcp-go`

## 📜 许可

本项目采用 [MIT License](LICENSE) 开源协议。
