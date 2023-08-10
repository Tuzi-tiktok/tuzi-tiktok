package cfg

import (
	"os"
)

const srcEnv = "TUZI_SRC"

func DeterminePath() string {
	env, ok := os.LookupEnv(srcEnv)
	if ok {
		return env
	} else {
		return ""
	}
}
