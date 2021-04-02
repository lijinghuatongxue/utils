package main

import (
	"github.com/sirupsen/logrus"
	"net"
)

func CheckIp(ip string) bool {
	if net.ParseIP(ip) == nil {
		logrus.Errorf("IP -> %s | Is False.", ip)
		return false
	} else {
		return true
	}
}

// 检查是否是IP格式
// Check Str it is in IP format
func main() {
	ip := "localhost"
	if CheckIp(ip) == true {
		logrus.Infof("IP -> %s | Is True.", ip)
	}
}
