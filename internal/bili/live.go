package bili

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// --- Request/Response Structs ---

type RoomInfoRequest struct {
	RoomID string `json:"room_id"`
}

type RoomInfoResponse struct {
	Room struct {
		RoomID   string `json:"room_id"`
		Title    string `json:"title"`
		Anchor   struct {
			Name string `json:"name"`
			UID  string `json:"uid"`
		} `json:"anchor"`
		Status   int    `json:"status"` // 1: Live, 0: Offline
		LiveTime string `json:"live_time"`
		Area     struct {
			Name string `json:"name"`
		} `json:"area"`
	} `json:"room"`
}

type PlayURLRequest struct {
	RoomID string `json:"room_id"`
}

type PlayURLResponse struct {
	Data struct {
		PlayURL string `json:"playurl"`
		Quality string `json:"quality"`
	} `json:"data"`
}

type AnchorInfoRequest struct {
	UID string `json:"uid"`
}

type AnchorInfoResponse struct {
	Data struct {
		User struct {
			Name     string `json:"name"`
			Level    int    `json:"level"`
			Signature string `json:"signature"`
			Followers int    `json:"followers"`
		} `json:"user"`
	} `json:"data"`
}

type LiveRecommendRequest struct {
	Page int `json:"page"`
}

type LiveRecommendResponse struct {
	Data struct {
		Rooms []struct {
			RoomID string `json:"room_id"`
			Title string `json:"title"`
			Anchor struct {
				Name string `json:"name"`
			} `json:"anchor"`
			Cover string `json:"cover"`
		} `json:"rooms"`
	} `json:"data"`
}

type RoomStatusBatchRequest struct {
	UIDs []string `json:"uids"`
}

type RoomStatusBatchResponse struct {
	Data struct {
		StatusList []struct {
			UID    string `json:"uid"`
			Status int    `json:"status"`
		} `json:"status_list"`
	} `json:"data"`
}

// --- Core Logic ---

func (c *BiliClient) GetLiveRoomInfo(req RoomInfoRequest) (*RoomInfoResponse, error) {
	params := map[string]string{"room_id": req.RoomID}
	data, err := c.doRequest("GET", "/room/v1/Room/get_info", params, true)
	if err != nil {
		return nil, err
	}

	var resp RoomInfoResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode room info: %w", err)
	}
	return &resp, nil
}

func (c *BiliClient) GetLivePlayURL(req PlayURLRequest) (*PlayURLResponse, error) {
	params := map[string]string{"room_id": req.RoomID}
	data, err := c.doRequest("GET", "/xlive/web-room/v2/index/getRoomPlayInfo", params, false)
	if err != nil {
		return nil, err
	}

	var resp PlayURLResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode play url: %w", err)
	}
	return &resp, nil
}

func (c *BiliClient) GetAnchorInfo(req AnchorInfoRequest) (*AnchorInfoResponse, error) {
	params := map[string]string{"uid": req.UID}
	data, err := c.doRequest("GET", "/live_user/v1/Master/info", params, true)
	if err != nil {
		return nil, err
	}

	var resp AnchorInfoResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode anchor info: %w", err)
	}
	return &resp, nil
}

func (c *BiliClient) GetLiveRecommend(req LiveRecommendRequest) (*LiveRecommendResponse, error) {
	params := map[string]string{"page": strconv.Itoa(req.Page)}
	data, err := c.doRequest("GET", "/xlive/web-interface/v1/webMain/getMoreRecList", params, false)
	if err != nil {
		return nil, err
	}

	var resp LiveRecommendResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode recommend list: %w", err)
	}
	return &resp, nil
}

func (c *BiliClient) GetRoomStatusBatch(req RoomStatusBatchRequest) (*RoomStatusBatchResponse, error) {
	params := map[string]string{"uids": strings.Join(req.UIDs, ",")}
	data, err := c.doRequest("GET", "/room/v1/Room/get_status_info_by_uids", params, true)
	if err != nil {
		return nil, err
	}

	var resp RoomStatusBatchResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, fmt.Errorf("failed to decode batch status: %w", err)
	}
	return &resp, nil
}
