package it

//  Imitation SPI mechanism
import (
	"strings"
	cfg "tuzi-tiktok/config"
	"tuzi-tiktok/logger"
	. "tuzi-tiktok/oss"
	//  Deps List
	_ "tuzi-tiktok/oss/lfs"
	_ "tuzi-tiktok/oss/minio"
)

const ossK = "oss"

// init Load SPI Config
func init() {
	v := cfg.VConfig.GetViper()
	sm := v.GetStringMap(ossK)
	if len(sm) != 1 {
		const err = "Please Check Your Oss Config"
		logger.Errorf(err)
		panic(err)
	}
	var fk string
	for fk = range sm {
	}
	sk := strings.Join([]string{ossK, fk}, ".")
	it := ImplType(fk)
	if err := Candidates[it].B(sk); err != nil {
		panic(err)
	}
	logger.Debugf("%v %v", fk, sk)
	SetUsing(it)
}

// RequiredInit Semantic call / No practical application
func RequiredInit() {
	//	TODO Nothing
}
