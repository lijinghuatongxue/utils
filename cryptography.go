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
