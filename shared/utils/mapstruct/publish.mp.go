package mapstruct

import (
	"tuzi-tiktok/gateway/biz/model/feed"
	"tuzi-tiktok/gateway/biz/model/publish"
)
import kpublish "tuzi-tiktok/kitex/kitex_gen/publish"

func ToPublishListResponse(k *kpublish.PublishListResponse) *publish.PublishListResponse {
	if k == nil {
		return nil
	}
	videoList := make([]*feed.Video, len(k.VideoList))
	for i := range k.VideoList {
		videoList[i] = ToVideo(k.VideoList[i])
	}
	return &publish.PublishListResponse{
		StatusCode: k.StatusCode,
		StatusMsg:  k.StatusMsg,
		VideoList:  videoList,
	}
}

func ToPublishResponse(k *kpublish.PublishResponse) *publish.PublishResponse {
	if k == nil {
		return nil
	}
	return &publish.PublishResponse{
		StatusCode: k.StatusCode,
		StatusMsg:  k.StatusMsg,
	}
}
