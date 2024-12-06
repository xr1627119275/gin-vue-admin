package highPort

import (
	"bufio"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/highPort"
	highPortReq "github.com/flipped-aurora/gin-vue-admin/server/model/highPort/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm/clause"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type HighRiskPortConfigService struct{}

// CreateHighRiskPortConfig 创建高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) CreateHighRiskPortConfig(HRPC *highPort.HighRiskPortConfig) (err error) {
	err = global.GVA_DB.Create(HRPC).Error
	return err
}

// DeleteHighRiskPortConfig 删除高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) DeleteHighRiskPortConfig(ID string) (err error) {
	err = global.GVA_DB.Delete(&highPort.HighRiskPortConfig{}, "id = ?", ID).Error
	return err
}

// DeleteHighRiskPortConfigByIds 批量删除高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) DeleteHighRiskPortConfigByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]highPort.HighRiskPortConfig{}, "id in ?", IDs).Error
	return err
}

// UpdateHighRiskPortConfig 更新高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) UpdateHighRiskPortConfig(HRPC highPort.HighRiskPortConfig) (err error) {
	err = global.GVA_DB.Model(&highPort.HighRiskPortConfig{}).Where("id = ?", HRPC.ID).Updates(&HRPC).Error
	return err
}

// GetHighRiskPortConfig 根据ID获取高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) GetHighRiskPortConfig(ID string) (HRPC highPort.HighRiskPortConfig, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&HRPC).Error
	return
}

// GetHighRiskPortLogsInfoList 分页获取高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) GetHighRiskPortLogsInfoList(info highPortReq.HighRiskPortConfigSearch) (list []highPort.HighRiskPortLog, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&highPort.HighRiskPortLog{}).Preload("PortConfig")
	var HRPCs []highPort.HighRiskPortLog
	if info.PortNumber != nil && *info.PortNumber != 0 {
		db = db.Joins("PortConfig")
		db.Clauses(clause.Eq{
			Column: "PortConfig.port_number",
			Value:  info.PortNumber,
		})
	}

	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&HRPCs).Error
	return HRPCs, total, err
}

// GetHighRiskPortConfigInfoList 分页获取高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) GetHighRiskPortConfigInfoList(info highPortReq.HighRiskPortConfigSearch) (list []highPort.HighRiskPortConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&highPort.HighRiskPortConfig{})
	var HRPCs []highPort.HighRiskPortConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["port_number"] = true
	orderMap["risk_level"] = true
	orderMap["order"] = true
	if orderMap[info.Sort] {
		OrderStr = "`" + info.Sort + "`"
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("`order` desc")
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&HRPCs).Error
	return HRPCs, total, err
}
func (HRPCService *HighRiskPortConfigService) GetHighRiskPortConfigPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
func (HRPCService *HighRiskPortConfigService) PortScan(params highPortReq.PortScanSearch) (err error, saveInfo highPort.HighRiskPortScan) {
	var cmd *exec.Cmd
	var execFileName = "kscan"
	//var runType = ""
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
	if params.Target != "" {
		args = append(args, "-t", params.Target)
	}
	if params.Port != "" {
		args = append(args, "-p", params.Port)
	}

	args = append(args, "-oJ", outFileName)

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
		return err, saveInfo
	}

	// Create a scanner to read the command's output
	scanner := bufio.NewScanner(stdout)
	db := global.GVA_DB.Model(&highPort.HighRiskPortScan{})
	saveInfo.TaskName = params.TaskName
	saveInfo.Port = params.Port
	saveInfo.Target = params.Target
	db.Save(&saveInfo)
	go func() {
		for scanner.Scan() {
			saveInfo.Info += scanner.Text() + "\n"
			utils.HighLogMessage[saveInfo.ID.String()] <- scanner.Text() + "\n"
			db.Save(&saveInfo)
		}

		// Wait for the command to finish
		if err = cmd.Wait(); err != nil {
			return
		}
		// 打开文件
		file, err := os.Open(outFileName)
		var fileResult string
		if err != nil {
			return
		}
		defer func() {
			utils.HighLogMessage[saveInfo.ID.String()] <- "{{over}}"
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
		saveInfo.JsonResult = fileResult
		saveInfo.Over = true
		db.Save(&saveInfo)
	}()

	return
	//scanInfo.JsonResult = fileResult
	//scanInfo.ScanOver = true
}

// GetHighRiskPortConfig 根据ID获取高危端口记录
// Author [yourname](https://github.com/yourname)
func (HRPCService *HighRiskPortConfigService) GetPortScan(ID string) (scan highPort.HighRiskPortScan, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&scan).Error
	return
}
