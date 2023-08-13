package main

import (
	message "tuzi-tiktok/kitex/kitex_gen/message/messageservice"
	"tuzi-tiktok/logger"
)

func main() {
	svr := message.NewServer(new(MessageServiceImpl))

	err := svr.Run()

	if err != nil {
		logger.Warn("Service Shutdown With Error: %v", err.Error())
	}
}
