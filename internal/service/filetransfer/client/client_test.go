package client

import (
	"github.com/segmentio/ksuid"
	"os"
	"testing"
	"tuzi-tiktok/logger"
)

func TestKsuid(t *testing.T) {
	k := ksuid.New()
	logger.Info(k.String())
}

func TestClient(t *testing.T) {
	transfer := NewTransfer()
	open, err := os.Open("E:\\Merge\\Tuzi-TikTok\\lfs\\tiktok\\1.png")
	if err != nil {
		t.Fatal(err)
	}
	put := transfer.Put("12anc#13.png", open)
	logger.Debugf("Res %v", put)
}
