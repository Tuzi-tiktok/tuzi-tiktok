// Code generated by hertz generator.

package favorite

import (
	"context"
	"tuzi-tiktok/gateway/biz/err/access"
	"tuzi-tiktok/gateway/biz/err/global"
	"tuzi-tiktok/gateway/biz/service"
	kfavorite "tuzi-tiktok/kitex/kitex_gen/favorite"
	"tuzi-tiktok/utils/mapstruct"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	favorite "tuzi-tiktok/gateway/biz/model/favorite"
)

// FavorVideo .
// @router /douyin/favorite/action/ [POST]
func FavorVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.FavoriteRequest
	var handler = "FavorVideo"
	// TODO Bind Action Type Enum 0 1
	err = c.Bind(&req)
	if err != nil {
		_ = c.Error(global.RequestParameterBindError.WithHandler(handler).WithWarn(err))
		return
	}

	access.DebugRecordRequest(c, req)

	R, err := service.ServiceSet.Favorite.FavorVideo(ctx, &kfavorite.FavoriteRequest{
		Token:      req.Token,
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
	})
	if err != nil {
		_ = c.Error(global.RPCClientCallError.WithHandler(handler).WithError(err))
		return
	}

	resp := mapstruct.ToFavoriteResponse(R)
	c.JSON(consts.StatusOK, resp)
}

// GetFavoriteList .
// @router /douyin/favorite/list/ [GET]
func GetFavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req favorite.FavoriteListRequest
	var handler = "GetFavoriteList"
	err = c.Bind(&req)
	if err != nil {
		_ = c.Error(global.RequestParameterBindError.WithHandler(handler).WithWarn(err))
		return
	}

	access.DebugRecordRequest(c, req)

	R, err := service.ServiceSet.Favorite.GetFavoriteList(ctx, &kfavorite.FavoriteListRequest{
		UserId: req.UserId,
		Token:  req.Token,
	})

	if err != nil {
		_ = c.Error(global.RPCClientCallError.WithHandler(handler).WithError(err))
		return
	}

	resp := mapstruct.ToFavoriteListResponse(R)

	c.JSON(consts.StatusOK, resp)
}
