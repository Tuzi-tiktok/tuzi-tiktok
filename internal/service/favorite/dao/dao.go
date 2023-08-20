package dao

import (
	"context"
	"errors"
	"strconv"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	"tuzi-tiktok/kitex/kitex_gen/favorite"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/redis"
	"tuzi-tiktok/utils/changes"
	consts "tuzi-tiktok/utils/consts/favorite"
)

var f = query.Favorite
var v = query.Video
var ctx = context.TODO()

// GetFavorList 得到点赞列表
func GetFavorList(UserId int64) (resp *favorite.FavoriteListResponse, err error) {
	resp = new(favorite.FavoriteListResponse)
	videos, err := f.Debug().Where(f.UID.Eq(UserId)).Find()
	if err != nil {
		return nil, err
	}
	for _, value := range videos {
		video, err := v.Where(v.ID.Eq(value.Vid)).First()
		if err != nil {

			return nil, err
		}
		VideoResp, err := changes.VideoRecord2videoResp(UserId, video)
		resp.VideoList = append(resp.VideoList, VideoResp)

	}

	return
}

// FavorAction FollowAction 点赞
func FavorAction(uid, vid int64) error {

	//查询点赞关系是否存在
	count, err := f.Where(f.UID.Eq(uid), f.Vid.Eq(vid), f.DeletedAt.IsNull()).Count()
	if err != nil {
		logger.Errorf("query favor record error", err.Error())
		return err
	}
	if count > 0 {
		logger.Infof("user: %d have liked", uid)
		return errors.New("不要重复点赞")
	}

	favor := model.Favorite{UID: uid, Vid: vid}
	logger.Infof("user:%d like video:%d", uid, vid)
	err = f.WithContext(ctx).Create(&favor)
	if err != nil {
		return err
	}
	//更新video点赞数量
	v := query.Video
	_, err = v.Where(v.ID.Eq(vid)).Update(v.FavoriteCount, v.FavoriteCount.Add(1))
	if err != nil {
		return err
	}
	return nil
}

// UnFavorAction UnFollowAction 取消点赞
func UnFavorAction(uid, vid int64) error {

	result, err := f.WithContext(ctx).Where(f.UID.Eq(uid), f.Vid.Eq(vid)).Delete()
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		logger.Infof("like record not exist")
		return nil
	}
	logger.Infof("user:%d unlike video:%d", uid, vid)
	//更新video点赞数量
	v := query.Video
	_, err = v.Where(v.ID.Eq(vid)).Update(v.FavoriteCount, v.FavoriteCount.Sub(1))
	if err != nil {
		return err
	}
	return nil
}

func UpdateLike(uid, vid int64, actionType int32) (resp *favorite.FavoriteResponse, err error) {
	resp = new(favorite.FavoriteResponse)
	//判断当前用户是否点过赞
	var key string = "video:liked:" + strconv.Itoa(int(vid))
	ok, err := redis.IRC.SIsMember(context.Background(), key, uid).Result()
	if err != nil {
		return nil, err
	}

	// actionType == 1 点赞
	if actionType == 1 {
		//用户点赞过了
		if ok {
			logger.Infof("user:%d has liked video:%d", uid, vid)
			return nil, errors.New("不要重复点赞")
		}
		//数据库点赞数+1
		result, err := v.Where(v.ID.Eq(vid)).Update(v.FavoriteCount, v.FavoriteCount.Add(1))
		if err != nil {
			return nil, err
		}
		if result.RowsAffected == 0 {
			logger.Infof("not video record")
			return nil, errors.New("not video record")
		}
		//点赞关系存入数据库
		favor := model.Favorite{UID: uid, Vid: vid}
		err = f.WithContext(ctx).Create(&favor)
		if err != nil {
			return nil, err
		}
		//保存用户到redis set集合
		err = redis.IRC.SAdd(context.Background(), key, uid).Err()
		if err != nil {
			return nil, err
		}
		//取消点赞 actionType == 2
	} else if actionType == 2 {
		if !ok {
			logger.Infof("user:%d like video:%d not exist", uid, vid)
			resp.StatusCode = consts.FavorRecordNotExist
			resp.StatusMsg = &consts.FavorRecordNotExistMsg
			return resp, errors.New("点赞关系不存在 无法删除")
		}
		//数据库点赞数-1
		v := query.Video
		result, err := v.Where(v.ID.Eq(vid)).Update(v.FavoriteCount, v.FavoriteCount.Sub(1))
		if err != nil {
			return nil, err
		}
		if result.RowsAffected == 0 {
			return nil, errors.New("video record not exist")
		}
		//删除点赞关系
		result, err = f.WithContext(ctx).Where(f.UID.Eq(uid), f.Vid.Eq(vid)).Delete()
		if err != nil {
			return nil, err
		}
		if result.RowsAffected == 0 {
			logger.Infof("user:%d and video:%d record not exist", uid, vid)
			return nil, errors.New("record not exist")
		}

		//把用户从redis的集合中移除
		err = redis.IRC.SRem(context.Background(), key, uid).Err()
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
