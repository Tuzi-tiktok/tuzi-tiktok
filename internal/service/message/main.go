package main

import (
	message "tuzi-tiktok/kitex/kitex_gen/message/messageservice"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/utils"
)

func main() {
	svr := message.NewServer(new(MessageServiceImpl), utils.NewServerOptions(utils.Message())...)

	err := svr.Run()

	if err != nil {
		logger.Warn("Service Shutdown With Error: %v", err.Error())
	}
}
