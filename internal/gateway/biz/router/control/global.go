package control

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"tuzi-tiktok/gateway/biz/err/global"
	"tuzi-tiktok/gateway/biz/service"
	"tuzi-tiktok/kitex/kitex_gen/auth"
)

type TokenEq struct {
	Token string `json:"token" form:"token" query:"token"`
}

const USERID = "USER_ID"

func Authentication() app.HandlerFunc {
	var handler = "AuthenticationGlobalFunc"
	return func(c context.Context, ctx *app.RequestContext) {
		var t TokenEq
		err := ctx.Bind(&t)
		if err != nil || t.Token == "" {
			_ = ctx.Error(global.TokenNotFound.WithHandler(handler))
			ctx.Abort()
			return
		}
		R, err := service.ServiceSet.Auth.TokenVerify(c, &auth.TokenVerifyRequest{
			Token: t.Token,
		})
		if err != nil {
			_ = ctx.Error(global.RPCClientCallError.WithHandler(handler).WithError(err))
			ctx.Abort()
		} else if R.StatusCode != 0 {
			_ = ctx.Error(global.InvalidTokenOrUnauthorized.WithHandler(handler))
			ctx.Abort()
		} else {
			ctx.Set(USERID, R.UserId)
		}
	}
}
