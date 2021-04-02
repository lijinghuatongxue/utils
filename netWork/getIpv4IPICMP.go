package main

import (
	"github.com/sirupsen/logrus"
	"github.com/tatsushid/go-fastping"
	"net"
	"time"
)

func GetIpv4IpICMP(Ipv4Addr string) {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", Ipv4Addr)
	if err != nil {
		logrus.Error(err)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		logrus.Infof("IP Addr: %s receive, 往返时延: %v\n", addr.String(), rtt)
	}
	err = p.Run()
	if err != nil {
		logrus.Error(err)
	}
}

// 获取IP或者域名的ping延迟，仅限于Ipv4
// Get IP or domain name Ping delay，IPv4 only
//func main() {
//	DomainName := "google.com"
//	//DomainName := "8.8.8.8"
//	GetIpv4IpICMP(DomainName)
//}
