package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func TestDlFile(t *testing.T) {
	// 文件保存路径
	SavePath := "./tmp/abc.html"
	// 下载地址
	Url := "https://www.blog.lijinghua.club"
	// 超时秒数
	TimeoutNum := 10
	utils.DlFile(SavePath, Url, time.Duration(TimeoutNum))
}

func TestGenerateQrCode(t *testing.T) {
	URL := "http://tel.lijinghua.club"
	QrCodePath := "./data/test-ftp.png"
	utils.GenerateQrCode(URL, QrCodePath, true)
}
func TestBase64(t *testing.T) {
	i := utils.Base64Encode("233")
	logrus.Infof("enCode str is %s", i)
	j := utils.Base64Decode([]byte("MjMz"))
	logrus.Infof("deCode str is %s", j)
}
func TestChkHTTPSCert(t *testing.T) {
	time := utils.ChkHTTPSCert("https://lijinghua.club", false)
	logrus.Warn(time)
}
func TestDlFtpFile(t *testing.T) {
	FtpPort := "21"
	var err error
	err = utils.DlFtpFile(FtpPort, "http://speedtest.tele2.net/1MB.zip", "./data/0415.zip", true)
	if err != nil {
		logrus.Error("download FtpFile err")
	}
}
