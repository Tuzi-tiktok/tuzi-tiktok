package rds

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"testing"
)

// TestClient https://redis.uptrace.dev/zh/
func TestClient(t *testing.T) {

	ping := IRC.Ping(context.TODO())
	log.Println(ping.Result())

}

func TestCfgKeys(t *testing.T) {
	result, err := IRC.Keys(context.TODO(), "*").Result()
	if err != nil {
		t.Error(err)
	}
	log.Println(result)
}
func TestNilHost(t *testing.T) {
	client := redis.NewClient(&redis.Options{})
	result, err := client.Keys(context.TODO(), "*").Result()
	if err != nil {
		t.Error(err)
	}
	log.Println(result)
}
