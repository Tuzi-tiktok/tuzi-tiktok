package changes

import (
	"errors"
	"gorm.io/gorm"
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
	//查询的用户是否关注了
	if userId != -1 {
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
func UserRecord2friendResp(reqId int64, user *model.User) (*relation.FriendUser, error) {

	userResp := new(relation.FriendUser)
	userResp.Id = user.ID
	userResp.Name = user.Username
	userResp.FollowCount = &user.FollowCount
	userResp.FollowerCount = &user.FollowerCount

	userResp.Avatar = user.Avatar
	userResp.BackgroundImage = user.BackgroundImage
	userResp.Signature = user.Signature
	//是否关注好友
	userResp.IsFollow = IsFollow(reqId, user.ID)
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

	userResp.Message, userResp.MsgType = GetLatestMsg(reqId, user.ID)

	return userResp, nil
}
func GetLatestMsg(reqId, friendId int64) (*string, int64) {

	m := query.Message

	result, err := m.Where(m.FormUserID.In(reqId, friendId), m.ToUserID.In(reqId, friendId)).Order(m.CreatedAt.Desc()).First()
	if err != nil {
		// check error ErrRecordNotFound
		if errors.Is(err, gorm.ErrRecordNotFound) {
			defaultMsg := "暂未发送信息"
			return &defaultMsg, 1
		}
		logger.Infof(err.Error())
	}

	var fromUserId int64
	fromUserId = *result.FormUserID
	var MsgType int64
	MsgType = 0
	if fromUserId == reqId {
		MsgType = 1
	}
	return result.Content, MsgType
}

// VideoRecord2videoResp 拿到video信息
func VideoRecord2videoResp(reqId int64, video *model.Video) (*feed.Video, error) {

	videoResp := new(feed.Video)
	videoResp.Id = video.ID
	u := query.User
	author, err := u.Where(u.ID.Eq(video.AuthorID)).First()
	if err != nil {
		logger.Infof(err.Error())
		return nil, err
	}
	videoResp.Author, err = UserRecord2userResp(reqId, author)
	if err != nil {
		logger.Infof(err.Error())
		return nil, err
	}

	videoResp.PlayUrl = video.PlayURL
	videoResp.CoverUrl = video.CoverURL
	videoResp.FavoriteCount = video.FavoriteCount
	videoResp.CommentCount = video.CommentCount

	// 查询用户是否点赞视频
	if reqId != -1 {
		f := query.Favorite
		favorRecord, err := f.Where(f.UID.Eq(reqId), f.Vid.Eq(video.ID)).Count()
		if err != nil {
			logger.Infof(err.Error())
			return nil, err
		}
		if favorRecord > 0 {
			videoResp.IsFavorite = true
		} else {
			videoResp.IsFavorite = false
		}
	} else {
		videoResp.IsFavorite = false
	}

	videoResp.Title = video.Title
	return videoResp, nil
}

func IsFollow(followerId, followingId int64) bool {
	r := query.Relation
	followRecord, err := r.Where(r.FollowerID.Eq(followingId), r.FollowingID.Eq(followingId)).Count()
	if err != nil {
		logger.Infof(err.Error())
	}
	if followRecord > 0 {
		return true
	}
	return false
}
