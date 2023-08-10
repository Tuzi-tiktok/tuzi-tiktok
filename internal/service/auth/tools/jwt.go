package tools

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"tuzi-tiktok/secret"
	"tuzi-tiktok/service/auth/config"
	"tuzi-tiktok/utils"
)

func NewToken(tokenPayload secret.TokenPayload, expAt time.Time) (string, error) {
	claims := secret.TokenClaims{
		Payload: tokenPayload,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    utils.Auth(),
			ExpiresAt: jwt.NewNumericDate(expAt), // TODO: token should be refreshed daily
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	key, _ := jwt.ParseECPrivateKeyFromPEM([]byte(config.SecretConfig.JWTPrivateKey))
	token, err := t.SignedString(key)
	return token, err
}

func ParseToken(token string) (claims secret.TokenClaims, err error) {
	t, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwt.ParseECPublicKeyFromPEM([]byte(config.SecretConfig.JWTPublicKey))
	}, jwt.WithValidMethods([]string{"ES256"}))

	if err != nil {
		return
	}

	if claims, ok := t.Claims.(*secret.TokenClaims); ok && t.Valid {
		return *claims, nil
	} else {
		return secret.TokenClaims{}, fmt.Errorf("invalid token")
	}
}
