package main

import (
	comment "tuzi-tiktok/kitex/kitex_gen/comment/commentservice"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/utils"
)

func main() {
	svr := comment.NewServer(new(CommentServiceImpl), utils.NewServerOptions(utils.Comment())...)

	err := svr.Run()

	if err != nil {
		logger.Warn("Service Shutdown With Error: %v", err.Error())
	}
}
