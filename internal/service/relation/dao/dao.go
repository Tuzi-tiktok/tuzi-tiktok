package dao

import (
	"context"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	"tuzi-tiktok/kitex/kitex_gen/relation"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/utils/changes"
)

var ctx = context.TODO()

// FollowAction 添加关注
func FollowAction(follower, following int64) error {
	//can not focus on yourself
	if follower == following {
		logger.Infof("user:%d can not focus on yourself", follower)
		return nil
	}
	r := query.Relation
	//查询关注关系是否存在
	count, err := r.Where(r.FollowerID.Eq(follower), r.FollowingID.Eq(following)).Count()
	if count > 0 {
		logger.Infof("user:%d have followed user:%d", follower, following)
		return nil
	}
	logger.Infof("user:%d follow user:%d", follower, following)
	relationRecord := model.Relation{FollowerID: int64(follower), FollowingID: int64(following)}
	err = r.WithContext(ctx).Create(&relationRecord)
	if err != nil {
		return err
	}
	return nil
}

// UnFollowAction 取消关注
func UnFollowAction(follower, following int64) error {
	//can not focus on yourself
	if follower == following {
		logger.Infof("user:%d can not focus on yourself", follower)
		return nil
	}
	r := query.Relation
	result, err := r.WithContext(ctx).Where(r.FollowerID.Eq(follower), r.FollowingID.Eq(following)).Delete()
	if result.RowsAffected == 0 {
		logger.Infof("user:%d and user:%d relation record not exist", follower, following)
		return nil
	}
	if err != nil {
		return err
	}
	logger.Infof("user:%d unfollow user:%d", follower, following)
	return nil
}

// GetFriendList 查找好友列表
func GetFriendList(usrId int64) (resp *relation.RelationFriendListResponse, err error) {
	resp = new(relation.RelationFriendListResponse)
	u := query.User
	r := query.Relation
	r2 := r.As("r2")
	result, err := r.WithContext(ctx).Select(r.FollowingID).LeftJoin(r2, r2.FollowerID.EqCol(r.FollowingID)).
		Where(r.FollowerID.EqCol(r2.FollowingID), r.FollowerID.Eq(usrId)).Find()
	if err != nil {
		return nil, err
	}
	
	for _, v := range result {
		user, err := u.Where(u.ID.Eq(v.FollowingID)).First()
		if err != nil {
			return nil, err
		}
		resp.UserList = append(resp.UserList, changes.UserRecord2friendResp(user))
	}

	return resp, nil

}
