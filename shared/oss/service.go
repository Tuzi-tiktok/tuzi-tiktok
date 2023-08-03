package oss

import (
	"errors"
	"io"
	"tuzi-tiktok/logger"
)

const defaultImpl = "lfs"

type ImplType string

// Attempted to use adapter mode, but not necessary
//func (t *ImplType) Supports(i ImplType) bool {
//	//Fault tolerance strings.ToLower(string(*t)) == strings.ToLower(string(i))
//	return i == *t
//}
//
//type Strategy interface {
//	Supports(i ImplType) bool
//}

type Constructor func() StorageTransmitter

type Binder func(string) error

type StorageTransmitter interface {
	Ping() error
	PutObject(string, io.Reader) error
	GetAddress(string) string
}

type CI struct {
	C Constructor
	B Binder
}

func Register(t ImplType, c Constructor, b Binder) error {
	if _, ok := Candidates[t]; ok {
		return errors.New("Error: Repetitive Candidate Check Your Register Name ")
	}
	Candidates[t] = &CI{c, b}
	return nil
}

func SetUsing(t ImplType) {
	var (
		ci *CI
		ok bool
	)
	if ci, ok = Candidates[t]; !ok {
		// TODO Replace logger
		logger.Debug("The corresponding configuration is not matched")
		ci = Candidates[defaultImpl]
	}
	sTransmitter = ci.C()
}
