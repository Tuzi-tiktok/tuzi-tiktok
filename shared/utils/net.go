package utils

import (
	"net"
)

func RandomAvailablePort() int {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port
}

func GetLocalAddrByRC(addr string) string {
	dial, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer dial.Close()
	return dial.LocalAddr().String()
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
