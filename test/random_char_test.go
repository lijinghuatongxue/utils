package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestRandomChar(t *testing.T) {
	logrus.Warn(utils.AlgorithmRandomCharacter(10))
	utils.AlgorithmRandomNum(99999)
}
