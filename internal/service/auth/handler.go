package main

import (
	"context"
	auth "tuzi-tiktok/kitex/kitex_gen/auth"
)

// AuthInfoServiceImpl implements the last service interface defined in the IDL.
type AuthInfoServiceImpl struct{}

// Login implements the AuthInfoServiceImpl interface.
func (s *AuthInfoServiceImpl) Login(ctx context.Context, req *auth.UserLoginRequest) (resp *auth.UserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Register implements the AuthInfoServiceImpl interface.
func (s *AuthInfoServiceImpl) Register(ctx context.Context, req *auth.UserRegisterRequest) (resp *auth.UserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserInfo implements the AuthInfoServiceImpl interface.
func (s *AuthInfoServiceImpl) GetUserInfo(ctx context.Context, req *auth.UserInfoRequest) (resp *auth.UserInfoResponse, err error) {
	// TODO: Your code here...
	return
}

// TokenVerify implements the AuthInfoServiceImpl interface.
func (s *AuthInfoServiceImpl) TokenVerify(ctx context.Context, req *auth.TokenVerifyRequest) (resp *auth.TokenVerifyResponse, err error) {
	// TODO: Your code here...
	return
}
