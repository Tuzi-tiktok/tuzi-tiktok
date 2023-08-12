package cfg

import (
	"os"
)

const (
	//APPSrcPath Source code location Simplified development, unnecessary
	APPSrcPath = "TUZI_SRC"
	// AppENVIRONMENT Application running environment
	AppENVIRONMENT = "TUZI_ENV"
	// APPConfigPath Path of the configuration file
	APPConfigPath = "TUZI_CFG"
	// APPSecretPath Path of the secret file
	APPSecretPath = "TUZI_SEC"
)

func DetermineSrcPath() string {
	env, ok := os.LookupEnv(APPSrcPath)
	if ok {
		return env
	} else {
		return ""
	}
}
func DetermineSecret() string {
	env, ok := os.LookupEnv(APPSecretPath)
	if ok {
		return env
	} else {
		return ""
	}
}
func DetermineConfig() string {
	env, ok := os.LookupEnv(APPConfigPath)
	if ok {
		return env
	} else {
		return ""
	}
}
