package main

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

var err error

func AuthClone(GitTmpDir, GitDir, AppName, Url, gitUser, GitPasswd string) error {
	// Tempdir to clone the repository
	_, err := ioutil.TempDir(GitTmpDir, AppName)
	if err != nil {
		logrus.Error(err)
		return err
	}
	GitDirAppName := GitDir + "/" + AppName
	err = os.RemoveAll(GitDirAppName) // clean up
	if err != nil {
		logrus.Error("git目标目录删除失败")
		return err
	}
	// Clones the repository into the given dir, just as a normal LinuxCMD clone does
	_, err = git.PlainClone(GitDirAppName, false, &git.CloneOptions{
		URL: Url,
		Auth: &http.BasicAuth{
			Username: gitUser,
			Password: GitPasswd,
		},
	})

	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func Clone(GitTmpDir, GitDir, AppName, Url string) error {
	// Tempdir to clone the repository
	_, err := ioutil.TempDir(GitTmpDir, AppName)
	if err != nil {
		logrus.Error(err)
		return err
	}
	GitDirAppName := GitDir + "/" + AppName
	err = os.RemoveAll(GitDirAppName) // clean up
	if err != nil {
		logrus.Error("git目标目录删除失败")
		return err
	}
	_, err = git.PlainClone(GitDirAppName, false, &git.CloneOptions{
		URL:      Url,
		Progress: os.Stdout, //如果想关闭输出，请注释这行
	})

	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

//func main() {
//	//gitUrl := "https://github.com/git-fixtures/basic.git"
//	gitUrl := "https://github.com/lijinghuatongxue/server-go.LinuxCMD"
//	// LinuxCMD 仓库存储目录
//	gitData := "./data"
//	// LinuxCMD 仓库临时存储目录
//	gitTmpData := "./tmp"
//	// 项目名字
//	AppName := "AppName"
//	gitUser := "lijinghuatongxue"
//	gitPasswd := "xxx"
//	// 无认证clone
//	err = Clone(gitTmpData, gitData, AppName, gitUrl)
//	if err != nil {
//		logrus.Error("Git Clone Err")
//	}
//	// 有认证clone
//	err = AuthClone(gitTmpData, gitData, AppName, gitUrl, gitUser, gitPasswd)
//	if err != nil {
//		logrus.Error("Git AuthClone Err")
//	}
//}
