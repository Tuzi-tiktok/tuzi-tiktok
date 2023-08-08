package tools

import (
	"crypto/sha256"
	"encoding/hex"
	"tuzi-tiktok/service/auth/config"
)

func HashPwd(password string) string {
	md5Ctx := sha256.New()
	// TODO: salt value should be different for each user, if you need to improve security
	md5Ctx.Write([]byte(password + config.SecretConfig.PwdSalt))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}
