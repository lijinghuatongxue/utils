package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestTar(t *testing.T) {
	// ----------------- 打包 -------------
	// 目标打包压缩后文件
	dst := "./tmp/a.tar"
	// 目标被打包文件1
	file1 := "./go.mod"
	// 目标被打包文件2
	file2 := "./README.md"
	if err := utils.Tar([]string{file1, file2}, dst); err != nil {
		logrus.Error(err)
	}
	// --------------- 解包 --------------
	// srcFile：目标tar文件 ，dsrDir ：解压目录
	if utils.Decompression("../data/a.tar", "../data") != nil {
		var err interface{}
		logrus.Error(err)
	}
}
