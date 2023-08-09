package mapstruct

import (
	"tuzi-tiktok/gateway/biz/model/favorite"
	"tuzi-tiktok/gateway/biz/model/feed"
)
import kfavorite "tuzi-tiktok/kitex/kitex_gen/favorite"

func ToFavoriteListResponse(k *kfavorite.FavoriteListResponse) *favorite.FavoriteListResponse {
	if k == nil {
		return nil
	}
	videoList := make([]*feed.Video, len(k.VideoList))
	for i := range k.VideoList {
		videoList[i] = ToVideo(k.VideoList[i])
	}

	return &favorite.FavoriteListResponse{
		StatusCode: k.StatusCode,
		StatusMsg:  k.StatusMsg,
		VideoList:  videoList,
	}
}

func ToFavoriteResponse(p *kfavorite.FavoriteListResponse) *favorite.FavoriteResponse {
	if p == nil {
		return nil
	}
	return &favorite.FavoriteResponse{
		StatusCode: p.StatusCode,
		StatusMsg:  p.StatusMsg,
	}
}
