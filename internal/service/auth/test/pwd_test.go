package test

import (
	"testing"
	"tuzi-tiktok/service/auth/tools"
)

func TestHashPwd(t *testing.T) {
	pwd := "123456"
	hashedPwd := tools.HashPwd(pwd)
	if len(hashedPwd) != 64 {
		t.Errorf("HashPwd length error, got: %d, want: %d", len(hashedPwd), 32)
	}
	if hashedPwd != "cae3ade7902020e665d86dfff89708098004d02a81712f2498f2c29d61f66b6e" {
		t.Errorf("HashPwd value error, got: %s", hashedPwd)
	}
}
