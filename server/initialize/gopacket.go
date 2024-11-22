package initialize

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize/packet"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/gopacket"
	"io"
	"log"
	"net"
	"runtime"
	"strings"
	"sync"
)

var eventCh = make(chan interface{}, 1024)
var tcpUdpEventCh = make(chan interface{}, 1024)
var device string

func getClientInfo() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("获取网卡信息失败:", err)
		return
	}
	for _, inter := range interfaces {
		fmt.Printf("接口名字: %v\n", inter.Name)
		fmt.Printf("MAC地址: %v\n", inter.HardwareAddr)
		addrs, err := inter.Addrs()
		if err != nil {
			fmt.Println("获取地址信息失败:", err)
			continue
		}
		for _, address := range addrs {
			// 检查ip地址判断是否回环地址
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					if ipnet.IP.String() != "" {
						fmt.Println(ipnet.IP.String())
					}
				}
			}
		}
		fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	}
}
func PacketHttpInit(wg *sync.WaitGroup) {
	if runtime.GOOS == "windows" {
		device = global.GVA_CONFIG.Scan.Device
	} else {
		device = "eth4"
	}
	getClientInfo()
	wg.Done()
	go handleR()
	go handleTcpUdpR()
	if err := packet.NewPacketHandle(context.Background(), device, eventCh, tcpUdpEventCh).Listen(); err != nil {
		log.Println(err.Error())
	}
}
func handleTcpUdpR() {
	//for i := range tcpUdpEventCh {
	//	data := i.(packet.TcpUdpEvent)
	//	fmt.Printf("%v", data)
	//}
}
func handleR() {
	for i := range eventCh {
		data := i.(packet.Event)
		var Req = data.Req
		var Resp = data.Resp

		req_bytes, _ := io.ReadAll(Req.Body)
		var ReqContent = string(req_bytes)

		var contentType = Resp.Header.Get("Content-Type")
		bytes, _ := io.ReadAll(Resp.Body)
		var ResContent = ""
		if strings.Contains(contentType, "text") || strings.Contains(contentType, "xml") || strings.Contains(contentType, "json") {
			ResContent = string(bytes)
		}
		if strings.ContainsFunc(Req.Host, func(r rune) bool {
			arr := []string{
				"qq.com",
				"192.168.10.1",
				"192.168.10.26",
			}
			for _, v := range arr {
				if strings.Contains(Req.Host, v) {
					return true
				}
			}
			return false

		}) {
			continue
		}
		var info = gopacket.HttpInfo{
			RemoteAddr: Req.RemoteAddr,
			Path:       Req.RequestURI,
			Host:       Req.Host,
			Headers:    Req.Header,
			ResData:    bytes,
			ResContent: ResContent,
			ReqData:    req_bytes,
			ReqContent: ReqContent,
			ResHeader:  Resp.Header,
			StatusCode: Resp.StatusCode,
			Method:     Req.Method,
			Proto:      Req.Proto,
		}
		global.GVA_DB.Create(&info)

		//log.Printf("%v"+
		//	"request uri: %s,"+
		//	"response status: %v",
		//	Req.Host,
		//	Req.URL.Port(),
		//	ResContent)
	}
}
