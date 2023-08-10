// Code generated by hertz generator.

package message

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"tuzi-tiktok/gateway/biz/err/global"
	"tuzi-tiktok/gateway/biz/model/message"
	"tuzi-tiktok/gateway/biz/service"
	kmessage "tuzi-tiktok/kitex/kitex_gen/message"
	"tuzi-tiktok/utils/mapstruct"
)

// GetMessageList .
// @router /douyin/message/chat/ [GET]
func GetMessageList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req message.MessageChatRequest
	var handler = "GetMessageList"
	err = c.Bind(&req)
	if err != nil {
		_ = c.Error(global.RequestParameterBindError.WithHandler(handler).WithWarn(err))
		return
	}

	R, err := service.ServiceSet.Message.GetMessageList(ctx, &kmessage.MessageChatRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserId,
		PreMsgTime: req.PreMsgTime,
	})
	if err != nil {
		_ = c.Error(global.RPCClientCallError.WithHandler(handler).WithError(err))
		return
	}

	resp := mapstruct.ToMessageChatResponse(R)
	c.JSON(consts.StatusOK, resp)
}

// MessageAction .
// @router /douyin/message/action/ [POST]
func MessageAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req message.MessageActionRequest
	var handler = "MessageAction"
	err = c.Bind(&req)
	if err != nil {
		_ = c.Error(global.RequestParameterBindError.WithHandler(handler).WithWarn(err))
		return
	}

	R, err := service.ServiceSet.Message.MessageAction(ctx, &kmessage.MessageActionRequest{
		Token:      req.Token,
		ToUserId:   req.ToUserId,
		ActionType: req.ActionType,
		Content:    req.Content,
	})
	if err != nil {
		_ = c.Error(global.RPCClientCallError.WithHandler(handler).WithError(err))
		return
	}
	resp := mapstruct.ToMessageActionResponse(R)

	c.JSON(consts.StatusOK, resp)
}
