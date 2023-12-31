package dao

import (
	"context"
	"errors"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	"tuzi-tiktok/kitex/kitex_gen/relation"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/utils/changes"
)

var ctx = context.TODO()

// FollowAction 添加关注
func FollowAction(follower, following int64) (resp *relation.RelationResponse, err error) {
	resp = new(relation.RelationResponse)
	r := query.Relation
	//查询关注关系是否存在
	count, err := r.Where(r.FollowerID.Eq(follower), r.FollowingID.Eq(following)).Count()
	if count > 0 {
		logger.Infof("user:%d have followed user:%d", follower, following)
		return nil, errors.New("不要重复关注")
	}
	logger.Infof("user:%d follow user:%d", follower, following)
	relationRecord := model.Relation{FollowerID: int64(follower), FollowingID: int64(following)}
	err = r.WithContext(ctx).Create(&relationRecord)
	if err != nil {
		return nil, err
	}

	u := query.User
	//following follower_count++
	result, err := u.Where(u.ID.Eq(following)).Update(u.FollowerCount, u.FollowerCount.Add(1))
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		logger.Infof("user:%d record not find", following)
		return nil, errors.New("follow action record not find")
	}
	//follower follow_count++
	result, err = u.Where(u.ID.Eq(follower)).Update(u.FollowCount, u.FollowCount.Add(1))
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		logger.Infof("user:%d record not find", following)
		return nil, errors.New("follow action record not find")
	}

	return nil, nil
}

// UnFollowAction 取消关注
func UnFollowAction(follower, following int64) error {

	r := query.Relation
	result, err := r.WithContext(ctx).Where(r.FollowerID.Eq(follower), r.FollowingID.Eq(following)).Delete()
	if result.RowsAffected == 0 {
		logger.Infof("user:%d and user:%d relation record not exist", follower, following)
		return errors.New("关注关系不存在 删除失败")
	}
	if err != nil {
		return err
	}
	logger.Infof("user:%d unfollow user:%d", follower, following)

	u := query.User
	//following follower_count--
	result, err = u.Where(u.ID.Eq(following)).Update(u.FollowerCount, u.FollowerCount.Sub(1))
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		logger.Infof("user:%d record not find", following)
		return errors.New("follow action record not find")
	}
	//follower follow_count--
	result, err = u.Where(u.ID.Eq(follower)).Update(u.FollowCount, u.FollowCount.Sub(1))
	if err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		logger.Infof("user:%d record not find", following)
		return errors.New("follow action record not find")
	}
	return nil
}

// GetFriendList 查找好友列表
func GetFriendList(usrId int64) (resp *relation.RelationFriendListResponse, err error) {
	resp = new(relation.RelationFriendListResponse)
	u := query.User
	r := query.Relation

	r2 := r.As("r2")
	result, err := r.WithContext(ctx).Debug().Select(r.FollowingID).LeftJoin(r2, r2.FollowerID.EqCol(r.FollowingID)).
		Where(r.FollowerID.EqCol(r2.FollowingID), r.FollowerID.Eq(usrId), r.DeletedAt.IsNull(), r2.DeletedAt.IsNull()).Find()
	if err != nil {
		return nil, err
	}

	for _, v := range result {
		user, err := u.Where(u.ID.Eq(v.FollowingID)).First()
		if err != nil {
			return nil, err
		}
		friendResp, err := changes.UserRecord2friendResp(usrId, user)
		if err != nil {
			return nil, err
		}
		resp.UserList = append(resp.UserList, friendResp)
	}

	return resp, nil

}
