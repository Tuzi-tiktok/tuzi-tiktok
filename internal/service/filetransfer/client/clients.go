package client

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	nmodel "github.com/nacos-group/nacos-sdk-go/model"
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
	err := utils.DefaultServiceSubscriber(utils.Transfer(), func(services []nmodel.SubscribeService, err error) {
		for _, s := range services {
			if s.Healthy && s.Enable {
				targetURL = fmt.Sprintf("http://%s:%d/", s.Ip, s.Port)
				break
			}
		}
		if len(services) == 0 {
			logger.Debugf("Service Candidate Is Empty")
			targetURL = ""
		} else {

			//logger.Debugf("Transfer Service Update Server Target URL %v", targetURL)

		}
	})
	if err != nil {
		logger.Error("ServiceSubscribe Error \n", err)
		panic(err)
	}
	logger.Debugf("Server Target URL  %v", targetURL)

}

type Transfer interface {
	Put(string, io.Reader) model.TransResult
}
type TransferImpl struct{}

func (t TransferImpl) Put(s string, reader io.Reader) (r model.TransResult) {
	if targetURL == "" {
		logger.Warnf("Service Candidate is Empty")
		r.Ok = false
		return
	}
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
