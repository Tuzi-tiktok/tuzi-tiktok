package main

import (
	"log"
	auth "tuzi-tiktok/kitex/kitex_gen/auth/authinfoservice"
)

func main() {

	svr := auth.NewServer(new(AuthInfoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
