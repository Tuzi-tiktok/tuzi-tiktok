package main

import (
	"log"
	relation "tuzi-tiktok/kitex/kitex_gen/relation/relationservice"
	"tuzi-tiktok/utils"
)

func main() {

	svr := relation.NewServer(new(RelationServiceImpl), utils.NewServerOptions(utils.Relation())...)

	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
