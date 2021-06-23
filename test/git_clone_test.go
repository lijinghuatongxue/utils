package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGitClone(t *testing.T) {
	gitUrl := "https://github.com/lijinghuatongxue/utils.git"
	// LinuxCMD 仓库存储目录
	gitData := "./data"
	// LinuxCMD 仓库临时存储目录
	gitTmpData := "./tmp"
	// 项目名字
	AppName := "AppName"

	// 无认证clone
	var err error
	err = utils.Clone(gitTmpData, gitData, AppName, gitUrl)
	if err != nil {
		logrus.Error("Git Clone Err")
	}
	gitUser := "lijinghuatongxue"
	gitPasswd := "xxx"
	// 有认证clone
	err = utils.AuthClone(gitTmpData, gitData, AppName, gitUrl, gitUser, gitPasswd)
	if err != nil {
		logrus.Error("Git AuthClone Err")
	}
}
