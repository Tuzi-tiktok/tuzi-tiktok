package test

import (
	"math"
	"testing"
	"time"
	"tuzi-tiktok/auth"
	"tuzi-tiktok/service/auth/tools"
)

func TestNewToken(t *testing.T) {
	payload := auth.TokenPayload{
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

	payload = auth.TokenPayload{
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
