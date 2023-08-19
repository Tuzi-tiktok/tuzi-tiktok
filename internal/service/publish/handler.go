package main

import (
	"context"
	"io"
	"strings"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	"tuzi-tiktok/kitex/kitex_gen/auth"
	"tuzi-tiktok/kitex/kitex_gen/feed"
	publish "tuzi-tiktok/kitex/kitex_gen/publish"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/service/auth/tools"
	"tuzi-tiktok/service/filetransfer/client"
	fm "tuzi-tiktok/service/filetransfer/model"
	"tuzi-tiktok/utils/consts"
	"tuzi-tiktok/utils/ffmpeg"
)

var transfer client.Transfer

var (
	qVideo    = query.Video
	qUser     = query.User
	qFavorite = query.Favorite
	qRelation = query.Relation
)

func init() {
	transfer = client.NewTransfer()
}

// UploadSnapShot Upload the video cover to Oss
func UploadSnapShot(shot io.Reader) fm.TransResult {
	pth := strings.Join([]string{"images", "1.png"}, "#")
	logger.Debug(pth)
	v := transfer.Put(pth, shot)
	return v
}

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishVideo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishVideo(ctx context.Context, req *publish.PublishRequest) (*publish.PublishResponse, error) {
	//   TODO  Token expires after video upload
	claims, err := tools.ParseToken(req.Token)

	if err != nil {
		logger.Debug(err)
		return &publish.PublishResponse{
			StatusMsg:  &consts.InvalidTokenMsg,
			StatusCode: consts.InvalidToken,
		}, nil
	}
	uid := claims.Payload.UID
	shot, err := ffmpeg.GetSnapShots(req.VideoUrl)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	result := UploadSnapShot(shot)
	if !result.Ok {
		logger.Debug("upload SnapShot error")
		return &publish.PublishResponse{
			StatusCode: consts.PublishUploadSnapShotError,
			StatusMsg:  &consts.PublishListErrorMsg,
		}, nil
	}

	video := &model.Video{
		Title:    req.Title,
		AuthorID: uid,
		CoverURL: result.Url,
		PlayURL:  req.VideoUrl,
	}
	err = qVideo.Create(video)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}

	return &publish.PublishResponse{
		StatusCode: consts.Success,
		StatusMsg:  &consts.SuccessMsg,
	}, nil
}

// GetPublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) GetPublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {

	claims, err := tools.ParseToken(req.Token)
	resp = new(publish.PublishListResponse)
	// 默认非游客
	isTourist := false
	if err != nil {
		isTourist = true
	}
	currentUid, targetId := claims.Payload.UID, req.UserId

	// UserId 对应发布的视频
	videos, err := qVideo.Where(qVideo.AuthorID.Eq(targetId)).Find()
	videosNums := int64(len(videos))
	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	// UserId 对应的个人信息
	users, err := qUser.Where(qUser.ID.Eq(targetId)).Find()
	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	if len(users) == 0 {
		return &publish.PublishListResponse{
			StatusMsg:  &consts.PublishTargetUserNotExistMsg,
			StatusCode: consts.PublishTargetUserNotExist,
		}, nil
	}

	user := users[0]
	isFollow := false
	if !isTourist {
		// 查询当前用户是否关注目标用户
		count, err := qRelation.Where(qRelation.FollowerID.Eq(currentUid), qRelation.FollowingID.Eq(targetId)).Count()
		if err != nil {
			logger.Warn(err)
			return nil, err
		}
		isFollow = count != 0
	}

	// 目标用户获赞总数和单个视频获赞数目
	var totalFavorited int64
	favoritedCounts := make([]int64, videosNums)
	for i := 0; i < len(videos); i++ {
		favoritedCount, err := qFavorite.Where(qFavorite.Vid.Eq(videos[i].ID)).Count()
		favoritedCounts[i] = favoritedCount
		if err != nil {
			return nil, err
		}
		totalFavorited += favoritedCount
	}

	if err != nil {
		return nil, err
	}

	// 目标用户点赞数目
	favoriteCount, err := qFavorite.Where(qFavorite.UID.Eq(targetId)).Count()
	if err != nil {
		logger.Warn(err)
		return nil, err
	}

	videoList := make([]*feed.Video, videosNums)
	for i := range videos {
		v := videos[i]
		isFavorite := false
		if !isTourist {
			count, err := qFavorite.Where(qFavorite.UID.Eq(currentUid), qFavorite.Vid.Eq(v.ID)).Count()
			if err != nil {
				return nil, err
			}
			// 当前用户是否点赞当前视频
			isFavorite = count != 0
		}
		followCount, followerCount := user.FollowCount, user.FollowerCount
		MergeStruct(favoritedCounts, videoList, i, v, isFavorite, user, followCount, followerCount,
			isFollow, totalFavorited, videosNums, favoriteCount)
	}
	return &publish.PublishListResponse{
		VideoList:  videoList,
		StatusMsg:  &consts.SuccessMsg,
		StatusCode: consts.Success,
	}, nil
}

func MergeStruct(favoritedCounts []int64, videoList []*feed.Video, i int, v *model.Video, isFavorite bool, user *model.User, followCount int64, followerCount int64, isFollow bool, totalFavorited int64, videosNums int64, favoriteCount int64) {
	videoList[i] = &feed.Video{
		Id:            v.ID,
		PlayUrl:       v.PlayURL,
		CoverUrl:      v.CoverURL,
		CommentCount:  v.CommentCount,
		Title:         v.Title,
		IsFavorite:    isFavorite,
		FavoriteCount: favoritedCounts[i],
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
