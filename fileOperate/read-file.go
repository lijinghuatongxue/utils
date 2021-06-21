package fileOperate

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func ReadFile2Str(FileName string, DetailedOutput bool) (string, error) {
	f, err := os.Open(FileName)
	if err != nil {
		if DetailedOutput {
			logrus.Error("read file fail", err)
		}
		return "null", err
	}
	defer f.Close()
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		if DetailedOutput {
			logrus.Error("read to fd fail", err)
		}
		return "null", err
	}
	return string(fd), nil
}
