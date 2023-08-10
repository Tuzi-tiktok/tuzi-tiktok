package err

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"tuzi-tiktok/gateway/biz/err/global"
	"tuzi-tiktok/logger"
)

// ErrorHandlerMiddleware Handle Global Error
func ErrorHandlerMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		logger.Info("Access")
		ctx.Next(c)
		err := ctx.Errors.Last()
		if err != nil && err.Err != nil {
			e := err.Err
			var exception *global.UniformException
			if errors.As(e, &exception) {
				logger.Warnf("Catch Internal Error is %v", exception.Error())
			}
		}
	}
}
