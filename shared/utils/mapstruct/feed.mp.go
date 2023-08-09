package mapstruct

import "tuzi-tiktok/gateway/biz/model/feed"
import kfeed "tuzi-tiktok/kitex/kitex_gen/feed"

func ToVideo(k *kfeed.Video) *feed.Video {
	if k == nil {
		return nil
	}
	return &feed.Video{
		Id:            k.Id,
		Author:        ToUser(k.Author),
		PlayUrl:       k.PlayUrl,
		CoverUrl:      k.CoverUrl,
		FavoriteCount: k.FavoriteCount,
		CommentCount:  k.CommentCount,
		IsFavorite:    k.IsFavorite,
		Title:         k.Title,
	}
}

func ToFeedResponse(k *kfeed.FeedResponse) *feed.FeedResponse {
	if k == nil {
		return nil
	}
	videoList := make([]*feed.Video, len(k.VideoList))
	for i := range k.VideoList {
		videoList[i] = ToVideo(k.VideoList[i])
	}
	return &feed.FeedResponse{
		StatusCode: k.StatusCode,
		StatusMsg:  k.StatusMsg,
		VideoList:  videoList,
		NextTime:   k.NextTime,
	}
}
