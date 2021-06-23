package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"net"
	"time"
)

func GetDomainIP(DNSServer, Domain string) string {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, network, DNSServer+":53")
		},
	}
	ip, _ := r.LookupHost(context.Background(), Domain)
	if len(ip) == 0 {
		logrus.Errorf("Domain Resolver Err! -> %s |DNSServer -> %s |IP -> %s", Domain, DNSServer, ip)
	} else {
		logrus.Infof("Domain -> %s |DNSServer -> %s |IP -> %s", Domain, DNSServer, ip[0])
		return ip[0]
	}
	return ""
}
