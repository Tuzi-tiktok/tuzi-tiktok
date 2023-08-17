package utils

import (
	"fmt"
	"net"
	"strings"
	cfg "tuzi-tiktok/config"
)

func RandomAvailablePort() int {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}

func GetLocalAddrSingle(addr string) string {
	dial, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer dial.Close()
	fmt.Printf(dial.LocalAddr().String())
	return dial.LocalAddr().String()
}

func GetLocalAddrByRC() string {
	dial, err := net.Dial("tcp", fmt.Sprintf("%v:%v", cfg.Registration.Host, cfg.Registration.Port))
	if err != nil {
		panic(err)
	}
	defer dial.Close()
	return strings.Split(dial.LocalAddr().String(), ":")[0]
}

func GetLocalAddr() string {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, addr := range addr {
		ipNet, isIpNet := addr.(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			ipv4 := ipNet.IP.To4()
			if ipv4 != nil {
				return ipv4.String()
			}
		}
	}
	return ""
}
