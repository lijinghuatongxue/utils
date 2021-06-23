package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetDomainIp(t *testing.T) {
	IP := utils.GetDomainIP("8.8.8.8", "lijinghua.club")
	logrus.Info(IP)
}
