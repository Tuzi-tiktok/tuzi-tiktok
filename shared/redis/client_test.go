package rds

import (
	"context"
	"log"
	"testing"
)

// TestClient https://redis.uptrace.dev/zh/
func TestClient(t *testing.T) {

	ping := IRC.Ping(context.TODO())
	log.Println(ping.Result())

}
