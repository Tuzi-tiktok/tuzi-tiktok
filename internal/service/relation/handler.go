package main

import (
	"context"
	"tuzi-tiktok/dao/query"
	relation "tuzi-tiktok/kitex/kitex_gen/relation"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/secret"
	"tuzi-tiktok/service/relation/dao"
	"tuzi-tiktok/utils/changes"
	consts "tuzi-tiktok/utils/consts/relation"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// FollowAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) FollowAction(ctx context.Context, req *relation.RelationRequest) (resp *relation.RelationResponse, err error) {

	resp = new(relation.RelationResponse)
	// check token & get uid
	claims, err := secret.ParseToken(req.Token)
	if err != nil {
		logger.Infof("failed to parse token, err: %v", err)
		resp.StatusCode = consts.RelationTokenParseFailed
		resp.StatusMsg = &consts.RelationTokenParseFailedMsg
		return resp, nil
	}
	uid := claims.Payload.UID
	//can not focus on yourself
	if uid == req.ToUserId {
		logger.Infof("user:%d can not focus on yourself", uid)
		resp.StatusCode = consts.RelationFollowFailed
		resp.StatusMsg = &consts.RelationFollowFailedMsg
		return resp, nil
	}

	logger.Infof("user:%d follow action user:%d", uid, req.ToUserId)
	if req.ActionType == 1 {
		//关注
		resp, err := dao.FollowAction(uid, req.ToUserId)
		if err != nil {
			logger.Infof("failed to follow action, err: %v", err)
			return nil, err
		}
		if resp != nil {
			return resp, nil
		}
	} else if req.ActionType == 2 {
		//取消关注
		err := dao.UnFollowAction(uid, req.ToUserId)
		if err != nil {
			logger.Infof("failed to unfollow action, err: %v", err)
			return nil, err
		}

	} else {
		//unknown action
		logger.Infof("relation unknown action")
		resp.StatusCode = consts.RelationUnKnownAction
		resp.StatusMsg = &consts.RelationUnKnownActionMsg
		return resp, nil
	}

	resp.StatusCode = consts.RelationSucceed
	resp.StatusMsg = &consts.RelationSucceedMsg
	return
}

// GetFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	logger.Infof("get user:%d follow list", req.UserId)
	// check token
	_, err = secret.ParseToken(req.Token)
	if err != nil {
		logger.Infof("failed to parse token, err: %v", err)
		resp.StatusCode = consts.RelationTokenParseFailed
		resp.StatusMsg = &consts.RelationTokenParseFailedMsg
		return resp, nil
	}

	resp = new(relation.RelationFollowListResponse)
	r := query.Relation
	u := query.User
	follows, err := r.Where(r.FollowerID.Eq(req.UserId)).Find()
	if err != nil {
		logger.Infof("failed to get follow, err: %v", err)
		return nil, err
	}
	for _, follow := range follows {
		user, err := u.Where(u.ID.Eq(follow.FollowingID)).First()
		if err != nil {
			logger.Infof("failed to query follow details, err: %v", err)
			return nil, err
		}
		userResp, err := changes.UserRecord2userResp(user)
		if err != nil {
			return nil, err
		}
		resp.UserList = append(resp.UserList, userResp)
	}

	resp.StatusCode = consts.RelationSucceed
	resp.StatusMsg = &consts.RelationSucceedMsg
	return
}

// GetFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	logger.Infof("get user:%d follower list", req.UserId)
	// check token
	_, err = secret.ParseToken(req.Token)
	if err != nil {
		logger.Infof("failed to parse token, err: %v", err)
		resp.StatusCode = consts.RelationTokenParseFailed
		resp.StatusMsg = &consts.RelationTokenParseFailedMsg
		return resp, nil
	}
	resp = new(relation.RelationFollowerListResponse)
	r := query.Relation
	u := query.User

	followers, err := r.Where(r.FollowingID.Eq(req.UserId)).Find()
	if err != nil {
		logger.Infof("failed to get followers, err: %v", err)
		return nil, err
	}
	for _, follower := range followers {
		user, err := u.Where(u.ID.Eq(follower.FollowerID)).First()
		if err != nil {
			logger.Infof("failed to get follower details, err: %v", err)
			return nil, err
		}
		userResp, err := changes.UserRecord2userResp(user)
		if err != nil {
			return nil, err
		}
		resp.UserList = append(resp.UserList, userResp)
	}

	resp.StatusCode = consts.RelationSucceed
	resp.StatusMsg = &consts.RelationSucceedMsg
	return
}

// GetFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) GetFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	logger.Infof("get user:%d friend list", req.UserId)
	// check token & get uid
	_, err = secret.ParseToken(req.Token)
	if err != nil {
		logger.Infof("failed to parse token, err: %v", err)
		return resp, nil
	}

	resp = new(relation.RelationFriendListResponse)
	resp, err = dao.GetFriendList(req.UserId)
	if err != nil {
		logger.Infof("failed to get friend list, err: %v", err)
		return nil, err
	}
	resp.StatusCode = consts.RelationSucceed
	resp.StatusMsg = &consts.RelationSucceedMsg
	return
}
