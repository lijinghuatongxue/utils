package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestRemoteSSH(t *testing.T) {
	cmd := "docker ps"
	ip := "192.168.0.123"
	port := "22"
	user := "root"
	RemoteCmdOutput, err := utils.RemoteCmd(ip, port, user, cmd, "")
	if err != true {
		logrus.Errorf("[util - remote-ssh] | ❌ false| Err ===》%s", err)
		return
	}
	logrus.Infof("cmd -> %s |Output -> \n%s", cmd, RemoteCmdOutput)
}
