// Code generated by Kitex v0.6.2. DO NOT EDIT.

package authinfoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	auth "tuzi-tiktok/kitex/kitex_gen/auth"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Login(ctx context.Context, Req *auth.UserLoginRequest, callOptions ...callopt.Option) (r *auth.UserLoginResponse, err error)
	Register(ctx context.Context, Req *auth.UserRegisterRequest, callOptions ...callopt.Option) (r *auth.UserRegisterResponse, err error)
	GetUserInfo(ctx context.Context, Req *auth.UserInfoRequest, callOptions ...callopt.Option) (r *auth.UserInfoResponse, err error)
	TokenVerify(ctx context.Context, Req *auth.TokenVerifyRequest, callOptions ...callopt.Option) (r *auth.TokenVerifyResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kAuthInfoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kAuthInfoServiceClient struct {
	*kClient
}

func (p *kAuthInfoServiceClient) Login(ctx context.Context, Req *auth.UserLoginRequest, callOptions ...callopt.Option) (r *auth.UserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, Req)
}

func (p *kAuthInfoServiceClient) Register(ctx context.Context, Req *auth.UserRegisterRequest, callOptions ...callopt.Option) (r *auth.UserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, Req)
}

func (p *kAuthInfoServiceClient) GetUserInfo(ctx context.Context, Req *auth.UserInfoRequest, callOptions ...callopt.Option) (r *auth.UserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, Req)
}

func (p *kAuthInfoServiceClient) TokenVerify(ctx context.Context, Req *auth.TokenVerifyRequest, callOptions ...callopt.Option) (r *auth.TokenVerifyResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.TokenVerify(ctx, Req)
}
