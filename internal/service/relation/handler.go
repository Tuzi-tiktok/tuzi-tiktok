package main

import (
	"context"
	"fmt"
	"log"
	"tuzi-tiktok/dao/query"
	relation "tuzi-tiktok/kitex/kitex_gen/relation"
	"tuzi-tiktok/service/relation/dao"
	"tuzi-tiktok/utils/changes"
	consts "tuzi-tiktok/utils/consts/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// FollowAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowAction(ctx context.Context, req *relation.RelationRequest) (resp *relation.RelationResponse, err error) {
	resp = new(relation.RelationResponse)
	if req.ActionType == 1 {
		//关注
		err := dao.FollowAction(1, req.ToUserId)
		if err != nil {
			resp.StatusCode = consts.RelationActionFailed
			resp.StatusMsg = &consts.RelationActionFailedMsg
			return resp, err
		}
	} else {
		//取消关注
		err := dao.UnFollowAction(1, req.ToUserId)
		if err != nil {
			resp.StatusMsg = &consts.RelationFollowFailedMsg
			resp.StatusMsg = &consts.RelationActionFailedMsg
			return resp, err
		}

	}
	resp.StatusCode = consts.RelationSucceed
	resp.StatusMsg = &consts.RelationSucceedMsg
	return
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {

	resp = new(relation.RelationFollowListResponse)
	r := query.Relation
	u := query.User
	follows, err := r.Where(r.FollowerID.Eq(req.UserId)).Find()
	if err != nil {
		log.Printf(err.Error())
		resp.StatusMsg = &consts.RelationFollowFailedMsg
		resp.StatusMsg = &consts.RelationActionFailedMsg
		return resp, err
	}
	for _, follow := range follows {
		user, err := u.Where(u.ID.Eq(follow.FollowingID)).First()
		if err != nil {
			log.Printf(err.Error())
			resp.StatusMsg = &consts.RelationFollowFailedMsg
			resp.StatusMsg = &consts.RelationActionFailedMsg
			return resp, err
		}
		resp.UserList = append(resp.UserList, changes.UserRecord2userResp(user))
	}

	resp.StatusCode = consts.RelationSucceed
	resp.StatusMsg = &consts.RelationSucceedMsg
	return
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {

	resp = new(relation.RelationFollowerListResponse)
	r := query.Relation
	u := query.User

	followers, err := r.Where(r.FollowingID.Eq(req.UserId)).Find()
	if err != nil {
		log.Println(err.Error())
		resp.StatusMsg = &consts.RelationFollowFailedMsg
		resp.StatusMsg = &consts.RelationActionFailedMsg
		return resp, err
	}
	for _, follower := range followers {
		user, err := u.Where(u.ID.Eq(follower.FollowerID)).First()
		if err != nil {
			log.Printf(err.Error())
			resp.StatusMsg = &consts.RelationFollowFailedMsg
			resp.StatusMsg = &consts.RelationActionFailedMsg
			return resp, err
		}
		resp.UserList = append(resp.UserList, changes.UserRecord2userResp(user))
	}

	resp.StatusCode = consts.RelationSucceed
	resp.StatusMsg = &consts.RelationSucceedMsg
	return
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {

	fmt.Println(req.UserId)
	fmt.Println(req.Token)
	resp = new(relation.RelationFriendListResponse)
	resp, err = dao.GetFriendList(req.UserId)
	if err != nil {
		resp.StatusMsg = &consts.RelationFollowFailedMsg
		resp.StatusMsg = &consts.RelationActionFailedMsg
		return nil, err
	}
	resp.StatusCode = consts.RelationSucceed
	resp.StatusMsg = &consts.RelationSucceedMsg
	return
}
