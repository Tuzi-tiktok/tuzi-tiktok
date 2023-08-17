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
	if err != nil {
		logger.Errorf("query favor record error", err.Error())
		return err
	}
	if count > 0 {
		logger.Infof("user: %d have liked", uid)
		return nil
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

//func UpdateLike(uid, vid int64, actionType int32) error {
//	logger.Infof("user:%d update like video:%d", uid, vid)
//
//	//单独记录视频本身的键的设计
//	var key string
//	key = "VideoID:" + strconv.Itoa(int(vid))
//	//记录用户点赞具体视频的键值设计
//	var userKey string
//	userKey = "VideoID_USERID:" + strconv.Itoa(int(vid)) + "-" + strconv.Itoa(int(uid))
//	exist, err := redis.IRC.Exists(context.Background(), userKey).Result()
//	if exist == 1 {
//		logger.Infof("user:%d have favor video:%d", uid, vid)
//		return nil
//	}
//	// 检测id字段是否存在
//	exist, err = redis.IRC.Exists(context.Background(), key).Result()
//	if err != nil {
//		logger.Errorf(err.Error())
//		return err
//	}
//
//	//视频点赞的键值存在
//	if exist == 1 {
//		result, err := redis.IRC.Get(context.Background(), key).Result()
//		if err != nil {
//			return err
//		}
//		likeNums, err := strconv.Atoi(result)
//		if err != nil {
//			return err
//		}
//		//用户点赞
//		if actionType == 1 {
//			logger.Infof("redis 关注")
//			//逻辑一:首先视频点赞总数的存储在Redis中
//			err = redis.IRC.Set(context.Background(), key, likeNums+1, 60*time.Minute).Err()
//			if err != nil {
//				return err
//			}
//			//逻辑二:在Redis中记录某一个具体用户的点赞情况
//			err := redis.IRC.Set(context.Background(), userKey, 1, 60*time.Minute).Err()
//			if err != nil {
//				return err
//			}
//
//		} else {
//			//逻辑一:首先视频点赞总数的存储在Redis中
//			logger.Infof("redis 取消关注")
//			err := redis.IRC.Set(context.Background(), key, likeNums-1, 60*time.Minute).Err()
//			if err != nil {
//				return err
//			}
//			//逻辑二:在Redis中记录某一个具体用户的点赞情况
//			err = redis.IRC.Del(context.Background(), userKey).Err()
//			if err != nil {
//				return err
//			}
//		}
//
//	} else {
//		/**
//		  这里是判断Redis中没有相关键值记录，即Redis中没有记录具体的视频点赞总数和某一个用户的点赞情况。
//
//		  这里一般是两个情况
//		  1.该视频从来没有被点赞过 这是第一次，第一次点赞的时候就会设置键值
//		  2.Redis中相关键值对过期
//		  **/
//
//		//数据库拿数据
//		v := query.Video
//		likeCount := 0
//		err = v.Select(v.FavoriteCount).Where(v.ID.Eq(vid)).Scan(&likeCount)
//		if err != nil {
//			return err
//		}
//
//		//用户点赞
//		if actionType == 1 {
//			//逻辑二:在Redis中记录某一个具体用户的点赞情况
//			err := redis.IRC.Set(context.Background(), userKey, 1, 60*time.Minute).Err()
//			if err != nil {
//				return err
//			}
//
//			//逻辑一:首先视频点赞总数的存储在Redis中
//			err = redis.IRC.Set(context.Background(), key, likeCount+1, 60*time.Minute).Err()
//			if err != nil {
//				return err
//			}
//
//		} else {
//			if likeCount == 0 {
//				logger.Infof("点赞总数为0，无法取消点赞")
//			} else {
//				//逻辑一:首先视频点赞总数的存储在Redis中
//				err := redis.IRC.Set(context.Background(), key, likeCount-1, 10*time.Minute).Err()
//				if err != nil {
//					return err
//				}
//				//逻辑二:在Redis中记录某一个具体用户的点赞情况
//				err = redis.IRC.Del(context.Background(), userKey).Err()
//				if err != nil {
//					return err
//				}
//			}
//		}
//	}
//
//	return nil
//}

func UpdateLike(uid, vid int64, actionType int32) error {
	//判断当前用户是否点过赞
	var key string = "video:liked:" + strconv.Itoa(int(vid))
	ok, err := redis.IRC.SIsMember(context.Background(), key, uid).Result()
	if err != nil {
		return err
	}
	// actionType == 1 点赞
	if actionType == 1 {

		//用户已经点赞
		if ok {
			logger.Infof("user:%d has liked video:%d", uid, vid)
			return nil
		}
		//数据库点赞数+1
		result, err := v.Where(v.ID.Eq(vid)).Update(v.FavoriteCount, v.FavoriteCount.Add(1))
		if err != nil {
			return err
		}
		if result.RowsAffected == 0 {
			logger.Infof("not video record")
			return errors.New("not video record")
		}
		//点赞关系存入数据库
		favor := model.Favorite{UID: uid, Vid: vid}
		err = f.WithContext(ctx).Create(&favor)
		if err != nil {
			return err
		}
		//保存用户到redis set集合
		err = redis.IRC.SAdd(context.Background(), key, uid).Err()
		if err != nil {
			return err
		}
		//取消点赞 actionType == 2
	} else if actionType == 2 {
		//用户原本没有点赞，不能够执行取消点赞的操作
		if !ok {
			logger.Infof("user:%d hasn't liked video:%d", uid, vid)
			return nil
		}
		//数据库点赞数-1
		v := query.Video
		_, err = v.Where(v.ID.Eq(vid)).Update(v.FavoriteCount, v.FavoriteCount.Sub(1))
		if err != nil {
			return err
		}
		//删除点赞关系
		result, err := f.WithContext(ctx).Where(f.UID.Eq(uid), f.Vid.Eq(vid)).Delete()
		if err != nil {
			return err
		}
		if result.RowsAffected == 0 {
			logger.Infof("user:%d and video:%d record not exist", uid, vid)
			return errors.New("record not exist")
		}

		//把用户从redis的集合中移除
		err = redis.IRC.SRem(context.Background(), key, uid).Err()
		if err != nil {
			return err
		}
	}

	return nil
}
