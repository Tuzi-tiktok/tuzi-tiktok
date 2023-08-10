package secret

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	cfg "tuzi-tiktok/config"
)

type TokenPayload struct {
	UID int64
}

type TokenClaims struct {
	Payload TokenPayload
	jwt.RegisteredClaims
}

const (
	jwtPubConfKey = "secret.JWTPublicKey"
)

var jWTPublicKey string

func init() {
	v := cfg.VConfig.GetViper()
	err := v.UnmarshalKey(jwtPubConfKey, &jWTPublicKey)
	if err != nil {
		panic(err)
	}

}

func ParseToken(token string) (claims TokenClaims, err error) {
	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseECPublicKeyFromPEM([]byte(jWTPublicKey))
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
