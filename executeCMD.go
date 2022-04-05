package utils

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

func LocalCMD(cmdStr string, Args []string) (err error, ProcessState *os.ProcessState, OutInfo string) {
	command := exec.Command(cmdStr, Args...)
	outInfo := bytes.Buffer{}
	command.Stdout = &outInfo
	err = command.Run()
	if err != nil {
		fmt.Println(err.Error())
		return err, nil, ""
	}
	if err = command.Wait(); err != nil {
		logrus.Error(err.Error())
	}
	return nil, command.ProcessState, outInfo.String()
}
