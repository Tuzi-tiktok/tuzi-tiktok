package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/url"
	cfg "tuzi-tiktok/config"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/oss/internal/define"
)

const Name = "minio"

// 注意先后顺序
func init() {

	_ = define.Register(Name, initialize, func(k string) error {
		err := cfg.VConfig.GetViper().UnmarshalKey(k, &c)
		if err == nil {
			base, err = url.JoinPath(c.ExternalURL, c.Bucket)
		}
		return err
	})
	logger.Debug("minio")
}

type impl struct {
	*minio.Client
}

func initialize() define.StorageTransmitter {
	var (
		err    error
		client *minio.Client
	)
	//TODO FIX Host Resolve
	//hosts, err := net.LookupHost(c.Endpoint)
	//if err != nil {
	//	log.Println("Error of Dns Resolve this Host")
	//	panic(err)
	//}
	//c.Endpoint = hosts[0]

	// Initialize minio client object.
	client, err = minio.New(c.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.AccessKey, c.SecretKey, ""),
		Secure: false,
	})
	if err != nil {
		logger.Error(err)
		panic(err)
	}
	_, err = client.ListBuckets(context.Background())
	if err != nil {
		panic(err)
	}
	return &impl{client}
}
