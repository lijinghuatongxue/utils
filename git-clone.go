package utils

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

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
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func Clone(GitTmpDir, GitDir, AppName, Url string) error {
	// Tempdir to clone the repository
	var err error
	isOk := IsDir(GitTmpDir)
	if isOk != true {
		var err error
		err = os.MkdirAll(GitTmpDir, 0755)
		if err != nil {
			logrus.Error(err)
		}
	}
	isOk = IsDir(GitDir)
	if isOk != true {
		var err error
		err = os.MkdirAll(GitDir, 0755)
		if err != nil {
			logrus.Error(err)
		}
	}
	_, err = ioutil.TempDir(GitTmpDir, AppName)
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
