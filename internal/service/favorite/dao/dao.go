package dao

import (
	"context"
	"errors"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	"tuzi-tiktok/kitex/kitex_gen/favorite"
	"tuzi-tiktok/utils/changes"
)

var f = query.Favorite
var v = query.Video
var ctx = context.TODO()

// GetFavorList 得到点赞列表
func GetFavorList(UserId int64) (resp *favorite.FavoriteListResponse, err error) {

	resp = new(favorite.FavoriteListResponse)
	videos, err := f.Where(f.UID.Eq(UserId)).Find()
	if err != nil {
		return nil, err
	}
	for _, value := range videos {
		video, err := v.Where(v.ID.Eq(value.Vid)).First()
		if err != nil {
			return nil, err
		}
		resp.VideoList = append(resp.VideoList, changes.VideoRecord2videoResp(video))

	}

	return
}

// FavorAction FollowAction 点赞
func FavorAction(uid, vid int64) error {

	//查询点赞关系是否存在
	count, err := f.Where(f.UID.Eq(uid), f.Vid.Eq(vid)).Count()
	if count > 0 {
		return errors.New("favored")
	}

	favor := model.Favorite{UID: uid, Vid: vid}
	err = f.WithContext(ctx).Create(&favor)
	if err != nil {
		return err
	}
	return nil
}

// UnFavorAction UnFollowAction 取消点赞
func UnFavorAction(uid, vid int64) error {

	result, err := f.WithContext(ctx).Where(f.UID.Eq(uid), f.Vid.Eq(vid)).Delete()
	if result.RowsAffected == 0 {
		return errors.New("unfavor failed")
	}
	if err != nil {
		return err
	}
	return nil
}
