package lfs

import (
	"io"
	"log"
	"sync"
	"tuzi-tiktok/oss"
)

var c config

var single = sync.Once{}

type config struct {
	Endpoint string
	Bucket   string
}
type impl struct{}

func (i *impl) Ping() error {
	// TODO Replace this logger
	log.Printf("I'm Impl LFS %v\n", c)
	return nil
}

func (i *impl) PutObject(reader io.Reader) (string, error) {
	return "", nil
}

func (i *impl) GetAddress(k string) string {
	return "http://phablet:9000/tiktok/" + k
}

func initialize() oss.StorageTransmitter {
	return &impl{}
}
