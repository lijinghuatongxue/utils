package meUtils

import (
	"bufio"
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

func WriteFile(FileName, contentString string, IsCover, DetailedOutput bool) bool {
	var f *os.File
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n, _ := w.WriteString(contentString)
	if IsCover {
		content := []byte(contentString)
		err := ioutil.WriteFile(FileName, content, 0644)
		if err != nil {
			logrus.Error(err)
			return false
		}
		if DetailedOutput {
			logrus.Infof("|文件 -> %s |覆盖写入%d个字节", FileName, n)
		}
	} else {
		var _ error
		if checkFileIsExist(FileName) { //如果文件存在
			f, _ = os.OpenFile(FileName, os.O_APPEND, 0666) //打开文件
		} else {
			f, _ = os.Create(FileName) //创建文件
			if DetailedOutput {
				logrus.Infof("｜文件不存在 -> %s |已创建.", FileName)
			}
		}
		file, err := os.OpenFile(FileName, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Error("文件打开失败", err)
			return false
		}
		//及时关闭file句柄
		defer file.Close()
		//写入文件时，使用带缓存的 *Writer
		write := bufio.NewWriter(file)
		_, err = write.WriteString(contentString)
		if err != nil {
			logrus.Error("文件缓存写入失败", err)
		}
		//Flush将缓存的文件真正写入到文件中
		err = write.Flush()
		if err != nil {
			logrus.Error("文件写入失败", err)
		}
		if DetailedOutput {
			logrus.Infof("|文件 -> %s |追加写入%d个字节", FileName, n)
		}
	}
	return true
}

// 文件读写，可追加，可覆盖
//func main() {
//	// 文件内容
//	var contentString = "测试1\n测试2\n测试1\n测试2\n测试1\n测试2\n"
//	// 文件位置
//	var filename = "./data/test.txt"
//	WriteFile(filename, contentString, false)
//}
