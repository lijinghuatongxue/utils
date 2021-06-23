package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func MakeChannelFunc() {
	logrus.Info(1)
}
func TestMakeChannel(t *testing.T) {
	utils.MakeChannel(10, MakeChannelFunc)
}
