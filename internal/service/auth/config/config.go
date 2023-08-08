package config

import (
	"github.com/spf13/viper"
)

const (
	configName = "secret"
	configType = "yaml"
	configPath = "."

	tempConfigPath = "C:\\Users\\Admin\\GolandProjects\\tuzi-tiktok\\internal\\service\\auth\\config"
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
	v.AddConfigPath(tempConfigPath)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.UnmarshalKey(secretConfigKey, &SecretConfig)
	if err != nil {
		panic(err)
	}
}
