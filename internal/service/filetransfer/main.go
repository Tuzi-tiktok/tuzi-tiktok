package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/network/standard"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Tips: https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/network-lib/
// main
func main() {
	h := server.New(server.WithTransport(standard.NewTransporter))
	h.Use(recovery.Recovery())

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	h.POST("/ping", Handle)
	h.Spin()
}
