package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"io"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/service/filetransfer/model"
	"tuzi-tiktok/utils"
)

var (
	cli       *client.Client
	targetURL string
)

func init() {
	var err error
	cli, err = client.NewClient()
	if err != nil {
		panic(err)
	}
	InitConfig()
}
func InitConfig() {
	instances, err := utils.DefaultServerSelector(utils.Transfer())
	if err != nil {
		logger.Error("Service discovery failure")
		panic(err)
	}
	if len(instances) == 0 {
		err := errors.New("There is no suitable service ")
		logger.Error(err)
		panic(err)

	}
	ip, port := instances[0].Ip, instances[0].Port
	targetURL = fmt.Sprintf("http://%s:%d/", ip, port)
	logger.Debugf("Server Target URL  %v", targetURL)
}

type Transfer interface {
	Put(string, io.Reader) model.TransResult
}
type TransferImpl struct{}

func (t TransferImpl) Put(s string, reader io.Reader) (r model.TransResult) {
	ctx := context.TODO()
	req, resp := &protocol.Request{}, &protocol.Response{}
	req.SetFileReader("data", s, reader)
	req.SetRequestURI(targetURL)
	req.SetMethod(consts.MethodPut)
	err := cli.Do(ctx, req, resp)
	if err != nil {
		r.Ok = false
		logger.Error("Transfer Client Do Error ", err)
		return
	}
	err = sonic.Unmarshal(resp.Body(), &r)
	if err != nil {
		r.Ok = false
		logger.Error("Json Unmarshal Error ", err)
		return
	}
	return
}

func NewTransfer() Transfer {
	return TransferImpl{}
}
