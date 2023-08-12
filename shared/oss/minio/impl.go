package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"mime"
	"net/url"
	"path"
	"tuzi-tiktok/logger"
)

var c config

type config struct {
	Endpoint    string
	Bucket      string
	AccessKey   string
	SecretKey   string
	ExternalURL string
}

func (i *impl) Ping() error {
	// TODO Replace this logger
	logger.Debug("I'm Impl Minio %v\n", c)
	bs, err := i.ListBuckets(context.Background())
	// TODO  DEBUG
	if err == nil {
		logger.Debug("Ping Start")
		for idx := range bs {
			logger.Debugf("%v", bs[idx])
		}
		logger.Debug("Ping End")
	}
	return err
}

func (i *impl) PutObject(k string, reader io.Reader) error {
	ctx := context.Background()
	//	DetectContentType Find By Extension
	ct := "application/octet-stream"
	if tp := mime.TypeByExtension(path.Ext(k)); tp != "" {
		ct = tp
	}
	options := minio.PutObjectOptions{
		ContentType: ct,
		PartSize:    1024 * 1024 * 5,
	}
	info, err := i.Client.PutObject(ctx, c.Bucket, k, reader, -1, options)
	logger.Debug(info.Key, info.Size)
	return err
}

var base string

func (i *impl) GetAddress(k string) string {
	if k == "" {
		return ""
	}
	result, _ := url.JoinPath(base, k)
	return result
}
