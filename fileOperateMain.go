package meUtils

import (
	"github.com/lijinghuatongxue/utils/fileOperate"
	"github.com/sirupsen/logrus"
)

func WfMain(FileName, contentString string, IsCover, DetailedOutput bool) bool {
	isOk := fileOperate.WriteFile(FileName, contentString, IsCover, DetailedOutput)
	if !isOk {
		logrus.Error("util | WriteFile operate err!")
		return false
	}
	return true
}

func ReadFile2StrMain(FileName string, DetailedOutput bool) (string, error) {
	StrOutput, isOk := fileOperate.ReadFile2Str(FileName, DetailedOutput)
	if isOk != nil {
		logrus.Error("util | ReadFile2Str operate err!")
		return "", isOk
	}
	return StrOutput, isOk
}
func ListFileDir(dirPath, suffix string) (files []string, err error) {
	ListRes, err := fileOperate.ListDir(dirPath, suffix)
	if err != nil {
		logrus.Errorf("util | ListFileDir operate err! \n%s", err)
	}
	return ListRes, err
	//for _, value := range ListRes{
	//	//fmt.Println(index, "\t",value)
	//	fmt.Println(value)
	//}
}
