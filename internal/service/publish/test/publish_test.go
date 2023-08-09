package test

import (
	"io"
	"strings"
	"sync"
	"testing"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	"tuzi-tiktok/kitex/kitex_gen/auth"
	"tuzi-tiktok/kitex/kitex_gen/feed"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/service/filetransfer/client"
	fm "tuzi-tiktok/service/filetransfer/model"
	"tuzi-tiktok/utils/ffmpeg"
)

var transfer client.Transfer

func init() {
	transfer = client.NewTransfer()
}

func TestPublishList(t *testing.T) {

	var (
		currentUid, targetId int64 = 3, 3
		qVideo                     = query.Video
		qUser                      = query.User
		qFavorite                  = query.Favorite
		qRelation                  = query.Relation
	)

	// UserId 对应发布的视频
	videos, err := qVideo.Where(qVideo.AuthorID.Eq(targetId)).Find()
	videosNums := int64(len(videos))
	if err != nil {
		logger.Warn(err)
		return
	}
	// UserId 对应的个人信息
	user, err := qUser.Where(qUser.ID.Eq(targetId)).First()
	if err != nil {
		logger.Warn(err)
		return
	}
	// 查询当前用户是否关注目标用户
	count, err := qRelation.Where(qRelation.FollowerID.Eq(currentUid), qRelation.FollowingID.Eq(targetId)).Count()
	isFollow := count != 0
	if err != nil {
		logger.Warn(err)
		return
	}
	// 目标用户获赞数目
	vids := make([]int64, videosNums)
	for i := range vids {
		vids[i] = videos[i].ID
	}
	totalFavorited, err := qFavorite.Where(qFavorite.Vid.In(vids...)).Count()
	if err != nil {
		return
	}

	// 目标用户点赞数目
	favoriteCount, err := qFavorite.Where(qFavorite.UID.Eq(targetId)).Count()
	if err != nil {
		logger.Warn(err)
		return
	}

	vs := make([]*feed.Video, videosNums)
	for i := range videos {
		v := videos[i]
		count, err := qFavorite.Where(qFavorite.UID.Eq(currentUid), qFavorite.Vid.Eq(v.ID)).Count()
		if err != nil {
			return
		}
		// 当前用户是否点赞当前视频
		isFavorite := count != 0
		followCount, followerCount := user.FollowCount, user.FollowerCount
		vs[i] = &feed.Video{
			Id:           v.ID,
			PlayUrl:      v.PlayURL,
			CoverUrl:     v.CoverURL,
			CommentCount: v.CommentCount,
			Title:        v.Title,
			IsFavorite:   isFavorite,
			Author: &auth.User{
				Id:              user.ID,
				Name:            user.Username,
				FollowCount:     &followCount,
				FollowerCount:   &followerCount,
				IsFollow:        isFollow,
				Avatar:          user.Avatar,
				BackgroundImage: user.BackgroundImage,
				Signature:       user.Signature,
				TotalFavorited:  &totalFavorited,
				WorkCount:       &videosNums,
				FavoriteCount:   &favoriteCount,
			},
		}
	}
	//logger.Infof("%v", vs)
}

func TestPublishListASync(t *testing.T) {

	var (
		currentUid, targetId int64 = 3, 3
		qVideo                     = query.Video
		qUser                      = query.User
		qFavorite                  = query.Favorite
		qRelation                  = query.Relation
	)
	parallelNums := 4
	errors := make(chan error, parallelNums)
	group := sync.WaitGroup{}
	group.Add(parallelNums)

	var (
		videosNums    int64
		videos        []*model.Video
		user          *model.User
		isFollow      bool
		favoriteCount int64
	)
	// UserId 对应发布的视频
	go func() {
		defer group.Done()
		var err error
		videos, err = qVideo.Where(qVideo.AuthorID.Eq(targetId)).Find()
		videosNums = int64(len(videos))
		if err != nil {
			logger.Warn(err)
			errors <- err
			return
		}
	}()
	// UserId 对应的个人信息
	go func() {
		defer group.Done()

		var err error
		user, err = qUser.Where(qUser.ID.Eq(targetId)).First()
		if err != nil {
			logger.Warn(err)
			errors <- err
			return
		}
	}()
	// 查询当前用户是否关注目标用户
	go func() {
		defer group.Done()
		count, err := qRelation.Where(qRelation.FollowerID.Eq(currentUid), qRelation.FollowingID.Eq(targetId)).Count()
		isFollow = count != 0
		if err != nil {
			logger.Warn(err)
			errors <- err
			return
		}
	}()
	// 目标用户点赞数目
	go func() {
		defer group.Done()
		var err error
		favoriteCount, err = qFavorite.Where(qFavorite.UID.Eq(targetId)).Count()
		if err != nil {
			logger.Warn(err)
			errors <- err
			return
		}
	}()
	group.Wait()
	close(errors)
	for err := range errors {
		logger.Error(err)
		return
	}

	// 目标用户获赞数目
	vids := make([]int64, videosNums)
	for i := range vids {
		vids[i] = videos[i].ID
	}
	totalFavorited, err := qFavorite.Where(qFavorite.Vid.In(vids...)).Count()
	if err != nil {
		return
	}

	vs := make([]*feed.Video, videosNums)
	for i := range videos {
		v := videos[i]
		count, err := qFavorite.Where(qFavorite.UID.Eq(currentUid), qFavorite.Vid.Eq(v.ID)).Count()
		if err != nil {
			return
		}
		// 当前用户是否点赞当前视频
		isFavorite := count != 0
		followCount, followerCount := user.FollowCount, user.FollowerCount
		vs[i] = &feed.Video{
			Id:           v.ID,
			PlayUrl:      v.PlayURL,
			CoverUrl:     v.CoverURL,
			CommentCount: v.CommentCount,
			Title:        v.Title,
			IsFavorite:   isFavorite,
			Author: &auth.User{
				Id:              user.ID,
				Name:            user.Username,
				FollowCount:     &followCount,
				FollowerCount:   &followerCount,
				IsFollow:        isFollow,
				Avatar:          user.Avatar,
				BackgroundImage: user.BackgroundImage,
				Signature:       user.Signature,
				TotalFavorited:  &totalFavorited,
				WorkCount:       &videosNums,
				FavoriteCount:   &favoriteCount,
			},
		}

	}
	//logger.Infof("%v", vs)
}

func BenchmarkPublishListASync(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TestPublishListASync(nil)
	}
}

func BenchmarkPublishList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TestPublishList(nil)
	}
}

// UploadSnapShot Upload the video cover to Oss
func UploadSnapShot(shot io.Reader) fm.TransResult {
	pth := strings.Join([]string{"images", "1.png"}, "#")
	logger.Debug(pth)

	v := transfer.Put(pth, shot)
	return v
}

func TestPublishVideo(t *testing.T) {

	uid := int64(3)
	vURL := "http://100.111.163.99:9000/images-storage/q.mp4"
	shot, err := ffmpeg.GetSnapShots(vURL)
	if err != nil {
		logger.Error(err)
		return
	}
	v := UploadSnapShot(shot)
	if !v.Ok {
		logger.Error(err)
		return
	}
	video := &model.Video{
		Title:    "A",
		AuthorID: uid,
		CoverURL: v.Url,
		PlayURL:  vURL,
	}
	vorm := query.Video
	err = vorm.Create(video)
	if err != nil {
		logger.Error(err)
		return
	}
}

func TestRepetitivePublishVideo(t *testing.T) {
	for i := 0; i < 0; i++ {
		t.Run("Single", TestPublishVideo)
	}
}
