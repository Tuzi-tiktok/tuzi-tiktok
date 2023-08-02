package oss

import (
	"errors"
	"io"
	"log"
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
	PutObject(reader io.Reader) (string, error)
	GetAddress(k string) string
}

// CI Endpoint-Bucket Pair
type CI struct {
	// Endpoint  Hos:Port
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
		log.Println("The corresponding configuration is not matched")
		ci = Candidates[defaultImpl]
	}
	sTransmitter = ci.C()
}
