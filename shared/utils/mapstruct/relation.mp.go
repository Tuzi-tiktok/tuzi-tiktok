package mapstruct

import (
	"tuzi-tiktok/gateway/biz/model/auth"
	"tuzi-tiktok/gateway/biz/model/relation"
)
import krelation "tuzi-tiktok/kitex/kitex_gen/relation"

func ToFriendUser(k *krelation.FriendUser) *relation.FriendUser {
	if k == nil {
		return nil
	}
	return &relation.FriendUser{
		Id:              k.Id,
		Name:            k.Name,
		FollowCount:     k.FollowCount,
		FollowerCount:   k.FollowerCount,
		IsFollow:        k.IsFollow,
		Avatar:          k.Avatar,
		BackgroundImage: k.BackgroundImage,
		Signature:       k.Signature,
		TotalFavorited:  k.TotalFavorited,
		WorkCount:       k.WorkCount,
		FavoriteCount:   k.FavoriteCount,
		Message:         k.Message,
		MsgType:         k.MsgType,
	}
}

func ToRelationResponse(k *krelation.RelationResponse) *relation.RelationResponse {
	if k == nil {
		return nil
	}
	return &relation.RelationResponse{
		StatusCode: k.StatusCode,
		StatusMsg:  k.StatusMsg,
	}
}

func ToRelationFriendListResponse(k *krelation.RelationFriendListResponse) *relation.RelationFriendListResponse {
	if k == nil {
		return nil
	}
	userList := make([]*relation.FriendUser, len(k.UserList))
	for i := range k.UserList {
		userList[i] = ToFriendUser(k.UserList[i])
	}
	return &relation.RelationFriendListResponse{
		StatusCode: k.StatusCode,
		StatusMsg:  k.StatusMsg,
		UserList:   userList,
	}
}

func ToRelationFollowerListResponse(k *krelation.RelationFollowerListResponse) *relation.RelationFollowerListResponse {
	if k == nil {
		return nil
	}
	userList := make([]*auth.User, len(k.UserList))
	for i := range k.UserList {
		userList[i] = ToUser(k.UserList[i])
	}
	return &relation.RelationFollowerListResponse{
		StatusCode: k.StatusCode,
		StatusMsg:  k.StatusMsg,
		UserList:   userList,
	}
}

func ToRelationFollowListResponse(k *krelation.RelationFollowListResponse) *relation.RelationFollowListResponse {
	if k == nil {
		return nil
	}
	userList := make([]*auth.User, len(k.UserList))
	for i := range k.UserList {
		userList[i] = ToUser(k.UserList[i])
	}
	return &relation.RelationFollowListResponse{
		StatusCode: k.StatusCode,
		StatusMsg:  k.StatusMsg,
		UserList:   userList,
	}
}
