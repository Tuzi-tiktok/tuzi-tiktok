package oss

import (
	"io"
	"tuzi-tiktok/logger"
)

var (
	Candidates   = make(map[ImplType]*CI, 2)
	sTransmitter StorageTransmitter
)

func init() {
	logger.Debug("all")
}

func Ping() error {
	return sTransmitter.Ping()
}

func PutObject(k string, reader io.Reader) error {
	return sTransmitter.PutObject(k, reader)
}

func GetAddress(k string) string {
	return sTransmitter.GetAddress(k)
}
