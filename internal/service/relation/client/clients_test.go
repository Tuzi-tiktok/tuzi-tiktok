package client

import (
	"context"
	"fmt"
	"log"
	"testing"
	krelation "tuzi-tiktok/kitex/kitex_gen/relation"
	"tuzi-tiktok/utils"
)

// 得到粉丝列表
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

// 得到关注列表
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
	req.Token = "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjp7IlVJRCI6N30sImlzcyI6ImF1dGgtYXBpIiwiZXhwIjoxNjkzMDM4MzU0fQ.Ye8IGvUMq9_7ySrF9U3caUXn8x4iNMshwK3QAIugDVNlHrNFrnqAARtOzsZVqN7E1S-ISNO257CkZYprkaSVWA"
	req.ToUserId = 6
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
	req.Token = "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjp7IlVJRCI6N30sImlzcyI6ImF1dGgtYXBpIiwiZXhwIjoxNjkzMDM4MzU0fQ.Ye8IGvUMq9_7ySrF9U3caUXn8x4iNMshwK3QAIugDVNlHrNFrnqAARtOzsZVqN7E1S-ISNO257CkZYprkaSVWA"
	req.ToUserId = 6
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

// Get FriendList
func TestGetFriendList(t *testing.T) {
	var req krelation.RelationFriendListRequest
	req.Token = "123"
	req.UserId = int64(7)

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
