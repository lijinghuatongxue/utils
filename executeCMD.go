package utils

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func LocalCMD(cmdStr string) (err error, ProcessState *os.ProcessState, OutInfo string) {
	command := exec.Command(cmdStr)
	outInfo := bytes.Buffer{}
	command.Stdout = &outInfo
	err = command.Start()
	if err != nil {
		fmt.Println(err.Error())
		return err, nil, ""
	}
	if err = command.Wait(); err != nil {
		logrus.Error(err.Error())
	}
	return nil, command.ProcessState, outInfo.String()
}
