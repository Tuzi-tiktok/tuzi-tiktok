package test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
)

// 查询粉丝
func TestQueryFollower(t *testing.T) {
	r := query.Relation
	u := query.User
	followers, err := r.Where(r.FollowingID.Eq(1)).Find()
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}
	for _, follower := range followers {
		user, err := u.Where(u.ID.Eq(follower.FollowerID)).First()
		if err != nil {
			log.Printf(err.Error())
		}
		fmt.Printf("user:%v", user.Username)
		fmt.Println()
	}

}

// 查询关注
func TestQueryFollow(t *testing.T) {
	r := query.Relation
	u := query.User
	follows, err := r.Where(r.FollowerID.Eq(3)).Find()
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}
	for _, follow := range follows {
		user, err := u.Where(u.ID.Eq(follow.FollowingID)).First()
		if err != nil {
			log.Printf(err.Error())
		}
		fmt.Printf("user:%v", user.Username)
		fmt.Println()
	}

}

// 添加关注
func TestFollowAction(t *testing.T) {
	follower := 11
	following := 1
	r := query.Relation
	relation := model.Relation{FollowerID: int64(follower), FollowingID: int64(following)}
	ctx := context.TODO()
	err := r.WithContext(ctx).Create(&relation)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

}

// 取消关注
func TestUnFollowAction(t *testing.T) {
	follower := int64(1)
	following := int64(3)
	r := query.Relation
	ctx := context.TODO()
	result, err := r.WithContext(ctx).Where(r.FollowerID.Eq(follower), r.FollowingID.Eq(following)).Delete()
	if result.RowsAffected == 0 {
		log.Printf("删除失败 没有该记录")
		t.Fail()
	}
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

}

// 获取好友列表
func TestGetFriendList(t *testing.T) {
	r := query.Relation
	r2 := r.As("r2")
	ctx := context.TODO()
	uid := int64(1)
	result, err := r.WithContext(ctx).Select(r.FollowingID).LeftJoin(r2, r2.FollowerID.EqCol(r.FollowingID)).
		Where(r.FollowerID.EqCol(r2.FollowingID), r.FollowerID.Eq(uid)).Find()
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

	for _, v := range result {
		fmt.Println(v.FollowingID)
	}

}
