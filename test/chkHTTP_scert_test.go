package test

import (
	"github.com/lijinghuatongxue/utils"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestChkHTTPSCert(t *testing.T) {
	time := utils.ChkHTTPSCert("https://lijinghua.club", false)
	logrus.Warn(time)
}
