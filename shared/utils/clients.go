//go:build wireinject
// +build wireinject

package utils

import (
	"github.com/google/wire"
	auth "tuzi-tiktok/kitex/kitex_gen/auth/authinfoservice"
	comment "tuzi-tiktok/kitex/kitex_gen/comment/commentservice"
	favorite "tuzi-tiktok/kitex/kitex_gen/favorite/favoriteservice"
	feed "tuzi-tiktok/kitex/kitex_gen/feed/feedservice"
	message "tuzi-tiktok/kitex/kitex_gen/message/messageservice"
	publish "tuzi-tiktok/kitex/kitex_gen/publish/publishservice"
	relation "tuzi-tiktok/kitex/kitex_gen/relation/relationservice"
)

func NewAuth() (auth.Client, error) {
	panic(wire.Build(auth.NewClient, Auth, NewClientOptions, NoopExtra))
}
func NewComment() (comment.Client, error) {
	panic(wire.Build(comment.NewClient, Comment, NewClientOptions, NoopExtra))
}
func NewFavorite() (favorite.Client, error) {
	panic(wire.Build(favorite.NewClient, Favorite, NewClientOptions, NoopExtra))
}
func NewFeed() (feed.Client, error) {
	panic(wire.Build(feed.NewClient, Feed, NewClientOptions, NoopExtra))
}
func NewMessage() (message.Client, error) {
	panic(wire.Build(message.NewClient, Message, NewClientOptions, NoopExtra))
}
func NewPublish() (publish.Client, error) {
	panic(wire.Build(publish.NewClient, Publish, NewClientOptions, NoopExtra))
}

func NewRelation() (relation.Client, error) {
	panic(wire.Build(relation.NewClient, Relation, NewClientOptions, NoopExtra))
}
