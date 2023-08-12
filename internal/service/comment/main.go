package main

import (
	"log"
	comment "tuzi-tiktok/kitex/kitex_gen/comment/commentservice"
	"tuzi-tiktok/utils"
)

func main() {
	svr := comment.NewServer(new(CommentServiceImpl), utils.NewServerOptions(utils.Comment())...)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
