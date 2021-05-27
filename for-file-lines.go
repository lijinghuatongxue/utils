package meUtils

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// 按行读取文件， 传入文件路径
func ForFileLines(FileName string) error {
	fi, err := os.Open(FileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		logrus.Info(string(a))
	}
	return nil
}

// 按行读取文件并执行函数，传入文件路径、执行函数名称
func ForFileLinesExecFunc(FileName string, FuncName func(string2 string)) error {
	fi, err := os.Open(FileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return err
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		FuncName(string(a))
	}
	return nil
}

//func PrintLog(Msg string)  {
//	logrus.Info(Msg)
//}

//func main() {
//	//var err error
//	//err = ForFileLines("./tmp/test.txt")
//	//if err != nil {
//	//	logrus.Error(err)
//	//}
//	ForFileLinesExecFunc("./data/tpl.txt",PrintLog)
//}
