package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFileOperate(t *testing.T) {
	var str, str1 []string
	str, _ = utils.FindDirAllFile("./", "")
	for _, s := range str {
		logrus.Info(s)
	}
	utils.WfMain("./tmp.txt", "abc", true, true)
	str1, _ = utils.ListFileDir("./", "")
	for _, s := range str1 {
		logrus.Info(s)
	}
	var str2 string
	str2, _ = utils.ReadFile2StrMain("/tmp.txt", true)
	logrus.Info(str2)
}
