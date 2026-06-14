package bili

import (
	"encoding/json"
	"fmt"
)

type UserProfileRequest struct {
	MID string `json:"mid"`
}

type UserCardRequest struct {
	MID string `json:"mid"`
}

type UserRelationStatRequest struct {
	MID string `json:"mid"`
}

type UserPublicDynamicsRequest struct {
	MID string `json:"mid"`
	Page int    `json:"page"`
}

type DynamicDetailRequest struct {
	DynamicID string `json:"dynamic_id"`
}

type ColumnInfoRequest struct {
	ColumnID string `json:"column_id"`
}

type ColumnListRequest struct {
	MID string `json:"mid"`
	Page int    `json:"page"`
}

type TopicDetailsRequest struct {
	TopicID string `json:"topic_id"`
}

type VideoCommentsRequest struct {
	OID  string `json:"oid"`
	Type string `json:"type"`
	Page int    `json:"page"`
}

func (c *BiliClient) GetUserProfile(req UserProfileRequest) (json.RawMessage, error) {
	params := map[string]string{"mid": req.MID}
	return c.doRequest("GET", "/x/web-interface/card", params, true)
}

func (c *BiliClient) GetUserCard(req UserCardRequest) (json.RawMessage, error) {
	params := map[string]string{"mid": req.MID}
	return c.doRequest("GET", "/x/web-interface/card", params, true)
}

func (c *BiliClient) GetUserRelationStat(req UserRelationStatRequest) (json.RawMessage, error) {
	params := map[string]string{"mid": req.MID}
	return c.doRequest("GET", "/x/web-interface/relation/stat", params, true)
}

func (c *BiliClient) GetUserPublicDynamics(req UserPublicDynamicsRequest) (json.RawMessage, error) {
	params := map[string]string{
		"mid":  req.MID,
		"page": fmt.Sprintf("%d", req.Page),
	}
	return c.doRequest("GET", "/x/polymer/web-dynamic/v1/get_dynamic_list", params, true)
}

func (c *BiliClient) GetDynamicDetail(req DynamicDetailRequest) (json.RawMessage, error) {
	params := map[string]string{"dynamic_id": req.DynamicID}
	return c.doRequest("GET", "/x/polymer/web-dynamic/v1/dynamic_detail", params, true)
}

func (c *BiliClient) GetColumnInfo(req ColumnInfoRequest) (json.RawMessage, error) {
	params := map[string]string{"column_id": req.ColumnID}
	return c.doRequest("GET", "/x/column/api/column/detail", params, true)
}

func (c *BiliClient) GetColumnList(req ColumnListRequest) (json.RawMessage, error) {
	params := map[string]string{
		"mid":  req.MID,
		"page": fmt.Sprintf("%d", req.Page),
	}
	return c.doRequest("GET", "/x/column/api/column/list", params, true)
}

func (c *BiliClient) GetTopicDetails(req TopicDetailsRequest) (json.RawMessage, error) {
	params := map[string]string{"topic_id": req.TopicID}
	return c.doRequest("GET", "/x/topic/api/topic/detail", params, true)
}

func (c *BiliClient) GetVideoComments(req VideoCommentsRequest) (json.RawMessage, error) {
	params := map[string]string{
		"oid":  req.OID,
		"type": req.Type,
		"pn":   fmt.Sprintf("%d", req.Page),
	}
	return c.doRequest("GET", "/x/web-interface/view.async.cmt", params, true)
}
