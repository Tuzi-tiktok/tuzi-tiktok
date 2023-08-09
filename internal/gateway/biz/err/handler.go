package err

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"tuzi-tiktok/logger"
)

// ErrorHandlerMiddleware Handle Global Error
func ErrorHandlerMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		err := ctx.Errors.Last()
		if err != nil && err.Err != nil {
			logger.Warnf("Catch Error is %v", err.Err)
		}
		ctx.Next(c)
	}
}
