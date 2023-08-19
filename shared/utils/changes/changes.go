package changes

import (
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	"tuzi-tiktok/kitex/kitex_gen/auth"
	"tuzi-tiktok/kitex/kitex_gen/feed"
	"tuzi-tiktok/kitex/kitex_gen/relation"
)

func UserRecord2userResp(user *model.User) (*auth.User, error) {

	userResp := new(auth.User)
	userResp.Id = user.ID
	userResp.Name = user.Username
	userResp.FollowCount = &user.FollowCount
	userResp.FollowerCount = &user.FollowerCount
	userResp.Avatar = user.Avatar
	userResp.BackgroundImage = user.BackgroundImage
	userResp.Signature = user.Signature
	userResp.IsFollow = true

	//发布视频数量
	v := query.Video
	workCount, err := v.Where(v.AuthorID.Eq(user.ID)).Count()
	if err != nil {
		return nil, err
	}
	userResp.WorkCount = &workCount

	//点赞数量
	f := query.Favorite
	favorCount, err := f.Where(f.UID.Eq(user.ID)).Count()
	if err != nil {
		return nil, err
	}
	userResp.FavoriteCount = &favorCount
	//获赞数量
	var totalFavor int64
	results, err := v.Where(v.AuthorID.Eq(user.ID)).Find()
	if err != nil {
		return nil, err
	}
	for i := range results {
		totalFavor += results[i].FavoriteCount
	}
	userResp.TotalFavorited = &totalFavor

	return userResp, nil
}
func UserRecord2friendResp(user *model.User) (*relation.FriendUser, error) {

	userResp := new(relation.FriendUser)
	userResp.Id = user.ID
	userResp.Name = user.Username
	userResp.FollowCount = &user.FollowCount
	userResp.FollowerCount = &user.FollowerCount
	userResp.IsFollow = true
	userResp.Avatar = user.Avatar
	userResp.BackgroundImage = user.BackgroundImage
	userResp.Signature = user.Signature

	//发布视频数量
	v := query.Video
	workCount, err := v.Where(v.AuthorID.Eq(user.ID)).Count()
	if err != nil {
		return nil, err
	}
	userResp.WorkCount = &workCount

	//点赞数量
	f := query.Favorite
	favorCount, err := f.Where(f.UID.Eq(user.ID)).Count()
	if err != nil {
		return nil, err
	}
	userResp.FavoriteCount = &favorCount
	//获赞数量
	results, err := v.Where(v.AuthorID.Eq(user.ID)).Find()
	if err != nil {
		return nil, err
	}
	var totalFavor int64
	for i := range results {
		totalFavor += results[i].FavoriteCount
	}
	userResp.TotalFavorited = &totalFavor
	return userResp, nil
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
