package utils

import (
	"crypto/tls"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

//获取相差时间
func getHourDiffer(startTime, endTime string, DetailedOutput bool) int64 {
	a, _ := time.Parse("2006-01-02 15:04:05", startTime)
	b, _ := time.Parse("2006-01-02 15:04:05", endTime)
	d := b.Sub(a)
	if DetailedOutput {
		logrus.Infof("距离域名过期还有 -> %d天 ｜%d小时 ｜%d分钟", int64(d.Hours()/24), int64(d.Hours()), int64(d.Minutes()))
	}
	return int64(d.Hours() / 24)
}

// 检测https域名的证书有效期
func ChkHTTPSCert(domainName string, DetailedOutput bool) int64 {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(domainName)
	defer resp.Body.Close()

	if err != nil {
		logrus.Error(domainName, " 请求失败")
		panic(err)
	}

	//fmt.Println(resp.TLS.PeerCertificates[0])
	certInfo := resp.TLS.PeerCertificates[0]
	if DetailedOutput {
		logrus.Info("过期时间:", certInfo.NotAfter)
		logrus.Info("组织信息:", certInfo.Subject)
	}

	nowTime := time.Now().Format("2006-01-02 15:04:05")
	endTime := certInfo.NotAfter.Format("2006-01-02 15:04:05")
	return getHourDiffer(nowTime, endTime, DetailedOutput)
}

//func main()  {
//	time := ChkHTTPSCert("https://lijinghua.club",false)
//	logrus.Warn(time)
//}
