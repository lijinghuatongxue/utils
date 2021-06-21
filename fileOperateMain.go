package meUtils

import (
	"github.com/lijinghuatongxue/utils/fileOperate"
	"github.com/sirupsen/logrus"
)

func WriteFile(FileName, contentString string, IsCover, DetailedOutput bool) bool {
	iok := fileOperate.WriteFile(FileName, contentString, IsCover, DetailedOutput)
	if !iok {
		logrus.Error("util file operate err!")
		return false
	}
	return true
}
