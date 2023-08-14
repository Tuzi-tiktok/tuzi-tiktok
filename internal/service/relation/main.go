package main

import (
	relation "tuzi-tiktok/kitex/kitex_gen/relation/relationservice"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/utils"
)

func main() {

	svr := relation.NewServer(new(RelationServiceImpl), utils.NewServerOptions(utils.Relation())...)

	err := svr.Run()
	if err != nil {
		logger.Warn("Service Shutdown With Error: %v", err.Error())
	}
}
