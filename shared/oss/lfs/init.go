package lfs

import (
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	cfg "tuzi-tiktok/config"
	"tuzi-tiktok/logger"
	"tuzi-tiktok/oss/internal/define"
)

const Name = "lfs"

var base string

// init Local File System
func init() {
	logger.Debug("lfs")
	//	TODO Load Config Form CC
	_ = define.Register(Name, initialize, func(k string) error {
		err := cfg.VConfig.GetViper().UnmarshalKey(k, &c)
		if err == nil {
			base, err = url.JoinPath(c.ExternalURL, c.Bucket)
		}
		return err
	})
}

type impl struct{}

func initialize() define.StorageTransmitter {
	if c.StoragePath == "" {
		c.StoragePath = os.TempDir()
	}
	if c.Bucket == "" {
		c.Bucket = "default"
	}
	logger.Info(c)
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
