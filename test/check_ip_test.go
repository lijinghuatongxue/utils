package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestCheckIpIsIPV4(t *testing.T) {
	ip := "8.8.8.8"
	if utils.CheckIp(ip) == true {
		logrus.Infof("IP -> %s | Is True.", ip)
	}
}
