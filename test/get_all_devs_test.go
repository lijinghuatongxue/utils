package test

import (
	"github.com/lijinghuatongxue/utils/NetWork"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetAllDevs(t *testing.T) {
	DevsList := NetWork.NewWorkGetAllDevs()
	for _, value := range DevsList {
		logrus.Infof("Device Name ：%s", value)
	}
}
