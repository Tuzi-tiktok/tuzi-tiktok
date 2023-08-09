package comment

import (
	"tuzi-tiktok/gateway/biz/model/auth"
)

type CommentRequest struct {
	Token       string  `json:"token" form:"token" query:"token"`                      // 用户鉴权token
	VideoId     int64   `json:"video_id" form:"video_id" query:"video_id"`             // 视频id
	ActionType  int32   `json:"action_type" form:"action_type" query:"action_type"`    // 1-发布评论，2-删除评论
	CommentText *string `json:"comment_text" form:"comment_text" query:"comment_text"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   *int64  `json:"comment_id" form:"comment_id" query:"comment_id"`       // 要删除的评论id，在action_type=2的时候使用
}

type CommentResponse struct {
	StatusCode int32    `json:"status_code" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `json:"status_msg" form:"status_msg" query:"status_msg"`    // 返回状态描述
	Comment    *Comment `json:"comment" form:"comment" query:"comment"`             // 评论成功返回评论内容，不需要重新拉取整个列表
}

type CommentListRequest struct {
	Token   string `json:"token" form:"token" query:"token"`          // 用户鉴权token
	VideoId int64  `json:"video_id" form:"video_id" query:"video_id"` // 视频id
}

type CommentListResponse struct {
	StatusCode  int32      `json:"status_code" form:"status_code" query:"status_code"`    // 状态码，0-成功，其他值-失败
	StatusMsg   *string    `json:"status_msg" form:"status_msg" query:"status_msg"`       // 返回状态描述
	CommentList []*Comment `json:"comment_list" form:"comment_list" query:"comment_list"` // 评论列表
}

type Comment struct {
	Id         int64      `json:"id" form:"id" query:"id"`                            // 视频评论id
	User       *auth.User `json:"user" form:"user" query:"user"`                      // 评论用户信息
	Content    string     `json:"content" form:"content" query:"content"`             // 评论内容
	CreateDate string     `json:"create_date" form:"create_date" query:"create_date"` // 评论发布日期，格式 mm-dd
}
