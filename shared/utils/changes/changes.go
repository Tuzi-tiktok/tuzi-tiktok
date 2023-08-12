package changes

import (
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/kitex/kitex_gen/auth"
	"tuzi-tiktok/kitex/kitex_gen/feed"
	"tuzi-tiktok/kitex/kitex_gen/relation"
)

func UserRecord2userResp(user *model.User) *auth.User {

	userResp := new(auth.User)
	userResp.Id = user.ID
	userResp.Name = user.Username
	userResp.FollowCount = &user.FollowCount
	userResp.FollowerCount = &user.FollowerCount
	userResp.Avatar = user.Avatar
	userResp.Signature = user.Signature

	return userResp
}
func UserRecord2friendResp(user *model.User) *relation.FriendUser {

	userResp := new(relation.FriendUser)
	userResp.Id = user.ID
	userResp.Name = user.Username
	userResp.FollowCount = &user.FollowCount
	userResp.FollowerCount = &user.FollowerCount
	userResp.Avatar = user.Avatar
	userResp.Signature = user.Signature

	return userResp
}
func VideoRecord2videoResp(video *model.Video) *feed.Video {

	videoResp := new(feed.Video)
	videoResp.Id = video.ID
	videoResp.CommentCount = video.CommentCount
	videoResp.CoverUrl = video.CoverURL
	videoResp.FavoriteCount = video.FavoriteCount
	videoResp.PlayUrl = video.PlayURL
	videoResp.Title = video.Title

	return videoResp
}
