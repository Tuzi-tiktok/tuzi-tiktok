package service

import (
	auth "tuzi-tiktok/kitex/kitex_gen/auth/authinfoservice"
	comment "tuzi-tiktok/kitex/kitex_gen/comment/commentservice"
	favorite "tuzi-tiktok/kitex/kitex_gen/favorite/favoriteservice"
	feed "tuzi-tiktok/kitex/kitex_gen/feed/feedservice"
	message "tuzi-tiktok/kitex/kitex_gen/message/messageservice"
	publish "tuzi-tiktok/kitex/kitex_gen/publish/publishservice"
	relation "tuzi-tiktok/kitex/kitex_gen/relation/relationservice"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/service/filetransfer/client"
	"tuzi-tiktok/utils"
)

var ServiceSet ClientsSet

func init() {
	initClientsSet()
}

type ClientsSet struct {
	Auth     auth.Client
	Comment  comment.Client
	Favorite favorite.Client
	Feed     feed.Client
	Message  message.Client
	Publish  publish.Client
	Transfer client.Transfer
	Relation relation.Client
}

func initClientsSet() {
	a, err := utils.NewAuth()
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	b, err := utils.NewComment()
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	c, err := utils.NewFavorite()
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	d, err := utils.NewFeed()
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	e, err := utils.NewMessage()
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	f, err := utils.NewPublish()
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	g, err := utils.NewRelation()
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	transfer := client.NewTransfer()
	ServiceSet = ClientsSet{
		Auth:     a,
		Comment:  b,
		Favorite: c,
		Feed:     d,
		Message:  e,
		Publish:  f,
		Relation: g,
		Transfer: transfer,
	}
}
