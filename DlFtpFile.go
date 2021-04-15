package meUtils

import (
	"github.com/jlaffaye/ftp"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

// download ftp file，support Ftp auth
func DlFtpFile(FtpPort, FtpFileUrl, SavePath string, DetailedOutput bool) error {
	//
	var FtpDomain string
	var FtpFilePath string
	_, FtpDomain, FtpFilePath = CuttingFtpFileAddress(FtpFileUrl, true)

	c, err := ftp.Dial(FtpDomain+":"+FtpPort, ftp.DialWithTimeout(10*time.Second))
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = c.Login("lijinghua", "123456")
	if err != nil {
		logrus.Error(err)
	}
	//c.ChangeDir("./data")
	res, err := c.Retr(FtpFilePath)
	if err != nil {
		logrus.Error(err)
		return err
	}

	defer res.Close()

	outFile, err := os.Create(SavePath)
	if err != nil {
		logrus.Error(err)
		return err
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, res)
	if err != nil {
		logrus.Error(err)
		return err
	}
	if DetailedOutput {
		logrus.Infof("util |Ftp文件下载成功 ｜ Ftp地址 -> %s |保存地址 -> %s", FtpFileUrl, SavePath)
	}
	return nil
}

func CuttingFtpFileAddress(FtpFileUrl string, DetailedOutput bool) (string, string, string) {
	// 截取到ftp
	ftpCount := strings.SplitN(FtpFileUrl, ":", len(FtpFileUrl))
	// 截取主域名
	domainTotal := strings.SplitN(ftpCount[1], "/", len(ftpCount[1]))
	// 获取文件路径
	FtpFilePath := strings.SplitN(ftpCount[1], "/", 4)
	if DetailedOutput {
		logrus.Infof("util |CuttingFtpFileAddress |protocol -> %s |Domain -> %s |FtpFileUrlPath -> %s |", ftpCount[0], domainTotal[2], FtpFilePath[3])
	}
	return ftpCount[0], domainTotal[2], FtpFilePath[3]
}

//func main()  {
//	FtpPort := "21"
//	var err error
//	err = DlFtpFile(FtpPort, "ftp://ali/1KB.zip", "./data/0415.zip",true)
//	if err != nil{
//		logrus.Error("download FtpFile err")
//	}
//}
