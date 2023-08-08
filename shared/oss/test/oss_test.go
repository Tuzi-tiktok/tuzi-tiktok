package test

import (
	"bytes"
	"os"
	"path"
	"sync"
	"testing"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/oss"
)

var testFilePath = "E:\\Merge\\Tuzi-TikTok\\lfs\\tiktok\\2.png"

func TestOss(t *testing.T) {
	logger.Debug(oss.Ping())
}

func TestPutObject(t *testing.T) {
	err := oss.PutObject("Hello/World.txt", bytes.NewReader([]byte("Hello,World!")))
	if err != nil {
		logger.Warn(err)
		t.Fail()
	} else {
		logger.Debug("Test Success")
	}
}

func TestLFSGetAddress(t *testing.T) {
	file, err := os.Open(testFilePath)
	if err != nil {
		panic(err)
	}
	k := "a/a.png"
	err = oss.PutObject(k, file)
	if err != nil {
		t.Fatal(err)
	}
	logger.Debug(oss.GetAddress(k))
	g := sync.WaitGroup{}
	g.Add(1)
	g.Wait()
}

func TestMinioGetAddress(t *testing.T) {
	file, err := os.Open(testFilePath)
	if err != nil {
		panic(err)
	}
	k := "a/a.png"
	err = oss.PutObject(k, file)
	if err != nil {
		t.Fatal(err)
	}
	logger.Debug(oss.GetAddress(k))
}

func TestExt(t *testing.T) {
	ext := path.Ext("a.png")
	logger.Info(ext)
}
