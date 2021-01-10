//package main
//
//import (
//	"fmt"
//	"github.com/sirupsen/logrus"
//	"io"
//	"os"
//)
//
//func checkFileIsExist(filename string) bool {
//	if _, err := os.Stat(filename); os.IsNotExist(err) {
//		return false
//	}
//	return true
//}
//func main() {
//	var contentString = "测试1\n测试2\n"
//	var filename = "./data/test.txt"
//	var f *os.File
//	var err1 error
//	if checkFileIsExist(filename) { //如果文件存在
//		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
//		fmt.Println("文件存在")
//	} else {
//		f, err1 = os.Create(filename) //创建文件
//		fmt.Println("文件不存在")
//	}
//	defer f.Close()
//	n, err1 := io.WriteString(f, contentString) //写入文件(字符串)
//	if err1 != nil {
//		logrus.Error(err1)
//	}
//	logrus.Infof("写入 %d 个字节n", n)
//}
package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func writeFile(FileName, contentString string, IsCover bool) bool {
	var f *os.File
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n, _ := w.WriteString(contentString)
	if IsCover == true {
		content := []byte(contentString)
		err := ioutil.WriteFile(FileName, content, 0644)
		if err != nil {
			logrus.Error(err)
			return false
		}
		logrus.Infof("覆盖写入%d个字节", n)
	} else {
		var _ error
		if checkFileIsExist(FileName) { //如果文件存在
			f, _ = os.OpenFile(FileName, os.O_APPEND, 0666) //打开文件
			logrus.Info("文件存在")
		} else {
			f, _ = os.Create(FileName) //创建文件
			logrus.Info("文件不存在")
		}
		file, err := os.OpenFile(FileName, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("文件打开失败", err)
			return false
		}
		//及时关闭file句柄
		defer file.Close()
		//写入文件时，使用带缓存的 *Writer
		write := bufio.NewWriter(file)
		write.WriteString(contentString)
		//Flush将缓存的文件真正写入到文件中
		write.Flush()
		logrus.Infof("追加写入%d个字节", n)
	}
	return true
}

// 文件读写，可追加，可覆盖
func main() {
	// 文件内容
	var contentString = "测试1\n测试2\n测试1\n测试2\n测试1\n测试2\n"
	// 文件位置
	var filename = "./data/test.txt"
	writeFile(filename, contentString, false)
}
