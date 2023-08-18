package cfg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	// APPRemoteMeta Remote configuration central metadata
	APPRemoteMeta = "TUZI_RFG"
	//APPSourcePath Source code location Simplified development, unnecessary
	APPSourcePath = "TUZI_SRC"
	// AppENVIRONMENT Application running environment
	AppENVIRONMENT = "TUZI_ENV"
	// APPConfigPath Path of the configuration file
	APPConfigPath = "TUZI_CFG"
	// APPSecretPath Path of the secret file
	APPSecretPath = "TUZI_SEC"
)

func DetermineEnvironment() {
	rfg := DetermineRFG()
	if rfg != "" {
		isRemoteMode = true
		ConfigEnv = "Remote Mode"
		ResolveRFG(rfg)
	} else {
		env, ok := os.LookupEnv(AppENVIRONMENT)
		if ok {
			ConfigEnv = env
			DefaultConfigName = fmt.Sprintf("%v-%v", DefaultConfigName, env)
		}
		CandidateConfigPath = DetermineSrcPath()
	}
	log.Println(fmt.Sprintf("- Load %v Config", strings.ToUpper(ConfigEnv)))
}

func DetermineSrcPath() string {
	env, ok := os.LookupEnv(APPSourcePath)
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

func DetermineRFG() string {
	env, ok := os.LookupEnv(APPRemoteMeta)
	if ok {
		return env
	} else {
		return ""
	}
}

func ResolveRFG(r string) {
	pattern := strings.Split(r, "?")
	if len(pattern) != 2 {
		panic(errors.New("Meta Rfg is Incorrect, Check your TUZI_RFG "))
	}
	addr, param := strings.Split(pattern[0], ":"), strings.Split(pattern[1], "&")
	if len(addr) != 2 {
		panic(errors.New("Meta Rfg is Incorrect, Check your TUZI_RFG "))
	}

	p, err := strconv.Atoi(addr[1])
	if err != nil {
		panic(errors.New("Meta Rfg is Incorrect, Check your TUZI_RFG :" + err.Error()))
	}
	Registration.Host = addr[0]
	Registration.Port = uint64(p)

	for i := range param {
		pair := strings.Split(param[i], "=")
		if len(addr) != 2 {
			panic(errors.New("Meta Rfg is Incorrect, Check your TUZI_RFG "))
		}
		switch pair[0] {
		case "n":
			Registration.NamespaceId = pair[1]
		case "g":
			Registration.Group = pair[1]
		case "d":
			Registration.DataId = pair[1]
		default:
			log.Println("Error No Match Flag  Check your TUZI_RFG")
		}
	}
	if Registration.Group == "" {
		Registration.Group = "DEFAULT_GROUP"
	}

}
