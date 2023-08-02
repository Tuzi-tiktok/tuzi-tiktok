// Code generated by Kitex v0.6.2. DO NOT EDIT.
package authinfoservice

import (
	server "github.com/cloudwego/kitex/server"
	auth "tuzi-tiktok/kitex/kitex_gen/auth"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler auth.AuthInfoService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}