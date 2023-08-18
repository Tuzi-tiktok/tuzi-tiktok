package err

import (
	"context"
	"errors"
	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/cloudwego/hertz/pkg/app"
	"tuzi-tiktok/gateway/biz/err/global"
	"tuzi-tiktok/logger"
)

const (
	code = "status_code"
	msg  = "status_msg"
)
const maxTraceId = 1000

// ErrorHandlerMiddleware Handle Global Error
func ErrorHandlerMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		tid := fastrand.Int31n(maxTraceId)
		logger.Debugf("T: %d Access %v", tid, ctx.FullPath())
		ctx.Next(c)
		err := ctx.Errors.Last()
		if err != nil && err.Err != nil {
			e := err.Err
			var exception *global.UniformException
			if !errors.As(e, &exception) {
				return
			}
			var ret = make(map[string]any, 2)
			switch exception.InternalStatusCode {
			case global.InvalidTokenOrUnauthorizedCode, global.TokenNotFoundCode:
				logger.Warnf("T: %d Required Token is Not Found or InvalidToken Handler :%v; msg :%v;", tid, exception.HandlerName, exception.StatusMessage)
				ret[msg] = "Authentication failure Check Your Token"
			default:
				logger.Warnf("T: %d Service Face %v Catch Internal Error is %v", tid, exception.HandlerName, exception.StatusMessage)
				ret[msg] = "Service Unavailable Check Server Log"
			}
			ret[code] = exception.InternalStatusCode
			ctx.JSON(exception.HttpStatusCode, ret)
		}
	}
}
