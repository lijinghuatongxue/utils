package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestWgetFile(t *testing.T) {
	var err error
	err = utils.DlFileUseWget("ftp://speedtest.tele2.net/../1KB.zip", "./../data/curl.zip")
	if err != nil {
		logrus.Error("Wget file error!")
	}
}
