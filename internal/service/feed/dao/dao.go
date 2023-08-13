package dao

import (
	"context"
	"time"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	"tuzi-tiktok/kitex/kitex_gen/auth"
	"tuzi-tiktok/kitex/kitex_gen/feed"
	"tuzi-tiktok/logger"
	consts "tuzi-tiktok/utils/consts/feed"
)

const (
	DefaultVideoListMax = 30
)

// todo: 这里最终决定直接通过调用query获取数据，不使用rpc的原因是需要传token值，后续看如何优化
var (
	qVideo    = query.Video
	qUser     = query.User
	qFavorite = query.Favorite
	qRelation = query.Relation
)

type QVideo struct{}

var Video = QVideo{}

type login struct {
	IsLogin bool
}

var p = login{}

// GetVideoListWithTime 根据本次时间逆序查找limit数量的video列表
func (QVideo) GetVideoListWithTime(ctx context.Context, uid int64, lTime int64, limit int) (videos []*feed.Video, nt time.Time, err error) {
	// 用户状态赋予login.IsLogin
	if uid == consts.NOUSERSTATE {
		p.IsLogin = false
	}

	// 转化lTime，时间戳转化为Time
	t := time.Unix(lTime, 0)

	// 单次返回视频列表为30，当超过30时，使用最大值
	// todo: 这里多了第二层判断返回视频列表数量的限制，可以考虑去掉
	if limit > 30 {
		limit = DefaultVideoListMax
	}

	// todo: 存在问题，当last_time超过了数据库中最早时间时，应当从最新的时间戳获取缺少的视频列表
	mVideos, err := qVideo.WithContext(ctx).Where(qVideo.CreatedAt.Lt(t), qVideo.DeletedAt.IsNull()).Order(qVideo.CreatedAt.Desc()).Limit(limit).Find()
	if err != nil {
		logger.Errorf("Error querying video list, err: %v", err)
		return nil, time.Now(), err
	}
	if len(mVideos) == 0 {
		logger.Info("The lastTime query is reset to time.Now().")
		mVideos, err = qVideo.WithContext(ctx).Where(qVideo.CreatedAt.Lt(time.Now()), qVideo.DeletedAt.IsNull()).Order(qVideo.CreatedAt.Desc()).Limit(limit).Find()
	}

	videos, nt, err = mVideo2fVideoMore(uid, t, mVideos)

	return
}

// countVideos 统计作品数量
func countVideos(aid int64) *int64 {
	count, err := qVideo.WithContext(context.Background()).Where(qVideo.AuthorID.Eq(aid), qVideo.DeletedAt.IsNull()).Count()
	if err != nil {
		logger.Errorf("Error Querying the number of works, err: %v", err)
		//todo: 修改？当发生错误时，返回0
		return nil
	}
	return &count
}

// getUserInfo 根据获取用户（作者）信息
func getUserInfoByAuthorID(aid int64) (u *model.User) {
	// todo: 从redis获取userInfo(?)
	u, err := qUser.WithContext(context.Background()).Where(qUser.ID.Eq(aid), qUser.DeletedAt.IsNull()).First()
	if err != nil {
		logger.Errorf("Error querying user, err: %v", err)
		return nil
	}
	return
}

// isFollower 判断是否关注该作者
func isFollower(uid int64, aid int64) bool {
	if !p.IsLogin {
		return false
	}

	find, err := qRelation.WithContext(context.Background()).Where(qRelation.FollowerID.Eq(uid), qRelation.FollowingID.Eq(aid), qRelation.DeletedAt.IsNull()).Count()
	if err != nil {
		logger.Errorf("Error querying whether the user follows the author, err: %v", err)
		return false
	}
	return find == 1
}

// isFavorite 判断用户是否点赞该视频
func isFavorite(uid int64, vid int64) bool {
	if !p.IsLogin {
		return false
	}

	find, err := qFavorite.WithContext(context.Background()).Where(qFavorite.UID.Eq(uid), qFavorite.Vid.Eq(vid), qFavorite.DeletedAt.IsNull()).Count()
	if err != nil {
		logger.Errorf("Error querying if the user has liked the video, err: %v", err)
		return false
	}
	return find == 1
}

// getUserFavorite 获取用户获取喜欢视频的数量
func getUserFavorite(aid int64) *int64 {
	count, err := qFavorite.WithContext(context.Background()).Where(qFavorite.UID.Eq(aid), qFavorite.DeletedAt.IsNull()).Count()
	if err != nil {
		logger.Errorf("Error querying the number of videos the user has liked, err: %v", err)
		return nil
	}
	return &count
}

// getVideoFavorite 获取视频获取点赞数量
func getVideoFavorite(vid int64) int64 {
	count, err := qFavorite.WithContext(context.Background()).Where(qFavorite.Vid.Eq(vid), qFavorite.DeletedAt.IsNull()).Count()
	if err != nil {
		logger.Errorf("Error querying the video to get the number of likes, err: %v", err)
		return 0
	}

	return count
}

// getTotalFavorite 获取作者获得点赞总数
func getTotalFavorite(aid int64) *int64 {
	// 获取作者作品vid列表
	vids, err := qVideo.WithContext(context.Background()).Select(qVideo.ID).Where(qVideo.AuthorID.Eq(aid), qVideo.DeletedAt.IsNull()).Find()
	if err != nil {
		logger.Errorf("Error querying all video ids of the account, err: %v", err)
		return nil
	}
	// 查询获得点赞数量
	var total int64
	for _, vid := range vids {
		total += getVideoFavorite(vid.ID)
	}
	return &total
}

// mVideo2fVideoOne 单个model.Video转化为feed.Video
func mVideo2fVideoOne(uid int64, m *model.Video) (f *feed.Video) {
	f = new(feed.Video)
	f.Id = m.ID
	f.Author = mUser2aUserOne(uid, getUserInfoByAuthorID(m.AuthorID))
	f.PlayUrl = m.PlayURL
	f.CoverUrl = m.CoverURL
	f.FavoriteCount = m.FavoriteCount
	f.CommentCount = m.CommentCount
	f.IsFavorite = isFavorite(uid, m.ID)
	f.Title = m.Title
	return
}

// mVideo2fVideoMore 切片model.Video转化为feed.Video
func mVideo2fVideoMore(uid int64, lTime time.Time, mv []*model.Video) ([]*feed.Video, time.Time, error) {
	fv := make([]*feed.Video, 0)
	nt := time.Now()

	for _, m := range mv {
		if nt.After(*m.CreatedAt) {
			nt = *m.CreatedAt
		}

		// todo: 感觉这里可以考虑开启协程
		f := mVideo2fVideoOne(uid, m)
		fv = append(fv, f)
	}
	return fv, nt, nil
}

// mUser2aUserOne 单个model.User转化为auth.User
func mUser2aUserOne(uid int64, u *model.User) (a *auth.User) {
	a = new(auth.User)
	a.Id = u.ID
	a.Name = u.Username
	a.FollowCount = &u.FollowCount
	a.FollowerCount = &u.FollowerCount
	a.Avatar = u.Avatar
	a.IsFollow = isFollower(uid, u.ID)
	a.BackgroundImage = u.BackgroundImage
	a.Signature = u.Signature
	a.TotalFavorited = getTotalFavorite(u.ID)
	a.WorkCount = countVideos(u.ID)
	a.FavoriteCount = getUserFavorite(u.ID)
	return
}
