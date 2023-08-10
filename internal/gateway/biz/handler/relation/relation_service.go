// Code generated by hertz generator.

package relation

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"tuzi-tiktok/gateway/biz/err/global"
	relation "tuzi-tiktok/gateway/biz/model/relation"
	"tuzi-tiktok/gateway/biz/service"
	krelation "tuzi-tiktok/kitex/kitex_gen/relation"
	"tuzi-tiktok/utils/mapstruct"
)

// FollowAction .
// @router /douyin/relation/action/ [POST]
func FollowAction(ctx context.Context, c *app.RequestContext) {
	var req relation.RelationRequest
	err := c.Bind(&req)
	var handler = "FollowAction"
	if err != nil {
		_ = c.Error(global.RequestParameterBindError.WithWarn(err).WithHandler(handler))
		return
	}
	R, err := service.ServiceSet.Relation.FollowAction(ctx, &krelation.RelationRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
	})
	if err != nil {
		_ = c.Error(global.RPCClientCallError.WithError(err).WithHandler(handler))
		return
	}
	resp := mapstruct.ToRelationResponse(R)

	c.JSON(consts.StatusOK, resp)
}

// GetFollowList .
// @router /douyin/relation/follow/list/ [GET]
func GetFollowList(ctx context.Context, c *app.RequestContext) {
	var req relation.RelationFollowListRequest
	var handler = "GetFollowList"
	err := c.Bind(&req)
	if err != nil {
		_ = c.Error(global.RequestParameterBindError.WithWarn(err).WithHandler(handler))
		return
	}
	R, err := service.ServiceSet.Relation.GetFollowList(ctx, &krelation.RelationFollowListRequest{
		Token:  req.Token,
		UserId: req.UserId,
	})
	if err != nil {
		_ = c.Error(global.RPCClientCallError.WithError(err).WithHandler(handler))
		return
	}
	resp := mapstruct.ToRelationFollowListResponse(R)

	c.JSON(consts.StatusOK, resp)
}

// GetFollowerList .
// @router /douyin/relation/follower/list/ [GET]
func GetFollowerList(ctx context.Context, c *app.RequestContext) {
	var req relation.RelationFollowerListRequest
	var handler = "GetFollowerList"
	err := c.Bind(&req)
	if err != nil {
		_ = c.Error(global.RequestParameterBindError.WithWarn(err).WithHandler(handler))
		return
	}
	R, err := service.ServiceSet.Relation.GetFollowerList(ctx, &krelation.RelationFollowerListRequest{
		Token:  req.Token,
		UserId: req.UserId,
	})
	if err != nil {
		_ = c.Error(global.RPCClientCallError.WithError(err).WithHandler(handler))
		return
	}
	resp := mapstruct.ToRelationFollowerListResponse(R)

	c.JSON(consts.StatusOK, resp)
}

// GetFriendList .
// @router /douyin/relation/friend/list/ [GET]
func GetFriendList(ctx context.Context, c *app.RequestContext) {
	var req relation.RelationFriendListRequest
	var handler = "GetFriendList"
	err := c.Bind(&req)
	if err != nil {
		_ = c.Error(global.RequestParameterBindError.WithWarn(err).WithHandler(handler))
		return
	}
	R, err := service.ServiceSet.Relation.GetFriendList(ctx, &krelation.RelationFriendListRequest{
		Token:  req.Token,
		UserId: req.UserId,
	})
	if err != nil {
		_ = c.Error(global.RPCClientCallError.WithError(err).WithHandler(handler))
		return
	}
	resp := mapstruct.ToRelationFriendListResponse(R)

	c.JSON(consts.StatusOK, resp)
}
