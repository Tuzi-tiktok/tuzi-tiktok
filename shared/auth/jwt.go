package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type TokenPayload struct {
	UID int64
}

type TokenClaims struct {
	Payload TokenPayload
	jwt.RegisteredClaims
}

const (
	configName    = "jwt"
	configType    = "yaml"
	configPath    = "."
	jwtPubConfKey = "JWTPublicKey"

	tempConfigPath = `C:\Users\Admin\GolandProjects\tuzi-tiktok\shared\auth`
)

var JWTPublicKey string

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
	err = v.UnmarshalKey(jwtPubConfKey, &JWTPublicKey)
	if err != nil {
		panic(err)
	}
}

func ParseToken(token string) (claims TokenClaims, err error) {
	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseECPublicKeyFromPEM([]byte(JWTPublicKey))
	}, jwt.WithValidMethods([]string{"ES256"}))

	if err != nil {
		return
	}

	if claims, ok := t.Claims.(*TokenClaims); ok && t.Valid {
		return *claims, nil
	} else {
		return TokenClaims{}, fmt.Errorf("invalid token")
	}
}
