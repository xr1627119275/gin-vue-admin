package webScan

import (
	"bufio"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webScan"
	webScanReq "github.com/flipped-aurora/gin-vue-admin/server/model/webScan/request"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

type WebScanService struct{}

var WebScanServiceApp = new(WebScanService)

// CreateWebScan 创建web扫描记录
// Author [piexlmax](https://github.com/piexlmax)
func (w_scanService *WebScanService) CreateWebScan(w_scan *webScan.WebScan) (err error) {
	err = global.GVA_DB.Create(w_scan).Error
	return err
}

// DeleteWebScan 删除web扫描记录
// Author [piexlmax](https://github.com/piexlmax)
func (w_scanService *WebScanService) DeleteWebScan(ID string) (err error) {
	err = global.GVA_DB.Delete(&webScan.WebScan{}, "id = ?", ID).Error
	return err
}

// DeleteWebScanByIds 批量删除web扫描记录
// Author [piexlmax](https://github.com/piexlmax)
func (w_scanService *WebScanService) DeleteWebScanByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]webScan.WebScan{}, "id in ?", IDs).Error
	return err
}

// UpdateWebScan 更新web扫描记录
// Author [piexlmax](https://github.com/piexlmax)
func (w_scanService *WebScanService) UpdateWebScan(w_scan webScan.WebScan) (err error) {
	err = global.GVA_DB.Model(&webScan.WebScan{}).Where("id = ?", w_scan.ID).Updates(&w_scan).Error
	return err
}

//// appendWebScanOutPut 最佳web扫描记录日志
//// Author [piexlmax](https://github.com/piexlmax)
//func (w_scanService *WebScanService) appendWebScanOutPut(w_scan webScan.WebScan) (err error) {
//	err = global.GVA_DB.Model(&webScan.WebScan{}).Where("id = ?", w_scan.ID).Updates(&w_scan).Error
//	return err
//}

// GetWebScan 根据ID获取web扫描记录
// Author [piexlmax](https://github.com/piexlmax)
func (w_scanService *WebScanService) GetWebScan(ID string) (w_scan webScan.WebScan, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&w_scan).Error
	return
}

// GetWebScanInfoList 分页获取web扫描记录
// Author [piexlmax](https://github.com/piexlmax)
func (w_scanService *WebScanService) GetWebScanInfoList(info webScanReq.WebScanSearch) (list []webScan.WebScan, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&webScan.WebScan{})
	var w_scans []webScan.WebScan
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.HOST != "" {
		db = db.Where("host LIKE ?", "%"+info.HOST+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&w_scans).Error
	return w_scans, total, err
}

func (w_scanService *WebScanService) RunScan(params webScanReq.RunWebScanParam, runType string) (err error) {
	var cmd *exec.Cmd
	var execFileName = "fscan"

	currentTime := time.Now()
	timestamp := currentTime.Unix()
	timestampStr := fmt.Sprintf("%d", timestamp)
	outFileName := "scanLogs" + string(os.PathSeparator) + timestampStr + ".json"
	// 如果scanLogs文件夹 不存在
	if _, err := os.Stat("scanLogs"); os.IsNotExist(err) {
		// 创建scanLogs文件夹
		_ = os.Mkdir("scanLogs", os.ModePerm)
	}
	var args []string
	if params.HOST != "" {
		args = append(args, "-h", params.HOST)
	}
	if params.PORT != "" {
		args = append(args, "-p", params.PORT)
	}
	for k, v := range params.Args {
		args = append(args, "-"+k)
		if v != "" {
			args = append(args, v.(string))
		}
	}
	args = append(args, "-o", outFileName)
	if runType == "TxPortMap" {
		execFileName = runType
		args = append(args, "-json", "")
	} else {
		args = append(args, "-json", "")
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
		return err
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return err
	}

	// Create a scanner to read the command's output
	scanner := bufio.NewScanner(stdout)
	scanInfo, err := w_scanService.GetWebScan(strconv.Itoa(int(params.WebScan.GVA_MODEL.ID)))
	if err != nil {
		return err
	}
	for scanner.Scan() {
		scanInfo.Result += scanner.Text() + "\n"
		w_scanService.UpdateWebScan(scanInfo)
	}

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {

		return err
	}
	// 打开文件
	file, err := os.Open(outFileName)
	var fileResult string
	if err != nil {
		return
	}
	defer func() {
		file.Close()
		os.Remove(outFileName)
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
	scanInfo.JsonResult = fileResult
	scanInfo.ScanOver = true
	err = w_scanService.UpdateWebScan(scanInfo)
	if err != nil {
		return err
	}
	return nil
}
