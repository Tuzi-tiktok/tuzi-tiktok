package test

import (
	"testing"
	"tuzi-tiktok/secret"
)

const (
	testToken = `eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJQYXlsb2FkIjp7IlVJRCI6OTIyMzM3MjAzNjg1NDc3NTgwN30sImlzcyI6ImF1dGgtYXBpIiwiZXhwIjo0ODQ1MTk1MjUxfQ.Io6-QH0XfEJQmHrqAafQG3v9bGbJDQjZbMFEpFtcv9eOGHHoubJFG8nReDlaWsf-nIpVr2mN1ALUTXVwEtJS-A`
)

func TestParseToken(t *testing.T) {
	token, err := secret.ParseToken(testToken)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(token)
		t.Logf("User ID: %v", token.Payload.UID)
	}
}
