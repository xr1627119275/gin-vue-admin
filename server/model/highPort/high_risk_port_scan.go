// 自动生成模板HighRiskPort
package highPort

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 高危端口 结构体  HighRiskPortScan
type HighRiskPortScan struct {
	global.GVA_MODEL_UUID
	Info       string `json:"info" gorm:"column:info;type:longtext"`
	JsonResult string `json:"jsonResult" gorm:"column:json_result;type:longtext"`
	Over       bool   `json:"over" gorm:"column:over;comment:是否扫描完成"`
}

// TableName 高危端口扫描记录 HighRiskPortScan自定义表名 high_risk_port_scan
func (HighRiskPortScan) TableName() string {
	return "high_risk_port_scan"
}
