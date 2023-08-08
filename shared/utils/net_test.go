package utils

import (
	"log"
	"testing"
)

func TestRandomAvailablePort(t *testing.T) {
	port := RandomAvailablePort()
	log.Printf("Random Port Is %v", port)
}

func TestGetLocalAddr(t *testing.T) {
	log.Printf("Local Addr  %v", GetLocalAddr())
}
func TestGetLocalAddrByRC(t *testing.T) {
	log.Printf("%v", GetLocalAddrByRC())
}
