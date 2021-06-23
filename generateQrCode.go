package utils

import (
	"github.com/sirupsen/logrus"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func generateQrCode(URL, QrCodePath string, DetailedOutput bool) {

	qrCode, _ := qr.Encode(URL, qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	file, _ := os.Create(QrCodePath)
	defer file.Close()

	if png.Encode(file, qrCode) != nil {
		logrus.Errorf("")
	}
	if DetailedOutput {
		logrus.Infof("URL: %s | QrCodePath: %s", URL, QrCodePath)
	}

}

//func main()  {
//	URL := "http://tel.lijinghua.club"
//	QrCodePath := "./qc.png"
//	generateQrCode(URL,QrCodePath,true)
//}
