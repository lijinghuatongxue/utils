package test

import (
	"github.com/lijinghuatongxue/utils"
	"testing"
)

func TestGenerateQrCode(t *testing.T) {
	URL := "http://tel.lijinghua.club"
	QrCodePath := "./data/test-ftp.png"
	utils.GenerateQrCode(URL, QrCodePath, true)
}
