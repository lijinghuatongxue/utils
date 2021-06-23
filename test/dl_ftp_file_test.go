package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestDlFtpFile(t *testing.T) {
	FtpPort := "21"
	var err error
	err = utils.DlFtpFile(FtpPort, "http://speedtest.tele2.net/1MB.zip", "../data/0415.zip", true)
	if err != nil {
		logrus.Error("download FtpFile err")
	}
}
