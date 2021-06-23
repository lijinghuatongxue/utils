package utils

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
