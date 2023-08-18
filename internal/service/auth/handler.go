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
	consts "tuzi-tiktok/utils/consts/auth"
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
		return nil, err
	}
	if len(u) == 0 {
		logger.Infof("user: %s not exist", req.Username)
		resp = &auth.UserLoginResponse{
			StatusCode: consts.AuthUserNotExist,
			StatusMsg:  &consts.AuthUserNotExistMsg,
		}
		return resp, nil
	}

	pwd := tools.HashPwd(req.Password)
	if pwd != u[0].Password {
		logger.Infof("user: %s password error", req.Username)
		resp = &auth.UserLoginResponse{
			StatusCode: consts.AuthWrongPwd,
			StatusMsg:  &consts.AuthWrongPwdMsg,
		}
		return resp, nil
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
		return nil, err
	}

	logger.Infof("login user: %s success", req.Username)
	resp = &auth.UserLoginResponse{
		StatusCode: consts.AuthSucceed,
		StatusMsg:  &consts.AuthSucceedMsg,
		UserId:     uid,
		Token:      token,
	}
	return resp, nil
}

// Register implements the AuthInfoServiceImpl interface.
func (s *AuthInfoServiceImpl) Register(ctx context.Context, req *auth.UserRegisterRequest) (resp *auth.UserRegisterResponse, err error) {
	logger.Infof("register user: %s", req.Username)

	// TODO: 异步获取随机头像/头图/签名
	imgCh := make(chan string, 2)
	go func() {
		ava, err := tools.GetRandomImage()
		if err != nil {
			logger.Errorf("failed to get random avatar, err: %v", err)
			imgCh <- ""
		}
		imgCh <- ava
	}()
	go func() {
		img, err := tools.GetRandomImage()
		if err != nil {
			logger.Errorf("failed to get random background image, err: %v", err)
			imgCh <- ""
		}
		imgCh <- img
	}()
	strCh := make(chan string, 1)
	go func() {
		sig, err := tools.GetRandomSentence()
		if err != nil {
			logger.Errorf("failed to get random sentence, err: %v", err)
			strCh <- ""
		}
		strCh <- sig
	}()

	//resp = &auth.UserRegisterResponse{}
	user := query.Q.User
	existedUser, err := user.WithContext(ctx).Where(user.Username.Eq(req.Username)).Select().Find()
	if err != nil {
		logger.Errorf("failed to query user by username: %s, err: %v", req.Username, err)
		return nil, err
	}

	if len(existedUser) > 0 {
		logger.Infof("user: %s has been registered", req.Username)
		resp = &auth.UserRegisterResponse{
			StatusCode: consts.AuthUserExisted,
			StatusMsg:  &consts.AuthUserExistedMsg,
		}
		return resp, nil
	}

	ava := <-imgCh
	img := <-imgCh
	sig := <-strCh
	newUser := model.User{
		Username: req.Username,
		Password: tools.HashPwd(req.Password),
		// TODO: 随机添加一些头像/头图/签名
		Avatar:          &ava,
		BackgroundImage: &img,
		Signature:       &sig,
	}
	err = user.WithContext(ctx).Clauses(clause.Returning{}).Create(&newUser)
	if err != nil {
		logger.Errorf("failed to save user: %s, err: %v", req.Username, err)
		return nil, err
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
		return nil, err
	}

	resp = &auth.UserRegisterResponse{
		StatusCode: consts.AuthSucceed,
		StatusMsg:  &consts.AuthSucceedMsg,
		UserId:     uid,
		Token:      token,
	}
	return resp, nil
}

// GetUserInfo implements the AuthInfoServiceImpl interface.
func (s *AuthInfoServiceImpl) GetUserInfo(ctx context.Context, req *auth.UserInfoRequest) (resp *auth.UserInfoResponse, err error) {
	logger.Infof("get user info, uid: %d", req.UserId)

	// check token
	c, err := secret.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("failed to parse token: %s, err: %v", req.Token, err)
		resp = &auth.UserInfoResponse{
			StatusCode: consts.AuthInvalidToken,
			StatusMsg:  &consts.AuthInvalidTokenMsg,
		}
		return resp, nil
	}

	logger.Infof("user %v seeking for user %v info", c.Payload.UID, req.UserId)

	// get user info
	user := query.Q.User
	u, err := user.WithContext(ctx).Where(user.ID.Eq(req.UserId)).Select().Find()
	if err != nil {
		logger.Errorf("failed to query user by uid: %d, err: %v", req.UserId, err)
		return nil, err
	}
	if len(u) == 0 {
		logger.Infof("user: %d not found", req.UserId)
		resp = &auth.UserInfoResponse{
			StatusCode: consts.AuthUserNotExist,
			StatusMsg:  &consts.AuthUserNotExistMsg,
		}
		return resp, nil
	}

	// check if the user is followed by the current user
	var isFollowed bool
	follow := query.Q.Relation
	f, err := follow.WithContext(ctx).Where(follow.FollowerID.Eq(c.Payload.UID), follow.FollowingID.Eq(req.UserId)).Select().Find()
	if err != nil {
		logger.Errorf("failed to query follow relation, follower: %d, following: %d, err: %v", c.Payload.UID, req.UserId, err)
		return nil, err
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
	resp = &auth.UserInfoResponse{
		StatusCode: consts.AuthSucceed,
		StatusMsg:  &consts.AuthSucceedMsg,
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
	return resp, nil
}

// TokenVerify implements the AuthInfoServiceImpl interface.
func (s *AuthInfoServiceImpl) TokenVerify(ctx context.Context, req *auth.TokenVerifyRequest) (resp *auth.TokenVerifyResponse, err error) {
	logger.Infof("verify token: %s", req.Token)

	token, err := secret.ParseToken(req.Token)
	if err != nil {
		logger.Errorf("failed to parse token: %s, err: %v", req.Token, err)
		resp = &auth.TokenVerifyResponse{
			StatusCode: consts.AuthInvalidToken,
			StatusMsg:  &consts.AuthInvalidTokenMsg,
		}
		return resp, nil
	}

	logger.Infof("token: %s is valid", req.Token)
	resp = &auth.TokenVerifyResponse{
		StatusCode: consts.AuthSucceed,
		StatusMsg:  &consts.AuthSucceedMsg,
		UserId:     token.Payload.UID,
	}
	return resp, nil
}
