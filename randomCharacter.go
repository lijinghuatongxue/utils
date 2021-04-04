package meUtils

import (
	"crypto/rand"
	"fmt"
)

func AlgorithmRandomCharacter(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

// 随机一个10位数字符串，包括数字和字符
// Random a 10 digit string, including numbers and characters
//func main()  {
//	logrus.Warn(AlgorithmRandomCharacter(10))
//}
