package minio

import (
	"context"
	"github.com/bytedance/gopkg/lang/fastrand"
	"github.com/minio/minio-go/v7"
	"io"
	"log"
	"strconv"
)

var c config

type config struct {
	Endpoint  string
	Bucket    string
	AccessKey string
	SecretKey string
}
type impl struct {
	*minio.Client
}

func (i *impl) Ping() error {
	// TODO Replace this logger
	log.Printf("I'm Impl Minio %v\n", c)
	bs, err := i.ListBuckets(context.Background())
	// TODO  DEBUG
	if err == nil {
		log.Println("Ping Start")
		for idx := range bs {
			log.Println(bs[idx].Name)
		}
		log.Println("Ping End")
	}
	return err
}

func (i *impl) PutObject(reader io.Reader) (string, error) {
	ctx := context.Background()
	options := minio.PutObjectOptions{
		ContentType: "image/png",
		PartSize:    1024 * 1024 * 5,
	}
	oName := strconv.Itoa(fastrand.Int()) + ".png"
	info, err := i.Client.PutObject(ctx, c.Bucket, oName, reader, -1, options)
	return info.Key, err
}

func (i *impl) GetAddress(k string) string {
	return "http://phablet:9000/tiktok/" + k
}
