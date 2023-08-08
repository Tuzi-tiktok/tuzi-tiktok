package main

import (
	"log"
	auth "tuzi-tiktok/kitex/kitex_gen/auth/authinfoservice"
	"tuzi-tiktok/utils"
)

func main() {

	svr := auth.NewServer(new(AuthInfoServiceImpl), utils.NewServerOptions(utils.Auth())...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
