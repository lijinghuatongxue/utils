package NetWork

import (
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

func DlFile(SavePath, Url string, TimeOutNum time.Duration) bool {
	newFile, err := os.Create(SavePath)
	if err != nil {
		logrus.Error(err)
		return false
	}
	defer newFile.Close()
	// 超时配置
	timeout := TimeOutNum * time.Second
	client := http.Client{
		Timeout: timeout,
	}
	response, err := client.Get(Url)
	if err != nil {
		logrus.Error(err)
		return false
	}
	defer response.Body.Close()
	// 将HTTP response Body中的内容写入到文件
	// Body满足reader接口，因此我们可以使用ioutil.Copy
	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		logrus.Error(err)
		return false
	}
	FileSize := numBytesWritten / 1024 / 1024
	if FileSize <= 0 {
		FileSize := "不足1"
		logrus.Infof("[util - Downloaded %s MB file.]", FileSize)
	} else {
		logrus.Infof("[util - Downloaded %d MB file.]", numBytesWritten/1024/1024)
	}
	return true
}

//func main() {
//	// 文件保存路径
//	SavePath := "./tmp/222.html"
//	// 下载地址
//	Url := "https://www.blog.lijinghua.club"
//	// 超时秒数
//	TimeoutNum := 2
//	DlFile(SavePath, Url, time.Duration(TimeoutNum))
//}
