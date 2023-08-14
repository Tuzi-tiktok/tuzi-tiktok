package main

import (
	publish "tuzi-tiktok/kitex/kitex_gen/publish/publishservice"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/utils"
)

func main() {
	svr := publish.NewServer(new(PublishServiceImpl), utils.NewServerOptions(utils.Publish())...)

	err := svr.Run()

	if err != nil {
		logger.Warn("Service Shutdown With Error: %v", err.Error())
	}
}
