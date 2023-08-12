package main

import (
	"context"
	"time"
	feed "tuzi-tiktok/kitex/kitex_gen/feed"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/secret"
	"tuzi-tiktok/service/feed/dao"
	consts "tuzi-tiktok/utils/consts/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetFeedList implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetFeedList(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	resp = new(feed.FeedResponse)

	var uid int64
	// check token & get uid
	claims, err := secret.ParseToken(*req.Token)
	if err != nil {
		logger.Errorf("failed to parse token, use a no-user state, err: %v", err)
		uid = consts.NOUSERSTATE
	} else {
		// 从token获取当前用户uid
		uid = claims.Payload.UID
	}

	// 判断是否传入获取LatestTime，否，使用当前时间
	if req.GetLatestTime() == 0 {
		now := time.Now().Unix()
		req.LatestTime = &now
	}

	var vl []*feed.Video
	var t time.Time
	vl, t, err = dao.Video.GetVideoListWithTime(ctx, uid, req.GetLatestTime(), consts.DEAULT_VIDEO_LIST_LIMIST)

	resp.VideoList = vl

	nt := t.Unix()
	resp.NextTime = &nt

	if err != nil {
		return resp, err
	}

	resp.StatusCode = consts.FEED_API_SUCCESS
	resp.StatusMsg = &consts.FEED_SUCCESS_MSG

	return
}
