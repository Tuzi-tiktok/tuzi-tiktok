package redis_test

import (
	"context"
	"fmt"
	"testing"
	"tuzi-tiktok/rds"
)

func TestRedis(t *testing.T) {
	count := "15-16-new-message-count"
	t2 := "15-16-pre-message-time"

	v1, err := rds.IRC.Get(context.Background(), count).Int()
	if err != nil {
		panic(err)
	}
	v2, err := rds.IRC.Get(context.Background(), t2).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("get %v is %v", count, v1)
	fmt.Println()
	fmt.Printf("get %v is %v", t2, v2)
}
