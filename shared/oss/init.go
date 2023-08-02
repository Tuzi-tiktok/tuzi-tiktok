package oss

import (
	"io"
	"log"
)

var (
	Candidates   = make(map[ImplType]*CI, 2)
	sTransmitter StorageTransmitter
)

func init() {
	log.Print("all")
}

func Ping() error {
	return sTransmitter.Ping()
}

func PutObject(reader io.Reader) (string, error) {
	return sTransmitter.PutObject(reader)
}

func GetAddress(k string) string {
	return sTransmitter.GetAddress(k)
}
