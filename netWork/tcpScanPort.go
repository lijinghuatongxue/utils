package main

import (
	"github.com/sirupsen/logrus"
	"net"
)

func CheckPort(ip net.IP, port int) error {
	tcpAddr := net.TCPAddr{
		IP:   ip,
		Port: port,
	}
	conn, err := net.DialTCP("tcp", nil, &tcpAddr)
	if err == nil {
		logrus.Infof("IP -> %s |Port scan passed! -> %d ", ip, port)
		if conn != nil {
			e := conn.Close()
			if e != nil {
				logrus.Error(e)
				return e
			}
		}
		return nil
	}
	logrus.Errorf("ip: %v port: %v \n", ip, port)
	return err

}
func checkIp(ip string) bool {
	if net.ParseIP(ip) == nil {
		logrus.Errorf("Incorrect IP address format！%s", ip)
		return false
	} else {
		return true
	}
}

// 扫描IP端口是否是passed
// Port scan passed
//func main() {
//	//ip := os.Args[1]
//	ip := "8.8.8.8"
//	var port int
//	port = 53
//	if checkIp(ip) {
//		if CheckPort(net.ParseIP(ip), port) != nil {
//			return
//		}
//	}
//}
