package test

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"os"
	"testing"
	. "tuzi-tiktok/config"
)

func TestConfig(t *testing.T) {
	log.Printf(LoggerConfig.Encoding)
	log.Printf("%#v", LoggerConfig)
}

func TestRemoteConfig(t *testing.T) {
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(Registration.NamespaceId),
		constant.WithTimeoutMs(5000),
		constant.WithLogLevel("debug"),
		constant.WithLogDir(os.TempDir()),
		constant.WithCacheDir(os.TempDir()),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogStdout(false),
	)
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: Registration.Host,
			Port:   Registration.Port,
		},
	}
	client, err := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfigs,
	})
	if err != nil {
		panic(err)
	}
	c, err := client.GetConfig(vo.ConfigParam{
		Group:  Registration.Group,
		DataId: Registration.DataId,
	})
	if err != nil {
		t.Error(c)
	}
	log.Println(c)
}
