package main

import (
	"log"
	feed "tuzi-tiktok/kitex/kitex_gen/feed/feedservice"
	"tuzi-tiktok/utils"
)

func main() {
	svr := feed.NewServer(new(FeedServiceImpl), utils.NewServerOptions(utils.Feed())...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
