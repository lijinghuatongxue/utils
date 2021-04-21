package meUtils

import (
	"crypto/rand"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/big"
)

// 随机一个10位数字符串，包括数字和字符
// Random a 10 digit string, including numbers and characters
func AlgorithmRandomCharacter(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

// 随机数字
// Generating random numbers
// Within parameter range
func AlgorithmRandomNum(n int) int64 {
	Num, _ := rand.Int(rand.Reader, big.NewInt(int64(n)))
	logrus.Info(Num.Int64())
	return Num.Int64()
}

//func main()  {
//	logrus.Warn(AlgorithmRandomCharacter(10))
//	AlgorithmRandomNum(99999)
//}
