package meUtils

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type StrOutput string

func ReadFile2Str(FileName string, DetailedOutput bool) (StrOutput, error) {
	f, err := os.Open(FileName)
	if err != nil {
		if DetailedOutput {
			logrus.Error("read file fail", err)
		}
		return "null", err
	}
	err = f.Close()
	if err != nil {
		if DetailedOutput {
			logrus.Error("close file fail", err)
		}
	}
	fd, err := ioutil.ReadAll(f)
	if err != nil {
		if DetailedOutput {
			logrus.Error("read to fd fail", err)
		}
		return "null", err
	}
	return StrOutput(fd), nil
}
