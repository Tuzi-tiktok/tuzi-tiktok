package main

import (
	"log"
	feed "tuzi-tiktok/kitex/kitex_gen/feed/feedservice"
)

func main() {
	svr := feed.NewServer(new(FeedServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
