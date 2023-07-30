// Code generated by hertz generator.

package publish

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	publish "tuzi-tiktok/gateway/biz/model/publish"
)

// PublishVideo .
// @router /douyin/publish/action/ [POST]
func PublishVideo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.PublishRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(publish.PublishResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetPublishList .
// @router /douyin/publish/list/ [GET]
func GetPublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req publish.PublishListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(publish.PublishListResponse)

	c.JSON(consts.StatusOK, resp)
}
