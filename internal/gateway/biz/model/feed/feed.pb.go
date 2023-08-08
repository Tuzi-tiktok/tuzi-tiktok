// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.24.0--rc2
// source: feed.proto

package feed

import "tuzi-tiktok/gateway/biz/model/auth"

type FeedRequest struct {
	LatestTime *int64  `json:"latest_time,omitempty" form:"latest_time" query:"latest_time"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `json:"token,omitempty" form:"token" query:"token"`                   // 可选参数，登录用户设置
}

type FeedResponse struct {
	StatusCode int32    `json:"status_code,omitempty" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `json:"status_msg,omitempty" form:"status_msg" query:"status_msg"`    // 返回状态描述
	VideoList  []*Video `json:"video_list,omitempty" form:"video_list" query:"video_list"`    // 视频列表
	NextTime   *int64   `json:"next_time,omitempty" form:"next_time" query:"next_time"`       // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

type Video struct {
	Id            int64      `json:"id,omitempty" form:"id" query:"id"`                                     // 视频唯一标识
	Author        *auth.User `json:"author,omitempty" form:"author" query:"author"`                         // 视频作者信息
	PlayUrl       string     `json:"play_url,omitempty" form:"play_url" query:"play_url"`                   // 视频播放地址
	CoverUrl      string     `json:"cover_url,omitempty" form:"cover_url" query:"cover_url"`                // 视频封面地址
	FavoriteCount int64      `json:"favorite_count,omitempty" form:"favorite_count" query:"favorite_count"` // 视频的点赞总数
	CommentCount  int64      `json:"comment_count,omitempty" form:"comment_count" query:"comment_count"`    // 视频的评论总数
	IsFavorite    bool       `json:"is_favorite,omitempty" form:"is_favorite" query:"is_favorite"`          // true-已点赞，false-未点赞
	Title         string     `json:"title,omitempty" form:"title" query:"title"`                            // 视频标题
}
