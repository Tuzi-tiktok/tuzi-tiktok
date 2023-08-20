package changes

import (
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	"tuzi-tiktok/kitex/kitex_gen/auth"
	"tuzi-tiktok/kitex/kitex_gen/feed"
	"tuzi-tiktok/kitex/kitex_gen/relation"
	"tuzi-tiktok/logger"
)

// UserRecord2userResp  userId 是req_userId
// 查询用户的记录转换为get follow / follower list的记录
func UserRecord2userResp(userId int64, user *model.User) (*auth.User, error) {

	userResp := new(auth.User)
	userResp.Id = user.ID
	userResp.Name = user.Username
	userResp.FollowCount = &user.FollowCount
	userResp.FollowerCount = &user.FollowerCount
	userResp.Avatar = user.Avatar
	userResp.BackgroundImage = user.BackgroundImage
	userResp.Signature = user.Signature
	//查询user是否关注
	r := query.Relation
	followRecord, err := r.Where(r.FollowerID.Eq(userId), r.FollowingID.Eq(user.ID)).Count()
	if err != nil {
		logger.Infof(err.Error())
		return nil, err
	}
	if followRecord > 0 {
		userResp.IsFollow = true
	} else {
		userResp.IsFollow = false
	}

	//发布视频数量
	v := query.Video
	workCount, err := v.Where(v.AuthorID.Eq(user.ID)).Count()
	if err != nil {
		logger.Infof(err.Error())
		return nil, err
	}
	userResp.WorkCount = &workCount

	//点赞数量
	f := query.Favorite
	favorCount, err := f.Where(f.UID.Eq(user.ID)).Count()
	if err != nil {
		logger.Infof(err.Error())
		return nil, err
	}
	userResp.FavoriteCount = &favorCount

	//获赞数量
	var totalFavor int64
	results, err := v.Where(v.AuthorID.Eq(user.ID)).Find()
	if err != nil {
		logger.Infof(err.Error())
		return nil, err
	}
	for i := range results {
		totalFavor += results[i].FavoriteCount
	}
	userResp.TotalFavorited = &totalFavor

	return userResp, nil
}

// UserRecord2friendResp 查询用户的记录转换为get friend list 的记录
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
		logger.Infof(err.Error())
		return nil, err
	}
	userResp.WorkCount = &workCount

	//点赞数量
	f := query.Favorite
	favorCount, err := f.Where(f.UID.Eq(user.ID)).Count()
	if err != nil {
		logger.Infof(err.Error())
		return nil, err
	}
	userResp.FavoriteCount = &favorCount
	//获赞数量
	results, err := v.Where(v.AuthorID.Eq(user.ID)).Find()
	if err != nil {
		logger.Infof(err.Error())
		return nil, err
	}
	var totalFavor int64
	for i := range results {
		totalFavor += results[i].FavoriteCount
	}
	userResp.TotalFavorited = &totalFavor
	return userResp, nil
}
func VideoRecord2videoResp(userId int64, video *model.Video) (*feed.Video, error) {

	videoResp := new(feed.Video)
	videoResp.Id = video.ID
	u := query.User
	author, err := u.Where(u.ID.Eq(video.AuthorID)).First()
	if err != nil {
		logger.Infof(err.Error())
		return nil, err
	}
	videoResp.Author, err = UserRecord2userResp(userId, author)
	if err != nil {
		logger.Infof(err.Error())
		return nil, err
	}

	videoResp.PlayUrl = video.PlayURL
	videoResp.CoverUrl = video.CoverURL
	videoResp.FavoriteCount = video.FavoriteCount
	videoResp.CommentCount = video.CommentCount

	// 查询用户是否点赞视频
	f := query.Favorite
	favorRecord, err := f.Where(f.UID.Eq(userId), f.Vid.Eq(video.ID)).Count()
	if err != nil {
		logger.Infof(err.Error())
		return nil, err
	}
	if favorRecord > 0 {
		videoResp.IsFavorite = true
	} else {
		videoResp.IsFavorite = false
	}

	videoResp.Title = video.Title

	return videoResp, nil
}
