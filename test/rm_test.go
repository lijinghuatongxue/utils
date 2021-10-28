package test

import (
	"github.com/lijinghuatongxue/utils"
	"testing"
)

func TestRM(t *testing.T) {
	// 删除文件或者文件夹,可选仅清空文件
	err := utils.RM("./../tmp", true)
	if err != nil {
		return
	}
	err = utils.RM("./../tmp", false)
	if err != nil {
		return
	}
}
