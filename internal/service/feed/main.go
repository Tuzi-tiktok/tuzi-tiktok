package main

import (
	feed "tuzi-tiktok/kitex/kitex_gen/feed/feedservice"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/utils"
)

func main() {
	svr := feed.NewServer(new(FeedServiceImpl), utils.NewServerOptions(utils.Feed())...)

	err := svr.Run()

	if err != nil {
		logger.Warn("Service Shutdown With Error: %v", err.Error())
	}
}
