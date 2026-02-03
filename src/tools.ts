import { z } from 'zod';
import * as api from './api.js';

export const TOOLS = {
  search: {
    name: 'search',
    description: 'Search for videos, users, etc. on Bilibili',
    inputSchema: z.object({
      keyword: z.string().describe('The search keyword'),
    }),
    handler: async (args: any) => {
      const res = await api.search(args.keyword);
      return {
        content: [{ type: 'text', text: JSON.stringify(res, null, 2) }],
      };
    },
  },
  get_video_info: {
    name: 'get_video_info',
    description: 'Get video details (title, description, stats)',
    inputSchema: z.object({
      bvid: z.string().describe('The video BVID (e.g., BV1xx411c7mD)'),
    }),
    handler: async (args: any) => {
      const res = await api.getVideoInfo(args.bvid);
      return {
        content: [{ type: 'text', text: JSON.stringify(res, null, 2) }],
      };
    },
  },
  get_video_player_info: {
    name: 'get_video_player_info',
    description: 'Get video player info, including subtitles/CC. Requires CID from get_video_info.',
    inputSchema: z.object({
      bvid: z.string().describe('The video BVID'),
      cid: z.number().describe('The video CID'),
    }),
    handler: async (args: any) => {
      const res = await api.getVideoPlayerInfo(args.bvid, args.cid);
      return {
        content: [{ type: 'text', text: JSON.stringify(res, null, 2) }],
      };
    },
  },
  get_user_info: {
    name: 'get_user_info',
    description: 'Get user profile information',
    inputSchema: z.object({
      mid: z.number().describe('The user ID (mid)'),
    }),
    handler: async (args: any) => {
      const res = await api.getUserInfo(args.mid);
      return {
        content: [{ type: 'text', text: JSON.stringify(res, null, 2) }],
      };
    },
  },
  get_user_dynamics: {
    name: 'get_user_dynamics',
    description: 'Get recent activities/dynamics of a user',
    inputSchema: z.object({
      host_mid: z.number().describe('The user ID (mid)'),
    }),
    handler: async (args: any) => {
      const res = await api.getUserDynamics(args.host_mid);
      return {
        content: [{ type: 'text', text: JSON.stringify(res, null, 2) }],
      };
    },
  },
  get_comments: {
    name: 'get_comments',
    description: 'Get comments for a video',
    inputSchema: z.object({
      oid: z.number().describe('The target ID (usually AID for videos)'),
      type: z.number().optional().default(1).describe('Type of target (1 for video)'),
    }),
    handler: async (args: any) => {
      const res = await api.getComments(args.oid, args.type);
      return {
        content: [{ type: 'text', text: JSON.stringify(res, null, 2) }],
      };
    },
  },
  get_popular_videos: {
    name: 'get_popular_videos',
    description: 'Get popular/trending videos',
    inputSchema: z.object({
      pn: z.number().optional().default(1).describe('Page number'),
      ps: z.number().optional().default(20).describe('Page size'),
    }),
    handler: async (args: any) => {
      const res = await api.getPopularVideos(args.pn, args.ps);
      return {
        content: [{ type: 'text', text: JSON.stringify(res, null, 2) }],
      };
    },
  },
};
