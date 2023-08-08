package utils

import (
	"fmt"
	"net"
	"testing"
)

func TestNewAuth(t *testing.T) {
	_, err := NewAuth()
	if err != nil {
		panic(err)
	}
}
func TestNetResolve(t *testing.T) {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		panic(err)
	}
	fmt.Println(addr.IP)
}
