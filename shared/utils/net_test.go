package utils

import (
	"log"
	"net"
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
func TestNsLookup(t *testing.T) {
	ip, err := net.LookupHost("192.168.2.1")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%v", ip)
}
