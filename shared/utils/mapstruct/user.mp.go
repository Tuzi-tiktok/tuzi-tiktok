package mapstruct

import (
	"tuzi-tiktok/gateway/biz/model/auth"
	pauth "tuzi-tiktok/kitex/kitex_gen/auth"
)

func ToUser(user *pauth.User) *auth.User {
	if user == nil {
		return nil
	}
	return &auth.User{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
}

func ToUserInfoResponse(p *pauth.UserInfoResponse) *auth.UserInfoResponse {
	if p == nil {
		return nil
	}
	return &auth.UserInfoResponse{
		StatusCode: p.StatusCode,
		StatusMsg:  p.StatusMsg,
		User:       ToUser(p.User),
	}
}

func ToUserRegisterResponse(p *pauth.UserRegisterResponse) *auth.UserRegisterResponse {
	if p == nil {
		return nil
	}
	return &auth.UserRegisterResponse{
		StatusCode: p.StatusCode,
		StatusMsg:  p.StatusMsg,
		UserId:     p.UserId,
		Token:      p.Token,
	}
}

func ToUserLoginResponse(p *pauth.UserLoginResponse) *auth.UserLoginResponse {
	if p == nil {
		return nil
	}
	return &auth.UserLoginResponse{
		StatusCode: p.StatusCode,
		StatusMsg:  p.StatusMsg,
		UserId:     p.UserId,
		Token:      p.Token,
	}
}
