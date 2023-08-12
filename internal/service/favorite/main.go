package main

import (
	"log"
	favorite "tuzi-tiktok/kitex/kitex_gen/favorite/favoriteservice"
	"tuzi-tiktok/utils"
)

func main() {
	svr := favorite.NewServer(new(FavoriteServiceImpl), utils.NewServerOptions("favorite-api")...)

	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
