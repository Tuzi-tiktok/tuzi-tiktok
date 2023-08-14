package main

import (
	"context"
	"fmt"
	favorite "tuzi-tiktok/kitex/kitex_gen/favorite"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/secret"
	"tuzi-tiktok/service/favorite/dao"
	consts "tuzi-tiktok/utils/consts/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavorVideo implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavorVideo(ctx context.Context, req *favorite.FavoriteRequest) (resp *favorite.FavoriteResponse, err error) {

	resp = new(favorite.FavoriteResponse)
	// check token & get uid
	claims, err := secret.ParseToken(req.Token)
	if err != nil {
		logger.Infof("failed to parse token, err: %v", err)
		return resp, nil
	}
	uid := claims.Payload.UID

	logger.Infof("user: %d favor action video: %d", uid, req.VideoId)
	if req.ActionType == 1 {
		//点赞操作
		err := dao.FavorAction(uid, req.VideoId)
		if err != nil {
			logger.Errorf("failed to favor action, err: %v", err)
			return nil, err
		}
	} else {
		//取消点赞操作
		err = dao.UnFavorAction(uid, req.VideoId)
		if err != nil {
			logger.Errorf("failed to cancel favor action, err: %v", err)
			return nil, err
		}
	}
	err = dao.UpdateLike(uid, req.VideoId, req.ActionType)
	if err != nil {
		logger.Errorf("favor redis error", err.Error())
		resp.StatusCode = consts.FavorGetFavorListError
		resp.StatusMsg = &consts.FavorGetListFailedMsg
		return
	}
	resp.StatusCode = consts.FavorSucceed
	resp.StatusMsg = &consts.FavorSucceedMsg
	return
}

// GetFavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) GetFavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {

	logger.Infof("get user:%d favorite list", req.UserId)
	_, err = secret.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("failed to parse token, err: %v", err)
		return resp, nil
	}
	fmt.Println(req.UserId)

	resp = new(favorite.FavoriteListResponse)
	resp, err = dao.GetFavorList(req.UserId)
	if err != nil {
		logger.Errorf("failed to get favor list, err: %v", err)
		resp.StatusCode = consts.FavorGetListFailed
		resp.StatusMsg = &consts.FavorGetListFailedMsg
		return resp, nil
	}
	resp.StatusCode = consts.FavorSucceed
	resp.StatusMsg = &consts.FavorSucceedMsg
	return
}
