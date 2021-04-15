package meUtils

import (
	"github.com/sirupsen/logrus"
	"os/exec"
)

func DlFileUseWget(FileUrl, SavePath string) error {
	cmd := exec.Command("wget", FileUrl, "-O", SavePath)
	cmdOut, err := cmd.CombinedOutput()
	if err != nil {
		logrus.Error(string(cmdOut))
		logrus.Error(err)
	}
	return nil
}

//func main() {
//	var err error
//	err = DlFileUseWget("ftp://speedtest.tele2.net/../1KB.zip", "./data/curl.zip")
//	if err != nil{
//		logrus.Error("Wget file error!")
//	}
//}
