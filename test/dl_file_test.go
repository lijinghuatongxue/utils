package test

import (
	"github.com/lijinghuatongxue/utils"
	"testing"
	"time"
)

func TestDlFile(t *testing.T) {
	// 文件保存路径
	SavePath := "./tmp/abc.html"
	// 下载地址
	Url := "https://www.blog.lijinghua.club"
	// 超时秒数
	TimeoutNum := 10
	utils.DlFile(SavePath, Url, time.Duration(TimeoutNum))
}
