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
	srcConfigPath = `internal\service\auth\config`
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
	v.AddConfigPath(path.Join(cfg.DeterminePath(), srcConfigPath))
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.UnmarshalKey(secretConfigKey, &SecretConfig)
	if err != nil {
		panic(err)
	}
}
