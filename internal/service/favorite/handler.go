package main

import (
	"context"
	"fmt"
	favorite "tuzi-tiktok/kitex/kitex_gen/favorite"
	"tuzi-tiktok/service/favorite/dao"
	"tuzi-tiktok/utils/consts"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavorVideo implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavorVideo(ctx context.Context, req *favorite.FavoriteRequest) (resp *favorite.FavoriteResponse, err error) {

	resp = new(favorite.FavoriteResponse)
	uid := int64(1)

	if req.ActionType == 1 {
		//点赞操作
		err := dao.FavorAction(uid, req.VideoId)
		if err != nil {
			return nil, err
		}
	} else {
		//取消点赞操作
		err = dao.UnFavorAction(uid, req.VideoId)
		if err != nil {
			return nil, err
		}
	}
	resp.StatusCode = consts.Success
	return
}

// GetFavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetFavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {

	fmt.Println(req.Token)
	fmt.Println(req.UserId)

	resp = new(favorite.FavoriteListResponse)
	resp, err = dao.GetFavorList(req.UserId)
	if err != nil {
		resp.StatusCode = consts.Success
		str := "查询喜欢列表失败"
		resp.StatusMsg = &str
		return resp, err
	}
	resp.StatusCode = consts.Success
	resp.StatusMsg = &consts.SuccessMsg
	return
}
