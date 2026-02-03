import axios from 'axios';
import { encWbi, getWbiKeys } from './utils/wbi.js';

const UA =
  'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36';

const client = axios.create({
  headers: {
    'User-Agent': UA,
    Referer: 'https://www.bilibili.com/',
  },
});

// 1. Search
export async function search(keyword: string) {
  const { img_key, sub_key } = await getWbiKeys();
  const params = {
    keyword,
    search_type: 'video', // default to video search
  };
  const query = await encWbi(params, img_key, sub_key);
  const res = await client.get(
    `https://api.bilibili.com/x/web-interface/wbi/search/all/v2?${query}`,
  );
  return res.data;
}

// 2. Video Info
export async function getVideoInfo(bvid: string) {
  const res = await client.get(`https://api.bilibili.com/x/web-interface/view?bvid=${bvid}`);
  return res.data;
}

// 2.1 Video Subtitles (Player Info)
export async function getVideoPlayerInfo(bvid: string, cid: number) {
  const res = await client.get(`https://api.bilibili.com/x/player/v2?bvid=${bvid}&cid=${cid}`);
  return res.data;
}

// 3. User Info
export async function getUserInfo(mid: number) {
  const { img_key, sub_key } = await getWbiKeys();
  const params = {
    mid,
  };
  const query = await encWbi(params, img_key, sub_key);
  const res = await client.get(`https://api.bilibili.com/x/space/wbi/acc/info?${query}`);
  return res.data;
}

// 4. Dynamic/Activity Feed
export async function getUserDynamics(host_mid: number) {
  const res = await client.get(
    `https://api.bilibili.com/x/polymer/web-dynamic/v1/feed/space?host_mid=${host_mid}`,
  );
  return res.data;
}

// 5. Comments
export async function getComments(oid: number, type: number = 1) {
  // type 1 is for video
  const res = await client.get(`https://api.bilibili.com/x/v2/reply?oid=${oid}&type=${type}`);
  return res.data;
}

// 6. Popular/Trending
export async function getPopularVideos(pn: number = 1, ps: number = 20) {
  const res = await client.get(
    `https://api.bilibili.com/x/web-interface/popular?pn=${pn}&ps=${ps}`,
  );
  return res.data;
}
