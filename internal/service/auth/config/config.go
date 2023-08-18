package config

import (
	"github.com/spf13/viper"
	"log"
	"path"
	cfg "tuzi-tiktok/config"
)

const (
	configName    = "secret"
	configType    = "yaml"
	configPath    = "."
	srcConfigPath = `internal/service/auth/config`
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

	if rfg := cfg.DetermineRFG(); rfg != "" {
		extendedRemote()
		return
	}

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
	log.Println("- Load Local Secret Key")
}
func extendedRemote() {
	err := cfg.VConfig.GetViper().UnmarshalKey(secretConfigKey, &SecretConfig)
	if err != nil {
		panic(err)
	}
	log.Println("- Load Remote Secret Key")
}
