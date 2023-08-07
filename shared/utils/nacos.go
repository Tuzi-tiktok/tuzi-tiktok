package utils

import (
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"net"
	"os"
	. "tuzi-tiktok/config"
	"tuzi-tiktok/logger"
)

func NewNacosClientParam() vo.NacosClientParam {
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(Registration.NamespaceId),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithUpdateCacheWhenEmpty(true),
		constant.WithLogDir(os.TempDir()),
		constant.WithCacheDir(os.TempDir()),
		constant.WithLogStdout(false),
	)
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: Registration.Host,
			Port:   Registration.Port,
		},
	}
	return vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfigs,
	}
}

func NewNacosNamingClient(para vo.NacosClientParam) naming_client.INamingClient {
	cli, err := clients.NewNamingClient(para)
	if err != nil {
		panic(err)
	}
	return cli
}

type ExtOption client.Option

func NewClientOptions(c ...ExtOption) []client.Option {

	//namingClient := NewNacosNamingClient(NewNacosClientParam())
	//return []client.Option{
	//	client.WithResolver(resolver.NewNacosResolver(namingClient)),
	//}
	namingClient := NewNacosNamingClient(NewNacosClientParam())
	options := make([]client.Option, len(c))
	for i := range options {
		options[i] = client.Option(c[i])
	}
	return append(options, []client.Option{
		client.WithResolver(resolver.NewNacosResolver(namingClient)),
	}...)
}

func NewServerOptions(serverName string) []server.Option {
	port := RandomAvailablePort()
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port))
	logger.Infof("Service Port  %v %v", serverName, port)
	if err != nil {
		panic(err)
	}
	namingClient := NewNacosNamingClient(NewNacosClientParam())
	options := []server.Option{
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: serverName,
		}),
		server.WithRegistry(registry.NewNacosRegistry(namingClient)),
		server.WithServiceAddr(addr),
	}
	return options

}
