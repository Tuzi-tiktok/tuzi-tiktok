package cfg

import (
	"bytes"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"log"
	"net"
	"os"
)

// localConfig TODO 用于解决远程配置 只能修改进行覆盖不能删除
var localConfig map[string]interface{}

var (
	isRemoteMode bool
)

func init() {
	DetermineEnvironment()
	LoadLocalConfig()
	LoadRemoteConfig()
	ComponentInitialization()
	log.Println("- Load Completed")
}

func LoadLocalConfig() {
	v := viper.New()
	SetDefault(v)
	v.SetConfigName(DefaultConfigName)
	v.SetConfigType(DefaultConfigType)
	v.AddConfigPath(DefaultConfigPath)
	v.AddConfigPath(DetermineConfig())
	v.AddConfigPath(CandidateConfigPath)
	if isRemoteMode {
		return
	}
	if err := v.ReadInConfig(); err != nil {
		log.Printf("Can't Find Config %v.%v", DefaultConfigName, DefaultConfigType)
		panic(err)
	}

}

// LoadRemoteConfig Load Basic Config
func LoadRemoteConfig() {
	log.Println(" - Parsing Config For Registry")
	if err := VConfig.viper.UnmarshalKey(registryK, &Registration); err != nil {
		panic(err)
	}
	LookupRegistryHost()

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
	log.Println(" - Parsing Config From Remote Center")
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
		panic(err)
	}
	if err := VConfig.viper.MergeConfig(bytes.NewReader([]byte(c))); err != nil {
		log.Printf("Error Parsed, Check Your Config Syntax %v ", err)
		panic(err)
	}

	err = client.ListenConfig(vo.ConfigParam{
		Group:  Registration.Group,
		DataId: Registration.DataId,
		OnChange: func(namespace, group, dataId, data string) {
			VConfig.rLock.RLock()
			defer VConfig.rLock.RUnlock()
			if err := VConfig.viper.MergeConfig(bytes.NewReader([]byte(data))); err != nil {
				log.Printf("Update Config Error %v ", err)
				return
			}
			log.Println("Nocas Config Update")
		},
	})
	if err != nil {
		panic(err)
	}

}

// LookupRegistryHost To Fix docker container resolve error
func LookupRegistryHost() {
	hosts, err := net.LookupHost(Registration.Host)
	if err != nil {
		log.Println("Error of Dns Resolve this Host")
		panic(err)
	}
	Registration.Host = hosts[0]
}

func ComponentInitialization() {
	log.Println(VConfig.GetViper().AllKeys())

	log.Println(" - Parsing Config For Zap Logger")
	if err := VConfig.viper.UnmarshalKey(logK, &LoggerConfig); err != nil {
		panic(err)
	}
}
