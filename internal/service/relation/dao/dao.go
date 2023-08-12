package dao

import (
	"context"
	"errors"
	"fmt"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	"tuzi-tiktok/kitex/kitex_gen/relation"
	"tuzi-tiktok/utils/changes"
)

var ctx = context.TODO()

// FollowAction 添加关注
func FollowAction(follower, following int64) error {

	r := query.Relation
	//查询关注关系是否存在
	count, err := r.Where(r.FollowerID.Eq(follower), r.FollowingID.Eq(following)).Count()
	if count > 0 {
		return errors.New("followed")
	}

	relation := model.Relation{FollowerID: int64(follower), FollowingID: int64(following)}
	err = r.WithContext(ctx).Create(&relation)
	if err != nil {
		return err
	}
	return nil
}

// UnFollowAction 取消关注
func UnFollowAction(follower, following int64) error {

	r := query.Relation
	result, err := r.WithContext(ctx).Where(r.FollowerID.Eq(follower), r.FollowingID.Eq(following)).Delete()
	if result.RowsAffected == 0 {
		return errors.New("unfollow failed")
	}
	if err != nil {
		return err
	}
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
		fmt.Println(v.FollowingID)
		user, err := u.Where(u.ID.Eq(v.FollowingID)).First()
		if err != nil {
			return nil, err
		}
		resp.UserList = append(resp.UserList, changes.UserRecord2friendResp(user))
	}

	return resp, nil

}
