package lfs

import (
	"net/http"
	"os"
	"path"
	"strings"
	cfg "tuzi-tiktok/config"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/oss"
)

const Name = "lfs"

// init Local File System
func init() {
	logger.Debug("lfs")
	//	TODO Load Config Form CC
	_ = oss.Register(Name, initialize, func(k string) error {
		return cfg.VConfig.GetViper().UnmarshalKey(k, &c)
	})
}

type impl struct{}

func initialize() oss.StorageTransmitter {
	if c.StoragePath == "" {
		c.StoragePath = os.TempDir()
	}
	if c.Bucket == "" {
		c.Bucket = "default"
	}
	logger.Info(c.StoragePath)
	go initFileServer()

	return &impl{}
}

// initFileServer
// TODO Path Match [ /{name}/ , /{name} ] All files in that folder are listed
//
//	http.Handle("/static", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
//		http.NotFound(writer, request)
//	}))
func initFileServer() {
	pref := strings.Join([]string{"/", c.Bucket, "/"}, "")
	// Concat String
	//buf := append(append([]byte{'/'}, c.Bucket...), []byte{'/'}...)
	fp := path.Join(c.StoragePath, c.Bucket)
	http.Handle(pref, http.StripPrefix(pref, http.FileServer(http.Dir(fp))))
	logger.Info("Local File Server Starting")
	err := http.ListenAndServe(c.Endpoint[strings.IndexRune(c.Endpoint, ':'):], nil)
	if err != nil {
		logger.Error(err)
		panic(err)
	}
}
