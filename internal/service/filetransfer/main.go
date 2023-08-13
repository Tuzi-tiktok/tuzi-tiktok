package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hu "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/utils"
)

// Tips: https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/network-lib/
// main
func register() uint64 {
	port := uint64(utils.RandomAvailablePort())
	err := utils.DefaultServerRegister(utils.Transfer(), port)
	if err != nil {
		panic(err)
	}
	return port
}

const MaxBodySize = 1024 * 1024 * 200

func main() {
	port := register()
	h := server.Default(
		server.WithTransport(standard.NewTransporter),
		server.WithMaxRequestBodySize(MaxBodySize),
		server.WithHostPorts(fmt.Sprintf(":%v", port)),
	)

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, hu.H{"message": "pong"})
	})

	h.PUT("/", Transfer)

	h.Spin()
	logger.Warn("Service Shutdown With Error")

}
