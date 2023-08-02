package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	cfg "tuzi-tiktok/config"
	"tuzi-tiktok/oss"
)

const Name = "minio"

// 注意先后顺序
func init() {
	//	TODO Load Config Form CC
	_ = oss.Register(Name, initialize, func(k string) error {
		return cfg.VConfig.GetViper().UnmarshalKey(k, &c)
	})
	log.Print("minio")
}

func initialize() oss.StorageTransmitter {
	var (
		err    error
		client *minio.Client
	)
	// Initialize minio client object.
	client, err = minio.New(c.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKey, c.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}
	_, err = client.ListBuckets(context.Background())
	if err != nil {
		panic(err)
	}
	return &impl{client}
}
