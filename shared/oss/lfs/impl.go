package lfs

import (
	"io"
	"net/url"
	"os"
	"path"
	"tuzi-tiktok/logger"
)

var c config

type config struct {
	Endpoint    string
	Bucket      string
	StoragePath string
	ExternalURL string
}

func (i *impl) Ping() error {
	// TODO Replace this logger
	logger.Debug("I'm Impl LFS ", c)
	return nil
}

// bufSz 2MB Buffer Size
const bufSz = 1024 * 1024 * 2

func (i *impl) PutObject(k string, reader io.Reader) error {
	dir, f := path.Split(k)
	dir = path.Join(c.StoragePath, c.Bucket, dir)
	err := os.MkdirAll(dir, 0666)
	if err != nil {
		return err
	}
	mode := os.O_RDWR | os.O_CREATE | os.O_TRUNC
	pth := path.Join(dir, f)
	file, err := os.OpenFile(pth, mode, 0666)
	defer file.Close()
	length, err := io.CopyBuffer(file, reader, make([]byte, bufSz))
	if err == nil {
		logger.Debugf("LFS Written %v %v", length, file.Name())
	}
	return err
}

func (i *impl) GetAddress(k string) string {
	if k == "" {
		return ""
	}
	result, _ := url.JoinPath(base, k)
	return result
}
