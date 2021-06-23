package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func ProgressBarFuncTest() {
	logrus.Info("abc")
}
func TestProgressBar(t *testing.T) {
	utils.ProgressBar("./utils", 67, ProgressBarFuncTest)
}
