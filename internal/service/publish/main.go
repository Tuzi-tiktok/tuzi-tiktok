package main

import (
	"log"
	publish "tuzi-tiktok/kitex/kitex_gen/publish/publishservice"
	"tuzi-tiktok/utils"
)

func main() {
	svr := publish.NewServer(new(PublishServiceImpl), utils.NewServerOptions(utils.Publish())...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
