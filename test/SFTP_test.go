package test

import (
	"github.com/lijinghuatongxue/utils"
	"testing"
)

func TestSFTP(t *testing.T) {
	err := utils.SftpDownload(true, "aaa", "passwd", "aaa", "./xx.pem", "./tmp/", "/serviceprovider/head_disable.png", 21)
	if err != nil {
		return
	}
	err = utils.SftpUpload(false, "aaa", "passwd", "aaa.com", "./xx.pem", "./11.png", "/serviceprovider/", 21)
	if err != nil {
		return
	}
}
