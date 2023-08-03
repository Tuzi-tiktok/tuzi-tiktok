package oss

import (
	"io"
	"tuzi-tiktok/logger"
	// Required Import
	_ "tuzi-tiktok/oss/internal"
	. "tuzi-tiktok/oss/internal/define"
)

func init() {
	logger.Debug("over story")
}

func Ping() error {
	return STransmitter.Ping()
}

func PutObject(k string, reader io.Reader) error {
	return STransmitter.PutObject(k, reader)
}

func GetAddress(k string) string {
	return STransmitter.GetAddress(k)
}
