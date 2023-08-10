package test

import (
	"math"
	"testing"
	"time"
	"tuzi-tiktok/service/auth/tools"
	"tuzi-tiktok/utils/secret"
)

func TestNewToken(t *testing.T) {
	payload := secret.TokenPayload{
		UID: math.MaxInt64,
	}
	token, err := tools.NewToken(payload, time.Now().Add(time.Hour))
	if err != nil {
		t.Error(err)
	}
	tokenClaims, err := tools.ParseToken(token)
	if err != nil {
		t.Error(err)
	}
	if tokenClaims.Payload.UID != math.MaxInt64 {
		t.Error("token parse error")
	}

	payload = secret.TokenPayload{
		UID: 0,
	}
	token, err = tools.NewToken(payload, time.Now())
	if err != nil {
		t.Error(err)
	}
	tokenClaims, err = tools.ParseToken(token)
	if err == nil {
		t.Error("token should be expired")
	}
}
