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
	"github.com/nacos-group/nacos-sdk-go/model"
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

func DefaultNacosClient() naming_client.INamingClient {
	return NewNacosNamingClient(NewNacosClientParam())
}

func DefaultServiceSelector(serverName string) ([]model.Instance, error) {
	cli := DefaultNacosClient()
	return cli.SelectInstances(vo.SelectInstancesParam{
		ServiceName: serverName,
		HealthyOnly: true,
	})
}
func DefaultServiceSubscriber(serverName string, callback func(services []model.SubscribeService, err error)) error {
	cli := DefaultNacosClient()
	err := cli.Subscribe(&vo.SubscribeParam{
		ServiceName:       serverName,
		SubscribeCallback: callback,
	})
	return err
}

func DefaultServerRegister(serverName string, port uint64) error {
	cli := DefaultNacosClient()
	_, err := cli.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          GetLocalAddr(),
		Port:        port,
		Weight:      10,
		ServiceName: serverName,
		Healthy:     true,
		Enable:      true,
		Ephemeral:   true,
	})
	return err
}

// ExtOption Kitex 适用
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
	logger.Infof("%v Service Port is: %v", serverName, port)
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
	server.RegisterShutdownHook(func() {
		logger.Infof("Service %v Start Shutdown", serverName)
	})
	return options

}
