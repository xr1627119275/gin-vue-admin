package main

import (
	"bufio"
	"fmt"
	webScanReq "github.com/flipped-aurora/gin-vue-admin/server/model/webScan/request"
	"os"
	"os/exec"
	"runtime"
	"testing"
	"time"
)

func TestWebScanApi_CreateWebScan(t *testing.T) {
	var params = webScanReq.RunWebScanParam{
		RunType: "TxPortMap",
	}
	var cmd *exec.Cmd
	var execFileName = "fscan"

	currentTime := time.Now()
	timestamp := currentTime.Unix()
	timestampStr := fmt.Sprintf("%d", timestamp)
	outFileName := "scanLogs" + string(os.PathSeparator) + timestampStr + ".json"
	var Args = map[string]string{
		"i":    "6.6.6.3",
		"p":    "5008",
		"json": "",
		"o":    outFileName,
	}
	// 如果scanLogs文件夹 不存在
	if _, err := os.Stat("scanLogs"); os.IsNotExist(err) {
		// 创建scanLogs文件夹
		_ = os.Mkdir("scanLogs", os.ModePerm)
	}
	var args []string

	for k, v := range Args {
		args = append(args, "-"+k)
		if v != "" {
			args = append(args, v)
		}
	}
	if params.RunType == "TxPortMap" {
		execFileName = params.RunType
	}
	// Determine the executable based on the OS
	if runtime.GOOS == "windows" {
		cmd = exec.Command(fmt.Sprintf("./%s.exe", execFileName), args...)
	} else {
		cmd = exec.Command(fmt.Sprintf("./%s", execFileName), args...)
	}

	// Get the stdout pipe
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return
	}

	// Create a scanner to read the command's output
	scanner := bufio.NewScanner(stdout)
	if err != nil {
		return
	}
	var result string
	for scanner.Scan() {
		result += scanner.Text()
	}
	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return
	}
	// 打开文件
	file, err := os.Open(Args["o"])
	var fileResult string
	if err != nil {
		return
	}
	defer func() {
		file.Close()
		os.Remove(Args["o"])
	}()
	// 创建一个文件的缓冲区
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if 0 == n {
			break
		}
		fileResult += string(buf[:n])
	}
	fmt.Println("fileResult:", fileResult)
	// 关闭文件
}

func TestParseFscan(t *testing.T) {
	//devices, err := pcap.FindAllDevs()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//for _, d := range devices {
	//
	//}
	//fmt.Println(devices)
}
