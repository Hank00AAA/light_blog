package net_util

import (
	"fmt"
	"net"
)

// GetIps
func GetIP() string {
	adds, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("GetIPErr :%v", err)
		return ""
	}

	for _, address := range adds {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}