package config

import (
	"github.com/spf13/viper"
	"path"
	cfg "tuzi-tiktok/config"
)

const (
	configName    = "secret"
	configType    = "yaml"
	configPath    = "."
	srcConfigPath = `C:\Users\Administrator\Desktop\tuzi-tiktok\internal\service\auth\config`
)

const (
	secretConfigKey = "secret"
)

var SecretConfig secretConfig

type secretConfig struct {
	PwdSalt       string
	JWTPrivateKey string
	JWTPublicKey  string
}

func init() {
	v := viper.New()
	v.SetConfigName(configName)
	v.SetConfigType(configType)
	v.AddConfigPath(configPath)
	v.AddConfigPath(cfg.DetermineSecret())
	v.AddConfigPath(path.Join(cfg.DetermineSrcPath(), srcConfigPath))
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.UnmarshalKey(secretConfigKey, &SecretConfig)
	if err != nil {
		panic(err)
	}
}
