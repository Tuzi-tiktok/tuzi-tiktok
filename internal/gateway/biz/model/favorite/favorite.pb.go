// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.24.0--rc2
// source: favorite.proto

package favorite

import (
	"tuzi-tiktok/gateway/biz/model/feed"
)

type FavoriteRequest struct {
	Token      string `json:"token" form:"token" query:"token"`                   // 用户鉴权token
	VideoId    int64  `json:"video_id" form:"video_id" query:"video_id"`          // 视频id
	ActionType int32  `json:"action_type" form:"action_type" query:"action_type"` // 1-点赞，2-取消点赞
}
type FavoriteResponse struct {
	StatusCode int32   `json:"status_code" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg" form:"status_msg" query:"status_msg"`    // 返回状态描述
}
type FavoriteListRequest struct {
	UserId int64  `json:"user_id" form:"user_id" query:"user_id"` // 用户id
	Token  string `json:"token" form:"token" query:"token"`       // 用户鉴权token
}

type FavoriteListResponse struct {
	StatusCode int32         `json:"status_code" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string       `json:"status_msg" form:"status_msg" query:"status_msg"`    // 返回状态描述
	VideoList  []*feed.Video `json:"video_list" form:"video_list" query:"video_list"`    // 用户点赞视频列表
}
