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
	uid = consts.NOUSERSTATE
	if req.Token != nil {
		// check token & get uid
		claims, err := secret.ParseToken(*req.Token)
		if err != nil {
			logger.Errorf("failed to parse token, use a no-user state, err: %v", err)
		} else {
			// 从token获取当前用户uid
			uid = claims.Payload.UID
			logger.Infof("success to get uid : %d ", uid)
		}
	}

	logger.Debugf("get param latest_time:  %v", req.GetLatestTime())

	t := transformTimeToSecond(req.GetLatestTime())

	var vl []*feed.Video
	vl, t, err = dao.Video.GetVideoListWithTime(ctx, dao.QueryOption{
		Uid:   uid,
		Ltime: t,
		Limit: 2,
	})
	if err != nil {
		resp.StatusCode = consts.FEED_API_ERROR
		resp.StatusMsg = &consts.FEED_FAIL_MSG

		return resp, err
	}

	resp.VideoList = vl

	// 获取毫米级时间戳
	nt := t.UnixMilli()
	resp.NextTime = &nt

	resp.StatusCode = consts.FEED_API_SUCCESS
	resp.StatusMsg = &consts.FEED_SUCCESS_MSG

	return
}

// transformTimeToSecond 转化为秒级时间
func transformTimeToSecond(ms int64) time.Time {
	if ms == 0 {
		return time.Now().Truncate(time.Second)
	}
	seconds := ms / 1000

	return time.Unix(seconds, 0)
}
