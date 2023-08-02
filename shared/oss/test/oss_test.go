package test

import (
	"fmt"
	"testing"
	"tuzi-tiktok/oss"
	_ "tuzi-tiktok/oss/it"
)

func TestOss(t *testing.T) {
	//defer func() {
	//	if r := recover(); r != nil {
	//		t.Error(r)
	//		t.Fail()
	//	}
	//}()
	//oss.SetUsing("minio")
	fmt.Println(oss.Ping())
}
