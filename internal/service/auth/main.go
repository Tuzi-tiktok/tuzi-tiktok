package main

import (
	auth "tuzi-tiktok/kitex/kitex_gen/auth/authinfoservice"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/utils"
)

func main() {

	svr := auth.NewServer(new(AuthInfoServiceImpl), utils.NewServerOptions(utils.Auth())...)

	err := svr.Run()

	if err != nil {
		logger.Warn("Service Shutdown With Error: %v", err.Error())
	}
}
