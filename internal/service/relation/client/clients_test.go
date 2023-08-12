package client

import (
	"context"
	"fmt"
	"log"
	"testing"
	krelation "tuzi-tiktok/kitex/kitex_gen/relation"
	"tuzi-tiktok/utils"
)

// 得到粉丝的信息
func TestGetFollowerList(t *testing.T) {
	var req krelation.RelationFollowerListRequest
	req.UserId = 1
	req.Token = "123"
	resp := new(krelation.RelationFollowerListResponse)
	clients, err := utils.NewRelation()
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}

	ctx := context.TODO()
	resp, err = clients.GetFollowerList(ctx, &req)
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(resp)
}

// 得到关注人的信息
func TestGetFollowList(t *testing.T) {
	var req krelation.RelationFollowListRequest
	req.UserId = 3
	req.Token = "123"

	resp := new(krelation.RelationFollowListResponse)
	clients, err := utils.NewRelation()
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}

	ctx := context.TODO()
	resp, err = clients.GetFollowList(ctx, &req)
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}

	fmt.Println(resp)
}

// 关注
func TestFollowAction(t *testing.T) {
	var req krelation.RelationRequest
	req.Token = "123"
	req.ToUserId = 10
	// 1是关注 2是取消关注
	req.ActionType = 1

	resp := new(krelation.RelationResponse)
	clients, err := utils.NewRelation()
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}

	ctx := context.TODO()
	resp, err = clients.FollowAction(ctx, &req)
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}

	fmt.Println(resp)
}

// 取消关注
func TestUnFollowAction(t *testing.T) {
	var req krelation.RelationRequest
	req.Token = "123"
	req.ToUserId = 10
	// 1是关注 2是取消关注
	req.ActionType = 2

	resp := new(krelation.RelationResponse)
	clients, err := utils.NewRelation()
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}

	ctx := context.TODO()
	resp, err = clients.FollowAction(ctx, &req)
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}

	fmt.Println(resp)
}

func TestGetFriendList(t *testing.T) {
	var req krelation.RelationFriendListRequest
	req.Token = "123"
	req.UserId = int64(1)

	resp := new(krelation.RelationFriendListResponse)
	clients, err := utils.NewRelation()
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}

	ctx := context.TODO()
	resp, err = clients.GetFriendList(ctx, &req)
	if err != nil {
		t.Errorf("%v", err)
		t.Fail()
	}

	fmt.Println(resp)
}
