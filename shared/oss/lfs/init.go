package lfs

import (
	"log"
	cfg "tuzi-tiktok/config"
	"tuzi-tiktok/oss"
)

const Name = "lfs"

// init Local File System
func init() {
	log.Print("lfs")

	//server := http.FileServer(http.Dir("."))
	//http.Handle(":8080", server)
	//err := http.ListenAndServe("", nil)
	//
	//if err != nil {
	//	panic(err)
	//}

	//	TODO Load Config Form CC
	_ = oss.Register(Name, initialize, func(k string) error {
		return cfg.VConfig.GetViper().UnmarshalKey(k, &c)
	})
}
