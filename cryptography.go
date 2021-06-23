package utils

import (
	"github.com/lijinghuatongxue/utils/EncryptionAndDecryption"
	"github.com/sirupsen/logrus"
)

func Base64Encode(encodeStr string) string {
	// encode
	deByte := EncryptionAndDecryption.Base64Encode([]byte(encodeStr))
	return string(deByte)
}

func Base64Decode(decodeStr []byte) string {
	// decode
	enByte, err := EncryptionAndDecryption.Base64Decode(decodeStr)
	if err != nil {
		logrus.Error(err.Error())
	}
	return string(enByte)
}

//func main()  {
//	i := Base64Encode("233")
//	logrus.Infof("enCode str is %s",i)
//	j := Base64Decode([]byte("MjMz"))
//	logrus.Infof("deCode str is %s",j)
//}
