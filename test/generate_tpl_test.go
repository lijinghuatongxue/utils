package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGenerateTpl(t *testing.T) {
	type DataStruct struct {
		ProjectName string
		CMD         string
		Program     string
	}
	// 实例化结构体
	foo := DataStruct{
		ProjectName: "AD-001",
		Program:     "spring-gateway",
		CMD:         "CMD11",
	}
	err, OutPut := utils.GenerateTpl("./../data/shuchu.txt", "./data/tpl.txt", &foo, false)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Info(OutPut)
}
