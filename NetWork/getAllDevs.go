package NetWork

import (
	"github.com/google/gopacket/pcap"
	"github.com/sirupsen/logrus"
	"log"
)

func NewWorkGetAllDevs() []string {
	// 得到所有的(网络)设备
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	// 打印设备信息
	//fmt.Println("Devices found:")
	var DevsList []string
	for _, device := range devices {
		DevsList = append(DevsList, device.Name)
		//fmt.Println("\nName: ", device.Name)
		//fmt.Println("Description: ", device.Description)
		//fmt.Println("Devices addresses: ", device.Description)
		//for _, address := range device.Addresses {
		//	fmt.Println("- IP address: ", address.IP)
		//	fmt.Println("- Subnet mask: ", address.Netmask)
		//}
	}
	logrus.Info(DevsList)
	return DevsList
}
