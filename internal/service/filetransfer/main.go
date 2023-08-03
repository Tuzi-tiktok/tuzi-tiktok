package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Tips: https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/network-lib/
// main
func main() {

	//rop := server.WithHostPorts(fmt.Sprintf(":%v", u.RandomAvailablePort()))
	h := server.Default(
		server.WithTransport(standard.NewTransporter),
	)

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})

	h.PUT("/", Transfer)
	h.Spin()
}
