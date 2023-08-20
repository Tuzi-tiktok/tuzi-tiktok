// Code generated by hertz generator.

package auth

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"tuzi-tiktok/gateway/biz/err/access"
	"tuzi-tiktok/gateway/biz/err/global"
	"tuzi-tiktok/gateway/biz/model/auth"
	"tuzi-tiktok/gateway/biz/service"
	kauth "tuzi-tiktok/kitex/kitex_gen/auth"
	"tuzi-tiktok/kitex/kitex_gen/auth/authinfoservice"
	"tuzi-tiktok/utils/mapstruct"
)

var authClient authinfoservice.Client

func init() {
	authClient = service.ServiceSet.Auth
}

// Login .
// @router /douyin/user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var (
		req         auth.UserLoginRequest
		handlerName = "Login"
	)

	err := c.Bind(&req)
	if err != nil {
		_ = c.Error(global.RequestParameterBindError.
			WithHandler(handlerName).
			WithWarn(err))
		return
	}

	if req.Username == "" || req.Password == "" || len(req.Password) > 32 {
		_ = c.Error(global.ParameterValidationError.
			WithHandler(handlerName).
			WithWarn(errors.New("username or password not validated ")))
		return
	}

	access.DebugRecordRequest(c, req)

	R, err := authClient.Login(ctx, &kauth.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		_ = c.Error(global.RPCClientCallError.
			WithHandler(handlerName).
			WithError(err))
		return
	}
	resp := mapstruct.ToUserLoginResponse(R)
	c.JSON(consts.StatusOK, resp)
}

// Register .
// @router /douyin/user/register/ [POST]
func Register(ctx context.Context, c *app.RequestContext) {
	var (
		req         auth.UserRegisterRequest
		handlerName = "Register"
	)
	err := c.Bind(&req)

	if err != nil {
		_ = c.Error(global.RequestParameterBindError.WithWarn(err))
		return
	}
	if req.Username == "" || req.Password == "" || len(req.Password) > 32 {
		_ = c.Error(global.ParameterValidationError.WithHandler(handlerName).WithWarn(err))
		return
	}

	access.DebugRecordRequest(c, req)

	R, err := authClient.Register(ctx, &kauth.UserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		_ = c.Error(global.RPCClientCallError.WithHandler(handlerName).WithError(err))
		return
	}
	resp := mapstruct.ToUserRegisterResponse(R)
	c.JSON(consts.StatusOK, resp)
}

// GetUserInfo .
// @router /douyin/user/ [GET]
func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var (
		req     auth.UserInfoRequest
		handler = "GetUserInfo"
	)
	err := c.Bind(&req)
	if err != nil {
		_ = c.Error(global.RequestParameterBindError.WithHandler(handler).WithWarn(err))
		return
	}

	access.DebugRecordRequest(c, req)

	R, err := authClient.GetUserInfo(ctx, &kauth.UserInfoRequest{
		UserId: req.UserId,
		Token:  req.Token,
	})
	if err != nil {
		_ = c.Error(global.RPCClientCallError.WithHandler(handler).WithError(err))
		return
	}
	resp := mapstruct.ToUserInfoResponse(R)

	c.JSON(consts.StatusOK, resp)
}
