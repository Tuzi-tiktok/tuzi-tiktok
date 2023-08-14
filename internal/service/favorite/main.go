package main

import (
	favorite "tuzi-tiktok/kitex/kitex_gen/favorite/favoriteservice"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/utils"
)

func main() {
	svr := favorite.NewServer(new(FavoriteServiceImpl), utils.NewServerOptions(utils.Favorite())...)

	err := svr.Run()
	if err != nil {
		logger.Warn("Service Shutdown With Error: %v", err.Error())
	}
}
