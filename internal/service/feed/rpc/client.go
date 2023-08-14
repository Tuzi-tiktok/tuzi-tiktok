package rpc

import (
	"github.com/cloudwego/kitex/client"
	auth "tuzi-tiktok/kitex/kitex_gen/auth/authinfoservice"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/utils"
)

type Client struct {
	Auth auth.Client
}

var C = &Client{}

func init() {
	var err error
	// todo: 从nacos获取auth服务
	//r, err := resolver.NewDefaultNacosResolver()
	//if err != nil {
	//	panic(err)
	//}
	//client, err := echo.NewClient("echo", client.WithResolver(r))
	//if err != nil {
	//	log.Fatal(err)
	//}
	C.Auth, err = auth.NewClient(utils.Auth(), client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		logger.Error("auth service init failed")
	}
}

//func (c *Client) GetUserInfo(id int) {
//	c.Auth.GetUserInfo(context.Background(), id, callopt.WithConnectTimeout(3*time.Second))
//}
