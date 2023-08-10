package main

import (
	"context"
	"gorm.io/gorm/clause"
	"time"
	"tuzi-tiktok/dao/model"
	"tuzi-tiktok/dao/query"
	auth "tuzi-tiktok/kitex/kitex_gen/auth"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/secret"
	"tuzi-tiktok/service/auth/tools"
)

// AuthInfoServiceImpl implements the last service interface defined in the IDL.
type AuthInfoServiceImpl struct{}

// Login implements the AuthInfoServiceImpl interface.
func (s *AuthInfoServiceImpl) Login(ctx context.Context, req *auth.UserLoginRequest) (resp *auth.UserLoginResponse, err error) {
	logger.Infof("login user: %s", req.Username)

	user := query.Q.User
	u, err := user.WithContext(ctx).Where(user.Username.Eq(req.Username)).Select().Find()
	if err != nil {
		logger.Errorf("failed to query user by username: %s, err: %v", req.Username, err)
		msg := tools.ServiceUnavailableMsg
		resp = &auth.UserLoginResponse{
			StatusCode: tools.ServiceUnavailable,
			StatusMsg:  &msg,
		}
		return
	}
	if len(u) == 0 {
		logger.Infof("user: %s not exist", req.Username)
		msg := tools.UserNotExistMsg
		resp = &auth.UserLoginResponse{
			StatusCode: tools.UserNotExist,
			StatusMsg:  &msg,
		}
		return
	}

	pwd := tools.HashPwd(req.Password)
	if pwd != u[0].Password {
		logger.Infof("user: %s password error", req.Username)
		msg := tools.WrongPwdMsg
		resp = &auth.UserLoginResponse{
			StatusCode: tools.WrongPwd,
			StatusMsg:  &msg,
		}
		return
	}

	uid := u[0].ID
	logger.Infof("user: %s, uid: %d", req.Username, uid)

	payload := secret.TokenPayload{
		UID: uid,
	}
	exp := time.Now().Add(time.Hour * 24 * 14)
	token, err := tools.NewToken(payload, exp)
	if err != nil {
		logger.Errorf("failed to generate token for user: %s, err: %v", req.Username, err)
		msg := tools.InternalServerErrorMsg
		resp = &auth.UserLoginResponse{
			StatusCode: tools.InternalServerError,
			StatusMsg:  &msg,
		}
		return
	}

	logger.Infof("login user: %s success", req.Username)
	msg := tools.SuccessMsg
	resp = &auth.UserLoginResponse{
		StatusCode: tools.Success,
		StatusMsg:  &msg,
		UserId:     uid,
		Token:      token,
	}

	return
}

// Register implements the AuthInfoServiceImpl interface.
func (s *AuthInfoServiceImpl) Register(ctx context.Context, req *auth.UserRegisterRequest) (resp *auth.UserRegisterResponse, err error) {
	logger.Infof("register user: %s", req.Username)

	//resp = &auth.UserRegisterResponse{}
	user := query.Q.User
	existedUser, err := user.WithContext(ctx).Where(user.Username.Eq(req.Username)).Select().Find()
	if err != nil {
		logger.Errorf("failed to query user by username: %s, err: %v", req.Username, err)
		msg := tools.ServiceUnavailableMsg
		resp = &auth.UserRegisterResponse{
			StatusCode: tools.ServiceUnavailable,
			StatusMsg:  &msg,
		}
		return
	}

	if len(existedUser) > 0 {
		logger.Infof("user: %s has been registered", req.Username)
		msg := tools.UserExistedMsg
		resp = &auth.UserRegisterResponse{
			StatusCode: tools.UserExisted,
			StatusMsg:  &msg,
		}
		return
	}

	pwd := tools.HashPwd(req.Password)
	newUser := model.User{
		Username: req.Username,
		Password: pwd,
	}
	err = user.WithContext(ctx).Clauses(clause.Returning{}).Create(&newUser)
	if err != nil {
		logger.Errorf("failed to save user: %s, err: %v", req.Username, err)
		msg := tools.ServiceUnavailableMsg
		resp = &auth.UserRegisterResponse{
			StatusCode: tools.ServiceUnavailable,
			StatusMsg:  &msg,
		}
		return
	}

	uid := newUser.ID
	logger.Infof("register user: %v success, uid: %v", req.Username, uid)

	payload := secret.TokenPayload{
		UID: uid,
	}
	exp := time.Now().Add(time.Hour * 24 * 14)
	token, err := tools.NewToken(payload, exp)
	if err != nil {
		logger.Errorf("failed to generate token for user: %s, err: %v", req.Username, err)
		msg := tools.InternalServerErrorMsg
		resp = &auth.UserRegisterResponse{
			StatusCode: tools.InternalServerError,
			StatusMsg:  &msg,
		}
		return
	}

	msg := tools.SuccessMsg
	resp = &auth.UserRegisterResponse{
		StatusCode: tools.Success,
		StatusMsg:  &msg,
		UserId:     uid,
		Token:      token,
	}

	return
}

// GetUserInfo implements the AuthInfoServiceImpl interface.
func (s *AuthInfoServiceImpl) GetUserInfo(ctx context.Context, req *auth.UserInfoRequest) (resp *auth.UserInfoResponse, err error) {
	logger.Infof("get user info, uid: %d", req.UserId)

	// check token
	c, err := secret.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("failed to parse token: %s, err: %v", req.Token, err)
		msg := tools.InvalidTokenMsg
		resp = &auth.UserInfoResponse{
			StatusCode: tools.InvalidToken,
			StatusMsg:  &msg,
		}
		return
	}

	logger.Infof("user %v seeking for user %v info", c.Payload.UID, req.UserId)

	// get user info
	user := query.Q.User
	u, err := user.WithContext(ctx).Where(user.ID.Eq(req.UserId)).Select().Find()
	if err != nil {
		logger.Errorf("failed to query user by uid: %d, err: %v", req.UserId, err)
		msg := tools.ServiceUnavailableMsg
		resp = &auth.UserInfoResponse{
			StatusCode: tools.ServiceUnavailable,
			StatusMsg:  &msg,
		}
		return
	}
	if len(u) == 0 {
		logger.Infof("user: %d not found", req.UserId)
		msg := tools.UserNotExistMsg
		resp = &auth.UserInfoResponse{
			StatusCode: tools.UserNotExist,
			StatusMsg:  &msg,
		}
		return
	}

	// check if the user is followed by the current user
	var isFollowed bool
	follow := query.Q.Relation
	f, err := follow.WithContext(ctx).Where(follow.FollowerID.Eq(c.Payload.UID), follow.FollowingID.Eq(req.UserId)).Select().Find()
	if err != nil {
		logger.Errorf("failed to query follow relation, follower: %d, following: %d, err: %v", c.Payload.UID, req.UserId, err)
		msg := tools.ServiceUnavailableMsg
		resp = &auth.UserInfoResponse{
			StatusCode: tools.ServiceUnavailable,
			StatusMsg:  &msg,
		}
		return
	}
	if len(f) > 0 {
		isFollowed = true
	} else {
		isFollowed = false
	}

	// check work count
	var workCount *int64
	work := query.Q.Video
	w, err := work.WithContext(ctx).Where(work.AuthorID.Eq(req.UserId)).Select().Count()
	if err != nil {
		logger.Errorf("failed to query work, user: %d, err: %v", req.UserId, err)
		workCount = nil
	}
	workCount = &w

	// check favorite count
	var favoriteCount *int64
	favorite := query.Q.Favorite
	fav, err := favorite.WithContext(ctx).Where(favorite.UID.Eq(req.UserId)).Select().Count()
	if err != nil {
		logger.Errorf("failed to query favorite, user: %d, err: %v", req.UserId, err)
		favoriteCount = nil
	}
	favoriteCount = &fav

	// check total number of favorites
	var totalFavorited *int64
	var count int64
	ws, err := work.WithContext(ctx).Where(work.AuthorID.Eq(req.UserId)).Select().Find()
	if err != nil {
		logger.Errorf("failed to query work, user: %d, err: %v", req.UserId, err)
		totalFavorited = nil
	} else {
		for _, v := range ws {
			count += v.FavoriteCount
		}
		totalFavorited = &count
	}

	logger.Infof("get user info success, uid: %d", req.UserId)
	msg := tools.SuccessMsg
	resp = &auth.UserInfoResponse{
		StatusCode: tools.Success,
		StatusMsg:  &msg,
		User: &auth.User{
			Id:              u[0].ID,
			Name:            u[0].Username,
			FollowCount:     &u[0].FollowCount,
			FollowerCount:   &u[0].FollowerCount,
			IsFollow:        isFollowed,
			Avatar:          u[0].Avatar,
			BackgroundImage: u[0].BackgroundImage,
			Signature:       u[0].Signature,
			TotalFavorited:  totalFavorited,
			WorkCount:       workCount,
			FavoriteCount:   favoriteCount,
		},
	}

	return
}

// TokenVerify implements the AuthInfoServiceImpl interface.
func (s *AuthInfoServiceImpl) TokenVerify(ctx context.Context, req *auth.TokenVerifyRequest) (resp *auth.TokenVerifyResponse, err error) {
	logger.Infof("verify token: %s", req.Token)

	token, err := secret.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("failed to parse token: %s, err: %v", req.Token, err)
		msg := tools.InvalidTokenMsg
		resp = &auth.TokenVerifyResponse{
			StatusCode: tools.InvalidToken,
			StatusMsg:  &msg,
		}
		return
	}

	logger.Infof("token: %s is valid", req.Token)
	msg := tools.SuccessMsg
	resp = &auth.TokenVerifyResponse{
		StatusCode: tools.Success,
		StatusMsg:  &msg,
		UserId:     token.Payload.UID,
	}

	return
}
