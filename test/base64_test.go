package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestBase64(t *testing.T) {
	i := utils.Base64Encode("233")
	logrus.Infof("enCode str is %s", i)
	j := utils.Base64Decode([]byte("MjMz"))
	logrus.Infof("deCode str is %s", j)
}
