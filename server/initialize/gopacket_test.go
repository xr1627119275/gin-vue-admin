package initialize

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
	"testing"
)

func TestPackageInitializeGorm_Injection(t *testing.T) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	// 打印可用的网络接口
	fmt.Println("可用网络接口:")
	for _, device := range devices {
		fmt.Printf("接口名称: %s\n", device.Name)
		for _, addr := range device.Addresses {
			fmt.Printf("IP地址: %s\n", addr.IP)
		}
		fmt.Printf("描述: %s\n", device.Description)
		fmt.Println("-------------------")
	}
}
