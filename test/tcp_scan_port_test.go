package test

import (
	"github.com/lijinghuatongxue/utils"
	"net"
	"testing"
)

func TestTcpScanPort(t *testing.T) {
	ip := "8.8.8.8"
	var port int
	port = 53
	if utils.CheckIp(ip) {
		if utils.CheckPort(net.ParseIP(ip), port) != nil {
			return
		}
	}
}
